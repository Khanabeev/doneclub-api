package auth

import (
	"doneclub-api/internal/domain/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type Storage struct {
}

type VerifiedResponse struct {
	UserId     int  `json:"user_id"`
	IsVerified bool `json:"is_verified"`
}

func (r Storage) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	u := buildVerifyURL(token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {
		m := VerifiedResponse{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			//fmt.Println("Error while decoding response from auth server:" + err.Error())
			return false
		}
		return m.IsVerified
	}
}

/*
  This will generate a url for token verification in the below format

  /auth/verify?token={token string}
              &routeName={current route name}
              &customer_id={customer id from the current route}
              &account_id={account id from current route if available}

  Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000&account_id=95470
*/
func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	host := os.Getenv("AUTH_SERVER")
	path := os.Getenv("AUTH_PATH")
	u := url.URL{Host: host, Path: path, Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewStorage() auth.Storage {
	return Storage{}
}
