package web

import (
	"models"
	"net/http"
	"controllers/common"
)

func ArticleCat(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		err := defaultHandler.RenderTemplate("acat.html", []string{"views/web/acat.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html"})
		if err != nil {
			return
		}
		data := make(map[string]interface {})
		// check login
		isUserLogin := false
		var user models.User
		if session.Get("user") != nil {
			user = (session.Get("user")).(models.User)
			isUserLogin = true
		} else {
			isUserLogin = false
		}
		// set login info
		data["User"] = user
		data["IsUserLogin"] = isUserLogin

		// current category
		r.ParseForm()
		get_c := r.Form.Get("c")
		currentCategory, err := models.GetArticleCategory(get_c)
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		// check available category
		if get_c != "一品菜谱" && get_c != "一品养生" && get_c != "一品故事" {
			defaultHandler.LogError("not allowed category")
			return
		}
		data["currentCategory"] = currentCategory
		// get all articles of the category
		articleList, err := models.GetArticleListByCategory(get_c)
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		data["articleList"] = articleList
		// get article ranking
		hotArticleList, err := models.GetHotArticles(get_c, "10")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		data["hotArticleList"] = hotArticleList

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}

func ArticleDetail(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		err := defaultHandler.RenderTemplate("article.html", []string{"views/web/article.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html"})
		if err != nil {
			return
		}
		data := make(map[string]interface {})
		// check login
		isUserLogin := false
		var user models.User
		if session.Get("user") != nil {
			user = (session.Get("user")).(models.User)
			isUserLogin = true
		} else {
			isUserLogin = false
		}
		// set login info
		data["User"] = user
		data["IsUserLogin"] = isUserLogin

		// current article
		r.ParseForm()
		get_id := r.Form.Get("id")
		article, err := models.GetArticle(get_id)
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if article != nil {
			data["article"] = article
			// current category
			r.ParseForm()
			currentCategory, err := models.GetArticleCategory(article.Category)
			if err != nil {
				defaultHandler.LogError(err)
				return
			}
			data["currentCategory"] = currentCategory
		} else {
			data["NotFound"] = "找不到该文章的信息"
		}
		// check available category
		if article.Category != "一品菜谱" && article.Category != "一品养生" && article.Category != "一品故事" {
			defaultHandler.LogError("not allowed category")
			return
		}
		// add article click
		err = models.IncreaseArticleReadTimes(get_id)
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		// get article ranking
		hotArticleList, err := models.GetHotArticles(article.Category, "10")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		data["hotArticleList"] = hotArticleList

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}
