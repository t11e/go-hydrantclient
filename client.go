package hydrantclient

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
	pc "github.com/t11e/go-pebbleclient"
)

type Client struct {
	c pc.Client
}

// New constructs a new client.
func New(client pc.Client) (*Client, error) {
	return &Client{client}, nil
}

func (client *Client) Query(query *Query, dataset string) (*ResultSet, error) {
	body, err := json.Marshal(query)
	if err != nil {
		return nil, errors.Wrap(err, "Could not marshal query to JSON")
	}

	result := ResultSet{}
	err = client.c.Post("/query/:dataset/json", &pc.RequestOptions{
		Params: pc.Params{
			"dataset": dataset,
		},
	}, bytes.NewReader(body), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
