package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/dev-mockingbird/repository"
	"gorm.io/gorm"
)

type Channel struct {
	Id              string         `json:"id" gorm:"type:CHAR(36);rimaryKey"`
	AvatarUrl       string         `json:"avatar_url" gorm:"type:VARCHAR(256)"`
	Description     string         `json:"description" gorm:"VARCHAR(256)"`
	Name            string         `json:"name" gorm:"type:VARCHAR(64)"`
	Group           string         `json:"group" gorm:"type:VARCHAR(64)"`
	WriteMaxRole    Role           `json:"write_max_role" gorm:"type:SMALLINT"`
	ReadMaxRole     Role           `json:"read_max_role" gorm:"type:SMALLINT"`
	LastMessageId   string         `json:"last_message_id" gorm:"type:CHAR(36)"`
	Messages        int64          `json:"messages" gorm:"type:BIGINT"`
	LastMessageType string         `json:"last_message_type" gorm:"type:VARCHAR(64)"`
	LastMessage     any            `json:"last_message" gorm:"type:TEXT"`
	LastMessageAt   sql.NullTime   `json:"last_message_at"`
	Subscribers     int64          `json:"subscribers" gorm:"type:BIGINT"`
	CreatorId       string         `json:"creator_id" gorm:"type:CHAR(36)"`
	CreatedAt       time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

type ChannelMatch interface {
	Id(ids ...string) repository.MatchOption
	UserId(ids ...string) repository.MatchOption
	ChannelId(ids ...string) repository.MatchOption
	Limit(limit int) repository.MatchOption
	Offset(offset int) repository.MatchOption
	Order(f ...string) repository.MatchOption
}

// ChannelRepository repository interface
type ChannelRepository interface {
	// Find
	Find(ctx context.Context, chs *[]*Channel, opts ...repository.MatchOption) error
	// First get the first one based on the match options
	First(ctx context.Context, ch *Channel, opts ...repository.MatchOption) error
	// Delete delete items with match options
	Delete(ctx context.Context, opts ...repository.MatchOption) error
	// UpdateFields update fields of item with match options
	UpdateFields(ctx context.Context, fields repository.Fields, opts ...repository.MatchOption) error
	// Update update single model
	Update(ctx context.Context, v *Channel) error
	// Count count items with match options
	Count(ctx context.Context, count *int64, opts ...repository.MatchOption) error
	// Create create items in repository
	Create(ctx context.Context, chs ...*Channel) error
}

// GetChannelRepository get the repository and match instance
func GetChannelRepository(opt any) (ChannelRepository, ChannelMatch) {
	return NewGormChannelRepository(opt.(*gorm.DB)), GormChannelMatch()
}
