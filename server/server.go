package server

import (
	"Blog/router"
	"log"
	"net/http"
)

var App = &MyServer{}

type MyServer struct {
}

func (*MyServer) Start(ip, port string) {
		// Web程序，http协议 ip port
		server := http.Server{
			Addr: ip + ":" + port,
		}
	
		// 路由
		router.Router()
	
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	
}