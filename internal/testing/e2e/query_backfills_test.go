package e2e

import (
	"context"
	"io"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"open-match.dev/open-match/internal/filter/testcases"
	"open-match.dev/open-match/pkg/pb"
)

const (
	firstBackfillGeneration = 1
)

func TestQueryBackfillsWithEmptyPool(t *testing.T) {
	om := newOM(t)
	stream, err := om.Query().QueryBackfills(context.Background(), &pb.QueryBackfillsRequest{Pool: nil})
	require.NoError(t, err)

	resp, err := stream.Recv()
	require.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
	require.Nil(t, resp)
}

func TestNoBackfills(t *testing.T) {
	om := newOM(t)
	stream, err := om.Query().QueryBackfills(context.Background(), &pb.QueryBackfillsRequest{Pool: &pb.Pool{}})
	require.NoError(t, err)

	resp, err := stream.Recv()
	require.Equal(t, io.EOF, err)
	require.Nil(t, resp)
}

func TestQueryBackfillsPaging(t *testing.T) {
	om := newOM(t)

	pageSize := 10 // TODO: read from config
	if pageSize < 1 {
		require.Fail(t, "invalid page size")
	}

	total := pageSize*5 + 1
	expectedIds := map[string]struct{}{}

	for i := 0; i < total; i++ {
		resp, err := om.Frontend().CreateBackfill(context.Background(), &pb.CreateBackfillRequest{Backfill: &pb.Backfill{}})
		require.NotNil(t, resp)
		require.NoError(t, err)

		expectedIds[resp.Id] = struct{}{}
	}

	stream, err := om.Query().QueryBackfills(context.Background(), &pb.QueryBackfillsRequest{Pool: &pb.Pool{}})
	require.NoError(t, err)

	foundIds := map[string]struct{}{}

	for i := 0; i < 5; i++ {
		var resp *pb.QueryBackfillsResponse
		resp, err = stream.Recv()
		require.NoError(t, err)
		require.Equal(t, pageSize, len(resp.Backfills))

		for _, backfill := range resp.Backfills {
			foundIds[backfill.Id] = struct{}{}
		}
	}

	resp, err := stream.Recv()
	require.NoError(t, err)
	require.Equal(t, len(resp.Backfills), 1)
	foundIds[resp.Backfills[0].Id] = struct{}{}

	require.Equal(t, expectedIds, foundIds)

	resp, err = stream.Recv()
	require.Equal(t, err, io.EOF)
	require.Nil(t, resp)
}

func TestBackfillFound(t *testing.T) {
	for _, tc := range testcases.IncludedTestCases() {
		tc := tc
		t.Run("QueryBackfills_"+tc.Name, func(t *testing.T) {
			if !backfillReturnedByQuery(t, tc) {
				require.Fail(t, "Expected to find backfill in pool but didn't.")
			}
		})
	}
}

func TestBackfillNotFound(t *testing.T) {
	for _, tc := range testcases.ExcludedTestCases() {
		tc := tc
		t.Run("QueryBackfills_"+tc.Name, func(t *testing.T) {
			if backfillReturnedByQuery(t, tc) {
				require.Fail(t, "Expected to not find backfill in pool but did.")
			}
		})
	}
}

