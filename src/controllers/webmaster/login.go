package webmaster

import (    _ "github.com/go-sql-driver/mysql"
	"html/template"
	"models"
	"net/http"
	"strings"
	"utils/authority"
	"utils/global"
	log "utils/logger"
	"utils/tools"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/login")

		t, err := template.ParseFiles("views/webmaster/login.html")
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
		data["login_result"] = r.Form.Get("msg")
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}

	} else if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/login")

		form_name := r.PostFormValue("Name")
		form_password := r.PostFormValue("Password")

		admin, err := models.GetAdminForLogin(form_name, form_password)
		if err != nil{
			log.Error(err.Error())
			return
		}
		if admin != nil {
			// check authorities
			if ok, msg := authority.Check(admin.Role, "登录"); !ok {
				http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
				return
			}

			log.Info(client_ip + " " + admin.Name + " logged")
			session.Set("admin", *admin)

			http.Redirect(w, r, "/webmaster/dashboard", http.StatusFound)
			return
		} else {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/login?msg=用户名或密码错误"), http.StatusFound)
			return
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " get /webmaster/logout")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with post method
		log.Info(client_ip + " get /webmaster/logout")

		log.Info(client_ip + " " + admin.Name + " logged out")
		session.Delete("admin")
		http.Redirect(w, r, "/webmaster/login", http.StatusFound)
		return
	}
}
