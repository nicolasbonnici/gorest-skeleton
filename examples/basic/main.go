package main

import (
	"log"

	"github.com/nicolasbonnici/gorest"
	"github.com/nicolasbonnici/gorest/pluginloader"

	authplugin "github.com/nicolasbonnici/gorest-auth"
	skeletonplugin "github.com/nicolasbonnici/gorest-skeleton"
)

func init() {
	pluginloader.RegisterPluginFactory("auth", authplugin.NewPlugin)
	pluginloader.RegisterPluginFactory("skeleton", skeletonplugin.NewPlugin)
}

func main() {
	cfg := gorest.Config{
		ConfigPath: ".",
	}

	log.Println("Starting GoREST with Skeleton Plugin example...")
	log.Println("The skeleton plugin provides CRUD operations at /api/skeleton")
	log.Println("Make sure to register and login first using the auth plugin endpoints")

	gorest.Start(cfg)
}
