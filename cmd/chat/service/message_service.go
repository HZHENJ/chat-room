package service

import (
	"chatting-room/cmd/api/biz/model/conversation"
	"chatting-room/cmd/chat/dal/db"
	"chatting-room/pkg/errno"
	"context"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{ctx: ctx}
}

func (s *MessageService) SendMessage(req *conversation.SendMessageRequest) (*conversation.SendMessageResponse, error) {
	conv, err := db.GetConversationById(s.ctx, req.ConversationId)
	if err != nil {
		return &conversation.SendMessageResponse{
			StatusCode: errno.ConversationNotExistErr.ErrCode,
			StatusMsg:  errno.ConversationNotExistErr.ErrMsg,
		}, err
	}

	if conv.Status != "active" {
		return &conversation.SendMessageResponse{
			StatusCode: errno.ConversationNotActiveErr.ErrCode,
			StatusMsg:  errno.ConversationNotActiveErr.ErrMsg,
		}, errno.ConversationNotActiveErr
	}

	// 分配消息序号
	seq := conv.NextSeq
	conv.NextSeq++

	msg := &db.Message{
		ConversationId: conv.ID,
		SenderId:       req.SenderId,
		Seq:            seq,
		Content:        req.Content,
	}

	if _, err := db.CreateMessage(s.ctx, msg); err != nil {
		return &conversation.SendMessageResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  errno.ServiceErr.ErrMsg,
		}, err
	}

	if err := db.UpdateConversation(s.ctx, conv); err != nil {
		return &conversation.SendMessageResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  errno.ServiceErr.ErrMsg,
		}, err
	}

	return &conversation.SendMessageResponse{
		MessageId:  msg.ID,
		Seq:        seq,
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	}, nil
}

func (s *MessageService) ListMessages(req *conversation.ListMessagesRequest) (*conversation.ListMessagesResponse, error) {
	msgs, err := db.GetMessagesByConversationId(s.ctx, req.ConversationId, req.Limit, req.Offset)
	if err != nil {
		return &conversation.ListMessagesResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  errno.ServiceErr.ErrMsg,
		}, err
	}

	return &conversation.ListMessagesResponse{
		Messages:   msgs,
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	}, nil
}
