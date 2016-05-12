package hydrantclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/t11e/go-pebbleclient"
)

type Client struct {
	client *pebbleclient.Client
}

// New constructs a new client.
func New(client *pebbleclient.Client) (*Client, error) {
	return &Client{client}, nil
}

// NewFromRequest constructs a new client from an HTTP request.
func NewFromRequest(options pebbleclient.ClientOptions, req *http.Request) (*Client, error) {
	if options.AppName == "" {
		options.AppName = "hydrant"
	}
	client, err := pebbleclient.NewFromRequest(options, req)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (client *Client) Query(query *Query, dataset string) (*ResultSet, error) {
	body, err := json.Marshal(query)
	if err != nil {
		return nil, errors.Wrap(err, "Could not marshal query to JSON")
	}

	result := ResultSet{}
	err = client.client.Post(fmt.Sprintf("/query/%s/json", dataset), pebbleclient.Body{
		Data:        body,
		ContentType: "application/json",
	}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
