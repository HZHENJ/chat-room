package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const MessageTableName = "messages"

type Message struct {
	ID             int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	ConversationId int64          `json:"conversation_id" gorm:"index;not null"`
	SenderId       int64          `json:"sender_id" gorm:"index;not null"`
	Seq            int64          `json:"seq" gorm:"not null"` // 消息序号，从 Conversation.NextSeq 生成
	Content        string         `json:"content" gorm:"type:text;not null"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Message) TableName() string {
	return MessageTableName
}

// CreateMessage
func CreateMessage(ctx context.Context, msg *Message) (id int64, err error) {
	err = DB.WithContext(ctx).Create(msg).Error
	if err != nil {
		return -1, err
	}
	return msg.ID, nil
}

// GetMessagesByConversationId
func GetMessagesByConversationId(ctx context.Context, convID int64, limit, offset int) ([]*Message, error) {
	var msgs []*Message
	err := DB.WithContext(ctx).
		Where("conversation_id = ?", convID).
		Order("seq ASC").
		Limit(limit).
		Offset(offset).
		Find(&msgs).Error
	return msgs, err
}
