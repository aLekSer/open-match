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

package e2e

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"open-match.dev/open-match/pkg/pb"
)

// TestCreateGetBackfill Create and Get Backfill test
func TestCreateGetBackfill(t *testing.T) {
	om := newOM(t)
	ctx := context.Background()

	bf := &pb.CreateBackfillRequest{Backfill: &pb.Backfill{SearchFields: &pb.SearchFields{
		StringArgs: map[string]string{
			"search": "me",
		},
	}}}
	b1, err := om.Frontend().CreateBackfill(ctx, bf)
	require.Nil(t, err)

	b2, err := om.Frontend().CreateBackfill(ctx, bf)
	require.Nil(t, err)

	// Different Ids should be generated
	assert.NotEqual(t, b1.Id, b2.Id)
	b1.Id = b2.Id
	b1.CreateTime = b2.CreateTime
	// Other than CreateTune abd Id fields, they should be equal
	assert.Equal(t, b1, b2)
	require.Nil(t, err)
	get, err := om.Frontend().GetBackfill(ctx, &pb.GetBackfillRequest{BackfillId: b1.Id})
	require.Nil(t, err)
	require.Equal(t, b1, get)
}

// TestUpdateBackfill Update Backfill test
func TestUpdateBackfill(t *testing.T) {
	om := newOM(t)
	ctx := context.Background()

	bf := &pb.Backfill{SearchFields: &pb.SearchFields{
		StringArgs: map[string]string{
			"search": "me",
		},
	},
	}
	bfCreated, err := om.Frontend().CreateBackfill(ctx, &pb.CreateBackfillRequest{Backfill: bf})
	require.Nil(t, err)

	bfCreated.SearchFields.StringArgs["key"] = "val"
	bfUpdated, err := om.Frontend().UpdateBackfill(ctx, &pb.UpdateBackfillRequest{BackfillTicket: bfCreated})
	require.Nil(t, err)

	// No changes to CreateTime
	bfCreated.CreateTime = bfUpdated.CreateTime

	require.Nil(t, err)
	get, err := om.Frontend().GetBackfill(ctx, &pb.GetBackfillRequest{BackfillId: bf.Id})
	require.Nil(t, err)
	require.Equal(t, bf, get)
}
