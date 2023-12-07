package main

import (
	"Blog/common"
	"Blog/server"
)

func init() {
	common.LoadTemplate()
}

func main() {
	// 程序入口
	server.App.Start("127.0.0.1", "8083")
}
