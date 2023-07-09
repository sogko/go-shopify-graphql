package main

import (
	"github.com/vinhluan/go-shopify-graphql"
)

func defaultClient() *shopify.Client {
	return shopify.NewDefaultClient()
}
