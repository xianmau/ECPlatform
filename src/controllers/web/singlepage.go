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

// 客户服务的页面
func Service(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/service")

		// render template
		t := template.New("service.html")
		t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
		t.Funcs(template.FuncMap{"ConvertToHtml": tools.ConvertToHtml})
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
		t, err := t.ParseFiles("views/web/service.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		article, err := models.GetArticleByTitle("客户服务")
		if err != nil {
			log.Error(err.Error())
			return
		}
		if article != nil {
			data["article"] = article
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

// 帮助的页面
func Help(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/help")

		// render template
		t := template.New("help.html")
		t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
		t.Funcs(template.FuncMap{"ConvertToHtml": tools.ConvertToHtml})
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
		t, err := t.ParseFiles("views/web/help.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		article, err := models.GetArticleByTitle("帮助")
		if err != nil {
			log.Error(err.Error())
			return
		}
		if article != nil {
			data["article"] = article
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

// 关于我们的页面
func About(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/about")

		// render template
		t := template.New("about.html")
		t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
		t.Funcs(template.FuncMap{"ConvertToHtml": tools.ConvertToHtml})
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
		t, err := t.ParseFiles("views/web/about.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		article, err := models.GetArticleByTitle("关于我们")
		if err != nil {
			log.Error(err.Error())
			return
		}
		if article != nil {
			data["article"] = article
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

// 供应商推广的页面
func Promotion(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/promotion")

		// render template
		t := template.New("promotion.html")
		t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
		t.Funcs(template.FuncMap{"ConvertToHtml": tools.ConvertToHtml})
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
		t, err := t.ParseFiles("views/web/promotion.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		article, err := models.GetArticleByTitle("供应商推广")
		if err != nil {
			log.Error(err.Error())
			return
		}
		if article != nil {
			data["article"] = article
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

// 隐私政策的页面
func Privacy(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/privacy")

		// render template
		t := template.New("privacy.html")
		t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
		t.Funcs(template.FuncMap{"ConvertToHtml": tools.ConvertToHtml})
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
		t, err := t.ParseFiles("views/web/privacy.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		article, err := models.GetArticleByTitle("隐私政策")
		if err != nil {
			log.Error(err.Error())
			return
		}
		if article != nil {
			data["article"] = article
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

// 友情链接的页面
func Link(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/link")

		// render template
		t := template.New("link.html")
		t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
		t.Funcs(template.FuncMap{"ConvertToHtml": tools.ConvertToHtml})
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
		t, err := t.ParseFiles("views/web/link.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

// 筛选规则的页面
func Screening(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])
	if xff_ip := r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		client_ip = xff_ip
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/screening")

		// render template
		t := template.New("screening.html")
		t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
		t.Funcs(template.FuncMap{"ConvertToHtml": tools.ConvertToHtml})
		t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
		t, err := t.ParseFiles("views/web/screening.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		article, err := models.GetArticleByTitle("筛选规则")
		if err != nil {
			log.Error(err.Error())
			return
		}
		if article != nil {
			data["article"] = article
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
