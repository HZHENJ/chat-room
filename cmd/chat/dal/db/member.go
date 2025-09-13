package db

import (
	"time"

	"gorm.io/gorm"
)

const MemberTableName = "members"

type Member struct {
	ID             int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	ConversationID int64  `json:"conversation_id" gorm:"index;not null;"`
	UserID         int64  `json:"user_id" gorm:"index;not null"`
	Role           string `json:"role" gorm:"type:varchar(20);not null"`
	Nickname       string `json:"nickname" gorm:"type:varchar(350)"`
	//MutedUntil     *time.Time     `json:"muted_until,omitempty"`
	ReadSeq   int64          `json:"read_seq" gorm:"default:0;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Member) TableName() string {
	return MemberTableName
}
