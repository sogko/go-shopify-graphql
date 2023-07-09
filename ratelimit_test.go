package shopify_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vinhluan/go-shopify-graphql"
)

func TestCalculateWaitTime(t *testing.T) {
	data := map[string]any{"cost": map[string]any{
		"requestedQueryCost": 101,
		"actualQueryCost":    46,
		"throttleStatus": map[string]any{
			"maximumAvailable":   1000,
			"currentlyAvailable": 954,
			"restoreRate":        50,
		},
	}}
	wait := shopify.CalculateWaitTime(data)
	assert.Zero(t, wait)

	data = map[string]any{"cost": map[string]any{
		"requestedQueryCost": 500,
		"actualQueryCost":    46,
		"throttleStatus": map[string]any{
			"maximumAvailable":   1000,
			"currentlyAvailable": 154,
			"restoreRate":        50,
		},
	}}
	wait = shopify.CalculateWaitTime(data)
	// math.Ceil((500 - 154) / 50) = 7
	assert.Equal(t, 7*time.Second, wait)
}
