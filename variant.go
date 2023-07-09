package shopify

import (
	"context"
	"fmt"

	"github.com/vinhluan/go-shopify-graphql/model"
)

//go:generate mockgen -destination=./mock/variant_service.go -package=mock . VariantService
type VariantService interface {
	Update(ctx context.Context, variant model.ProductVariantInput) error
}

type VariantServiceOp struct {
	client *Client
}

var _ VariantService = &VariantServiceOp{}

type mutationProductVariantUpdate struct {
	ProductVariantUpdateResult struct {
		UserErrors []model.UserError `json:"userErrors,omitempty"`
	} `graphql:"productVariantUpdate(input: $input)" json:"productVariantUpdate"`
}

func (s *VariantServiceOp) Update(ctx context.Context, variant model.ProductVariantInput) error {
	m := mutationProductVariantUpdate{}

	vars := map[string]interface{}{
		"input": variant,
	}
	err := s.client.Mutate(ctx, &m, vars)
	if err != nil {
		return fmt.Errorf("mutation: %w", err)
	}

	if len(m.ProductVariantUpdateResult.UserErrors) > 0 {
		return fmt.Errorf("%+v", m.ProductVariantUpdateResult.UserErrors)
	}

	return nil
}
