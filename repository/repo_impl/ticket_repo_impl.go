package repo_impl

import (
	"cinema/model"
	"cinema/repository"

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
