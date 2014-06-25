package webmaster

import (
	"html/template"
	"net/http"
	log "utils/logger"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/webmaster/dashboard.html")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	data := make(map[string]interface{})
	data["Name"] = "Xianmau"
	data["Email"] = "xianmaulin@gmail.com"

	err = t.Execute(w, data)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Info("a client visit /webmaster/dashboard")
}
