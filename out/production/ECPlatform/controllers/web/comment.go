package web

import (
	"models"
	"net/http"
	"strings"
	"utils/global"
	log "utils/logger"
)


func CommentCreate(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	var user models.User
	if session.Get("user") != nil {
		user = (session.Get("user")).(models.User)
	} else {
		log.Info(client_ip + " access /web/comment/create")
		// return faild
		//w.Write([]byte("faild"))
		//return
		user = *(new(models.User))
		user.Name = "游客"
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /web/comment/create")

		name := user.Name
		form_goodsid := r.PostFormValue("goodsid")
		form_content := r.PostFormValue("content")

		err := models.CreateComment(form_goodsid, name, form_content, "0")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}
