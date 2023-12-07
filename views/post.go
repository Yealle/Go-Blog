package views

import (
	"Blog/common"
	"Blog/service"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail

	// 获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")

	// 7.html

	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteErr(w, errors.New("路径不匹配"))
		return
	}

	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteErr(w, errors.New("查询出错"))
		return
	}

	detail.WriteData(w, postRes)

}
