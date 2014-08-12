package web

import (
	"html/template"
	"net/http"
	"strings"
	"utils/global"
	log "utils/logger"
)

func Detail(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/detail")

		// render template
		t, err := template.ParseFiles("views/web/detail.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		t, err = t.ParseFiles("views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
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