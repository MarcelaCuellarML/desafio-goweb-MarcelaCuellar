package tickets

import "github.com/MarcelaCuellarML/desafio-goweb-MarcelaCuellar/internal/domain"

type serviceTickets struct {
	repo TicketRepository
}

func NewService(r TicketRepository) TicketsService {
	return &serviceTickets{
		repo: r,
	}
}

type TicketsService interface {
	GetAll() ([]domain.Ticket, error)
	GetCountTicketByDestination(destination string) (int, error)
	GetAverageByDate(destination, date string) (int, error)
}

func (s *serviceTickets) GetAll() ([]domain.Ticket, error) {
	ps, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *serviceTickets) GetCountTicketByDestination(destination string) (int, error) {
	ps, err := s.repo.GetCountTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return ps, nil
}

func (s *serviceTickets) GetAverageByDate(destination, date string) (int, error) {
	ps, err := s.repo.GetAverageByDate(destination, date)
	if err != nil {
		return 0, err
	}
	return ps, nil
}
