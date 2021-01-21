package director

import (
	"context"
	"testing"

	"github.com/tj/assert"
)

func TestSomething(t *testing.T) {
	ctx := context.Background()
	assert.False(t, isContextDone(ctx))
}
