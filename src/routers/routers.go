package routers

import (
	"controllers/web"
	"controllers/webmaster"
	"controllers/upload"
	"net/http"
)

// 注册路由
func Register() {
	// 注册静态文件
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	registerWebmaster()

	registerWeb()
}

func registerWeb() {
	http.HandleFunc("/web", web.Home)
	http.HandleFunc("/web/home", web.Home)

	http.HandleFunc("/web/join", web.Join)
}

func registerWebmaster() {
	// 错误处理
	http.HandleFunc("/webmaster/errorpage", webmaster.ErrorPage)

	// 仪表板
	http.HandleFunc("/webmaster/dashboard", webmaster.Dashboard)
	http.HandleFunc("/webmaster", webmaster.Dashboard)

	// 商品分类
	http.HandleFunc("/webmaster/goodscategory", webmaster.GoodsCategory)
	http.HandleFunc("/webmaster/goodscategory/create", webmaster.GoodsCategoryCreate)
	http.HandleFunc("/webmaster/goodscategory/edit", webmaster.GoodsCategoryEdit)
	http.HandleFunc("/webmaster/goodscategory/delete", webmaster.GoodsCategoryDelete)
	// 商品
	http.HandleFunc("/webmaster/goods", webmaster.Goods)
	http.HandleFunc("/webmaster/goods/create", webmaster.GoodsCreate)
	http.HandleFunc("/webmaster/goods/edit", webmaster.GoodsEdit)
	http.HandleFunc("/webmaster/goods/delete", webmaster.GoodsDelete)

	// 文章分类
	http.HandleFunc("/webmaster/articlecategory", webmaster.ArticleCategory)
	http.HandleFunc("/webmaster/articlecategory/create", webmaster.ArticleCategoryCreate)
	http.HandleFunc("/webmaster/articlecategory/edit", webmaster.ArticleCategoryEdit)
	http.HandleFunc("/webmaster/articlecategory/delete", webmaster.ArticleCategoryDelete)
	// 文章
	http.HandleFunc("/webmaster/article", webmaster.Article)
	http.HandleFunc("/webmaster/article/create", webmaster.ArticleCreate)
	http.HandleFunc("/webmaster/article/edit", webmaster.ArticleEdit)
	http.HandleFunc("/webmaster/article/delete", webmaster.ArticleDelete)


	// 登录和注销
	http.HandleFunc("/webmaster/login", webmaster.Login)
	http.HandleFunc("/webmaster/logout", webmaster.Logout)

	// 上传文件相关
	http.HandleFunc("/uploadprocess/webmaster/ckeditor", upload.CKEditorWebmaster)
	http.HandleFunc("/uploadprocess/webmaster/uploadify", upload.UploadifyWebmaster)
}
