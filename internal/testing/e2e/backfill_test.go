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
	createdBf, err := om.Frontend().CreateBackfill(ctx, &pb.CreateBackfillRequest{Backfill: bf})
	require.Nil(t, err)

	createdBf.SearchFields.StringArgs["key"] = "val"
	updatedBf, err := om.Frontend().UpdateBackfill(ctx, &pb.UpdateBackfillRequest{BackfillTicket: createdBf})
	require.Nil(t, err)

	// No changes to CreateTime
	assert.Equal(t, createdBf.CreateTime.GetNanos(), updatedBf.CreateTime.GetNanos())

	get, err := om.Frontend().GetBackfill(ctx, &pb.GetBackfillRequest{BackfillId: createdBf.Id})
	require.Nil(t, err)
	require.Equal(t, createdBf.SearchFields.StringArgs, get.SearchFields.StringArgs)
	_, err = om.Frontend().DeleteBackfill(ctx, &pb.DeleteBackfillRequest{BackfillId: createdBf.Id})
	require.Nil(t, err)

	get, err = om.Frontend().GetBackfill(ctx, &pb.GetBackfillRequest{BackfillId: createdBf.Id})
	require.Error(t, err, "Backfill id: bup6oduvvhfj2n8o86b0 not foun4d")
	require.Nil(t, get)
}

// TestAcknowledgeBackfill Update Backfill test
func TestAcknowledgeBackfill(t *testing.T) {
	om := newOM(t)
	ctx := context.Background()

	bf := &pb.Backfill{SearchFields: &pb.SearchFields{
		StringArgs: map[string]string{
			"search": "me",
		},
	},
	}
	createdBf, err := om.Frontend().CreateBackfill(ctx, &pb.CreateBackfillRequest{Backfill: bf})
	require.Nil(t, err)

	acknowledgedBf, err := om.Frontend().AcknowledgeBackfill(ctx, &pb.AcknowledgeBackfillRequest{BackfillId: createdBf.Id, Assignment: &pb.Assignment{Connection: "127.0.0.1:54000"}})
	require.Equal(t, acknowledgedBf.Id, createdBf.Id)
}
