package main

import (
	"os"

	"github.com/sogko/go-shopify-graphql"
)

func clientWithToken() *shopify.Client {
	return shopify.NewClientWithToken(os.Getenv("STORE_ACCESS_TOKEN"), os.Getenv("STORE_NAME"))
}
