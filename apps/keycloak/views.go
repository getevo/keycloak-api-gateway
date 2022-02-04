package keycloak
/*
* Copyright © 2022 Allan Nava <EVO TEAM>
* Created 04/02/2022
* Updated 04/02/2022
*
 */
import (
	"fmt"

	"github.com/getevo/evo"
	"github.com/getevo/evo/lib/log"
)

func RefreshToken(request *evo.Request) {
	if request.Unauthorized() {
		return
	}
	if true || request.Get("last-modified") != "true" {
		refresh := request.FormValue("token")

		jwt, err := GCloakClient.RefreshToken(refresh, Client, ClientSecret, Realm)

		if err != nil {
			request.WriteResponse(err)
			log.Error(err)
			log.Error(refresh, Client, ClientSecret, Realm)
		} else {
			//b, err := json.Marshal(jwt)
			request.WriteResponse(true, jwt)
		}
	} else {
		request.WriteResponse(false, fmt.Errorf("no need of update"))
	}
}

func AuthLogin(request *evo.Request) {
	//log.DebugF("auth-login called")
	username := request.FormValue("username")
	password := request.FormValue("password")
	if username == "" {
		request.WriteResponse(false, fmt.Errorf("username is empty"))
		return
	}
	if password == "" {
		request.WriteResponse(false, fmt.Errorf("password is empty"))
		return
	}
	jwt, err := GCloakClient.Login(Client, ClientSecret, Realm, username, password)
	//
	if err != nil {
		// need to change the status code 400
		request.WriteResponse(err)
		log.Error(err)
		log.Error(jwt, Client, ClientSecret, Realm)
		return
	}
	//log.DebugF("login success")
	resp := JSONResponse{*jwt, "success"}
	request.Status(200).JSON(resp)
}

//
func AuthLogout(request *evo.Request) {
	log.DebugF("auth-login called")
	token := request.FormValue("token")
	if token == "" {
		request.WriteResponse(false, fmt.Errorf("token is empty"))
		return
	}
	//
	err := GCloakClient.Logout(Client, ClientSecret, Realm, token)
	//
	if err != nil {
		request.WriteResponse(err)
		log.Error(err)
		return
	}
	//
	request.WriteResponse(true, "logout successfull")
}
