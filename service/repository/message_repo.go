package repository

import (
	"context"
	"github.com/dev-mockingbird/repository"
	"gorm.io/gorm"
)

type gormMessageRepository struct {
	db *gorm.DB
}

type gormMessageMatch struct {}

func (gormMessageMatch) Id (ids ...string) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.IN("id", ids)
	}
}

func (gormMessageMatch) Limit (limit int) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.SetLimit(limit)
	}
}

func (gormMessageMatch) Offset (offset int) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.SetOffset(offset)
	}
}

func GormMessageMatch() MessageMatch {
	return &gormMessageMatch{}
}

var _ MessageRepository = &gormMessageRepository{}

func NewGormMessageRepository(db *gorm.DB) MessageRepository {
	return &gormMessageRepository{db: db}
}

func (s *gormMessageRepository) Create(ctx context.Context, chs ...*Message) error {
	return repository.New(s.db).Create(ctx, chs)
}

func (s *gormMessageRepository) Count(ctx context.Context, count *int64, opts ...repository.MatchOption) error {
	return repository.New(s.db, &Message{}).Count(ctx, count, opts...)
}

func (s *gormMessageRepository) Find(ctx context.Context, chs *[]*Message, opts ...repository.MatchOption) error {
	return repository.New(s.db).Find(ctx, chs, opts...)
}

func (s *gormMessageRepository) First(ctx context.Context, ch *Message, opts ...repository.MatchOption) error {
	return repository.New(s.db).First(ctx, ch, opts...)
}

func (s *gormMessageRepository) Delete(ctx context.Context, opts ...repository.MatchOption) error {
	return repository.New(s.db, &Message{}).Delete(ctx, opts...)
}

func (s *gormMessageRepository) UpdateFields(ctx context.Context, fields repository.Fields, opts ...repository.MatchOption) error {
	return repository.New(s.db, &Message{}).UpdateFields(ctx, fields, opts...)
}

func (s *gormMessageRepository) Update(ctx context.Context, v *Message) error {
	return repository.New(s.db, &Message{}).Update(ctx, v)
}