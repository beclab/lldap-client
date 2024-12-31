package memory

import (
	"github.com/beclab/lldap-client/pkg/cache"
	"k8s.io/klog/v2"
	"sync"
	"time"
)

type TokenCache struct {
	token        string
	refreshToken string
	expiresAt    time.Time
	mu           sync.RWMutex
}

var (
	tokenCache *TokenCache
	once       sync.Once
)

func New() cache.TokenCacheInterface {
	once.Do(func() {
		tokenCache = &TokenCache{}
	})
	return tokenCache
}

func (c *TokenCache) Set(token string, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.token = token
	c.expiresAt = time.Now().Add(ttl)
	return nil
}

func (c *TokenCache) Get() (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	klog.Infof("get token.....111111")
	if time.Now().After(c.expiresAt) {
		klog.Infof("token expired....xxxxx")
	}
	if c.token == "" || time.Now().After(c.expiresAt) {

		return "", false
	}
	return c.token, true
}

func (c *TokenCache) SetRefreshToken(refreshToken string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.refreshToken = refreshToken
	return nil
}

func (c *TokenCache) GetRefreshToken() (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.refreshToken == "" {
		return "", false
	}
	return c.refreshToken, true
}

func (c *TokenCache) Delete() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.token = ""
	c.refreshToken = ""
	c.expiresAt = time.Time{}
	return nil
}
