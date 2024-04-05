package types

import (
	"fmt"

	"github.com/machinebox/graphql"
)

type RegionsRes struct {
	Regions []struct {
		Name string `json:"name"`
	} `json:"regions"`
}

func (cli *RailwayClient) GetRegions() (*RegionsRes, error) {
	query := `
		query Query {
  			regions {
    			name
  			}
		}
	`
	req := graphql.NewRequest(query)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cli.Token))
	var res *RegionsRes
	if err := cli.GraphQL.Run(cli.Context, req, &res); err != nil {
		return nil, err
	}
	return res, nil
}
