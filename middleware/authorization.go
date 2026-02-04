package middleware

import (
	"errors"
	"example/hello/api"
	"go/token"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("unauthorized access")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface

		database, err = tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.InternalServerErrorHandler(w, err)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails, err = (*database).GetUserLoginDetails(username)
		if err != nil {
			log.Error(err)
			api.RequestErrorHandler(w, err)
			return
		}

		if loginDetails == nil || token != (*loginDetails).AuthToken {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
	})
}