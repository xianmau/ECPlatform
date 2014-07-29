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

func Shop(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/shop")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/shop")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商店管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/shop/list.html")
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
		shopList, err := models.GetShopList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["ShopList"] = shopList

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

func ShopCreate(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/shop/create")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/shop/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商店管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/shop/create.html")
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
		// 商店类型
		data["ShopKindMap"] = global.ShopKind
		// 商店状态
		data["ShopStatusMap"] = global.ShopStatus

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	} else if r.Method == "POST" {
		// deal with get method
		log.Info(client_ip + " post /webmaster/shop/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商店管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")
		form_username := r.PostFormValue("username")
		form_kind := r.PostFormValue("kind")
		//form_introduce := r.PostFormValue("introduce")
		//form_applystatement := r.PostFormValue("applystatement")
		form_status := r.PostFormValue("status")

		err := models.CreateShop(form_name, form_username, form_kind, "", "", form_status)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func ShopEdit(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/shop/edit")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/shop/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商店管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/shop/edit.html")
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
		// 商店类型
		data["ShopKindMap"] = global.ShopKind
		// 商店状态
		data["ShopStatusMap"] = global.ShopStatus
		// 当前用户
		r.ParseForm()
		get_name := r.Form.Get("name")
		shop, err := models.GetShop(get_name)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if shop == nil {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=商店不存在"), http.StatusFound)
			return
		}
		data["Shop"] = shop

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	} else if r.Method == "POST" {
		// deal with get method
		log.Info(client_ip + " post /webmaster/shop/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商店管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")
		form_username := r.PostFormValue("username")
		form_kind := r.PostFormValue("kind")
		form_introduce := r.PostFormValue("introduce")
		form_applystatement := r.PostFormValue("applystatement")
		form_status := r.PostFormValue("status")

		err := models.EditShop(form_name, form_username, form_kind, form_introduce, form_applystatement, form_status)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func ShopDelete(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/shop/delete")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/shop/delete")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商店管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")

		err := models.DeleteShop(form_name)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

