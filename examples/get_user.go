package main

import (
	"github.com/beclab/lldap-client/pkg/cache/memory"
	"github.com/beclab/lldap-client/pkg/client"
	"github.com/beclab/lldap-client/pkg/config"
	apierrors "github.com/beclab/lldap-client/pkg/errors"
	"os"

	"context"
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

	u, err := graphqlClient.Users().Get(context.Background(), "admin")

	if err != nil && apierrors.IsNotFound(err) {
		klog.Infof("user is not exists")
		os.Exit(0)
	}
	klog.Infof("Get User: %#v", u)

}
