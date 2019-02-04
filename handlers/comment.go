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

// AddComment adds comment to a post by id
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
