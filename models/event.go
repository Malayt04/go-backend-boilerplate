package models

import (
	"context"
	"time"
)

type Event struct {
	ID        uint      `json:"id", gorm:"primarykey"`
	Name      string    `json:"name", gorm:"text;not null"`
	Location  string    `json:"location", gorm:"text;not null"`
	Date      time.Time `json:"date", gorm:"not null"`
	CreatedAt time.Time `json:"created_at", gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at", gorm:"autoUpdateTime"`
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, id uint) (*Event, error)
	CreateOne(ctx context.Context, event *Event) error
	UpdateEvent(ctx context.Context, event *Event) error
	DeleteEvent(ctx context.Context, id uint) error
}