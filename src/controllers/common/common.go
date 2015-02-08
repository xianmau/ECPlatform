package common

import (
	"html/template"
	"models"
	"net/http"
	"utils/global"
	log "utils/glog"
	"utils/session"
	"utils/tools"
)

type DefaultHandler struct {
	w          http.ResponseWriter
	r          *http.Request
	t          *template.Template
	ClientIP  string
	CommonPage []string
}

// return a default handler
func NewDefaultHandler(w http.ResponseWriter, r *http.Request) *DefaultHandler {
	defaultHandler := new(DefaultHandler)
	defaultHandler.w = w
	defaultHandler.r = r
	defaultHandler.CommonPage = []string{
		"views/web/styles.html",
		"views/web/scripts.html",
		"views/web/headerpart.html",
		"views/web/footerpart.html",
	}
	return defaultHandler
}

// get the client's session and record client's request info
func (d *DefaultHandler) Prepare() *session.Session {
	// prepare session
	session := global.Sessions.Prepare(d.w, d.r)
	// get client ip, it may be a proxy ip
	d.ClientIP = d.r.RemoteAddr //string([]byte(d.r.RemoteAddr)[0:strings.LastIndex(d.r.RemoteAddr, ":")])
	if xff_ip := d.r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		d.ClientIP = xff_ip
	}
	// print a log
	log.Infoln(d.ClientIP, d.r.Method, d.r.URL.Path)
	return session
}

// render template
func (d *DefaultHandler) RenderTemplate(page string, pages []string) error {
	var err error
	// render template
	d.t = template.New(page)
	d.t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
	d.t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
	d.t.Funcs(template.FuncMap{"ConvertToHtml": tools.ConvertToHtml})
	d.t, err = d.t.ParseFiles(pages...)
	if err != nil {
		log.Errorln(d.ClientIP, "render template error:", page, err)
		return err
	}
	return nil
}

// execute template
func (d *DefaultHandler) ExecuteTemplate(data map[string]interface{}) {
	// execute template
	err := d.t.Execute(d.w, data)
	if err != nil {
		log.Errorln(d.ClientIP, "execute template error:", d.t.Name(), err)
		return
	}
}

// write error log
func (d *DefaultHandler) LogError(v ...interface{}) {
	log.Errorln(d.ClientIP, "handling error:", v)
}

// check user login and set user info
func (d *DefaultHandler) CheckUserLogin(se *session.Session, data map[string]interface {}) {
	isUserLogin := false
	var user models.User
	if se.Get("user") != nil {
		user = (se.Get("user")).(models.User)
		isUserLogin = true
	} else {
		isUserLogin = false
	}
	data["User"] = user
	data["IsUserLogin"] = isUserLogin
}
