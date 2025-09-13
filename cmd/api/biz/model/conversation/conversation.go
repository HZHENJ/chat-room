package conversation

// CreateConversationRequest
type CreateConversationRequest struct {
	Title   string `json:"title"`    // conversation title
	OwnerId int64  `json:"owner_id"` // owner id now is a mvp version, later we can get id from token
}

// CreateConversationResponse
type CreateConversationResponse struct {
	ConversationId int64  `json:"conversation_id"` // new conversation id
	StatusCode     int32  `json:"status_code"`     // status code
	StatusMsg      string `json:"status_msg"`      // status message
}

// AddMemberRequest
type AddMemberRequest struct {
	ConversationId int64 `json:"conversation_id"` // conversation id
	UserId         int64 `json:"user_id"`         // member id
}

// AddMemberResponse
type AddMemberResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
