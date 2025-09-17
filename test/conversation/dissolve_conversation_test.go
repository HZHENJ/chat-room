package conversation_test

import (
	"chatting-room/cmd/chat/dal/db"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	apiConv "chatting-room/cmd/api/biz/model/conversation"
	"chatting-room/pkg/errno"
)

// TestDissolveConversation 测试群聊解散功能
func TestDissolveConversation(t *testing.T) {
	// create a conversation
	createReq := &apiConv.CreateConversationRequest{
		Title:   "Dissolve Test Group",
		OwnerId: DemoUser.ID,
	}
	createResp, err := ConversationService.CreateConversation(createReq)
	assert.NoError(t, err)
	assert.Equal(t, int32(0), createResp.StatusCode)
	convID := createResp.ConversationId
	t.Logf("Created conversation: %d", convID)

	// normal conversation: owner dissolve conversation
	dissolveReq := &apiConv.DissolveConversationRequest{
		ConversationId: convID,
		OwnerId:        DemoUser.ID,
	}
	resp, err := ConversationService.DissolveConversation(dissolveReq)
	assert.NoError(t, err)
	assert.Equal(t, int32(0), resp.StatusCode)
	assert.Equal(t, "Success", resp.StatusMsg)
	t.Logf("Dissolved conversation: %+v", resp)

	// dissolve again should put error ConversationNotActiveErr
	respAgain, err := ConversationService.DissolveConversation(dissolveReq)
	assert.Error(t, err)
	assert.Equal(t, errno.ConversationNotActiveErr.ErrCode, respAgain.StatusCode)
	assert.Equal(t, errno.ConversationNotActiveErr.ErrMsg, respAgain.StatusMsg)
	t.Logf("Dissolve again response: %+v", respAgain)

	// member (not owner) try to dissolve conversation
	createReq2 := &apiConv.CreateConversationRequest{
		Title:   "Second Test Group",
		OwnerId: DemoUser.ID,
	}
	createResp2, err := ConversationService.CreateConversation(createReq2)
	assert.NoError(t, err)
	convID2 := createResp2.ConversationId

	dissolveReqWrongUser := &apiConv.DissolveConversationRequest{
		ConversationId: convID2,
		OwnerId:        9999, // not owner
	}
	respWrongUser, err := ConversationService.DissolveConversation(dissolveReqWrongUser)
	assert.Error(t, err)
	assert.Equal(t, errno.ConversationPermissionDeniedErr.ErrCode, respWrongUser.StatusCode)
	assert.Equal(t, errno.ConversationPermissionDeniedErr.ErrMsg, respWrongUser.StatusMsg)
	t.Logf("Non-owner dissolve response: %+v", respWrongUser)

	// dissolve conversation which is not exist -> NotExist
	dissolveReqInvalid := &apiConv.DissolveConversationRequest{
		ConversationId: 999999, // assume not exist
		OwnerId:        DemoUser.ID,
	}
	respInvalid, err := ConversationService.DissolveConversation(dissolveReqInvalid)
	assert.Error(t, err)
	assert.Equal(t, errno.ConversationNotExistErr.ErrCode, respInvalid.StatusCode)
	assert.Equal(t, errno.ConversationNotExistErr.ErrMsg, respInvalid.StatusMsg)
	t.Logf("Invalid conversation dissolve response: %+v", respInvalid)

	// verify status in database about conversation is "dissolved"
	conv, err := db.GetConversationById(ctx, convID)
	assert.NoError(t, err)
	assert.Equal(t, "dissolved", conv.Status)
	assert.WithinDuration(t, time.Now(), *conv.DissolvedAt, time.Second*5)
	t.Logf("Conversation status after dissolve: %+v", conv)
}
