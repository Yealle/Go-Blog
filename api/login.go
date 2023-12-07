package api

import (
	"Blog/common"
	"Blog/service"
	"net/http"
)

func (*APIHandler) Login(w http.ResponseWriter, r *http.Request) {
	// 接收用户名和密码 返回对应的json数据

	// 接收
	params := common.GetRequestJsonParam(r)

	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}

	// 返回
	common.Success(w, loginRes)

}
