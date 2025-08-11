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

	_, err = graphqlClient.Groups().Create(context.TODO(), "testgroup", "admin")
	if err != nil && !apierrors.IsAlreadyExists(err) {
		panic(err)
	}
	err = graphqlClient.Groups().CreateAttribute(context.TODO(), "aaa", generated.AttributeTypeString, false, false)
	if err != nil {
		panic(err)
	}
	gg, err := graphqlClient.Groups().GetByName(context.TODO(), "testgroup")
	if err != nil {
		panic(err)
	}
	err = graphqlClient.Groups().Update(context.TODO(), generated.UpdateGroupInput{
		Id:          gg.Id,
		DisplayName: gg.DisplayName,
		//RemoveAttributes: []string{"aaa"},
		InsertAttributes: []generated.AttributeValueInput{
			{
				Name:  "aaa",
				Value: []string{"aaaa"},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	gg, err = graphqlClient.Groups().GetByName(context.TODO(), "testgroup")
	if err != nil {
		panic(err)
	}

	klog.Infof("Get groups: %#v", gg)

}
