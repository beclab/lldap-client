package transport

import (
	"errors"
	"fmt"
	"github.com/beclab/lldap-client/pkg/auth"
	"github.com/beclab/lldap-client/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"k8s.io/klog/v2"
	"net/http"
	"time"
)

func HTTPWrappersForConfig(cfg *config.Config, rt http.RoundTripper) (http.RoundTripper, error) {
	rt = NewBearerAuthRoundTripper(cfg, rt)
	return rt, nil
}

type bearerAuthRoundTripper struct {
	cfg    *config.Config
	source oauth2.TokenSource
	rt     http.RoundTripper
}

// NewBearerAuthRoundTripper adds the provided bearer token to a request
// unless the authorization header has already been set.
func NewBearerAuthRoundTripper(cfg *config.Config, rt http.RoundTripper) http.RoundTripper {
	return &bearerAuthRoundTripper{cfg, nil, rt}
}

func (rt *bearerAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(req.Header.Get("Authorization")) != 0 {
		return rt.rt.RoundTrip(req)
	}
	req = CloneRequest(req)
	_, ok := rt.cfg.TokenCache.Get()
	if !ok {
		klog.Infof("not found token")
		accessToken, err := rt.ensureValidToken()
		if err != nil {
			return nil, err
		}
		rt.cfg.BearerToken = accessToken
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rt.cfg.BearerToken))

	return rt.rt.RoundTrip(req)
}

func (rt *bearerAuthRoundTripper) ensureValidToken() (string, error) {
	token, err := rt.tryRefreshToken()
	if err != nil {
		return "", nil
	}
	return token, nil
}

func (rt *bearerAuthRoundTripper) tryRefreshToken() (string, error) {
	now := time.Now().Unix()
	//_, tokenExists := rt.cfg.TokenCache.Get()
	refreshToken, refreshTokenExists := rt.cfg.TokenCache.GetRefreshToken()
	var err error
	var retToken, retRefreshToken string
	if refreshTokenExists {
		klog.Infof("try to refresh token....")
		retToken, err = auth.Refresh(rt.cfg.Host, refreshToken)
		if err != nil {
			klog.Infof("refresh token err=%v", err)
		}
	}
	if err != nil || !refreshTokenExists {
		klog.Infof("try login...")
		resp, err := auth.Login(rt.cfg.Host, rt.cfg.Username, rt.cfg.Password)
		if err != nil {
			return "", err
		}
		retToken, retRefreshToken = resp.Token, resp.RefreshToken
	}
	exp, err := getTokenExp(retToken)
	if err != nil {
		return "", err
	}
	if retToken != "" {
		if exp-now <= 0 {
			return "", errors.New("token already expired")
		}
		klog.Infof("set token...")
		klog.Infof("exp:%d,now:%d,exp-now:%d", exp, now, exp-now)
		err = rt.cfg.TokenCache.Set(retToken, time.Duration(exp-now-5)*time.Second)
		if err != nil {
			return "", err
		}
	}
	if retRefreshToken != "" {
		klog.Infof("set refresh token...")

		err = rt.cfg.TokenCache.SetRefreshToken(retRefreshToken)
		if err != nil {
			return "", err
		}
	}
	rt.cfg.BearerToken = retToken
	return retToken, nil
}

func getTokenExp(tokenString string) (int64, error) {
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(tokenString, &jwt.StandardClaims{})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}
	return claims.ExpiresAt, nil
}

func CloneRequest(req *http.Request) *http.Request {
	r := new(http.Request)

	// shallow clone
	*r = *req

	// deep copy headers
	r.Header = CloneHeader(req.Header)

	return r
}

// CloneHeader creates a deep copy of an http.Header.
func CloneHeader(in http.Header) http.Header {
	out := make(http.Header, len(in))
	for key, values := range in {
		newValues := make([]string, len(values))
		copy(newValues, values)
		out[key] = newValues
	}
	return out
}
