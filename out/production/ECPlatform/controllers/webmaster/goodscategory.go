package webmaster

import (
	"encoding/json"
	"html/template"
	"models"
	"net/http"
	"strings"
	"utils/authority"
	"utils/global"
	log "utils/logger"
	"utils/tools"
)

func GoodsCategory(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/goodscategory")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/goodscategory")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/goodscategory/list.html")
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
		categoryList, err := models.GetGoodsCategoryList()
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

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

func GoodsCategoryCreate(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/goodscategory/create")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/goodscategory/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")
		form_parent := r.PostFormValue("parent")
		form_ordering := r.PostFormValue("ordering")

		err := models.CreateGoodsCategory(form_name, form_parent, form_ordering)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func GoodsCategoryEdit(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/goodscategory/edit")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/goodscategory/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")
		form_parent := r.PostFormValue("parent")
		form_ordering := r.PostFormValue("ordering")

		err := models.EditGoodsCategory(form_name, form_parent, form_ordering)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func GoodsCategoryDelete(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/goodscategory/delete")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/goodscategory/delete")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_name := r.PostFormValue("name")

		err := models.DeleteGoodsCategory(form_name)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}
