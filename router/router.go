package router

import (
	"Blog/api"
	"Blog/views"
	"net/http"
)

type IndexData struct {
}

func index() {

}

func Router() {

	// 1.访问页面 Views 2.访问数据 3.返回静态数据
	http.HandleFunc("/", views.HTML.Index)

	// http://localhost:8083/c/1 1参数 分类的id
	http.HandleFunc("/c/", views.HTML.Category)

	// 登录
	http.HandleFunc("/login", views.HTML.Login)

	// 文章页面 http://127.0.0.1:8083/p/7.html
	http.HandleFunc("/p/", views.HTML.Detail)

	// 写文章 http://127.0.0.1:8083/writing
	http.HandleFunc("/writing/", views.HTML.Writing)

	// 归档 http://127.0.0.1:8083/pigeonhole
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)

	http.HandleFunc("/api/v1/post", api.API.SaveandPost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)

	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)

	// 图片
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)

	http.HandleFunc("/api/v1/login", api.API.Login)

	// 静态资源配置
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
