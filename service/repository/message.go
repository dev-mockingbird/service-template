package repository

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/dev-mockingbird/repository"
	"gorm.io/gorm"
)

type MsgType string

const (
	MsgText  = MsgType("text")
	MsgMedia = MsgType("media")
)

type Message struct {
	Id        string         `json:"id" gorm:"type:CHAR(36);primaryKey"`
	ChannelId string         `json:"channel_id" gorm:"type:CHAR(36)"`
	CreatorId string         `json:"creator_id" gorm:"type:CHAR(36)"`
	Checksum  string         `json:"checksum" gorm:"type:CHAR(36)"`
	Type      MsgType        `json:"type" gorm:"type:VARCHAR(16)"`
	Content   any            `json:"content" gorm:"type:TEXT"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (m *Message) SetCheckSum(lastMsgCheckSum string) error {
	sum := md5.Sum([]byte(m.Id + lastMsgCheckSum))
	bs, err := hex.DecodeString(string(sum[:]))
	if err != nil {
		return err
	}
	m.Checksum = string(bs)
	return nil
}

type MessageMatch interface {
	Id(ids ...string) repository.MatchOption
	Limit(limit int) repository.MatchOption
	Offset(offset int) repository.MatchOption
}

// MessageRepository repository interface
type MessageRepository interface {
	// Find
	Find(ctx context.Context, chs *[]*Message, opts ...repository.MatchOption) error
	// First get the first one based on the match options
	First(ctx context.Context, ch *Message, opts ...repository.MatchOption) error
	// Delete delete items with match options
	Delete(ctx context.Context, opts ...repository.MatchOption) error
	// UpdateFields update fields of item with match options
	UpdateFields(ctx context.Context, fields repository.Fields, opts ...repository.MatchOption) error
	// Update update single model
	Update(ctx context.Context, v *Message) error
	// Count count items with match options
	Count(ctx context.Context, count *int64, opts ...repository.MatchOption) error
	// Create create items in repository
	Create(ctx context.Context, chs ...*Message) error
}

// GetMessageRepository get the repository and match instance
func GetMessageRepository(opt any) (MessageRepository, MessageMatch) {
	return NewGormMessageRepository(opt.(*gorm.DB)), GormMessageMatch()
}
