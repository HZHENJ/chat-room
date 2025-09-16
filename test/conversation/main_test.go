package conversation_test

import (
	conversation_dal "chatting-room/cmd/chat/dal"
	conversation_service "chatting-room/cmd/chat/service"
	"context"
	"os"
	"testing"
)

type DemoUserType struct {
	ID       int64
	UserName string
	Avatar   string
}

var (
	ctx                 = context.Background()
	ConversationService *conversation_service.ConversationService

	// test user
	DemoUser = DemoUserType{
		ID:       1,
		UserName: "test1",
	}
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")

	// 初始化数据库
	conversation_dal.Init()

	// 初始化 Service
	ConversationService = conversation_service.NewConversationService(ctx)

	// 运行测试
	code := m.Run()
	os.Exit(code)
}
