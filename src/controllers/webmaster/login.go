package webmaster

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"models"
	"net/http"
	"strings"
	"utils/authority"
	"utils/global"
	log "utils/logger"
	"utils/tools"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/login")

		t, err := template.ParseFiles("views/webmaster/login.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		t, err = t.ParseFiles("views/webmaster/styles.html", "views/webmaster/scripts.html")
		if err != nil {
			log.Error(err.Error())
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Error(err.Error())
			return
		}

	} else if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /webmaster/login")

		form_name := r.PostFormValue("Name")
		form_password := r.PostFormValue("Password")

		db, err := sql.Open("mysql", global.Config.Get("conn_str"))
		defer db.Close()
		if err != nil {
			log.Error(err.Error())
			return
		}
		rows, err := db.Query("select `tb_admin`.*,`tb_role`.`Authority` from `tb_admin` left join `tb_role` on `tb_admin`.`Role`=`tb_role`.`Name` where `tb_admin`.`Name`=? and `tb_admin`.`Password`=?", form_name, form_password)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if rows.Next() {
			var admin models.Admin
			var Name string
			var Password string
			var RoleName string
			var RoleAuthority sql.NullString
			if err := rows.Scan(&Name, &Password, &RoleName, &RoleAuthority); err != nil {
				log.Error(err.Error())
				return
			}
			admin.Name = Name
			admin.Password = Password
			admin.Role.Name = RoleName
			if RoleAuthority.Valid {
				if err := json.Unmarshal([]byte(RoleAuthority.String), &admin.Role.Authority); err != nil {
					log.Error(err.Error())
					return
				}
			}

			// check authorities
			if ok, msg := authority.Check(admin.Role, "登录"); !ok {
				http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusMovedPermanently)
				return
			}

			log.Info(client_ip + " " + admin.Name + " logged")
			session.Set("admin", admin)
			http.Redirect(w, r, "/webmaster", http.StatusMovedPermanently)
			return
		} else {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/login?msg=用户名或密码错误"), http.StatusMovedPermanently)
			return
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " get /webmaster/logout")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusMovedPermanently)
		return
	}

	if r.Method == "GET" {
		// deal with post method
		log.Info(client_ip + " get /webmaster/logout")

		log.Info(client_ip + " " + admin.Name + " logged out")
		session.Delete("admin")
		http.Redirect(w, r, "/webmaster/login", http.StatusMovedPermanently)
		return
	}
}
