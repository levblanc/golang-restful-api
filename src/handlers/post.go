package handlers

import (
	"errors"
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

/**
 *
 * @api {post} /post/create create a post
 * @apiName CreatePost
 * @apiGroup Post
 * @apiVersion  0.1.0
 *
 *
 * @apiParam  {String} content post content
 *
 * @apiSuccess (200) {String} status success status
 * @apiSuccess (200) {Object} data post data
 *
 * @apiParamExample  Request-Example:
	{
		content : "my AWESOME post!"
	}
 *
 *
 * @apiSuccessExample Success-Response:
	HTTP/1.1 200 OK
	{
		"status": "success",
		"data": {
			"id": "bhc5sih5vl33qmk8p5t0",
			"content": "my AWESOME post!",
			"creatorId": "bhc4lnh5vl33qmk8p5r0",
			"createdAt": "2019-02-04T23:46:18.156289+08:00",
			"createdTime": "2019-02-04 23:46:18",
			"modifiedAt": "2019-02-04T23:46:18.156289+08:00",
			"modifiedTime": "2019-02-04 23:46:18",
			"comments": []
		}
	}
 *
 *
*/
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

/**
 *
 * @api {get} /post/{id} get a post by id
 * @apiName GetPost
 * @apiGroup Post
 * @apiVersion  0.1.0
 *
 *
 * @apiParam  {String} id post id
 *
 * @apiSuccess (200) {String} status success status
 * @apiSuccess (200) {Object} data post data
 *
 * @apiParamExample  Request-Example:
	{
		id : "bhc5sih5vl33qmk8p5t0"
	}
 *
 *
 * @apiSuccessExample Success-Response:
	HTTP/1.1 200 OK
	{
		"status": "success",
		"data": {
			"id": "bhc5sih5vl33qmk8p5t0",
			"content": "my AWESOME post!",
			"creator": "john doe",
			"creatorId": "bhc4lnh5vl33qmk8p5r0",
			"createdAt": "2019-02-04T15:46:18.156Z",
			"createdTime": "2019-02-04 23:46:18",
			"modifiedAt": "2019-02-04T15:46:18.156Z",
			"modifiedTime": "2019-02-04 23:46:18",
			"comments": []
		}
	}
 *
 *
*/
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

/**
 *
 * @api {get} /post/all get post list in descending order
 * @apiName GetAllPosts
 * @apiGroup Post
 * @apiVersion  0.1.0
 *
 * @apiSuccess (200) {String} status success status
 * @apiSuccess (200) {Object} data post list
 *
 * @apiSuccessExample Success-Response:
	HTTP/1.1 200 OK
	{
		"status": "success",
		"data": [
			{
				"id": "bhc5vup5vl33qmk8p5u0",
				"content": "my GREAT GREAT post!",
				"creator": "john doe",
				"creatorId": "bhc4lnh5vl33qmk8p5r0",
				"createdAt": "2019-02-04T15:53:31.474Z",
				"createdTime": "2019-02-04 23:53:31",
				"modifiedAt": "2019-02-04T15:53:31.474Z",
				"modifiedTime": "2019-02-04 23:53:31",
				"comments": []
			},
			{
				"id": "bhc5sih5vl33qmk8p5t0",
				"content": "my AWESOME post!",
				"creator": "john doe",
				"creatorId": "bhc4lnh5vl33qmk8p5r0",
				"createdAt": "2019-02-04T15:46:18.156Z",
				"createdTime": "2019-02-04 23:46:18",
				"modifiedAt": "2019-02-04T15:46:18.156Z",
				"modifiedTime": "2019-02-04 23:46:18",
				"comments": []
			}
		]
	}
 *
 *
*/
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

func isPostCreator(w http.ResponseWriter, req *http.Request, id xid.ID) error {
	var post models.Post

	creatorID := ctx.Get(req, constants.ContextUserID).(xid.ID)

	err := db.Post.Find(bson.M{"id": id}).One(&post)

	if err != nil {
		return err
	}

	if post.CreatorID != creatorID {
		error := errors.New("not post creator")
		return error
	}

	return nil
}

/**
 *
 * @api {patch} /post/update update post content by id
 * @apiName UpdatePist
 * @apiGroup Post
 * @apiVersion  0.1.0
 *
 *
 * @apiParam  {String} id post id
 *
 * @apiSuccess (200) {String} status success status
 * @apiSuccess (200) {Object} data success data
 *
 * @apiParamExample   Request-Example:
	{
		id : "bhc5vup5vl33qmk8p5u0",
		content: "my GREAT GREAT GREAT post"
	}
 *
 *
 * @apiSuccessExample jSuccess-Response:
	HTTP/1.1 200 OK
	{
		"status": "success",
		"data": {
			"message": "Post update success!"
		}
	}
 *
 *
*/
func UpdatePost(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil || !validator.ValidContentType(req) {
		response.ReqParamError(w)
		return
	}

	postID, err := xid.FromString(req.FormValue("id"))

	if err != nil {
		response.SendError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	if err = isPostCreator(w, req, postID); err != nil {
		response.SendError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	postContent := req.FormValue("content")

	if postContent == "" {
		response.SendError(
			w,
			http.StatusBadRequest,
			"Update content cannot be empty!",
		)
		return
	}

	t := time.Now()
	err = db.Post.Update(
		bson.M{"id": postID},
		bson.M{"$set": bson.M{
			"content":      postContent,
			"modifiedAt":   t,
			"modifiedTime": format.Time(t),
		}},
	)

	if err != nil {
		response.SendError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	data := struct {
		Message string `json:"message"`
	}{
		Message: "Post update success!",
	}

	response.SendData(w, data)
}

/**
 *
 * @api {delete} /post/delete/{id} delete post by id
 * @apiName DeletePost
 * @apiGroup Post
 * @apiVersion  0.1.0
 *
 *
 * @apiParam  {String} id post id
 *
 * @apiSuccess (200) {String} status success status
 * @apiSuccess (200) {Object} data success data
 *
 * @apiParamExample  {type} Request-Example:
 * {
 *     id : "bhc5vup5vl33qmk8p5u0"
 * }
 *
 *
 * @apiSuccessExample {type} Success-Response:
	HTTP/1.1 200 OK
	{
		"status": "success",
		"data": {
			"message": "Delete success!"
		}
	}
 *
 *
*/
func DeletePost(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	postID, _ := xid.FromString(params["id"])

	if err := isPostCreator(w, req, postID); err != nil {
		response.SendError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	err := db.Post.Remove(bson.M{"id": postID})

	if err != nil {
		response.SendError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	data := struct {
		Message string `json:"message"`
	}{
		Message: "Delete success!",
	}

	response.SendData(w, data)
}
