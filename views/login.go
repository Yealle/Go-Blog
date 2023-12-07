package views

import (
	"Blog/common"
	"config"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	//
	// var indexData IndexData
	// indexData.Title = "Go-Blog-index"
	// indexData.Desc = "Start!"
	// 拿到模板
	login := common.Template.Login

	// 填充数据
	login.WriteData(w, config.Cfg.Viewer)

}
