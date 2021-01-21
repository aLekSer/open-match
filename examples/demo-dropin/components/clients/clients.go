// Copyright 2019 Google LLC
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

package clients

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"

	"open-match.dev/open-match/examples/demo-dropin/components"
	"open-match.dev/open-match/examples/demo-dropin/updater"
	"open-match.dev/open-match/pkg/pb"
)

const (
	playersPerMatch = 30000
	openSlotsKey    = "open-slots"
	matchName       = "backfill-matchfunction"
)

func Run(ds *components.DemoShared) {
	u := updater.NewNested(ds.Ctx, ds.Update)

	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("fakeplayer_%d", i)
		go func() {
			for !isContextDone(ds.Ctx) {
				runScenario(ds.Ctx, name, u.ForField(name), i)
			}
		}()
	}
	name := fmt.Sprintf("backfill_%d", 0)
	go func() {
		for !isContextDone(ds.Ctx) {
			runBackfillScenario(ds.Ctx, name, u.ForField(name), 0)
		}
	}()
}

func isContextDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

type status struct {
	Status     string
	Assignment *pb.Assignment
	Backfills  []pb.Backfill
}

var once = true

var bfIds = map[string]string{}

func runBackfillScenario(ctx context.Context, name string, update updater.SetFunc, i int) {
	regions := []string{"East", "North"}

	defer func() {
		r := recover()
		if r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
			update(status{Status: fmt.Sprintf("Encountered error: %s", err.Error())})
			time.Sleep(time.Second * 10)
		}
	}()

	// See https://open-match.dev/site/docs/guides/api/
	conn, err := grpc.Dial("open-match-frontend.open-match.svc.cluster.local:50504", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fe := pb.NewFrontendServiceClient(conn)

	if once {

		for _, i := range regions {

			backfill := pb.Backfill{
				SearchFields: &pb.SearchFields{
					StringArgs: map[string]string{
						"region": i,
					},
				},
			}
			setOpenSlots(&backfill, playersPerMatch)
			bf := &pb.CreateBackfillRequest{
				Backfill: &backfill,
			}
			resp, err := fe.CreateBackfill(ctx, bf)
			if err != nil {
				panic(fmt.Sprintf("h %s", err.Error()))
			}
			bfIds[i] = resp.Id

		}
		once = false
	}
	bfs := []pb.Backfill{}
	for _, i := range regions {
		bf, err := fe.GetBackfill(ctx, &pb.GetBackfillRequest{BackfillId: bfIds[i]})
		if err != nil {
			panic(err)
		}
		bfs = append(bfs, *bf)

		bf, err = fe.AcknowledgeBackfill(ctx, &pb.AcknowledgeBackfillRequest{BackfillId: bfIds[i], Assignment: &pb.Assignment{Connection: fmt.Sprintf("backfill.%s", i)}})
		if err != nil {
			panic(err)
		}
		bfs = append(bfs, *bf)
	}

	s := status{}
	s.Backfills = bfs
	update(s)
}

func runScenario(ctx context.Context, name string, update updater.SetFunc, i int) {
	defer func() {
		r := recover()
		if r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}

			update(status{Status: fmt.Sprintf("Encountered error: %s", err.Error())})
			time.Sleep(time.Second * 10)
		}
	}()

	s := status{}

	//////////////////////////////////////////////////////////////////////////////
	s.Status = "Main Menu"
	update(s)

	time.Sleep(time.Duration(rand.Int63()) % (time.Second * 15))

	//////////////////////////////////////////////////////////////////////////////
	s.Status = "Connecting to Open Match frontend"
	update(s)

	// See https://open-match.dev/site/docs/guides/api/
	conn, err := grpc.Dial("open-match-frontend.open-match.svc.cluster.local:50504", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fe := pb.NewFrontendServiceClient(conn)
	helper(ctx, fe, i, update, 0)
	s.Status = "Switching the World GameServer ..."
	s.Assignment = nil
	update(s)
	time.Sleep(2 * time.Second)
	helper(ctx, fe, i, update, 1)
}

func helper(ctx context.Context, fe pb.FrontendServiceClient, i int, update updater.SetFunc, change int) {
	s := status{}
	//////////////////////////////////////////////////////////////////////////////
	s.Status = "Creating Open Match Ticket"
	update(s)

	var ticketId string
	var region string

	{
		if i%2 == change {
			region = "North"
		} else {
			region = "East"
		}
		req := &pb.CreateTicketRequest{
			Ticket: &pb.Ticket{
				SearchFields: &pb.SearchFields{
					StringArgs: map[string]string{
						"region": region,
					},
				},
			},
		}

		resp, err := fe.CreateTicket(ctx, req)
		if err != nil {
			panic(err)
		}
		ticketId = resp.Id
	}

	//////////////////////////////////////////////////////////////////////////////
	s.Status = fmt.Sprintf("Waiting match with ticket Id %s", ticketId)
	update(s)

	var assignment *pb.Assignment
	{
		req := &pb.WatchAssignmentsRequest{
			TicketId: ticketId,
		}

		stream, err := fe.WatchAssignments(ctx, req)
		for assignment.GetConnection() == "" {
			resp, err := stream.Recv()
			if err != nil {
				// For now we don't expect to get EOF, so that's still an error worthy of panic.
				panic(err)
			}

			assignment = resp.Assignment
		}

		err = stream.CloseSend()
		if err != nil {
			panic(err)
		}
	}

	//////////////////////////////////////////////////////////////////////////////
	s.Status = "Sleeping (pretend this is playing a match...)"
	s.Assignment = assignment
	update(s)

	time.Sleep(time.Second * 10)
}

func setOpenSlots(b *pb.Backfill, val int32) error {
	if b.Extensions == nil {
		b.Extensions = make(map[string]*any.Any)
	}

	any, err := ptypes.MarshalAny(&wrappers.Int32Value{Value: val})
	if err != nil {
		return err
	}

	b.Extensions[openSlotsKey] = any
	return nil
}
