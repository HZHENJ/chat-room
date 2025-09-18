package conversation

import "chatting-room/cmd/chat/dal/db"

// SendMessage
type SendMessageRequest struct {
	ConversationId int64  `json:"conversation_id"`
	SenderId       int64  `json:"sender_id"`
	Content        string `json:"content"`
}

type SendMessageResponse struct {
	MessageId  int64  `json:"message_id"`
	Seq        int64  `json:"seq"`
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// ListMessages
type ListMessagesRequest struct {
	ConversationId int64 `json:"conversation_id"`
	Limit          int   `json:"limit"`
	Offset         int   `json:"offset"`
}

type ListMessagesResponse struct {
	Messages   []*db.Message `json:"messages"`
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
}
