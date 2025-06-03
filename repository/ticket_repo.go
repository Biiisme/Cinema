package repository

import (
	"cinema/model"
)

type TicketRepo interface {
	SaveTicket(ticket model.Ticket) error
}
