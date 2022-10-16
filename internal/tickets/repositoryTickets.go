package tickets

import (
	"github.com/MarcelaCuellarML/desafio-goweb-MarcelaCuellar/internal/domain"
	"github.com/MarcelaCuellarML/desafio-goweb-MarcelaCuellar/pkg/store"
)

type TicketRepository interface {
	GetAll() ([]domain.Ticket, error)
	GetCountTicketByDestination(destination string) (int, error)
	GetAverageByDate(destination, date string) (int, error)
}

var tickets []domain.Ticket

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) TicketRepository {
	return &repository{
		db: db,
	}
}

func (rp *repository) GetAll() ([]domain.Ticket, error) {

	err := rp.db.Read(&tickets)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (rp *repository) GetCountTicketByDestination(destination string) (int, error) {
	err := rp.db.Read(&tickets)
	if err != nil {
		return 0, err
	}
	var count int
	for _, ticketFilt := range tickets {
		if ticketFilt.Country == destination {
			count = count + 1
		}
	}

	return count, nil
}

func (rp *repository) GetAverageByDate(destination, date string) (int, error) {
	err := rp.db.Read(&tickets)
	if err != nil {
		return 0, err
	}
	var count int
	var countDate int
	for _, ticketFilt := range tickets {
		if ticketFilt.Country == date {
			countDate = countDate + 1
			if ticketFilt.Time == destination {
				count = count + 1
			}
		}
	}
	average := countDate / count
	return average, nil
}
