package web

import (
	"models"
	"net/http"
	"controllers/common"
)

// 客户服务的页面
func Service(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		err := defaultHandler.RenderTemplate("service.html", []string{
			"views/web/service.html",
			"views/web/styles.html",
			"views/web/scripts.html",
			"views/web/headerpart.html",
			"views/web/footerpart.html",
		})
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

		article, err := models.GetArticleByTitle("客户服务")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if article != nil {
			data["article"] = article
		} else {
			data["NotFound"] = "找不到该文章的信息"
		}

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}

// 帮助的页面
func Help(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		tpls := append([]string{"views/web/help.html"}, defaultHandler.CommonPage...)
		err := defaultHandler.RenderTemplate("help.html", tpls)
		if err != nil {
			return
		}
		data := make(map[string]interface {})
		// check login
		defaultHandler.CheckUserLogin(session, data)

		article, err := models.GetArticleByTitle("帮助")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if article != nil {
			data["article"] = article
		} else {
			data["NotFound"] = "找不到该文章的信息"
		}

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}

// 关于我们的页面
func About(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		tpls := append([]string{"views/web/about.html"}, defaultHandler.CommonPage...)
		err := defaultHandler.RenderTemplate("about.html", tpls)
		if err != nil {
			return
		}
		data := make(map[string]interface {})
		// check login
		defaultHandler.CheckUserLogin(session, data)

		article, err := models.GetArticleByTitle("关于我们")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if article != nil {
			data["article"] = article
		} else {
			data["NotFound"] = "找不到该文章的信息"
		}

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}

// 供应商推广的页面
func Promotion(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		tpls := append([]string{"views/web/promotion.html"}, defaultHandler.CommonPage...)
		err := defaultHandler.RenderTemplate("promotion.html", tpls)
		if err != nil {
			return
		}
		data := make(map[string]interface {})
		// check login
		defaultHandler.CheckUserLogin(session, data)

		article, err := models.GetArticleByTitle("供应商推广")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if article != nil {
			data["article"] = article
		} else {
			data["NotFound"] = "找不到该文章的信息"
		}

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}

// 隐私政策的页面
func Privacy(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		tpls := append([]string{"views/web/privacy.html"}, defaultHandler.CommonPage...)
		err := defaultHandler.RenderTemplate("privacy.html", tpls)
		if err != nil {
			return
		}
		data := make(map[string]interface {})
		// check login
		defaultHandler.CheckUserLogin(session, data)

		article, err := models.GetArticleByTitle("隐私政策")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if article != nil {
			data["article"] = article
		} else {
			data["NotFound"] = "找不到该文章的信息"
		}

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}

// 友情链接的页面
func Link(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		tpls := append([]string{"views/web/link.html"}, defaultHandler.CommonPage...)
		err := defaultHandler.RenderTemplate("link.html", tpls)
		if err != nil {
			return
		}
		data := make(map[string]interface {})
		// check login
		defaultHandler.CheckUserLogin(session, data)

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}

// 筛选规则的页面
func Screening(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		tpls := append([]string{"views/web/screening.html"}, defaultHandler.CommonPage...)
		err := defaultHandler.RenderTemplate("screening.html", tpls)
		if err != nil {
			return
		}
		data := make(map[string]interface {})
		// check login
		defaultHandler.CheckUserLogin(session, data)

		article, err := models.GetArticleByTitle("筛选规则")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if article != nil {
			data["article"] = article
		} else {
			data["NotFound"] = "找不到该文章的信息"
		}

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}
