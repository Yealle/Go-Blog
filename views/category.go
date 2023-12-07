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

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteErr(w, errors.New("路径不匹配"))
		return
	}

	// 分页
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败", err)
		categoryTemplate.WriteErr(w, errors.New("系统错误，请联系开发人员"))

		return
	}

	pageStr := r.Form.Get("page")
	page := 1
	if pageStr == "" {
		pageStr = "1"
	}

	page, _ = strconv.Atoi(pageStr)

	// 每页显示的数量

	pageSize := 10

	categoryResponse, err := service.GetPostByCategoryId(cId, page, pageSize)

	if err != nil {
		categoryTemplate.WriteErr(w, err)
		return
	}

	categoryTemplate.WriteData(w, categoryResponse)

}
