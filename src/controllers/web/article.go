package web

import (
	"html/template"
	"models"
	"net/http"
	"strings"
	"utils/global"
	log "utils/logger"
)

func ArticleCat(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/acat")

		// render template
		t := template.New("acat.html")
		t, err := t.ParseFiles("views/web/acat.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		// 当前分类
		r.ParseForm()
		get_c := r.Form.Get("c")
		currentCategory, err := models.GetArticleCategory(get_c)
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["currentCategory"] = currentCategory

		// 获取当前分类下的所有文章
		articleList, err := models.GetArticleListByCategory(get_c)
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["articleList"] = articleList
		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

func ArticleDetail(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/article")

		// render template
		t := template.New("article.html")
		t, err := t.ParseFiles("views/web/article.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		// 当前文章
		r.ParseForm()
		get_id := r.Form.Get("id")
		article, err := models.GetArticle(get_id)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if article != nil {
			data["article"] = article
			// 当前分类
			r.ParseForm()
			currentCategory, err := models.GetArticleCategory(article.Category)
			if err != nil {
				log.Error(err.Error())
				return
			}
			data["currentCategory"] = currentCategory
		} else {
			data["NotFound"] = "找不到该文章的信息"
		}

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}
