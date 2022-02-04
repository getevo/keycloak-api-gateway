package main

import (
	"fmt"
	"keycloak-api-gateway/apps/keycloak"
	"keycloak-api-gateway/apps/settings"

	"github.com/getevo/evo"
)

var config = struct {
	DSN     map[string]string `yaml:"dsn"`
	Include map[string]bool   `yaml:"include"`
}{}

func main() {
	//
	evo.Setup()
	var cfg = evo.GetConfig()
	//
	if cfg != nil && cfg.Log.WriteFile {
		fmt.Print("\n cfg != nil \n ")
		//log.SetFile(cfg.Log.Path)
		//log.Verbosity = log.ParseLevel(cfg.Log.Level)
	}
	//
	evo.Get("/health", health)
	//
	settings.Register()
	keycloak.Register(settings.Settings.Keycloak.Server, settings.Settings.Keycloak.Realm, settings.Settings.Keycloak.Client)
	//
	// run the evo stuff
	evo.Run()
	//
}

func health(request *evo.Request) {
	request.Status(200)
	//request.WriteResponse(true, redis.MyAddress)
}

func app(app string) bool {
	if v, ok := config.Include[app]; ok {
		return v
	}
	return false
}
