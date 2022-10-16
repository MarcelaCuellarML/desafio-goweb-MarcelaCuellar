package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MarcelaCuellarML/desafio-goweb-MarcelaCuellar/internal/tickets"
)

type RequestPatch struct {
	destination int `json:"destino" binding:"required"`
	hourIn      int `json:"horaDesde" binding:"required"`
	hourOut     int `json:"horaHasta" binding:"required"`
}

type Service struct {
	service tickets.TicketsService
}

func NewTickets(s tickets.TicketsService) *Service {
	return &Service{
		service: s,
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetCountTicketByDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("destino")
		date := c.Param("fecha")

		avg, err := s.service.GetAverageByDate(destination, date)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
