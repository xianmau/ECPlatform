package webmaster

import (
	"encoding/json"
	"html/template"
	"models"
	"net/http"
	"strings"
	"utils/authority"
	"utils/global"
	log "utils/glog"
	"utils/tools"
)

func SettingHomePage(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/setting/homepage")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/setting/homepage")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "网站设置"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/setting/homepage.html")
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
		webHomeInfoList, err := models.GetWebHomeInfoList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["webHomeInfoList"] = webHomeInfoList

		jsonWebHomeInfoList, err := json.Marshal(webHomeInfoList)
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["jsonWebHomeInfoList"] = string(jsonWebHomeInfoList)

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

func SettingHomePageCreate(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/setting/homepage/create")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/setting/homepage/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "网站设置"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_floor1 := r.PostFormValue("floor1")
		form_floor2 := r.PostFormValue("floor2")
		form_floor3 := r.PostFormValue("floor3")
		form_floor4 := r.PostFormValue("floor4")
		form_floor5 := r.PostFormValue("floor5")
		form_floor6 := r.PostFormValue("floor6")
		form_floor7 := r.PostFormValue("floor7")
		form_floor8 := r.PostFormValue("floor8")

		err := models.CreateWebHomeInfo(form_floor1, form_floor2, form_floor3, form_floor4, form_floor5, form_floor6, form_floor7, form_floor8)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func SettingHomePageEdit(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/setting/homepage/edit")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/setting/homepage/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "网站设置"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_id := r.PostFormValue("id")
		form_floor1 := r.PostFormValue("floor1")
		form_floor2 := r.PostFormValue("floor2")
		form_floor3 := r.PostFormValue("floor3")
		form_floor4 := r.PostFormValue("floor4")
		form_floor5 := r.PostFormValue("floor5")
		form_floor6 := r.PostFormValue("floor6")
		form_floor7 := r.PostFormValue("floor7")
		form_floor8 := r.PostFormValue("floor8")

		err := models.EditWebHomeInfo(form_id, form_floor1, form_floor2, form_floor3, form_floor4, form_floor5, form_floor6, form_floor7, form_floor8)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func SettingHomePageDelete(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/setting/homepage/delete")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/setting/homepage/delete")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "网站设置"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_id := r.PostFormValue("id")

		err := models.DeleteWebHomeInfo(form_id)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}
