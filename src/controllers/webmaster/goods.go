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

func Goods(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/goods")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/goods")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/goods/list.html")
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
		goodslist, err := models.GetGoodsList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["GoodsList"] = goodslist

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

func GoodsCreate(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/goods/create")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/goods/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/goods/create.html")
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
		data["CategoryList"] = categoryList
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
	} else if r.Method == "POST" {
		// deal with get method
		log.Info(client_ip + " post /webmaster/goods/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_title := r.PostFormValue("title")
		form_category := r.PostFormValue("category")
		form_content := r.PostFormValue("content")
		form_origin := r.PostFormValue("origin")
		form_unit := r.PostFormValue("unit")
		form_price := r.PostFormValue("price")
		form_shop := r.PostFormValue("shop")
		form_buylink := r.PostFormValue("buylink")
		form_images := r.PostFormValue("images")
		form_certificates := r.PostFormValue("certificates")
		form_status := r.PostFormValue("status")

		err := models.CreateGoods(form_title, form_category, form_content, form_origin, form_unit, form_price, form_shop, form_buylink, form_images, form_certificates, form_status)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func GoodsEdit(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/goods/edit")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/goods/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/goods/edit.html")
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
		data["CategoryList"] = categoryList
		jsonCategoryList, err := json.Marshal(categoryList)
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["JsonCategoryList"] = string(jsonCategoryList)
		r.ParseForm()
		get_id := r.Form.Get("id")
		goods, err := models.GetGoods(get_id)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if goods == nil {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=商品不存在"), http.StatusFound)
			return
		}
		data["Goods"] = goods

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	} else if r.Method == "POST" {
		// deal with get method
		log.Info(client_ip + " post /webmaster/goods/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_id := r.PostFormValue("id")
		form_title := r.PostFormValue("title")
		form_category := r.PostFormValue("category")
		form_content := r.PostFormValue("content")
		form_origin := r.PostFormValue("origin")
		form_unit := r.PostFormValue("unit")
		form_price := r.PostFormValue("price")
		form_shop := r.PostFormValue("shop")
		form_buylink := r.PostFormValue("buylink")
		form_images := r.PostFormValue("images")
		form_certificates := r.PostFormValue("certificates")
		form_status := r.PostFormValue("status")

		err := models.EditGoods(form_id, form_title, form_category, form_content, form_origin, form_unit, form_price, form_shop, form_buylink, form_images, form_certificates, form_status)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func GoodsDelete(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /webmaster/goods/delete")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with get method
		log.Info(client_ip + " post /webmaster/goods/delete")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_id := r.PostFormValue("id")

		err := models.DeleteGoods(form_id)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}
