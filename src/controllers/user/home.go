package user

import (
	"html/template"
	"net/http"
	"strings"
	"utils/global"
	log "utils/logger"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
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

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}
