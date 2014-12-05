package user

import (
	"html/template"
	"net/http"
	"strings"
	"utils/global"
	log "utils/logger"
	"models"
	"utils/tools"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}
	isUserLogin := false
	var user models.User
	if session.Get("user") != nil {
		user = (session.Get("user")).(models.User)
		isUserLogin = true
	} else {
		isUserLogin = false
		http.Redirect(w, r, tools.UrlEncode("/user/login"), http.StatusFound)
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /user/home")

		// render template
		t, err := template.ParseFiles("views/user/home.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		t, err = t.ParseFiles("views/user/styles.html", "views/user/scripts.html", "views/user/headerpart.html", "views/user/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		// 记录登录信息
		data["isUserLogin"] = isUserLogin
		data["user"] = user

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}


