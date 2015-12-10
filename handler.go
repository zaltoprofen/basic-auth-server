package basicauth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

type Authenticator func(string, string) bool

var Realm string

func deny(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", Realm))
	w.WriteHeader(401)
}

func BasicAuth(handlerFunc http.HandlerFunc, auth Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		basic := strings.TrimPrefix(authHeader, "Basic ")
		if basic == authHeader {
			deny(w)
			return
		}
		data, err := base64.StdEncoding.DecodeString(basic)
		if err != nil {
			deny(w)
			return
		}

		userAndPass := strings.SplitN(string(data), ":", 2)
		if len(userAndPass) == 2 && auth(userAndPass[0], userAndPass[1]) {
			handlerFunc(w, r)
		} else {
			deny(w)
		}
	}
}
