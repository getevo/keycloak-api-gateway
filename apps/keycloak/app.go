package keycloak
/*
* Copyright © 2022 Allan Nava <EVO TEAM>
* Created 04/02/2022
* Updated 04/02/2022
*
 */

// @doc type 		app
// @doc name		keycloak
// @doc description database keycloak reader
// @doc author		reza | allan 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"keycloak-api-gateway/apps/settings"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v5"
	"github.com/getevo/evo"
	"github.com/getevo/evo/lib/log"
	"github.com/getevo/evo/menu"
	"gopkg.in/square/go-jose.v2"
)

var Server string
var Realm string
var Client string
var Certificates jose.JSONWebKeySet
var GCloakClient gocloak.GoCloak
var AdminJWT *gocloak.JWT
var ClientSecret string
var ClientID string

//
func Register(server, realm, client string) {
	Server = server
	Realm = realm
	Client = client

	evo.Register(App{})
	GCloakClient = gocloak.NewClient(server)

	var err error
	AdminJWT, err = GCloakClient.LoginAdmin(settings.Settings.Keycloak.Username, settings.Settings.Keycloak.Password, realm)
	if err != nil {
		log.Critical(err)
	}
	clients, err := GCloakClient.GetClients(AdminJWT.AccessToken, Realm, gocloak.GetClientsParams{})
	if err != nil {
		log.Critical(err)
	}

	for _, client := range clients {
		if client.ClientID != nil && *client.ClientID == Client {
			ClientID = *client.ID
			fmt.Printf("clientId: %v , Client: %v \n ", ClientID, Client)
			cr, err := GCloakClient.GetClientSecret(AdminJWT.AccessToken, Realm, ClientID)

			if err == nil {
				fmt.Println(*cr.Value)
				ClientSecret = *cr.Value
			} else {
				log.Critical(err)
			}
			break
		}
	}

}

type App struct{}

//                   	Register    	the bible
func (App) Register() {
	fmt.Println("Keycloak	Registered")

}

// WhenReady called after setup all apps
func (App) WhenReady() {
	//
	resp, err := http.Get(strings.TrimRight(settings.Settings.Keycloak.Server, "/") + "/auth/realms/" + settings.Settings.Keycloak.Realm + "/protocol/openid-connect/certs")
	if err != nil {
		log.Error("Unable connect to keycloak server")
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Error("Unable connect parse keycloak cert")
		log.Fatal(err)
	}
	Certificates = jose.JSONWebKeySet{}
	err = json.Unmarshal(body, &Certificates)
	if len(Certificates.Keys) == 0 {
		fmt.Println(string(body))
		log.Fatal("invalid keycloak cert")
	}
	if err != nil {
		log.Error("Unable connect parse keycloak cert")
		log.Fatal(err)
	}
	evo.SetUserInterface(User{})

}

// Router setup routers
func (App) Router() {
	evo.Get("me", func(request *evo.Request) {
		request.WriteResponse(request.User)
	})
	//
	var errorResponse = JSONErrorResponse{"method 'GET' not allowed."}
	//
	evo.Post("api/auth-login", AuthLogin)
	evo.Get("api/auth-login", func(request *evo.Request) {
		// capire dal frontend usare lo standard di errore di EVO
		request.Status(400).JSON(errorResponse)
	})
	evo.Post("api/auth-logout", AuthLogout)
	evo.Get("api/auth-logout", func(request *evo.Request) {
		// capire dal frontend usare lo standard di errore di EVO
		request.Status(400).JSON(errorResponse)
	})
	//
	evo.Post("api/auth-refresh-token", RefreshToken)
	evo.Get("api/refresh-token", func(request *evo.Request) {
		// capire dal frontend usare lo standard di errore di EVO
		request.Status(400).JSON(errorResponse)
	})
}

// Permissions setup permissions of app
func (App) Permissions() []evo.Permission { return []evo.Permission{} }

// Menus setup menus
func (App) Menus() []menu.Menu {
	return []menu.Menu{}
}

func (App) Pack() {}
