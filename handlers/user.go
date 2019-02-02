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
	var existUser models.User

	err := req.ParseForm()

	if err != nil {
		response.ReqParamError(w)
		return
	}

	user := models.User{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	if user.Username == "" {
		response.SendError(
			w,
			http.StatusBadRequest,
			"Username is missing!",
		)
		return
	}

	err = db.User.Find(bson.M{"username": user.Username}).One(&existUser)

	if existUser.Username == user.Username {
		response.SendError(
			w,
			http.StatusBadRequest,
			"Username taken, try again!",
		)
		return
	}

	if user.Password == "" {
		response.SendError(
			w,
			http.StatusBadRequest,
			"Password is missing!",
		)
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
		response.SendError(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
		return
	}

	// empty password to omit the field in response
	user.Password = ""

	response.SendData(w, user)
}

// GetUser gets user info by user id
func GetUser(w http.ResponseWriter, req *http.Request) {
	var user models.User

	params := mux.Vars(req)

	id, err := xid.FromString(params["id"])
	logger.Fatal(err)

	err = db.User.Find(bson.M{"userId": id}).One(&user)

	if err != nil {
		response.SendError(
			w,
			http.StatusNotFound,
			"Cannot find user!",
		)
		return
	}

	// empty password to omit the field in response
	user.Password = ""

	response.SendData(w, user)
}

// Login handles user login process
// it creates cookie and user session
func Login(w http.ResponseWriter, req *http.Request) {
	var found models.User

	err := req.ParseForm()

	if err != nil {
		response.ReqParamError(w)
		return
	}

	user := models.User{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	if user.Username == "" {
		response.SendError(
			w,
			http.StatusBadRequest,
			"Username is missing!",
		)
		return
	}

	if user.Password == "" {
		response.SendError(
			w,
			http.StatusBadRequest,
			"Password is missing!",
		)
		return
	}

	username := user.Username
	password := user.Password

	err = db.User.Find(bson.M{"username": username}).One(&found)

	if err != nil {
		response.SendError(
			w,
			http.StatusNotFound,
			"User not found!",
		)
		return
	}

	hashedPassword := found.Password

	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)

	if err != nil {
		response.SendError(
			w,
			http.StatusUnauthorized,
			"Password is not correct!",
		)
		return
	}

	// empty password to omit the field in response
	found.Password = ""

	sid := xid.New().String()
	auth.CreateCookie(w, sid)

	// create session error
	err = auth.CreateSession(sid, found.Username)

	if err != nil {
		response.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to create user session!",
		)
		return
	}

	response.SendData(w, found)
}

// Logout handles user logout
// expires user cookie and session
func Logout(w http.ResponseWriter, req *http.Request) {
	var data response.Success

	cookie, _ := req.Cookie("mstream-session")

	auth.ExpireCookie(w)
	err := auth.ExpireSession(cookie.Value)

	// expire session error
	if err != nil {
		response.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to expire user session!",
		)
		return
	}

	data.Data = struct {
		Message string `json:"message"`
	}{
		Message: "Logout success!",
	}

	response.SendData(w, data)
}
