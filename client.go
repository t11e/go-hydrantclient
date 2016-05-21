package hydrantclient

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
	pc "github.com/t11e/go-pebbleclient"
)

//go:generate go run vendor/github.com/vektra/mockery/cmd/mockery/mockery.go -name=Client -case=underscore

type Client interface {
	Query(query *Query, dataset string) (*ResultSet, error)
}

type client struct {
	c pc.Client
}

// New constructs a new client.
func New(pebbleClient pc.Client) (Client, error) {
	return &client{pebbleClient.Options(pc.Options{
		ServiceName: "hydrant",
		ApiVersion:  1,
	})}, nil
}

func (c *client) Query(query *Query, dataset string) (*ResultSet, error) {
	body, err := json.Marshal(query)
	if err != nil {
		return nil, errors.Wrap(err, "Could not marshal query to JSON")
	}

	result := ResultSet{}
	err = c.c.Post("/query/:dataset/json", &pc.RequestOptions{
		Params: pc.Params{
			"dataset": dataset,
		},
	}, bytes.NewReader(body), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
