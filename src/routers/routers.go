package routers

import (
	"controllers/webmaster"
	"net/http"
)

// 注册路由
func Register() {
	// 注册静态文件
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))))

	// 错误处理
	http.HandleFunc("/webmaster/errorpage", webmaster.ErrorPage)

	// 仪表板
	http.HandleFunc("/webmaster", webmaster.Dashboard)
	http.HandleFunc("/webmaster/dashboard", webmaster.Dashboard)

	// 商品分类
	http.HandleFunc("webmaster/goodscategory", webmaster.GoodsCategory)
	//http.HandleFunc("webmaster/goodscategory/create", webmaster.GoodsCategoryCreate)
	//http.HandleFunc("webmaster/goodscategory/edit", webmaster.GoodsCategoryEdit)
	//http.HandleFunc("webmaster/goodscategory/delete", webmaster.GoodsCategoryDelete)
	// 商品
	//http.HandleFunc("webmaster/goods", webmaster.Goods)
	//http.HandleFunc("webmaster/goods/create", webmaster.GoodsCreate)
	//http.HandleFunc("webmaster/goods/edit", webmaster.GoodsEdit)
	//http.HandleFunc("webmaster/goods/delete", webmaster.GoodsDelete)

	// 登录和注销
	http.HandleFunc("/webmaster/login", webmaster.Login)
	http.HandleFunc("/webmaster/logout", webmaster.Logout)

}
