package simaland

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	tokenHealth = time.Hour
)

// JWT Auth
// Авторизация средствами JWT токена

func newJWTAuth(client *Client, login, pass string) (*jwt_auth, error) {
	new := &jwt_auth{
		login:    login,
		password: pass,
		mutex:    &sync.Mutex{},
		cl:       client,
	}

	err := new.updateToken()
	go new.autoupdate()
	return new, err
}

type jwt_auth struct {
	login, password string

	cl    *Client
	mutex *sync.Mutex

	token string
}

func (j jwt_auth) autoupdate() {
	t := time.NewTimer(tokenHealth)

	for {
		<-t.C

		err := j.updateToken()
		if err != nil {
			if j.cl.log != nil {
				j.cl.log.Logf("[ERROR] update token failed with error: %v", err)
				j.cl.log.Logf("[DEBUG] next update after 10s")
			}

			t.Reset(time.Second * 10)
			continue
		}

		j.cl.log.Logf("[DEBUG] update token success")
		t.Reset(time.Hour)
	}
}

func (j jwt_auth) updateToken() error {
	var data struct {
		JWT string `json:"jwt"`
	}

	err := j.cl.get("auth", &data,
		SetHeader("Authorization", "Basic "+base64.RawStdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", j.login, j.password)))),
	)
	if err != nil {
		return err
	}

	j.mutex.Lock()
	defer j.mutex.Unlock()
	j.token = data.JWT
	return nil
}

func (j *jwt_auth) modificator() RequestModificator {
	return func(r *http.Request) {
		j.mutex.Lock()
		defer j.mutex.Unlock()

		r.Header.Set("Authorization", "Bearer "+j.token)
	}
}
