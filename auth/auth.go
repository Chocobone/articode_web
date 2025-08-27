//우선순위 밀림
package auth

import(
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/oauth2"
	"golang.gor/x/oauth2/google"
)

type User struct{
	Name  string  `json:"name"`
	Email string  `json:"email"`
}

const (
    CallBackURL = "http://localhost:1333/auth/callback"
  
    UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
    ScopeEmail          = "https://www.googleapis.com/auth/userinfo.email"
    ScopeProfile        = "https://www.googleapis.com/auth/userinfo.profile"
)

var OAuthConf *oauth2.Config

func init() {
	OAuthConf = &oauth2.Config{
		ClientID:	  "google client id",
		ClientSecret: "google client secret",
		RedirectURL:  CallBackURL,
		Scopes:		  []string{ScopeEmail, ScopeProfile},
		Endpoint:	  google.Endpoint,
	}
}

func GetLoginURL(state string) string {
	return OAuthConf.AuthCodeURL(state)
}