package clients

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSomething(t *testing.T) {
	ctx := context.Background()
	require.False(t, isContextDone(ctx))
}
