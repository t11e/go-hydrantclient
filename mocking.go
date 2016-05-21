package hydrantclient

import (
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (c *MockClient) Query(query *Query, dataset string) (*ResultSet, error) {
	args := c.Called(query, dataset)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ResultSet), nil
}
