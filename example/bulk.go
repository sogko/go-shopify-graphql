package main

import (
	"context"
	"fmt"

	"github.com/sogko/go-shopify-graphql"
	"github.com/sogko/go-shopify-graphql/model"
)

func bulk(client *shopify.Client) {
	q := `
	{
		products{
			edges {
				node {
					id
					variants {
						edges {
							node {
								id
								media{
									edges {
										node {
											... on MediaImage {
												id
												image {
													url
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}`

	products := []*model.Product{}
	err := client.BulkOperation.BulkQuery(context.Background(), q, &products)
	if err != nil {
		panic(err)
	}

	// Print out the result
	for _, p := range products {
		for _, v := range p.Variants.Edges {
			for _, m := range v.Node.Media.Edges {
				fmt.Println(m.Node.(*model.MediaImage).Image.URL)
			}
		}
	}
}
