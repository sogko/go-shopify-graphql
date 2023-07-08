package shopify

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/r0busta/graphql"
	log "github.com/sirupsen/logrus"
)

const (
	shopifyBaseDomain = "myshopify.com"

	defaultAPIProtocol       = "https"
	defaultAPIBasePath       = "admin/api"
	defaultAPIEndpoint       = "graphql.json"
	defaultShopifyAPIVersion = "2023-04"
	defaultHttpTimeout       = time.Second * 10
)

type Client struct {
	gql         graphql.GraphQL
	accessToken string
	apiKey      string
	apiBasePath string
	retries     int
	timeout     time.Duration
	transport   http.RoundTripper

	Product       ProductService
	Variant       VariantService
	Inventory     InventoryService
	Collection    CollectionService
	Order         OrderService
	Fulfillment   FulfillmentService
	Location      LocationService
	Metafield     MetafieldService
	BulkOperation BulkOperationService
}

func NewClient(shopName string, opts ...Option) *Client {
	c := &Client{
		apiBasePath: defaultAPIBasePath,
		timeout:     defaultHttpTimeout,
		transport:   http.DefaultTransport,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.gql == nil {
		url := buildAPIEndpoint(shopName, c.apiBasePath)
		httpClient := &http.Client{
			Timeout: c.timeout,
			Transport: &transport{
				accessToken:  c.accessToken,
				apiKey:       c.apiKey,
				apiBasePath:  c.apiBasePath,
				roundTripper: c.transport,
			},
		}
		c.gql = graphql.NewClient(url, httpClient)
	}

	c.Product = &ProductServiceOp{client: c}
	c.Variant = &VariantServiceOp{client: c}
	c.Inventory = &InventoryServiceOp{client: c}
	c.Collection = &CollectionServiceOp{client: c}
	c.Order = &OrderServiceOp{client: c}
	c.Fulfillment = &FulfillmentServiceOp{client: c}
	c.Location = &LocationServiceOp{client: c}
	c.Metafield = &MetafieldServiceOp{client: c}
	c.BulkOperation = &BulkOperationServiceOp{client: c}

	return c
}

func NewDefaultClient() *Client {
	apiKey := os.Getenv("STORE_API_KEY")
	accessToken := os.Getenv("STORE_PASSWORD")
	storeName := os.Getenv("STORE_NAME")
	if apiKey == "" || accessToken == "" || storeName == "" {
		log.Fatalln("Shopify Admin API Key and/or Password (aka access token) and/or store name not set")
	}

	return NewClient(storeName, WithPrivateAppAuth(apiKey, accessToken), WithVersion(defaultShopifyAPIVersion))
}

func NewClientWithToken(accessToken string, storeName string) *Client {
	if accessToken == "" || storeName == "" {
		log.Fatalln("Shopify Admin API access token and/or store name not set")
	}

	return NewClient(storeName, WithToken(accessToken), WithVersion(defaultShopifyAPIVersion))
}

func buildAPIEndpoint(shopName string, apiPathPrefix string) string {
	return fmt.Sprintf("%s://%s.%s/%s/%s", defaultAPIProtocol, shopName, shopifyBaseDomain, apiPathPrefix, defaultAPIEndpoint)
}

func (c *Client) GraphQLClient() graphql.GraphQL {
	return c.gql
}

func (c *Client) mutate(ctx context.Context, m interface{}, variables map[string]interface{}) error {
	err := c.gql.Mutate(ctx, m, variables)
	return err
}

func (c *Client) mutateString(ctx context.Context, m string, variables map[string]interface{}, out interface{}) error {
	err := c.gql.MutateString(ctx, m, variables, out)
	return err
}

func (c *Client) query(ctx context.Context, q interface{}, variables map[string]interface{}) error {
	err := c.gql.Query(ctx, q, variables)
	return err
}

func (c *Client) queryString(ctx context.Context, q string, variables map[string]interface{}, out interface{}) error {
	err := c.gql.QueryString(ctx, q, variables, out)
	return err
}
