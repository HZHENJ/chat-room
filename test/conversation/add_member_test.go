package conversation_test

import (
	apiConv "chatting-room/cmd/api/biz/model/conversation"
	"chatting-room/pkg/errno"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test AddMember
func TestAddMember(t *testing.T) {
	// Assume that there is conversation which id is 8 and user which id is 2
	req := &apiConv.AddMemberRequest{
		ConversationId: 8,
		UserId:         2,
	}

	// normal situation
	resp, err := ConversationService.AddMember(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(0), resp.StatusCode)
	assert.Equal(t, "Success", resp.StatusMsg)
	t.Logf("AddMember response: %+v", resp)

	// join in a conversation duplicately should be failed
	respDup, err := ConversationService.AddMember(req)
	assert.Error(t, err)
	assert.NotNil(t, respDup)
	assert.Equal(t, errno.UserAlreadyInConversationErr.ErrCode, respDup.StatusCode)
	assert.Equal(t, errno.UserAlreadyInConversationErr.ErrMsg, respDup.StatusMsg)
	t.Logf("Duplicate AddMember response: %+v", respDup)

	// join in a conversation which is not exist
	reqInvalidConv := &apiConv.AddMemberRequest{
		ConversationId: 99999, // assume is not exist
		UserId:         2,
	}
	respInvalid, err := ConversationService.AddMember(reqInvalidConv)
	assert.Error(t, err)
	assert.NotNil(t, respInvalid)
	assert.Equal(t, errno.ConversationNotExistErr.ErrCode, respInvalid.StatusCode)
	t.Logf("Invalid Conversation AddMember response: %+v", respInvalid)

	// join in a conversation which is
	reqDissolved := &apiConv.AddMemberRequest{
		ConversationId: 9, // assume conversation 9 is dissolved
		UserId:         3,
	}
	respDissolved, err := ConversationService.AddMember(reqDissolved)
	assert.Error(t, err)
	assert.NotNil(t, respDissolved)
	assert.Equal(t, errno.ConversationNotActiveErr.ErrCode, respDissolved.StatusCode)
	t.Logf("Dissolved Conversation AddMember response: %+v", respDissolved)
}
