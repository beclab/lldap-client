package main

import (
	"context"
	"github.com/beclab/lldap-client/pkg/cache/memory"
	"github.com/beclab/lldap-client/pkg/client"
	"github.com/beclab/lldap-client/pkg/config"
	apierrors "github.com/beclab/lldap-client/pkg/errors"
	"k8s.io/klog/v2"
)

func main() {
	cfg := config.Config{
		Host:       "http://127.0.0.1:17170",
		Username:   "admin",
		Password:   "adminpasswor",
		TokenCache: memory.New(),
	}
	graphqlClient, err := client.New(&cfg)
	if err != nil {
		klog.Infof("create graphqlClient err=%v", err)
	}

	_, err = graphqlClient.Groups().Create(context.Background(), "group001")

	if err != nil && apierrors.IsAlreadyExists(err) {
		klog.Infof("group already exists")

	}
	//g, err := graphqlClient.Groups().Get(context.Background(), groups.Id)
	//if err != nil {
	//	panic(err)
	//}
	//
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

	//g, err := graphqlClient.Groups().Get(context.Background(), groups.Id)
	//if err != nil {
	//	panic(err)
	//}

	klog.Infof("Get groups: %#v", g)

}
