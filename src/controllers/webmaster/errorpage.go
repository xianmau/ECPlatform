package webmaster

import (
	"html/template"
	"net/http"
	"strings"
	"utils/global"
	log "utils/glog"
)

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/errorpage")

		t, err := template.ParseFiles("views/webmaster/errorpage.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		t, err = t.ParseFiles("views/webmaster/styles.html", "views/webmaster/scripts.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		data := make(map[string]interface{})
		r.ParseForm()
		data["Msg"] = r.Form.Get("msg")
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}
