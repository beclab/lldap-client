package client

import (
	"context"
	"github.com/beclab/lldap-client/pkg/generated"
)

type UserInterface interface {
	Get(ctx context.Context, name string) (*generated.GetUserDetailsUser, error)
	Create(ctx context.Context, user *generated.CreateUserInput, password string) (*generated.CreateUserResponse, error)
	Delete(ctx context.Context, name string) error
	List(ctx context.Context) ([]generated.ListUsersQueryUsersUser, error)

	ResetPassword(ctx context.Context, username, password string) error
}

type GroupInterface interface {
	Get(ctx context.Context, id int) (*generated.GetGroupDetailsGroup, error)
	Create(ctx context.Context, name string) (*generated.CreateGroupCreateGroup, error)
	Delete(ctx context.Context, id int) (*generated.DeleteGroupQueryDeleteGroupSuccess, error)
	List(ctx context.Context) ([]generated.GetGroupListGroupsGroup, error)
	GetByName(ctx context.Context, name string) (*generated.GetGroupDetailsByNameGroupByNameGroup, error)

	AddUser(ctx context.Context, username string, groupID int) error
}
