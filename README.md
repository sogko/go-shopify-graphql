# go-shopify-graphql

**Preface:** This is a fork from https://github.com/r0busta/go-shopify-graphql which extended features (retry capability, optionally set timeout, transport)

A simple client using the Shopify GraphQL Admin API.

## Getting started

Hello World example

### 0. Setup

```bash
export STORE_API_KEY=<private_app_api_key>
export STORE_PASSWORD=<private_app_access_token>
export STORE_NAME=<store_name>
```

### 1. Program

```go
package main

import (
	"context"
	"fmt"

	shopify "github.com/vinhluan/go-shopify-graphql"
)

func main() {
	// Create client
	client := shopify.NewDefaultClient()

	// Get all collections
	collections, err := client.Collection.ListAll(context.Background())
	if err != nil {
		panic(err)
	}

	// Print out the result
	for _, c := range collections {
		fmt.Println(c.Handle)
	}
}
```

### 3. Run

```bash
go run .
```
