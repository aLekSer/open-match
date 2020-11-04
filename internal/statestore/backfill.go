// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package statestore

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"open-match.dev/open-match/pkg/pb"
)

// CreateBackfill creates a new Backfill in the state storage. If the id already exists, it will be overwritten.
func (rb *redisBackend) CreateBackfill(ctx context.Context, backfill *pb.Backfill) error {
	redisConn, err := rb.redisPool.GetContext(ctx)
	if err != nil {
		return status.Errorf(codes.Unavailable, "CreateBackfill, id: %s, failed to connect to redis: %v", backfill.GetId(), err)
	}
	defer handleConnectionClose(&redisConn)

	value, err := proto.Marshal(backfill)
	if err != nil {
		err = errors.Wrapf(err, "failed to marshal the backfill proto, id: %s", backfill.GetId())
		return status.Errorf(codes.Internal, "%v", err)
	}

	_, err = redisConn.Do("SET", backfill.GetId(), value)
	if err != nil {
		err = errors.Wrapf(err, "failed to set the value for backfill, id: %s", backfill.GetId())
		return status.Errorf(codes.Internal, "%v", err)
	}

	return nil
}

// GetBackfill gets the Backfill with the specified id from state storage. This method fails if the Backfill does not exist.
func (rb *redisBackend) GetBackfill(ctx context.Context, id string) (*pb.Backfill, error) {
	redisConn, err := rb.redisPool.GetContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "GetBackfill, id: %s, failed to connect to redis: %v", id, err)
	}
	defer handleConnectionClose(&redisConn)

	return getBackfill(redisConn, id)
}

// DeleteBackfill removes the Backfill with the specified id from state storage. This method succeeds if the Backfill does not exist.
func (rb *redisBackend) DeleteBackfill(ctx context.Context, id string) error {
	redisConn, err := rb.redisPool.GetContext(ctx)
	if err != nil {
		return status.Errorf(codes.Unavailable, "DeleteBackfill, id: %s, failed to connect to redis: %v", id, err)
	}
	defer handleConnectionClose(&redisConn)

	_, err = redisConn.Do("DEL", id)
	if err != nil {
		err = errors.Wrapf(err, "failed to delete the backfill from state storage, id: %s", id)
		return status.Errorf(codes.Internal, "%v", err)
	}

	return nil
}

// UpdateBackfill updates an existing Backfill with a new data. Caller has to provide a custom updateFunc if this function is called not for the game server.
func (rb *redisBackend) UpdateBackfill(ctx context.Context, isGS bool, backfill *pb.Backfill, updateFunc func(current *pb.Backfill, new *pb.Backfill) (*pb.Backfill, error)) (*pb.Backfill, error) {
	var backfillBytes []byte
	var err error
	redisConn, err := rb.redisPool.GetContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "UpdateBackfill, id: %s, failed to connect to redis: %v", backfill.GetId(), err)
	}
	defer handleConnectionClose(&redisConn)

	if isGS {
		backfillBytes, err = proto.Marshal(backfill)
		if err != nil {
			err = errors.Wrapf(err, "failed to marshal the backfill proto, id: %s", backfill.GetId())
			return nil, status.Errorf(codes.Internal, "%v", err)
		}

		err = redisConn.Send("SET", backfill.GetId(), backfillBytes)
		if err != nil {
			err = errors.Wrapf(err, "failed to set the value for backfill, id: %s", backfill.GetId())
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
	} else {
		if updateFunc == nil {
			return nil, status.Errorf(codes.Internal, "nil updateFunc provided")
		}

		_, err = redisConn.Do("WATCH", backfill.GetId())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to WATCH, id: %s", backfill.GetId())
		}

		currentBackfill, err := getBackfill(redisConn, backfill.GetId())
		if err != nil {
			return nil, err
		}

		// update current backfill by invoking an updateFunc which is implemented on the caller side
		backfillToSet, err := updateFunc(currentBackfill, backfill)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		backfillBytes, err = proto.Marshal(backfillToSet)
		if err != nil {
			err = errors.Wrapf(err, "failed to marshal the backfill proto, id: %s", backfillToSet.GetId())
			return nil, status.Errorf(codes.Internal, "%v", err)
		}

		err = redisConn.Send("MULTI")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to MULTI, id: %s", backfill.GetId())
		}

		err = redisConn.Send("SET", backfill.GetId(), backfillBytes)
		if err != nil {
			err = errors.Wrapf(err, "failed to set the value for backfill, id: %s", backfillToSet.GetId())
			return nil, status.Errorf(codes.Internal, "%v", err)
		}

		_, err = redis.Values(redisConn.Do("EXEC"))
		if err != nil {
			if err == redis.ErrNil {
				redisLogger.WithField("backfillID", backfillToSet.GetId()).Debug("Backfill was modified, transaction aborted. Will try again...")
				return nil, err
			}

			err = errors.Wrapf(err, "failed to set the value for backfill, id: %s", backfillToSet.GetId())
			return nil, status.Errorf(codes.Internal, "%v", err)
		}

		return backfillToSet, nil
	}

	return backfill, nil
}

func getBackfill(redisConn redis.Conn, id string) (*pb.Backfill, error) {
	value, err := redis.Bytes(redisConn.Do("GET", id))
	if err != nil {
		// Return NotFound if redigo did not find the backfill in storage.
		if err == redis.ErrNil {
			return nil, status.Errorf(codes.NotFound, "Backfill id: %s not found", id)
		}

		err = errors.Wrapf(err, "failed to get the backfill from state storage, id: %s", id)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	if value == nil {
		return nil, status.Errorf(codes.NotFound, "Backfill id: %s not found", id)
	}

	backfill := &pb.Backfill{}
	err = proto.Unmarshal(value, backfill)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal the backfill proto, id: %s", id)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return backfill, nil
}
