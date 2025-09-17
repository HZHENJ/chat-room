package conversation

import (
	"chatting-room/cmd/api/biz/model/conversation"
	"chatting-room/cmd/chat/service"
	"chatting-room/pkg/errno"
	"chatting-room/pkg/utils/conv"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateConversation
// @router /conversation/create [POST]
func CreateConversation(ctx context.Context, c *app.RequestContext) {
	var req conversation.CreateConversationRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(consts.StatusBadRequest, conv.ToHertzBaseResponse(errno.ParamErr))
	}
	resp, err := service.NewConversationService(ctx).CreateConversation(&req)
	if err != nil {
		c.JSON(consts.StatusOK, conv.ToHertzBaseResponse(err))
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// AddMember
// @route /conversation/add_member [POST]
func AddMember(ctx context.Context, c *app.RequestContext) {
	var req conversation.AddMemberRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(consts.StatusBadRequest, conv.ToHertzBaseResponse(errno.ParamErr))
		return
	}
	resp, err := service.NewConversationService(ctx).AddMember(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, conv.ToHertzBaseResponse(err))
		return
	}
	c.JSON(consts.StatusOK, resp)
}
