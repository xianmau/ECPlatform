package web

import (
	"html/template"
	"net/http"
	"strings"
	"utils/global"
	log "utils/logger"
	"utils/tools"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	// 对首页来说需要特殊判断一下URL
	checkUrl := strings.ToLower(r.URL.Path)
	if checkUrl != "/" && checkUrl !="/index.html" && checkUrl != "/index.htm" && checkUrl != "/web" && checkUrl != "/web/home" {
		http.NotFound(w, r)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/home")

		// render template
		t := template.New("home.html")
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
		t, err := t.ParseFiles("views/web/home.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
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
