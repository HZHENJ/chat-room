package service

import (
	"chatting-room/cmd/api/biz/model/conversation"
	"chatting-room/cmd/chat/dal/db"
	"chatting-room/pkg/errno"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ConversationService struct {
	ctx context.Context
}

func NewConversationService(ctx context.Context) *ConversationService {
	return &ConversationService{ctx: ctx}
}

// CreateConversation create conversation
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

// AddMember add member
func (s *ConversationService) AddMember(request *conversation.AddMemberRequest) (
	response *conversation.AddMemberResponse,
	err error) {
	/*
		user can search conversation id or uuid or called room id to join in this room
		In the first version (MVP), if user want to join in some conversation, they can
		join in directly.
		In the second version, if user need to apply for joining the conversation, and
		the owner can allow/reject this request
	*/

	// verify user id & conversation id
	// whether there is user and conversation
	if request.UserId == 0 || request.ConversationId == 0 {
		return &conversation.AddMemberResponse{
			StatusCode: errno.ParamErr.ErrCode,
			StatusMsg:  "missing user id or conversation id",
		}, errors.New("user id or conversation id is empty")
	}

	// whether there is conversation
	conv, err := db.GetConversationById(s.ctx, request.ConversationId)
	if err != nil {
		return &conversation.AddMemberResponse{
			StatusCode: errno.ConversationNotExistErr.ErrCode,
			StatusMsg:  errno.ConversationNotExistErr.ErrMsg,
		}, err
	}

	// whether this is user
	// todo get from user service

	// Add more verification ...
	// whether status is active
	// todo status should be a static variable
	if conv.Status != "active" {
		return &conversation.AddMemberResponse{
			StatusCode: errno.ConversationNotActiveErr.ErrCode,
			StatusMsg:  errno.ConversationNotActiveErr.ErrMsg,
		}, errno.ConversationNotActiveErr
	}

	// whether user is in conversation
	if _, err := db.GetMember(s.ctx, request.ConversationId, request.UserId); err == nil {
		// member already in the conversation
		return &conversation.AddMemberResponse{
			StatusCode: errno.UserAlreadyInConversationErr.ErrCode,
			StatusMsg:  errno.UserAlreadyInConversationErr.ErrMsg,
		}, errno.UserAlreadyInConversationErr
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// not Not Found other error
		return &conversation.AddMemberResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  errno.ServiceErr.ErrMsg,
		}, err
	}

	// insert new member
	/*
		In this version we use default name like "nickname" and "member" as default value in
		Role and Nickname field, in the next step we need to fetch UserName as first nickname
	*/
	newMember := &db.Member{
		ConversationID: request.ConversationId,
		UserID:         request.UserId,
		Role:           "member",
		Nickname:       "nickname",
	}

	id, err := db.AddMember(s.ctx, newMember)
	if err != nil {
		return &conversation.AddMemberResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  errno.ServiceErr.ErrMsg,
		}, err
	}
	return &conversation.AddMemberResponse{
		MemberId:   id,
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  errno.Success.ErrMsg,
	}, nil
}
