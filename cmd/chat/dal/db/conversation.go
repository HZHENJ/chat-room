package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const ConversationTableName = "conversations"

type Conversation struct {
	ID          int64          `json:"id" gorm:"primary_key;auto_increment"`
	Uuid        string         `json:"uuid" gorm:"type:varchar(150);not null;uniqueIndex"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	OwnerID     int64          `json:"owner_id" gorm:"index;not null"`
	Status      string         `json:"status" gorm:"type:varchar(20);not null"` // active | dissolved
	NextSeq     int64          `json:"next_seq" gorm:"default:1;not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DissolvedAt *time.Time     `json:"dissolved_at,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Conversation) TableName() string {
	return ConversationTableName
}

// BeforeCreate Create UUID before insert data in db
func (c *Conversation) BeforeCreate(tx *gorm.DB) (err error) {
	if c.Uuid == "" {
		c.Uuid = uuid.NewString()
	}
	return
}

// CreateConversation create a new conversation
func CreateConversation(ctx context.Context, conversation *Conversation) (id int64, err error) {
	err = DB.WithContext(ctx).Create(conversation).Error
	if err != nil {
		return -1, err
	}
	id = conversation.ID
	return
}

// GetConversationById get conversation detail by conversation id
func GetConversationById(ctx context.Context, id int64) (conversation *Conversation, err error) {
	err = DB.WithContext(ctx).Where("id = ?", id).First(&conversation).Error
	if err != nil {
		return nil, err
	}
	return
}

// AddMember add member into conversation
func AddMember(ctx context.Context, member *Member) (id int64, err error) {
	err = DB.WithContext(ctx).Create(member).Error
	if err != nil {
		return -1, err
	}
	id = member.ID
	return
}

// UpdateConversation update conversation
func UpdateConversation(ctx context.Context, conversation *Conversation) (err error) {
	err = DB.WithContext(ctx).Save(conversation).Error
	return
}
