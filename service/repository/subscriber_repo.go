package repository

import (
	"context"
	"reflect"

	"github.com/dev-mockingbird/repository"
	"gorm.io/gorm"
)

type gormSubscriberRepository struct {
	db *gorm.DB
}

type gormSubscriberMatch struct{}

func (gormSubscriberMatch) Id(ids ...string) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.IN("id", ids)
	}
}

func (gormSubscriberMatch) ChannelId(id any) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		vo := reflect.ValueOf(id)
		if vo.Kind() == reflect.Slice {
			opts.IN("subscribers.channel_id", id)
			return
		}
		opts.EQ("subscribers.channel_id", id)
	}
}

func (gormSubscriberMatch) UserId(ids ...string) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.IN("user_id", ids)
	}
}

func (gormSubscriberMatch) Online(b bool) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.EQ("online", func() int {
			if b {
				return 1
			}
			return 0
		}())
	}
}

func (gormSubscriberMatch) Limit(limit int) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.SetLimit(limit)
	}
}

func (gormSubscriberMatch) Offset(offset int) repository.MatchOption {
	return func(opts *repository.MatchOptions) {
		opts.SetOffset(offset)
	}
}

func GormSubscriberMatch() SubscriberMatch {
	return &gormSubscriberMatch{}
}

var _ SubscriberRepository = &gormSubscriberRepository{}

func NewGormSubscriberRepository(db *gorm.DB) SubscriberRepository {
	return &gormSubscriberRepository{db: db}
}

func (s *gormSubscriberRepository) Create(ctx context.Context, chs ...*Subscriber) error {
	return repository.New(s.db).Create(ctx, chs)
}

func (s *gormSubscriberRepository) Count(ctx context.Context, count *int64, opts ...repository.MatchOption) error {
	return repository.New(s.db, &Subscriber{}).Count(ctx, count, opts...)
}

func (s *gormSubscriberRepository) Find(ctx context.Context, chs *[]*Subscriber, opts ...repository.MatchOption) error {
	return repository.New(s.db).Find(ctx, chs, opts...)
}

func (s *gormSubscriberRepository) First(ctx context.Context, ch *Subscriber, opts ...repository.MatchOption) error {
	return repository.New(s.db).First(ctx, ch, opts...)
}

func (s *gormSubscriberRepository) Delete(ctx context.Context, opts ...repository.MatchOption) error {
	return repository.New(s.db, &Subscriber{}).Delete(ctx, opts...)
}

func (s *gormSubscriberRepository) UpdateFields(ctx context.Context, fields repository.Fields, opts ...repository.MatchOption) error {
	return repository.New(s.db, &Subscriber{}).UpdateFields(ctx, fields, opts...)
}

func (s *gormSubscriberRepository) Update(ctx context.Context, v *Subscriber) error {
	return repository.New(s.db, &Subscriber{}).Update(ctx, v)
}
