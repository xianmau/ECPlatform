package common

import (
	"net/http"
	"html/template"
	log "utils/logger"
	"utils/session"
	"utils/tools"
	"utils/global"
	"strings"
)

type DefaultHandler struct {
	w http.ResponseWriter
	r *http.Request
	t *template.Template
	client_ip string
}

// return a default handler
func NewDefaultHandler(w http.ResponseWriter, r *http.Request) *DefaultHandler{
	defaultHandler := new(DefaultHandler)
	defaultHandler.w = w
	defaultHandler.r = r
	return defaultHandler
}

// get the client's session and record client's request info
func (d *DefaultHandler)Prepare() (*session.Session){
	// prepare session
	session := global.Sessions.Prepare(d.w, d.r)
	// get client ip, it may be a proxy ip
	d.client_ip = string([]byte(d.r.RemoteAddr)[0:strings.LastIndex(d.r.RemoteAddr, ":")])
	if xff_ip := d.r.Header.Get("X-Forwarded-For"); xff_ip != "" {
		d.client_ip = xff_ip
	}
	// print a log
	log.Info(d.client_ip, d.r.Method, d.r.URL.Path)
	return session
}

// render template
func (d *DefaultHandler)RenderTemplate(page string, pages []string) (error){
	var err error
	// render template
	d.t = template.New(page)
	d.t.Funcs(template.FuncMap{"GetJsonData": tools.GetJsonData})
	d.t.Funcs(template.FuncMap{"UrlEncode": tools.UrlEncode})
	d.t.Funcs(template.FuncMap{"ConvertToHtml": tools.ConvertToHtml})
	d.t, err = d.t.ParseFiles(pages...)
	if err != nil {
		log.Error(d.client_ip, "render template error:", page, err)
		return err
	}
	return nil
}

// execute template
func (d *DefaultHandler)ExecuteTemplate(data map[string]interface {}){
	// execute template
	err := d.t.Execute(d.w, data)
	if err != nil {
		log.Error(d.client_ip, "execute template error:", d.t.Name(), err)
		return
	}
}

// write error log
func (d *DefaultHandler)LogError(v ...interface {}){
	log.Error(d.client_ip, "handling error:", v)
}
