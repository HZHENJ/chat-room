package main

import (
	"chatting-room/cmd/api/biz/router"
	"chatting-room/cmd/chat/dal"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {

	// Initialize database
	dal.Init()

	// Initialize Hertz server
	h := server.Default()

	// register route
	router.Register(h)

	// start server
	h.Spin()
}