func TestBackfillQueryAfterMMFUpdate(t *testing.T) {
	ctx := context.Background()
	om := newOM(t)
	backfill := &pb.Backfill{
		SearchFields: &pb.SearchFields{
			StringArgs: map[string]string{
				"field": "value",
			},
		},
	}
	pool := &pb.Pool{
		StringEqualsFilters: []*pb.StringEqualsFilter{{StringArg: "field", Value: "value"}},
	}
	match := &pb.Match{
		MatchId: "1",
		Tickets: []*pb.Ticket{},
	}
	{
		resp, err := om.Frontend().CreateBackfill(ctx, &pb.CreateBackfillRequest{Backfill: backfill})
		require.Nil(t, err)
		require.NotNil(t, resp)
		match.Backfill = resp
	}

	om.SetMMF(func(ctx context.Context, profile *pb.MatchProfile, out chan<- *pb.Match) error {
		out <- match
		return nil
	})

	om.SetEvaluator(func(ctx context.Context, in <-chan *pb.Match, out chan<- string) error {
		p, ok := <-in
		require.True(t, ok)

		out <- p.MatchId
		return nil
	})
	{
		stream, err := om.Backend().FetchMatches(ctx, &pb.FetchMatchesRequest{
			Config:  om.MMFConfigGRPC(),
			Profile: &pb.MatchProfile{},
		})
		require.Nil(t, err)

		resp, err := stream.Recv()
		require.Nil(t, err)
		require.True(t, proto.Equal(match, resp.Match))

		resp, err = stream.Recv()
		require.Nil(t, resp)
		require.Equal(t, io.EOF, err)
	}
	{
		stream, err := om.Query().QueryBackfills(context.Background(), &pb.QueryBackfillsRequest{Pool: pool})
		require.NoError(t, err)
		resp, err := stream.Recv()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, int64(firstBackfillGeneration), resp.Backfills[0].Generation)

		resp, err = stream.Recv()
		require.Equal(t, io.EOF, err)
		require.Nil(t, resp)
	}
}

func TestBackfillQueryAfterGSUpdate(t *testing.T) {
	om := newOM(t)
	backfill := &pb.Backfill{
		SearchFields: &pb.SearchFields{
			StringArgs: map[string]string{
				"field": "value",
			},
		},
	}
	pool := &pb.Pool{
		StringEqualsFilters: []*pb.StringEqualsFilter{{StringArg: "field", Value: "value"}},
	}
	{
		resp, err := om.Frontend().CreateBackfill(context.Background(), &pb.CreateBackfillRequest{Backfill: backfill})
		require.NotNil(t, resp)
		require.NoError(t, err)
		require.Equal(t, int64(firstBackfillGeneration), resp.Generation)
	}
	{
		stream, err := om.Query().QueryBackfills(context.Background(), &pb.QueryBackfillsRequest{Pool: pool})
		require.NoError(t, err)
		resp, err := stream.Recv()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, int64(firstBackfillGeneration), resp.Backfills[0].Generation)
		backfill = resp.Backfills[0]

		resp, err = stream.Recv()
		require.Equal(t, io.EOF, err)
		require.Nil(t, resp)
	}
	{
		resp, err := om.Frontend().UpdateBackfill(context.Background(), &pb.UpdateBackfillRequest{Backfill: backfill})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, int64(firstBackfillGeneration+1), resp.Generation)
	}
	{
		stream, err := om.Query().QueryBackfills(context.Background(), &pb.QueryBackfillsRequest{Pool: pool})
		require.NoError(t, err)
		resp, err := stream.Recv()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, int64(firstBackfillGeneration+1), resp.Backfills[0].Generation)

		resp, err = stream.Recv()
		require.Equal(t, io.EOF, err)
		require.Nil(t, resp)
	}
}

func backfillReturnedByQuery(t *testing.T, tc testcases.TestCase) (found bool) {
	om := newOM(t)

	{
		backfill := pb.Backfill{
			SearchFields: tc.SearchFields,
		}
		resp, err := om.Frontend().CreateBackfill(context.Background(), &pb.CreateBackfillRequest{Backfill: &backfill})

		require.NotNil(t, resp)
		require.NoError(t, err)
	}

	stream, err := om.Query().QueryBackfills(context.Background(), &pb.QueryBackfillsRequest{Pool: tc.Pool})
	require.NoError(t, err)

	backfills := []*pb.Backfill{}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)

		backfills = append(backfills, resp.Backfills...)
	}

	if len(backfills) > 1 {
		require.Fail(t, "More than one backfill found")
	}

	return len(backfills) == 1
}
