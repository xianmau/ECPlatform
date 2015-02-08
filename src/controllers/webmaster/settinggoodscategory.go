package webmaster

import (
	"html/template"
	"models"
	"net/http"
	"strings"
	"utils/authority"
	"utils/global"
	log "utils/glog"
	"utils/tools"
	"encoding/json"
)

func SettingGoodsCategory(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/setting/gcatpage")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/setting/gcatpage")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "网站设置"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/setting/gcatpage.html")
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
		goodsCategoryInfoList, err := models.GetGoodsCategoryInfoList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["goodsCategoryInfoList"] = goodsCategoryInfoList

		jsonGoodsCategoryInfoList, err := json.Marshal(goodsCategoryInfoList)
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["jsonGoodsCategoryInfoList"] = string(jsonGoodsCategoryInfoList)

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}


func SettingGoodsCategoryCreate(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/setting/gcatpage/create")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/setting/gcatpage/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "网站设置"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")
		form_info := r.PostFormValue("info")
		form_image := r.PostFormValue("image")

		err := models.CreateGoodsCategoryInfo(form_name, form_info, form_image)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func SettingGoodsCategoryEdit(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/setting/gcatpage/edit")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/setting/gcatpage/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "网站设置"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}


		form_name := r.PostFormValue("name")
		form_info := r.PostFormValue("info")
		form_image := r.PostFormValue("image")

		err := models.EditGoodsCategoryInfo(form_name, form_info, form_image)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func SettingGoodsCategoryDelete(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/setting/gcatpage/delete")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/setting/gcatpage/delete")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "网站设置"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")

		err := models.DeleteGoodsCategoryInfo(form_name)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}
