package cache

import "time"

type TokenCacheInterface interface {
	Set(token string, ttl time.Duration) error
	Get() (string, bool)
	Delete() error
	SetRefreshToken(refreshToken string) error
	GetRefreshToken() (string, bool)
}
