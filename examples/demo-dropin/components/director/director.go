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

package director

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"

	"open-match.dev/open-match/examples/demo-dropin/components"
	"open-match.dev/open-match/examples/demo-dropin/updater"
	"open-match.dev/open-match/pkg/pb"
)

const (
	regionArg = "region"
	host      = "om-function.open-match-demo.svc.cluster.local"
)

func Run(ds *components.DemoShared) {
	for !isContextDone(ds.Ctx) {
		run(ds)
	}
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
	Status        string
	LatestMatches []*pb.Match
}

// Used to store previous matches between calls
var latestMatches1 []*pb.Match
var latestMatches2 []*pb.Match

func run(ds *components.DemoShared) {
	u := updater.NewNested(ds.Ctx, ds.Update)
	update := u.ForField("Director 1 subprocess")
	update2 := u.ForField("Director 2 subprocess")
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
	s.Status = "Connecting to backend"
	update(s)

	// See https://open-match.dev/site/docs/guides/api/
	conn, err := grpc.Dial("open-match-backend.open-match.svc.cluster.local:50505", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	be := pb.NewBackendServiceClient(conn)

	//////////////////////////////////////////////////////////////////////////////
	s.Status = "Match Match: Sending Request"
	s.LatestMatches = latestMatches1
	update(s)

	go func() {
		s := status{}
		var matches []*pb.Match
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
		req := &pb.FetchMatchesRequest{
			Config: &pb.FunctionConfig{
				Host: host,
				Port: 50502,
				Type: pb.FunctionConfig_GRPC,
			},
			Profile: &pb.MatchProfile{
				Name: "North",
				Pools: []*pb.Pool{
					{
						Name: "1",
						StringEqualsFilters: []*pb.StringEqualsFilter{
							{
								StringArg: regionArg,
								Value:     "North",
							},
						},
					},
				},
			},
		}

		stream, err := be.FetchMatches(ds.Ctx, req)
		if err != nil {
			panic(err)
		}

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			matches = append(matches, resp.GetMatch())
			latestMatches1 = matches
		}
		//////////////////////////////////////////////////////////////////////////////
		s.Status = "Matches Found"
		s.LatestMatches = latestMatches1
		update(s)
	}()
	{
		s := status{}
		var matches []*pb.Match
		req2 := &pb.FetchMatchesRequest{
			Config: &pb.FunctionConfig{
				Host: host,
				Port: 50502,
				Type: pb.FunctionConfig_GRPC,
			},
			Profile: &pb.MatchProfile{
				Name: "East",
				Pools: []*pb.Pool{
					{
						Name: "2",
						StringEqualsFilters: []*pb.StringEqualsFilter{
							{
								StringArg: regionArg,
								Value:     "East",
							},
						},
					},
				},
			},
		}

		stream, err := be.FetchMatches(ds.Ctx, req2)
		if err != nil {
			panic(err)
		}

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			matches = append(matches, resp.GetMatch())
			latestMatches2 = matches
		}
		//////////////////////////////////////////////////////////////////////////////
		s.Status = "Matches Found"
		s.LatestMatches = latestMatches2
		update2(s)
	}

	/*
		//////////////////////////////////////////////////////////////////////////////
		s.Status = "Assigning Players"
		update(s)

		for _, match := range matches {
			ids := []string{}

			for _, t := range match.Tickets {
				ids = append(ids, t.Id)
			}

			req := &pb.AssignTicketsRequest{
				Assignments: []*pb.AssignmentGroup{
					{
						TicketIds: ids,
						Assignment: &pb.Assignment{
							Connection: fmt.Sprintf("%d.%d.%d.%d:2222", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)),
						},
					},
				},
			}

			resp, err := be.AssignTickets(ds.Ctx, req)
			if err != nil {
				panic(err)
			}

			_ = resp
		}
	*/

	//////////////////////////////////////////////////////////////////////////////
	s.Status = "Sleeping"
	update(s)

	time.Sleep(time.Second * 5)
}
