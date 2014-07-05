package webmaster

import (
	"database/sql"
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

func GoodsCategory(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " get /webmaster/goodscategory")
		http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg=未登录"), http.StatusFound)
		return
	}

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /webmaster/goodscategory")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "商品管理"); !ok {
			http.Redirect(w, r, tools.UrlEncode("/webmaster/errorpage?msg="+msg), http.StatusFound)
			return
		}

		// render template
		t, err := template.ParseFiles("views/webmaster/goodscategory/list.html")
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

		db, err := sql.Open("mysql", global.Config.Get("conn_str"))
		defer db.Close()
		if err != nil {
			log.Error(err.Error())
			return
		}
		rows, err := db.Query("select * from `tb_goods_category` where true")
		defer rows.Close()
		if err != nil {
			log.Error(err.Error())
			return
		}
		categorylist := []models.GoodsCategory{}
		for rows.Next() {
			var goodscategory models.GoodsCategory
			var Title string
			var Parent string
			var Ordering int
			if err := rows.Scan(&Title, &Parent, &Ordering); err != nil {
				log.Error(err.Error())
				return
			}
			goodscategory.Title = Title
			goodscategory.Parent = Parent
			goodscategory.Ordering = Ordering
			categorylist = append(categorylist, goodscategory)
		}
		data["CategoryList"] = categorylist
		log.Trace(categorylist)
		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}
