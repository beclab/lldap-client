package config

import (
	"github.com/beclab/lldap-client/pkg/cache"
	"github.com/beclab/lldap-client/pkg/cache/memory"
	"net/http"
	"time"
)

type Config struct {
	Host string

	Username string
	Password string

	// Bearer token for authentication
	BearerToken string
	Timeout     time.Duration

	Transport http.RoundTripper

	TokenCache cache.TokenCacheInterface
}

func NewConfig() *Config {
	return &Config{
		Timeout:    5 * time.Second,
		TokenCache: memory.New(),
	}
}

func (c *Config) HasBasicAuth() bool {
	return len(c.Username) != 0
}

func (c *Config) HasTokenAuth() bool {
	return len(c.BearerToken) != 0
}
