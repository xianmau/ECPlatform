package webmaster

import (
	"html/template"
	"log"
	"net/http"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/webmaster/dashboard.html")
	if err != nil {
		log.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Name"] = "Xianmau"
	data["Email"] = "xianmaulin@gmail.com"

	err = t.Execute(w, data)
	if err != nil {
		log.Println(err)
		return
	}
}
