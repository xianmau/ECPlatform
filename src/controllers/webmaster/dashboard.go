package webmaster

import (
	"html/template"
	"net/http"
	"utils/global"
	log "utils/logger"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	session := global.Sessions.Prepare(w, r)
	session.Set("admin", "xm")

	t, err := template.ParseFiles("views/webmaster/dashboard.html")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	t, err = t.ParseFiles("views/webmaster/styles.html", "views/webmaster/scripts.html", "views/webmaster/leftside.html")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	data := make(map[string]interface{})
	data["Name"] = "Xianmau"
	data["Email"] = "xianmaulin@gmail.com"
	data["Css"] = "body{padding:0;}"

	err = t.Execute(w, data)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Info("a client visit /webmaster/dashboard")
}
