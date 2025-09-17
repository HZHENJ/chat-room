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
		{
			_add := _conv.Group("/add", _addMw()...)
			_add.POST("/", append(_add0Mw(), conversation.AddMember)...)
		}
		{
			_dissolve := _conv.Group("/dissolve", _dissolveMw()...)
			_dissolve.POST("/", append(_dissolve0Mw(), conversation.DissolveConversation)...)
		}
	}
}
