package main

import (
	"os"

	"github.com/r0busta/go-shopify-graphql/v7"
)

func clientWithVersion() *shopify.Client {
	return shopify.NewClient(os.Getenv("STORE_NAME"), shopify.WithToken(os.Getenv("STORE_ACCESS_TOKEN")), shopify.WithVersion("2022-10"))
}
