package handlers

import (
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/levblanc/golang-restful-api/constants"
	"github.com/levblanc/golang-restful-api/db"
	"github.com/levblanc/golang-restful-api/models"
	"github.com/levblanc/golang-restful-api/utils/ctx"
	"github.com/levblanc/golang-restful-api/utils/format"
	"github.com/levblanc/golang-restful-api/utils/response"
	"github.com/levblanc/golang-restful-api/utils/validator"
	"github.com/rs/xid"
)

// func getCommentCreator(list []*models.Comment) []*models.Comment {
// 	var creator models.User
// 	var cache = make(map[xid.ID]string)

// 	for _, comment := range list {
// 		if _, ok := cache[comment.CreatorID]; ok {
// 			comment.Creator = cache[comment.CreatorID]
// 		} else {
// 			err := db.User.Find(bson.M{"id": comment.CreatorID}).One(&creator)

// 			if err != nil {
// 				cache[comment.CreatorID] = ""
// 			} else {
// 				comment.Creator = creator.Username
// 				cache[comment.CreatorID] = creator.Username
// 			}
// 		}
// 	}

// 	return list
// }

/**
 * @api {post} /comment/add add comment to a post by id
 * @apiName AddComment
 * @apiGroup Comment
 * @apiVersion  0.1.0
 *
 *
 * @apiParam  {String} postId target post id
 * @apiParam  {String} content comment contents
 *
 * @apiSuccess (200) {String} status success status
 * @apiSuccess (200) {Object} data success message
 *
 * @apiParamExample  Request-Example:
	{
		postId : "bhc5sih5vl33qmk8p5t0",
		content: "nice writeup!"
	}
 *
 *
 * @apiSuccessExample Success-Response:
	HTTP/1.1 200 OK
	{
		"status" : "success",
		"data" : {
			"message": "Add comment success!"
		}
	}
 *
 *
*/
func AddComment(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil || !validator.ValidContentType(req) {
		response.ReqParamError(w)
		return
	}

	postID, err := xid.FromString(req.FormValue("postId"))

	if err != nil {
		response.SendError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	var post *models.Post
	err = db.Post.Find(bson.M{"id": postID}).One(&post)

	if err != nil {
		response.SendError(
			w,
			http.StatusBadRequest,
			"Post not found!",
		)
		return
	}

	content := req.FormValue("content")

	if content == "" {
		response.SendError(
			w,
			http.StatusBadRequest,
			"Comment content cannot be empty!",
		)
		return
	}

	t := time.Now()
	comment := models.Comment{
		ID:          xid.New(),
		Content:     content,
		CreatorID:   ctx.Get(req, constants.ContextUserID).(xid.ID),
		CreatedAt:   t,
		CreatedTime: format.Time(t),
	}

	err = db.Post.Update(
		bson.M{"id": postID},
		bson.M{"$set": bson.M{
			"comments": append(post.Comments, comment),
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
		Message: "Add comment success!",
	}

	response.SendData(w, data)
}
