package repository

import (
	"cinema/model"
)

type TicketRepo interface {
	SaveTicket(ticket model.Ticket) error
	FindTicket(userId string) ([]model.Ticket, error)
}
