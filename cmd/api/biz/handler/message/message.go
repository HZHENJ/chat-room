package message

import (
	"context"

	apiMsg "chatting-room/cmd/api/biz/model/conversation"
	"chatting-room/cmd/chat/service"
	"chatting-room/pkg/errno"
	"chatting-room/pkg/utils/conv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendMessage
// @router /message/send [POST]
func SendMessage(ctx context.Context, c *app.RequestContext) {
	var req apiMsg.SendMessageRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, conv.ToHertzBaseResponse(errno.ParamErr))
		return
	}

	resp, err := service.NewMessageService(ctx).SendMessage(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, conv.ToHertzBaseResponse(err))
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// ListMessages
// @router /message/list [GET]
func ListMessages(ctx context.Context, c *app.RequestContext) {
	var req apiMsg.ListMessagesRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, conv.ToHertzBaseResponse(errno.ParamErr))
		return
	}

	resp, err := service.NewMessageService(ctx).ListMessages(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, conv.ToHertzBaseResponse(err))
		return
	}
	c.JSON(consts.StatusOK, resp)
}
