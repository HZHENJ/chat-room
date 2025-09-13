package service

import (
	"chatting-room/cmd/api/biz/model/conversation"
	"chatting-room/cmd/chat/dal/db"
	"chatting-room/pkg/errno"
	"context"
	"errors"
	"time"
)

type ConversationService struct {
	ctx context.Context
}

func NewConversationService(ctx context.Context) *ConversationService {
	return &ConversationService{ctx: ctx}
}

// CreateConversation
func (s *ConversationService) CreateConversation(request *conversation.CreateConversationRequest) (
	response *conversation.CreateConversationResponse,
	err error) {

	// verify title and owner id
	if request.Title == "" {
		return &conversation.CreateConversationResponse{
			StatusCode: errno.ParamErr.ErrCode,
			StatusMsg:  "title cannot be empty",
		}, errors.New("title is empty")
	}

	if request.OwnerId == 0 {
		return &conversation.CreateConversationResponse{
			StatusCode: errno.ParamErr.ErrCode,
			StatusMsg:  "owner_id cannot be empty",
		}, errors.New("owner_id is empty")
	}

	conv := &db.Conversation{
		Title:   request.Title,
		OwnerID: request.OwnerId,
		Status:  "active",
		NextSeq: 1,
		// todo 找一种优雅的方式换成utc+8
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := db.CreateConversation(s.ctx, conv)
	if err != nil {
		return &conversation.CreateConversationResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  err.Error(),
		}, err
	}
	return &conversation.CreateConversationResponse{
		ConversationId: id,
		StatusCode:     errno.Success.ErrCode,
		StatusMsg:      errno.Success.ErrMsg,
	}, nil
}
