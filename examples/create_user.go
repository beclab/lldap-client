package main

import (
	"context"
	"github.com/beclab/lldap-client/pkg/generated"

	//"github.com/beclab/lldap-client"
	"github.com/beclab/lldap-client/pkg/cache/memory"
	"github.com/beclab/lldap-client/pkg/client"
	"github.com/beclab/lldap-client/pkg/config"
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
		Id:    "test002",
		Email: "test002@gmail.com",
	}

	_, err = graphqlClient.Users().Create(context.Background(), &user, "12345678")

	if err != nil {
		//fmt.Println("user already exists")
		panic(err)
	}

	u2, err := graphqlClient.Users().Get(context.Background(), "test002")
	if err != nil {
		panic(err)
	}

	klog.Infof("Get User: %#v", u2)

}
