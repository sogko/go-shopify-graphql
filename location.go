package shopify

import (
	"context"
	"fmt"

	"github.com/vinhluan/go-shopify-graphql/model"
)

//go:generate mockgen -destination=./mock/location_service.go -package=mock . LocationService
type LocationService interface {
	Get(ctx context.Context, id string) (*model.Location, error)
}

type LocationServiceOp struct {
	client *Client
}

var _ LocationService = &LocationServiceOp{}

func (s *LocationServiceOp) Get(ctx context.Context, id string) (*model.Location, error) {
	q := `query location($id: ID!) {
		location(id: $id){
			id
			name
		}
	}`

	vars := map[string]interface{}{
		"id": id,
	}

	var out struct {
		*model.Location `json:"location"`
	}
	err := s.client.QueryString(ctx, q, vars, &out)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return out.Location, nil
}
