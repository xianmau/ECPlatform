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

func Login(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /user/login")

		t := template.New("login.html")
		t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
		t, err := t.ParseFiles("views/user/login.html", "views/user/styles.html", "views/user/scripts.html", "views/user/headerpart.html", "views/user/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		data := make(map[string]interface{})
		r.ParseForm()
		data["login_result"] = r.Form.Get("msg")
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}

	} else if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /user/login")

		form_name := r.PostFormValue("Name")
		form_password := r.PostFormValue("Password")

		user, err := models.GetUserForLogin(form_name, form_password)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if user != nil {
			// check authorities
			if user.Status == 0 {
				msg := "帐号暂不可用，请联系管理员！"
				http.Redirect(w, r, tools.UrlEncode("/user/errorpage?msg="+msg), http.StatusFound)
				return
			}

			log.Info(client_ip + " " + user.Name + " logged")
			session.Set("user", *user)

			http.Redirect(w, r, "/user/home", http.StatusFound)
			return
		} else {
			msg := "用户名或密码错误"
			http.Redirect(w, r, tools.UrlEncode("/user/login?msg="+msg), http.StatusFound)
			return
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " get /user/logout")
		msg := "未登录"
		http.Redirect(w, r, tools.UrlEncode("/user/errorpage?msg="+msg), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /user/logout")

		log.Info(client_ip + " " + user.Name + " logged out")
		session.Delete("user")
		http.Redirect(w, r, "/user/login", http.StatusFound)
		return
	}
}
