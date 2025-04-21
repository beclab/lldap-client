package main

import (
	"context"
	"github.com/beclab/lldap-client/pkg/cache/memory"
	"github.com/beclab/lldap-client/pkg/client"
	"github.com/beclab/lldap-client/pkg/config"
	apierrors "github.com/beclab/lldap-client/pkg/errors"
	"github.com/beclab/lldap-client/pkg/generated"
	"k8s.io/klog/v2"
)

func main() {
	cfg := config.Config{
		Host:       "http://127.0.0.1:17170",
		Username:   "admin",
		Password:   "adminpassword",
		TokenCache: memory.New(),
	}
	graphqlClient, err := client.New(&cfg)
	if err != nil {
		klog.Infof("create graphqlClient err=%v", err)
	}

	user := generated.CreateUserInput{
		Id:    "test001",
		Email: "test001@gmail.com",
	}

	_, err = graphqlClient.Users().Create(context.Background(), &user, "userpassword")
	if err != nil {
		klog.Infof("get user test001 failed %v", err)
	}

	_, err = graphqlClient.Groups().Create(context.Background(), "group001")

	if err != nil && apierrors.IsAlreadyExists(err) {
		klog.Infof("group already exists")

	}

	g, err := graphqlClient.Groups().GetByName(context.Background(), "group001")
	if err != nil {
		panic(err)
	}

	klog.Infof("gg : %v", g)

	err = graphqlClient.Groups().AddUser(context.Background(), "test001", g.Id)
	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			klog.Infof("membeship: %v", err)
		}
	}

	u, err := graphqlClient.Users().Get(context.Background(), "test001")
	if err != nil {
		klog.Infof("get user test001 failed %v", err)
	}
	klog.Infof("user test001 groups: %v", u.Groups)

	err = graphqlClient.Groups().RemoveUser(context.Background(), "test001", g.Id)
	if err != nil {
		klog.Errorf("remove user from group failed %v", err)
	}

	u, err = graphqlClient.Users().Get(context.Background(), "test001")
	if err != nil {
		klog.Infof("get user test001 failed %v", err)
	}
	klog.Infof("user test001 groups: %v", u.Groups)

	klog.Infof("Get groups: %#v", g)

}
