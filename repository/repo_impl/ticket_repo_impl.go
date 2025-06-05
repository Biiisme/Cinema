package repo_impl

import (
	"cinema/model"
	"cinema/repository"
	"log"

	"gorm.io/gorm"
)

type TicketRepoImpl struct {
	db *gorm.DB
}

func NewTicketRepo(db *gorm.DB) repository.TicketRepo {
	return &TicketRepoImpl{db: db}
}

func (r *TicketRepoImpl) SaveTicket(ticket model.Ticket) error {
	return r.db.Create(&ticket).Error
}

func (r *TicketRepoImpl) FindTicket(userId string) ([]model.Ticket, error) {
	var tickets []model.Ticket
	// Get all films
	if err := r.db.Where("user_id=?", userId).Find(&tickets).Error; err != nil {
		log.Println("Error find ticket:", err)
		return nil, err
	}

	return tickets, nil
}
