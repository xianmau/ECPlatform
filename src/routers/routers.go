package routers

import (
	"controllers/webmaster"
	"net/http"
)

// 注册路由
func Register() {
	http.HandleFunc("/webmaster", webmaster.GetDashboard)
	http.HandleFunc("/webmaster/dashboard", webmaster.GetDashboard)

	http.HandleFunc("/webmaster/login", webmaster.Login)

}
