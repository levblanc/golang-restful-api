package handlers

import (
	"net/http"
	"time"

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
