package repository

import (
	"context"
	"time"

	"github.com/dev-mockingbird/repository"
	"gorm.io/gorm"
)

type Role int

const (
	RoleOwner   = Role(0)
	RoleManager = Role(1)
	RoleMember  = Role(2)
	RoleGuest   = Role(3)
	RoleAny     = Role(10)
)

type Subscriber struct {
	Id                    string         `json:"id" gorm:"type:CHAR(36);primaryKey"`
	UserId                string         `json:"user_id" gorm:"type:CHAR(36)"`
	ChannelId             string         `json:"channel_id" gorm:"type:CHAR(36)"`
	Nickname              string         `json:"nickname,omitempty" gorm:"type:VARCHAR(64)"`
	AvatarUrl             string         `json:"avatar_url,omitempty" gorm:"type:VARCHAR(1024)"`
	Online                bool           `json:"online" gorm:"type:SMALLINT"`
	Role                  Role           `json:"role" gorm:"type:SMALLINT"`
	LastReceivedMessageId string         `json:"last_received_message_id" gorm:"type:CHAR(36)"`
	LastReadMessageId     string         `json:"last_read_message_id" gorm:"type:CHAR(36)"`
	CreatedAt             time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
}

type SubscriberMatch interface {
	Id(ids ...string) repository.MatchOption
	UserId(ids ...string) repository.MatchOption
	Online(bool) repository.MatchOption
	ChannelId(id any) repository.MatchOption
	Limit(limit int) repository.MatchOption
	Offset(offset int) repository.MatchOption
}

// SubscriberRepository repository interface
type SubscriberRepository interface {
	// Find
	Find(ctx context.Context, chs *[]*Subscriber, opts ...repository.MatchOption) error
	// First get the first one based on the match options
	First(ctx context.Context, ch *Subscriber, opts ...repository.MatchOption) error
	// Delete delete items with match options
	Delete(ctx context.Context, opts ...repository.MatchOption) error
	// UpdateFields update fields of item with match options
	UpdateFields(ctx context.Context, fields repository.Fields, opts ...repository.MatchOption) error
	// Update update single model
	Update(ctx context.Context, v *Subscriber) error
	// Count count items with match options
	Count(ctx context.Context, count *int64, opts ...repository.MatchOption) error
	// Create create items in repository
	Create(ctx context.Context, chs ...*Subscriber) error
}

// GetSubscriberRepository get the repository and match instance
func GetSubscriberRepository(opt any) (SubscriberRepository, SubscriberMatch) {
	return NewGormSubscriberRepository(opt.(*gorm.DB)), GormSubscriberMatch()
}
