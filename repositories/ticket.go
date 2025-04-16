package repositories

import (
	"context"

	"github.com/go-backend/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) CreateTicket(context context.Context, ticketData *models.Ticket) (*models.Ticket, error) {

	ticket := &models.Ticket{
		EventID: ticketData.EventID,
		UserID: ticketData.UserID,
		Entered: ticketData.Entered,
	}

	res := r.db.Model(&models.Ticket{}).Create(&ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil

}

func (r *TicketRepository) GetTickets(context context.Context) ([]*models.Ticket, error) {

	tickets := []*models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil

}

func (r *TicketRepository) GetTicketById(context context.Context, eventId uint, userId uint) (*models.Ticket, error) {

	ticket := &models.Ticket{}
	
	res := r.db.Model(&models.Ticket{}).Where("event_id = ? AND user_id = ?", eventId, userId).First(&ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil

}

func (r *TicketRepository) GetTicketByUserId(context context.Context, id uint) (*models.Ticket, error) {

	ticket := &models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Where("user_id = ?", id).First(&ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil

}

func (r *TicketRepository) GetTicketByEventId(context context.Context, id uint) (*models.Ticket, error) {

	ticket := &models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Where("event_id = ?", id).First(&ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil

}

func (r *TicketRepository) UpdateTicket(context context.Context, ticket *models.Ticket) error {

	res := r.db.Model(&models.Ticket{}).Where("id = ?", ticket.ID).Updates(ticket)

	if res.Error != nil {
		return res.Error
	}

	return nil

}


