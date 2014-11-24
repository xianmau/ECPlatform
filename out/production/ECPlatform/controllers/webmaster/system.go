package webmaster

import (
	"html/template"
	"models"
	"net/http"
	"strings"
	"utils/authority"
	"utils/global"
	log "utils/logger"
	"utils/tools"
	"encoding/json"
)

func System(w http.ResponseWriter, r *http.Request) {
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
		log.Info(client_ip + " access /webmaster/system")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}
	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/system")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "系统管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/system.html")
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
		data["AuthorityList"] = global.AuthorityList
		roleList, err := models.GetRoleList()
		if err != nil{
			log.Error(err.Error())
			return
		}
		data["RoleList"] = roleList
		adminList, err := models.GetAdminList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["AdminList"] = adminList
		originList, err := models.GetOriginList()
		if err != nil{
			log.Error(err.Error())
			return
		}
		data["OriginList"] = originList
		authList, err := json.Marshal(global.AuthorityList)
		if err != nil{
			log.Error(err.Error())
			return
		}
		data["AuthList"] = string(authList)

		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}
