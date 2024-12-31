package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/beclab/lldap-client/pkg/generated"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type users struct {
	client *Client
}

func NewUserRoute(c *Client) *users {
	return &users{
		client: c,
	}
}

func (u *users) Get(ctx context.Context, name string) (*generated.GetUserDetailsUser, error) {
	var resp *generated.GetUserDetailsResponse
	resp, err := generated.GetUserDetails(ctx, u.client, name)
	if err != nil {
		return nil, err
	}
	return &resp.User, nil
}

func (u *users) Create(ctx context.Context, user *generated.CreateUserInput, password string) (*generated.CreateUserResponse, error) {
	if user == nil {
		return nil, errors.New("empty user")
	}
	if len(password) < 8 {
		return nil, errors.New("password length must > 8 char")
	}
	resp, err := generated.CreateUser(ctx, u.client, *user)
	if err != nil {
		return nil, err
	}
	err = u.ResetPassword(ctx, user.Id, password)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *users) Delete(ctx context.Context, name string) error {
	//var resp *generated.DeleteUserQueryResponse
	_, err := generated.DeleteUserQuery(ctx, u.client, name)
	if err != nil {
		return err
	}
	return nil
}

func (u *users) List(ctx context.Context) ([]generated.ListUsersQueryUsersUser, error) {
	var resp *generated.ListUsersQueryResponse
	resp, err := generated.ListUsersQuery(ctx, u.client, nil)
	if err != nil {
		return nil, err
	}
	return resp.Users, nil
}

type resetPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *users) ResetPassword(ctx context.Context, username, password string) error {
	creds := resetPassword{
		Username: username,
		Password: password,
	}
	url := fmt.Sprintf("%s/auth/simple/register", u.client.cfg.Host)
	client := resty.New()
	resp, err := client.SetTimeout(5*time.Second).R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+u.client.cfg.BearerToken).
		SetBody(creds).Post(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.String())
	}
	return nil
}
