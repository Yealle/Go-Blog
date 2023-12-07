package views

import (
	"Blog/common"
	"Blog/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	//
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes := service.FindPostPigeonhole()

	// 填充数据
	pigeonhole.WriteData(w, pigeonholeRes)

}
