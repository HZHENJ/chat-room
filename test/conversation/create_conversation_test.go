package conversation_test

import (
	apiConv "chatting-room/cmd/api/biz/model/conversation"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试 CreateConversation
func TestCreateConversation(t *testing.T) {
	req := &apiConv.CreateConversationRequest{
		Title:   "Test Group",
		OwnerId: DemoUser.ID,
	}

	resp, err := ConversationService.CreateConversation(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(0), resp.StatusCode)
	assert.Equal(t, "Success", resp.StatusMsg)
	assert.True(t, resp.ConversationId > 0)
	t.Logf("Created Conversation: %+v", resp)
}
