package main

import (
	"os"

	"github.com/sogko/go-shopify-graphql"
)

func clientWithVersion() *shopify.Client {
	return shopify.NewClient(os.Getenv("STORE_NAME"), shopify.WithToken(os.Getenv("STORE_ACCESS_TOKEN")), shopify.WithVersion("2022-10"))
}
