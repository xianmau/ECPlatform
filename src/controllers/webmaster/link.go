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

func Link(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/link")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/link")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "链接管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/link/list.html")
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
		linkList, err := models.GetLinkList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["LinkList"] = linkList

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

func LinkCreate(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/link/create")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/link/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "链接管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/link/create.html")
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
		// 分类，json格式
		categoryList, err := models.GetLinkCategoryList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		jsonCategoryList, err := json.Marshal(categoryList)
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["JsonCategoryList"] = string(jsonCategoryList)
		// 链接状态
		data["LinkStatusMap"] = global.LinkStatus

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	} else if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/link/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "链接管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_title := r.PostFormValue("title")
		form_category := r.PostFormValue("category")
		form_linkUrl := r.PostFormValue("linkUrl")
		form_imageUrl := r.PostFormValue("imageUrl")
		form_status := r.PostFormValue("status")

		err := models.CreateLink(form_title, form_category, form_linkUrl, form_imageUrl, form_status)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func LinkEdit(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/link/edit")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/link/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "链接管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/link/edit.html")
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
		// 分类，json格式
		categoryList, err := models.GetLinkCategoryList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		jsonCategoryList, err := json.Marshal(categoryList)
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["JsonCategoryList"] = string(jsonCategoryList)
		// 链接状态
		data["LinkStatusMap"] = global.LinkStatus
		// 当前链接
		r.ParseForm()
		get_id := r.Form.Get("id")
		link, err := models.GetLink(get_id)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if link == nil {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=链接不存在"), http.StatusFound)
			return
		}
		data["Link"] = link

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	} else if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/link/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "链接管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_id := r.PostFormValue("id")
		form_title := r.PostFormValue("title")
		form_category := r.PostFormValue("category")
		form_linkUrl := r.PostFormValue("linkUrl")
		form_imageUrl := r.PostFormValue("imageUrl")
		form_status := r.PostFormValue("status")

		err := models.EditLink(form_id, form_title, form_category, form_linkUrl, form_imageUrl, form_status)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func LinkDelete(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/link/delete")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/link/delete")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "链接管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_id := r.PostFormValue("id")

		err := models.DeleteLink(form_id)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}
