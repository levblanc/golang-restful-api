package handlers

import (
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/levblanc/golang-restful-api/constants"
	"github.com/levblanc/golang-restful-api/db"
	"github.com/levblanc/golang-restful-api/models"
	"github.com/levblanc/golang-restful-api/utils/auth"
	"github.com/levblanc/golang-restful-api/utils/format"
	"github.com/levblanc/golang-restful-api/utils/response"
	"github.com/levblanc/golang-restful-api/utils/validator"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

// Signup handles sign up requests
/**
 *
 * @api {post} /signup user sign up
 * @apiName Signup
 * @apiGroup User
 * @apiVersion  0.1.0
 *
 *
 * @apiParam  {String} username username
 *
 * @apiSuccess (200) {string} status success status
 * @apiSuccess (200) {object} data user data
 *
 * @apiParamExample  {type} Request-Example:
	{
		"username": "john doe",
		"password": "123123123"
	}
 *
 *
 * @apiSuccessExample {type} Success-Response:
	HTTP/1.1 200 OK
	{
		"status": "success",
		"data": {
			"userId": "bhc4lnh5vl33qmk8p5r0",
			"username": "john doe",
			"createdAt": "2019-02-04T22:23:26.06252+08:00",
			"createdTime": "2019-02-04 22:23:26"
		}
	}
 *
 *
*/
func Signup(w http.ResponseWriter, req *http.Request) {
	var existUser models.User

	err := req.ParseForm()

	if err != nil || !validator.ValidContentType(req) {
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
	if err != nil {
		response.SendError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

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

/**
 *
 * @api {post} /login user login
 * @apiName Login
 * @apiGroup User
 * @apiVersion  0.1.0
 *
 *
 * @apiParam  {String} username username
 * @apiParam  {String} password password
 *
 * @apiSuccess (200) {String} status success status
 * @apiSuccess (200) {Object} success data
 *
 * @apiParamExample  Request-Example:
	{
		username : "john doe",
		password: "123123123"
	}
 *
 *
 * @apiSuccessExample Success-Response:
	HTTP/1.1 200 OK
	{
		"status": "success",
		"data": {
			"userId": "bhc4lnh5vl33qmk8p5r0",
			"username": "john doe",
			"createdAt": "2019-02-04T14:23:26.062Z",
			"createdTime": "2019-02-04 22:23:26"
		}
	}
 *
 *
*/
func Login(w http.ResponseWriter, req *http.Request) {
	var found models.User

	err := req.ParseForm()

	if err != nil || !validator.ValidContentType(req) {
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
	err = auth.CreateSession(sid, found.UserID)

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
/**
 *
 * @api {post} /logout user logout
 * @apiName Logout
 * @apiGroup User
 * @apiVersion  0.1.0
 *
 *
 *
 * @apiSuccess (200) {String} status success status
 * @apiSuccess (200) {Object} data success data
 *
 *
 *
 * @apiSuccessExample {Object} Success-Response:
	HTTP/1.1 200 OK
	{
		"status": "success",
		"data": {
			"message": "Logout success!"
		}
	}
 *
 *
*/
func Logout(w http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie(constants.CookieName)

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

	data := struct {
		Message string `json:"message"`
	}{
		Message: "Logout success!",
	}

	response.SendData(w, data)
}

/**
 *
 * @api {get} /user/{id} get user by id
 * @apiName GetUser
 * @apiGroup User
 * @apiVersion  0.1.0
 *
 *
 * @apiParam  {String} id user id
 *
 * @apiSuccess (200) {String} status success status
 * @apiSuccess (200) {Object} data user data
 *
 * @apiParamExample  Request-Example:
	{
		id : "bhc4lnh5vl33qmk8p5r0"
	}
 *
 *
 * @apiSuccessExample Success-Response:
	HTTP/1.1 200 OK
	{
		"status": "success",
		"data": {
			"userId": "bhc4lnh5vl33qmk8p5r0",
			"username": "john doe",
			"createdAt": "2019-02-04T14:23:26.062Z",
			"createdTime": "2019-02-04 22:23:26"
		}
	}
 *
 *
*/
func GetUser(w http.ResponseWriter, req *http.Request) {
	var user models.User

	params := mux.Vars(req)

	id, err := xid.FromString(params["id"])
	if err != nil {
		response.SendError(
			w,
			http.StatusBadRequest,
			"Param id error!",
		)
		return
	}

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

// GetAllUsers gets all user info
/**
 *
 * @api {get} /user/all get user list
 * @apiName GetAllUsers
 * @apiGroup User
 * @apiVersion  0.1.0
 *
 *
 *
 * @apiSuccess (200) {String} status success status
 * @apiSuccess (200) {Object} user list
 *
 *
 *
 * @apiSuccessExample  Success-Response:
	HTTP/1.1 200 OK
	{
		"status": "success",
		"data": [
			{
				"userId": "bhc4lnh5vl33qmk8p5r0",
				"username": "john doe",
				"createdAt": "2019-02-04T14:23:26.062Z",
				"createdTime": "2019-02-04 22:23:26"
			}
		]
	}
 *
 *
*/
func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	var userList []*models.User

	err := db.User.Find(nil).Sort("-userId").All(&userList)

	if err != nil {
		response.SendError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	for _, user := range userList {
		// empty password to omit the field in response
		user.Password = ""
	}

	response.SendData(w, userList)
}
