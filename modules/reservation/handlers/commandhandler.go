package handlers

import (
	"context"
	"cqrs-demo/modules/reservation/model/commands"
	"cqrs-demo/modules/reservation/repository"
)

type commandHandler struct {
	repo repository.CommandRepository
}

func NewCommandHandler(repo repository.CommandRepository) *commandHandler {
	return &commandHandler{
		repo: repo,
	}
}

func (handler *commandHandler) CreateReservation(ctx context.Context,
	create *commands.CreateReservation) {
	//
	// some business logic ?
	//
	err := handler.repo.CreateReservation(ctx, create)
	if err != nil {

	}

}
