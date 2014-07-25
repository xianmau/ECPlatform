package webmaster

import (
	"html/template"
	"models"
	"net/http"
	"strings"
	"utils/authority"
	"utils/global"
	log "utils/logger"
	"utils/tools"
)

func User(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/user")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " post /webmaster/user")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "用户管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/user/list.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		t, err = t.ParseFiles("views/webmaster/leftside.html", "views/webmaster/styles.html", "views/webmaster/scripts.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})
		data["Admin"] = admin
		userList, err := models.GetUserList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["UserList"] = userList

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/user/create")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/user/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "用户管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/user/create.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		t, err = t.ParseFiles("views/webmaster/leftside.html", "views/webmaster/styles.html", "views/webmaster/scripts.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})
		data["Admin"] = admin
		// 用户级别
		data["UserLevelMap"] = global.UserLevel
		// 用户状态
		data["UserStatusMap"] = global.UserStatus

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	} else if r.Method == "POST" {
		// deal with get method
		log.Info(client_ip + " post /webmaster/user/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "用户管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")
		form_password := r.PostFormValue("password")
		form_level := r.PostFormValue("level")
		form_status := r.PostFormValue("status")

		err := models.CreateUser(form_name, form_password, form_level, form_status)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func UserEdit(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/user/edit")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/user/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "用户管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/user/edit.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		t, err = t.ParseFiles("views/webmaster/leftside.html", "views/webmaster/styles.html", "views/webmaster/scripts.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})
		data["Admin"] = admin
		// 用户级别
		data["UserLevelMap"] = global.UserLevel
		// 用户状态
		data["UserStatusMap"] = global.UserStatus
		// 当前用户
		r.ParseForm()
		get_name := r.Form.Get("name")
		user, err := models.GetUser(get_name)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if user == nil {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=用户不存在"), http.StatusFound)
			return
		}
		data["User"] = user

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	} else if r.Method == "POST" {
		// deal with get method
		log.Info(client_ip + " post /webmaster/user/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "用户管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")
		form_level := r.PostFormValue("level")
		form_status := r.PostFormValue("status")

		err := models.EditUserByAdmin(form_name, form_level, form_status)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func UserResetPassword(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/user/resetpassword")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/user/resetpassword")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "用户管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")
		form_password := r.PostFormValue("password")

		err := models.EditUserPassword(form_name, form_password)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/user/delete")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/user/delete")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "用户管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")

		err := models.DeleteUser(form_name)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/user/detail")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/user/detail")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "用户管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/user/detail.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		t, err = t.ParseFiles("views/webmaster/leftside.html", "views/webmaster/styles.html", "views/webmaster/scripts.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})
		data["Admin"] = admin
		// 用户级别
		data["UserLevelMap"] = global.UserLevel
		// 用户状态
		data["UserStatusMap"] = global.UserStatus
		// 当前用户
		r.ParseForm()
		get_name := r.Form.Get("name")
		user, err := models.GetUser(get_name)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if user == nil {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=用户不存在"), http.StatusFound)
			return
		}
		data["User"] = user

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}
