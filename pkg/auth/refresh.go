package auth

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

func Refresh(baseURL, refreshToken string) (string, error) {
	url := fmt.Sprintf("%s/auth/refresh", baseURL)
	type RefreshTokenResponse struct {
		Token string `json:"token"` // 假设返回一个 token
	}
	client := resty.New()
	resp, err := client.SetTimeout(5*time.Second).R().
		SetHeader("Content-Type", "application/json").
		//SetHeader("Authorization", "Bearer "+token).
		SetHeader("refresh-token", refreshToken).
		SetResult(&RefreshTokenResponse{}).
		Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != http.StatusOK {
		return "", errors.New(resp.String())
	}
	ret := resp.Result().(*RefreshTokenResponse)
	return ret.Token, nil
}
