package repository

import (
	"context"

	"github.com/dev-mockingbird/repository"
	"gorm.io/gorm"
)

type gormChannelRepository struct {
	db *gorm.DB
}

type gormChannelMatch struct{}

func (gormChannelMatch) Id(ids ...string) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.IN("id", ids)
	}
}

func (gormChannelMatch) UserId(ids ...string) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.IN("subscribers.user_id", ids)
	}
}

func (gormChannelMatch) ChannelId(ids ...string) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.IN("channel_id", ids)
	}
}

func (gormChannelMatch) Limit(limit int) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.SetLimit(limit)
	}
}

func (gormChannelMatch) Order(fs ...string) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.SetSort(fs...)
	}
}

func (gormChannelMatch) Offset(offset int) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.SetOffset(offset)
	}
}

func GormChannelMatch() ChannelMatch {
	return &gormChannelMatch{}
}

var _ ChannelRepository = &gormChannelRepository{}

func NewGormChannelRepository(db *gorm.DB) ChannelRepository {
	return &gormChannelRepository{db: db}
}

func (s *gormChannelRepository) Create(ctx context.Context, chs ...*Channel) error {
	return repository.New(s.db).Create(ctx, chs)
}

func (s *gormChannelRepository) Count(ctx context.Context, count *int64, opts ...repository.MatchOption) error {
	return repository.New(s.db, &Channel{}).Count(ctx, count, opts...)
}

func (s *gormChannelRepository) Find(ctx context.Context, chs *[]*Channel, opts ...repository.MatchOption) error {
	match := gormSubscriberMatch{}
	m := repository.M(chs, &Channel{}).With(&Subscriber{}, match.ChannelId(repository.Field("channels.id")))
	return repository.New(s.db).Find(ctx, m, opts...)
}

func (s *gormChannelRepository) First(ctx context.Context, ch *Channel, opts ...repository.MatchOption) error {
	return repository.New(s.db).First(ctx, ch, opts...)
}

func (s *gormChannelRepository) Delete(ctx context.Context, opts ...repository.MatchOption) error {
	return repository.New(s.db, &Channel{}).Delete(ctx, opts...)
}

func (s *gormChannelRepository) UpdateFields(ctx context.Context, fields repository.Fields, opts ...repository.MatchOption) error {
	return repository.New(s.db, &Channel{}).UpdateFields(ctx, fields, opts...)
}

func (s *gormChannelRepository) Update(ctx context.Context, v *Channel) error {
	return repository.New(s.db, &Channel{}).Update(ctx, v)
}
