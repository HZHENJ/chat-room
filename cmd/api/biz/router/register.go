package router

import (
	"chatting-room/cmd/api/biz/router/conversation"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(r *server.Hertz) {
	conversation.Register(r)
}
