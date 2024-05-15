package handlers

import (
	"context"
	"cqrs-demo/modules/reservation/repository"
	"log"
)

type readHandler struct {
	repo repository.QueryRepository
}

func NewReadHandler(repo repository.CommandRepository) *commandHandler {
	return &commandHandler{
		repo: repo,
	}
}

func (handler *readHandler) FindReservationById(ctx context.Context, id int) {
	// some business logic

	res, err := handler.repo.SelectReservation(ctx, map[string]interface{}{"id": id})
	if err != nil {
		// Tra ket qua ve tro client
	}
	log.Println(res)
	// Tra ket qua ve tro client
}
