package keycloak

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
	"time"

	"github.com/getevo/evo"
	"github.com/getevo/evo/lib/data"
	"github.com/getevo/evo/lib/log"
	"gopkg.in/square/go-jose.v2/jwt"
)

type User struct{}

func (p User) Save(u *evo.User) error {
	log.Info("implement	me")
	return nil
}

func (p User) HasPerm(u *evo.User, v string) bool {
	log.Info("implement me")
	return true
}

func (p User) HasRole(u *evo.User, v interface{}) bool {
	log.Info("implement me")
	return true
}

func (p User) Image(u *evo.User) string {
	return "files/profile/profile-" + fmt.Sprint(u.ID) + ".jpg"
}

func (p User) SetPassword(u *evo.User, password string) error {
	log.Info("implement me")
	return nil
}

func (p User) SetGroup(u *evo.User, group interface{}) error {
	log.Info("implement me")
	return nil
}

func (p User) AfterFind(u *evo.User) error {

	return nil
}

func (p User) SyncPermissions(app string, perms evo.Permissions) {
	log.Info("implement me")
}

// SetGroup set user group
func (p User) FromRequest(request *evo.Request) {
	request.User = &evo.User{Anonymous: true}
	accessToken := request.Get("Authorization")

	if accessToken == "" {
		accessToken = request.Cookies("Authorization")
	}
	if accessToken == "" {
		accessToken = request.FormValue("Authorization")
	}
	accessToken = strings.TrimSpace(accessToken)

	if len(accessToken) > 10 {
		if strings.ToLower(accessToken[0:6]) == "bearer" {
			accessToken = accessToken[7:]
		}

		token, err := jwt.ParseSigned(accessToken)

		if err != nil {
			log.Error(err)
			//request.WriteResponse(false, fmt.Errorf("unauthorized"), 401)
			return
		}
		var claims data.Map
		err = token.Claims(Certificates.Keys[0], &claims)

		if err != nil {
			log.Error(err)
			//request.WriteResponse(false, fmt.Errorf("unauthorized"), 401)
			return
		}
		if claims.Get("email") != nil {
			request.User.Email = claims.Get("email").(string)
			request.User.Anonymous = false
			if claims.Get("user_id") != nil {
				request.User.ID = uint(claims.Get("user_id").(float64))
			} else {
				var hash = fnv.New32()
				hash.Write([]byte(request.User.Email))
				request.User.ID = uint(hash.Sum32())
			}
			request.User.Params = claims

			if claims.Get("given_name") != nil {
				request.User.GivenName = claims.Get("given_name").(string)
			}
			if claims.Get("family_name") != nil {
				request.User.FamilyName = claims.Get("family_name").(string)
			}

		} else {
			request.User.Anonymous = true
		}

	}
}

func SetUserParam(request *evo.Request, key string, value string) error {
	user, err := GCloakClient.GetUserByID(AdminJWT.AccessToken, Realm, request.User.Params.Get("sub").(string))
	var version = strconv.FormatInt(time.Now().Unix(), 16)
	if err == nil {
		user.Attributes[key] = []string{value}
		user.Attributes["upd"] = []string{version}
		return GCloakClient.UpdateUser(request.Get("Authorization"), Realm, *user)
	}
	return err
}

func UpdateUserJWTByMap(request *evo.Request, data map[string]interface{}) {
	if request.Unauthorized() {
		return
	}
	/*var redisKey = "user.id." + fmt.Sprint(request.User.ID)
	var version string
	err := redis.Get(redisKey, &version)
	if err != nil || version == "" {
		version = strconv.FormatInt(time.Now().Unix(), 16)
		err = redis.Set(redisKey, version, 5*time.Minute)
		if err != nil {
			log.Error(err)
		}
		user, err := GCloakClient.GetUserByID(AdminJWT.AccessToken, Realm, request.User.Params.Get("sub").(string))
		if err == nil {
			user.Attributes["upd"] = []string{version}
			for k, v := range data {
				user.Attributes[k] = []string{fmt.Sprint(v)}
			}
			err = GCloakClient.UpdateUser(AdminJWT.AccessToken, Realm, *user)
			if err != nil {
				log.Error("unable to access keycloak")
				log.Error(err)
			}
			request.User.Params["upd"] = version
		} else {
			log.Error("unable to access keycloak")
			log.Error(err)
		}
	} else {
		if request.User.Params.Get("upd") == nil || version != request.User.Params.Get("upd") {
			request.Set("last-modified", "true")
			//request.Set("last-modified","false")

		}
	}*/
}

func UpdateUserJWT(request *evo.Request, key string, value interface{}) {
	if request.Unauthorized() {
		return
	}
	/*var redisKey = "user.sub." + fmt.Sprint(request.User.ID)
	var version string
	err := redis.Get(redisKey, &version)
	if err != nil || version == "" {
		version = strconv.FormatInt(time.Now().Unix(), 16)
		err = redis.Set(redisKey, version, 5*time.Minute)
		if err != nil {
			log.Error(err)
		}
		user, err := GCloakClient.GetUserByID(AdminJWT.AccessToken, Realm, request.User.Params.Get("sub").(string))
		if err == nil {
			user.Attributes["upd"] = []string{version}
			user.Attributes[key] = []string{fmt.Sprint(value)}
			err = GCloakClient.UpdateUser(AdminJWT.AccessToken, Realm, *user)
			if err != nil {
				log.Error("unable to access keycloak")
				log.Error(err)
			}
			request.User.Params["upd"] = version
		} else {
			log.Error("unable to access keycloak")
			log.Error(err)
		}
	} else {
		if request.User.Params.Get("upd") == nil || version != request.User.Params.Get("upd") {
			//request.Set("last-modified","true")
			request.Set("last-modified", "false")
		}
	}*/
}
