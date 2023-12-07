package views

import (
	"Blog/common"
	"Blog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	//
	// var indexData IndexData
	// indexData.Title = "Go-Blog-index"
	// indexData.Desc = "Start!"
	// 拿到模板
	index := common.Template.Index

	// 数据库查询

	// 分页
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败", err)
		return
	}

	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	// 每页显示的数量
	pageSize := 10

	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")

	hr, err := service.GetIndexInfo(slug, page, pageSize)

	if err != nil {
		log.Println("Index获取数据出错", err)
		index.WriteErr(w, errors.New("系统错误，请联系开发人员"))
	}
	// 填充数据
	index.WriteData(w, hr)

}
