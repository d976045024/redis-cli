package util

import (
	"context"
	"testing"

	"github.com/redis-cli/pkg"
	"github.com/redis-cli/pkg/constants"
	"github.com/redis-cli/pkg/pointer"
	"github.com/stretchr/testify/assert"
)

func TestGetSet(t *testing.T) {
	r, err := pkg.NewRedisRunner(make(map[string]any))
	assert.NoError(t, err)
	ctx := context.Background()
	query, err := r.Query(ctx, pointer.String("SET"), "foo", "bar")
	assert.NoError(t, err)
	assert.Equal(t, query.Status, constants.REDIS_STATUS_SUCCESS)

	response, err := r.Query(ctx, pointer.String("GET"), "foo")
	assert.NoError(t, err)
	assert.Equal(t, response.Data, "bar")
	assert.Equal(t, response.Status, constants.REDIS_STATUS_SUCCESS)
}
