package client

import (
	"context"
	"github.com/beclab/lldap-client/pkg/generated"
)

type groups struct {
	client *Client
}

func NewGroupRoute(c *Client) *groups {
	return &groups{
		client: c,
	}
}

func (g *groups) Get(ctx context.Context, id int) (*generated.GetGroupDetailsGroup, error) {
	var resp *generated.GetGroupDetailsResponse
	resp, err := generated.GetGroupDetails(ctx, g.client, id)
	if err != nil {
		return nil, err
	}
	return &resp.Group, nil
}

func (g *groups) Create(ctx context.Context, name string, creator string) (*generated.CreateGroupCreateGroup, error) {
	var resp *generated.CreateGroupResponse
	resp, err := generated.CreateGroup(ctx, g.client, name, creator)
	if err != nil {
		return nil, err
	}
	return &resp.CreateGroup, nil
}

func (g *groups) Delete(ctx context.Context, id int) (*generated.DeleteGroupQueryDeleteGroupSuccess, error) {
	var resp *generated.DeleteGroupQueryResponse
	resp, err := generated.DeleteGroupQuery(ctx, g.client, id)
	if err != nil {
		return nil, err
	}
	return &resp.DeleteGroup, nil
}

func (g *groups) List(ctx context.Context) ([]generated.GetGroupListGroupsGroup, error) {
	var resp *generated.GetGroupListResponse
	resp, err := generated.GetGroupList(ctx, g.client)
	if err != nil {
		return nil, err
	}
	return resp.Groups, nil
}

func (g *groups) GetByName(ctx context.Context, name string) (*generated.GetGroupDetailsByNameGroupByNameGroup, error) {
	var resp *generated.GetGroupDetailsByNameResponse
	resp, err := generated.GetGroupDetailsByName(ctx, g.client, name)
	if err != nil {
		return nil, err
	}
	return &resp.GroupByName, nil
}

func (g *groups) AddUser(ctx context.Context, username string, groupID int) error {
	//var resp *generated.AddUserToGroupResponse
	_, err := generated.AddUserToGroup(ctx, g.client, username, groupID)
	if err != nil {
		return err
	}
	return nil
}

func (g *groups) RemoveUser(ctx context.Context, username string, groupID int) error {
	_, err := generated.RemoveUserFromGroup(ctx, g.client, username, groupID)
	if err != nil {
		return err
	}
	return nil
}

func (g *groups) Update(ctx context.Context, updateInput generated.UpdateGroupInput) error {
	_, err := generated.UpdateGroup(ctx, g.client, updateInput)
	if err != nil {
		return err
	}
	return nil
}

func (g *groups) CreateAttribute(ctx context.Context, name string,
	attributeType generated.AttributeType, isList, isVisible bool) error {
	_, err := generated.CreateGroupAttribute(ctx, g.client, name, attributeType, isList, isVisible)
	if err != nil {
		return err
	}
	return nil
}

func (g *groups) GetAttributeSchema(ctx context.Context) (*generated.GetGroupAttributesSchemaResponse, error) {
	res, err := generated.GetGroupAttributesSchema(ctx, g.client)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *groups) DeleteAttribute(ctx context.Context, name string) error {
	_, err := generated.DeleteGroupAttributeQuery(ctx, g.client, name)
	if err != nil {
		return err
	}
	return nil
}
