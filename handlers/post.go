package handlers

import (
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/levblanc/golang-restful-api/constants"
	"github.com/levblanc/golang-restful-api/db"
	"github.com/levblanc/golang-restful-api/models"
	"github.com/levblanc/golang-restful-api/utils/ctx"
	"github.com/levblanc/golang-restful-api/utils/format"
	"github.com/levblanc/golang-restful-api/utils/response"
	"github.com/levblanc/golang-restful-api/utils/validator"
	"github.com/rs/xid"
)

// CreatePost creates a post record
func CreatePost(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil || !validator.ValidContentType(req) {
		response.ReqParamError(w)
		return
	}

	t := time.Now()
	creatorID := ctx.Get(req, constants.ContextUserID).(xid.ID)

	post := models.Post{
		ID:           xid.New(),
		Content:      req.FormValue("content"),
		CreatorID:    creatorID,
		CreatedAt:    t,
		CreatedTime:  format.Time(t),
		ModifiedAt:   t,
		ModifiedTime: format.Time(t),
		IsDeleted:    false,
		Comments:     []models.Comment{},
	}

	err = db.Post.Insert(post)

	if err != nil {
		response.SendError(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
		return
	}

	response.SendData(w, post)
}

// GetPost gets single post of target id
func GetPost(w http.ResponseWriter, req *http.Request) {
	var post models.Post
	var creator models.User

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

	err = db.Post.Find(bson.M{"id": id}).One(&post)

	if err != nil {
		response.SendError(
			w,
			http.StatusNotFound,
			"Post not found!",
		)
		return
	}

	err = db.User.Find(bson.M{"userId": post.CreatorID}).One(&creator)

	if err != nil {
		response.SendError(
			w,
			http.StatusNotFound,
			"Post creator not found!",
		)
		return
	}

	post.Creator = creator.Username

	response.SendData(w, post)
}

// GetAllPosts gets all post sorted in descending order
func GetAllPosts(w http.ResponseWriter, req *http.Request) {
	var postList []*models.Post
	var creator models.User
	var cache = make(map[xid.ID]string)

	err := db.Post.Find(nil).Sort("-id").All(&postList)

	if err != nil {
		response.SendError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	for _, post := range postList {
		if _, ok := cache[post.CreatorID]; ok {
			post.Creator = cache[post.CreatorID]
		} else {
			err = db.User.Find(bson.M{"userId": post.CreatorID}).One(&creator)

			if err != nil {
				cache[post.CreatorID] = ""
			} else {
				post.Creator = creator.Username
				cache[post.CreatorID] = creator.Username
			}
		}
	}

	response.SendData(w, postList)
}
