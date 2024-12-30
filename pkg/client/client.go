package client

import (
	"bytetrade.io/web3os/lldap-client/pkg/config"
	"bytetrade.io/web3os/lldap-client/pkg/transport"
	"github.com/Khan/genqlient/graphql"
	"net/http"
)

type Client struct {
	cfg *config.Config
	graphql.Client
	UserInterface
	GroupInterface
	//Transport http.RoundTripper
}

func (c *Client) Users() UserInterface {
	return NewUserRoute(c)
}

func (c *Client) Groups() GroupInterface {
	return NewGroupRoute(c)
}

func New(cfg *config.Config) (*Client, error) {
	rt, err := transport.New(cfg)
	if err != nil {
		return nil, err
	}
	httpClient := http.Client{
		Transport: rt,
	}
	graphqlClient := graphql.NewClient(cfg.Host+"/api/graphql", &httpClient)
	c := &Client{
		cfg:    cfg,
		Client: graphqlClient,
	}
	return c, nil
}
