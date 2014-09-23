package user

import (
	"html/template"
	"net/http"
	"strings"
	"utils/global"
	log "utils/logger"
	"utils/tools"
	"models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /user/register")

		// render template
		t := template.New("register.html")
		t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
		t, err := t.ParseFiles("views/user/register.html", "views/user/styles.html", "views/user/scripts.html", "views/user/headerpart.html", "views/user/footerpart.html")
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
	} else if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /user/register")

		form_name := r.PostFormValue("Name")
		form_password := r.PostFormValue("Password")

		user, err := models.GetUserForLogin(form_name, form_password)

		if err != nil {
			log.Error(err.Error())
			return
		}
		if user != nil {
			msg := "该邮箱已经注册过了"
			w.Write([]byte(msg))
			return
		}
		if err = models.RegisterUser(form_name, form_password); err != nil{
			w.Write([]byte("注册失败，请稍后重试"))
			return
		} else {
			w.Write([]byte("success"))
			return
		}
	}
}


