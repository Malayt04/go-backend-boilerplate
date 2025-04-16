package repositories

import (
	"context"

	"github.com/go-backend/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) CreateEvent(context context.Context, eventData *models.Event) (*models.Event, error){

	event := &models.Event{
		Name: eventData.Name,
		Location: eventData.Location,
		Date: eventData.Date,
	}

	res := r.db.Model(&models.Event{}).Create(&eventData)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil

}

func (r *EventRepository) GetEvents(context context.Context) ([]*models.Event, error) {

	events := []*models.Event{}

	res := r.db.Model(&models.Event{}).Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil

}

func (r *EventRepository) GetEventById(context context.Context, id uint) (*models.Event, error) {

	event := &models.Event{}

	res := r.db.Model(&models.Event{}).First(&event, id)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil

}

func (r *EventRepository) UpdateEvent(context context.Context, event *models.Event) error {

	res := r.db.Model(&models.Event{}).Where("id = ?", event.ID).Updates(event)

	if res.Error != nil {
		return res.Error
	}

	return nil

}

func (r *EventRepository) DeleteEvent(context context.Context, id uint) error {

	res := r.db.Model(&models.Event{}).Delete(&models.Event{}, id)

	if res.Error != nil {
		return res.Error
	}

	return nil

}
