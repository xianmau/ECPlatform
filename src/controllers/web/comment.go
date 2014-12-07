package web

import (
	"models"
	"net/http"
	"controllers/common"
)


func CommentCreate(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "POST" {
		var user models.User
		if session.Get("user") != nil {
			user = (session.Get("user")).(models.User)
		} else {
			user = *(new(models.User))
			user.Name = "游客"
		}
		form_goodsid := r.PostFormValue("goodsid")
		form_content := r.PostFormValue("content")

		err := models.CreateComment(form_goodsid, user.Name, form_content, "0")
		if err != nil {
			defaultHandler.LogError(err)
			w.Write([]byte(""))
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}
