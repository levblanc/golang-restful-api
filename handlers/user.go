package handlers

import (
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/levblanc/golang-restful-api/db"
	"github.com/levblanc/golang-restful-api/models"
	"github.com/levblanc/golang-restful-api/utils/auth"
	"github.com/levblanc/golang-restful-api/utils/format"
	"github.com/levblanc/golang-restful-api/utils/logger"
	"github.com/levblanc/golang-restful-api/utils/response"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

// Signup handles sign up requests
func Signup(w http.ResponseWriter, req *http.Request) {
	var error response.Error
	var existUser models.User

	err := req.ParseForm()

	if err != nil {
		error.Status = response.StatusError
		error.Message = "Post method params error! One possible reason is your request Content-Type is not set to application/x-www-form-urlencoded. Set it and try again."

		response.Send(w, http.StatusBadRequest, error)
		return
	}

	user := models.User{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	if user.Username == "" {
		error.Status = response.StatusError
		error.Message = "Username is missing!"

		response.Send(w, http.StatusBadRequest, error)
		return
	}

	err = db.User.Find(bson.M{"username": user.Username}).One(&existUser)

	if existUser.Username == user.Username {
		error.Status = response.StatusError
		error.Message = "Username taken, try again!"

		response.Send(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Status = response.StatusError
		error.Message = "Password is missing!"

		response.Send(w, http.StatusBadRequest, error)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	logger.Fatal(err)

	user.Password = string(hash)
	user.UserID = xid.New()

	t := time.Now()
	user.CreatedAt = t
	user.CreatedTime = format.Time(t)

	err = db.User.Insert(user)

	if err != nil {
		error.Status = response.StatusError
		error.Message = http.StatusText(http.StatusInternalServerError)

		response.Send(w, http.StatusInternalServerError, error)
		return
	}

	// empty password to omit the field in response
	user.Password = ""

	response.Send(w, http.StatusOK, user)
}

// GetUser gets user info by user id
func GetUser(w http.ResponseWriter, req *http.Request) {
	var error response.Error
	var user models.User

	params := mux.Vars(req)

	id, err := xid.FromString(params["id"])
	logger.Fatal(err)

	err = db.User.Find(bson.M{"userId": id}).One(&user)

	if err != nil {
		error.Status = response.StatusError
		error.Message = "Cannot find user!"

		response.Send(w, http.StatusNotFound, error)
		return
	}

	// empty password to omit the field in response
	user.Password = ""

	response.Send(w, http.StatusOK, user)
}

// Login handles user login process
// it creates cookie and user session
func Login(w http.ResponseWriter, req *http.Request) {
	var error response.Error
	var found models.User

	err := req.ParseForm()

	if err != nil {
		error.Status = response.StatusError
		error.Message = "Post method params error! One possible reason is your request Content-Type is not set to application/x-www-form-urlencoded. Set it and try again."

		response.Send(w, http.StatusBadRequest, error)
		return
	}

	user := models.User{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	if user.Username == "" {
		error.Status = response.StatusError
		error.Message = "Username is missing!"

		response.Send(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Status = response.StatusError
		error.Message = "Password is missing!"

		response.Send(w, http.StatusBadRequest, error)
		return
	}

	username := user.Username
	password := user.Password

	err = db.User.Find(bson.M{"username": username}).One(&found)

	if err != nil {
		error.Status = response.StatusError
		error.Message = "User not found!"

		response.Send(w, http.StatusNotFound, error)
		return
	}

	hashedPassword := found.Password

	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)

	if err != nil {
		error.Status = response.StatusError
		error.Message = "Password is not correct!"

		response.Send(w, http.StatusUnauthorized, error)
		return
	}

	// empty password to omit the field in response
	found.Password = ""

	sid := xid.New().String()
	auth.CreateCookie(w, sid)

	// create session error
	err = auth.CreateSession(sid, found.Username)

	if err != nil {
		error.Status = response.StatusError
		error.Message = "Failed to create user session!"

		response.Send(w, http.StatusInternalServerError, error)
		return
	}

	response.Send(w, http.StatusOK, found)
}

// Logout handles user logout
// expires user cookie and session
func Logout(w http.ResponseWriter, req *http.Request) {
	var error response.Error
	var data response.Success

	cookie, _ := req.Cookie("mstream-session")

	auth.ExpireCookie(w)
	err := auth.ExpireSession(cookie.Value)

	// expire session error
	if err != nil {
		error.Status = response.StatusError
		error.Message = "Failed to expire user session!"

		response.Send(w, http.StatusInternalServerError, error)
		return
	}

	data.Status = response.StatusSuccess
	data.Data = struct {
		Message string `json:"message"`
	}{
		Message: "Logout success!",
	}

	response.Send(w, http.StatusOK, data)
}
