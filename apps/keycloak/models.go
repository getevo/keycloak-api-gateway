package keycloak

import "github.com/Nerzal/gocloak/v5"

type JSONResponse struct {
	Jwt      gocloak.JWT `json:"token"`
	Response string      `json:"response"`
}
type JSONErrorResponse struct {
	Detail string `json:"detail"`
}
