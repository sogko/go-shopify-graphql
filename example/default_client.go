package main

import (
	"github.com/sogko/go-shopify-graphql"
)

func defaultClient() *shopify.Client {
	return shopify.NewDefaultClient()
}
