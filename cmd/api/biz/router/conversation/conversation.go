package conversation

import (
	"chatting-room/cmd/api/biz/handler/conversation"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(r *server.Hertz) {
	root := r.Group("/", rootMw()...)
	{
		_conv := root.Group("/conversation", _convMw()...)
		{
			_create := _conv.Group("/create", _createMw()...)
			_create.POST("/", append(_create0Mw(), conversation.CreateConversation)...)
		}
	}
}
