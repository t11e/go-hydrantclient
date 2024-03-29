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

// Register registers us in a connector.
func Register(connector *pc.Connector) {
	connector.Register((*Client)(nil), func(client pc.Client) (pc.Service, error) {
		return New(client)
	})
}

// New constructs a new client.
func New(pebbleClient pc.Client) (Client, error) {
	return &client{pebbleClient.WithOptions(pc.Options{
		ServiceName: "hydrant",
		APIVersion:  1,
	})}, nil
}

func (c *client) Query(query *Query, dataset string) (*ResultSet, error) {
	body, err := json.Marshal(query)
	if err != nil {
		return nil, errors.Wrap(err, "Could not marshal query to JSON")
	}

	result := ResultSet{}
	err = c.c.Post("/query/:dataset", &pc.RequestOptions{
		Params: pc.Params{
			"dataset": dataset,
		},
	}, bytes.NewReader(body), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
