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

func Article(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/article")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/article")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "文章管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/article/list.html")
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
		articleList, err := models.GetArticleList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["ArticleList"] = articleList

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

func ArticleCreate(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/article/create")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/article/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "文章管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/article/create.html")
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
		categoryList, err := models.GetArticleCategoryList()
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
		// 文章状态
		data["ArticleStatusMap"] = global.ArticleStatus

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	} else if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/article/create")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "文章管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_title := r.PostFormValue("title")
		form_category := r.PostFormValue("category")
		form_abstract := r.PostFormValue("abstract")
		form_image := r.PostFormValue("image")
		form_content := r.PostFormValue("content")
		form_status := r.PostFormValue("status")
		form_remark := r.PostFormValue("remark")

		err := models.CreateArticle(form_title, form_category, form_abstract, form_image, form_content, form_status, form_remark)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func ArticleEdit(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/article/edit")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/article/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "文章管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/article/edit.html")
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
		categoryList, err := models.GetArticleCategoryList()
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
		// 文章状态
		data["ArticleStatusMap"] = global.ArticleStatus
		// 当前文章
		r.ParseForm()
		get_id := r.Form.Get("id")
		article, err := models.GetArticle(get_id)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if article == nil {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=文章不存在"), http.StatusFound)
			return
		}
		data["Article"] = article

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	} else if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/article/edit")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "文章管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_id := r.PostFormValue("id")
		form_title := r.PostFormValue("title")
		form_category := r.PostFormValue("category")
		form_abstract := r.PostFormValue("abstract")
		form_image := r.PostFormValue("image")
		form_content := r.PostFormValue("content")
		form_status := r.PostFormValue("status")
		form_remark := r.PostFormValue("remark")

		err := models.EditArticle(form_id, form_title, form_category, form_abstract, form_image, form_content, form_status, form_remark)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}

func ArticleDelete(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/article/delete")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/article/delete")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "文章管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		form_id := r.PostFormValue("id")

		err := models.DeleteArticle(form_id)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}
