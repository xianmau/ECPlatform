package web

import (
	"html/template"
	"models"
	"net/http"
	"strings"
	"utils/global"
	log "utils/logger"
	"utils/tools"
)

func ArticleCat(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/acat")

		// render template
		t := template.New("acat.html")
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
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
		// 判断一下是否为允许在前台显示的栏目
		if get_c != "健康食谱" {
			log.Error("not allowed category")
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
		// 获取排行文章
		hotArticleList, err := models.GetHotArticles(get_c, "10")
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["hotArticleList"] = hotArticleList

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
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/article")

		// render template
		t := template.New("article.html")
		t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
		t.Funcs(template.FuncMap{"ConvertToHtml": tools.ConvertToHtml})
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
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
		// 判断一下是否为允许在前台显示的栏目
		if article.Category != "健康食谱" {
			log.Error("not allowed category")
			return
		}
		// 文章阅读次数增加
		err = models.IncreaseArticleReadTimes(get_id)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// 获取排行文章
		hotArticleList, err := models.GetHotArticles(article.Category, "10")
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["hotArticleList"] = hotArticleList

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}
