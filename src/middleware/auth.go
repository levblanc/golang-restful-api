package middleware

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/levblanc/golang-restful-api/constants"
	"github.com/levblanc/golang-restful-api/db"
	"github.com/levblanc/golang-restful-api/models"
	"github.com/levblanc/golang-restful-api/utils/ctx"
	"github.com/levblanc/golang-restful-api/utils/response"
)

// Auth checks whether a user's session is valid
// if not, returns unAuthorized status code
func Auth(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var session models.Session
		var error response.Error

		cookie, err := req.Cookie(constants.CookieName)

		if err != nil {
			error.Status = response.StatusError
			error.Message = http.StatusText(http.StatusUnauthorized)

			response.Send(w, http.StatusUnauthorized, error)
			return
		}

		err = db.UserSession.Find(bson.M{"sid": cookie.Value}).One(&session)

		if err != nil {
			error.Status = response.StatusError
			error.Message = http.StatusText(http.StatusUnauthorized)

			response.Send(w, http.StatusUnauthorized, error)
			return
		}

		next(w, ctx.Set(req, constants.ContextUserID, session.UserID))
	})
}
