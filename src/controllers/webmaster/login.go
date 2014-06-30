package webmaster

import (
	"database/sql"
	"html/template"
	"net/http"
	"strings"
	"utils/global"
	log "utils/logger"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	session.Get("admin")
	if r.Method == "GET" {
		// deal with get method
		t, err := template.ParseFiles("views/webmaster/login.html")
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		t, err = t.ParseFiles("views/webmaster/styles.html", "views/webmaster/scripts.html")
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		data := make(map[string]interface{})

		err = t.Execute(w, data)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		client_ip := strings.Split(r.RemoteAddr, ":")[0]
		log.Info(client_ip + " get /webmaster/login")

	} else if r.Method == "POST" {
		// deal with post method

		form_name := r.PostFormValue("Name")
		form_password := r.PostFormValue("Password")

		db, err := sql.Open("mysql", global.Config.Get("conn_str"))
		defer db.Close()
		if err != nil {
			log.Error(err.Error())
			return
		}
		rows, err := db.Query("select * from `tb_admin` where Name=? and Password=?", form_name, form_password)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if rows.Next() {

		}

		log.Trace(form_name + "&" + form_password)

		http.Redirect(w, r, "/webmaster/login", http.StatusMovedPermanently)

		client_ip := strings.Split(r.RemoteAddr, ":")[0]
		log.Info(client_ip + " post /webmaster/login")
	}
}
