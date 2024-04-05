package types

import (
	"context"

	"github.com/machinebox/graphql"
)

const PUBLIC_API = "https://backboard.railway.app/graphql/v2"

type RailwayClient struct {
	GraphQL  *graphql.Client
	Context  context.Context
	Cancel   context.CancelFunc
	Endpoint string
	Token    string
}

func NewRailwayClient(token string) *RailwayClient {
	ctx, cancel := context.WithCancel(context.Background())
	cli := &RailwayClient{
		Endpoint: PUBLIC_API,
		Token:    token,
		Context:  ctx,
		Cancel:   cancel,
	}
	gql := graphql.NewClient(cli.Endpoint)
	cli.GraphQL = gql
	return cli
}
