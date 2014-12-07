package web

import (
	"net/http"
	"models"
	"controllers/common"
)

func Message(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		err := defaultHandler.RenderTemplate("message.html", []string{"views/web/message.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html"})
		if err != nil {
			return
		}
		data := make(map[string]interface {})
		// check login
		isUserLogin := false
		var user models.User
		if session.Get("user") != nil {
			user = (session.Get("user")).(models.User)
			isUserLogin = true
		} else {
			isUserLogin = false
		}
		// set login info
		data["User"] = user
		data["IsUserLogin"] = isUserLogin

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	defaultHandler.Prepare()

	if r.Method == "POST" {
		form_name := r.PostFormValue("name")
		form_email := r.PostFormValue("email")
		form_content := r.PostFormValue("content")

		err := models.CreateMessage(form_name, form_email, form_content)
		if err != nil {
			defaultHandler.LogError(err)
			w.Write([]byte(""))
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}
