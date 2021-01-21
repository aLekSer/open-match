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
	"log"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pkg/errors"
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

	for i := 0; i < 4; i++ {
		name := fmt.Sprintf("fakeplayer_%d", i)
		go func() {
			for !isContextDone(ds.Ctx) {
				runScenario(ds.Ctx, name, u.ForField(name), i)
			}
		}()
	}
	go func() {
		for !isContextDone(ds.Ctx) {
			runBackfillScenario(ds.Ctx, 0, u)
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
	Ticket     pb.Ticket
}

type backfillStatus struct {
	Status    string
	Backfill  pb.Backfill
	OpenSlots int32
}

var once = true

var bfIds = map[string]string{}

var updaters = map[string]updater.SetFunc{}

func runBackfillScenario(ctx context.Context, i int, u *updater.Updater) {
	regions := []string{"East", "North"}

	defer func() {
		r := recover()
		if r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
			update := u.ForField("Backfills")
			log.Printf("Encountered error in BackfillScenario: %s", err.Error())
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
			updaters[i] = u.ForField("backfill_" + i)
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
				panic(errors.Wrapf(err, "create BF error"))
			}
			log.Printf("Backfill %s was created %v", i, resp)
			bfIds[i] = resp.Id

		}
		once = false
	}
	s := backfillStatus{}
	for _, i := range regions {
		bf, err := fe.AcknowledgeBackfill(ctx, &pb.AcknowledgeBackfillRequest{BackfillId: bfIds[i], Assignment: &pb.Assignment{Connection: fmt.Sprintf("backfill.%s", i)}})
		if err != nil {
			panic(errors.Wrapf(err, "p1"))
		}
		os, err := getOpenSlots(bf)
		if err != nil {
			panic(errors.Wrapf(err, "p2"))
		}
		s.OpenSlots = os
		if bf != nil {
			log.Printf("bf id %s , open slots: %d", bf.Id, os)
			s.Backfill = *bf
		}
		updaters[i](s)
	}
}

func runScenario(ctx context.Context, name string, update updater.SetFunc, i int) {
	defer func() {
		r := recover()
		if r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}

			log.Printf("Encountered error: %s", err.Error())
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
		s.Ticket = *resp
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
		if err != nil {
			panic(err)
		}
		for assignment.GetConnection() == "" {
			resp, err := stream.Recv()
			if err != nil {
				// For now we don't expect to get EOF, so that's still an error worthy of panic.
				log.Println("assignment err:", err)
				panic(err)
			}
			log.Println("New assignment ", assignment)

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

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10_000)))
	bf, err := fe.GetBackfill(ctx, &pb.GetBackfillRequest{BackfillId: bfIds[region]})
	if err != nil {
		panic(errors.Wrapf(err, "%v", bf))
	}
	n, err := getOpenSlots(bf)
	if err != nil {
		panic(err)
	}
	n++
	err = setOpenSlots(bf, n)
	if err != nil {
		panic(err)
	}
	bf, err = fe.UpdateBackfill(ctx, &pb.UpdateBackfillRequest{Backfill: bf})
	if err != nil {
		panic(errors.Wrapf(err, "%v", bf))
	}
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

func getOpenSlots(b *pb.Backfill) (int32, error) {
	if b == nil {
		return 0, fmt.Errorf("expected backfill is not nil")
	}

	if b.Extensions != nil {
		if any, ok := b.Extensions[openSlotsKey]; ok {
			var val wrappers.Int32Value
			err := ptypes.UnmarshalAny(any, &val)
			if err != nil {
				return 0, err
			}

			return val.Value, nil
		}
	}

	return playersPerMatch, nil
}
