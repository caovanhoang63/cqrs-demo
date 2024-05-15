package repository

import (
	"context"
	"cqrs-demo/modules/reservation/model/commands"
	"cqrs-demo/modules/reservation/model/reservation"
)

type CommandRepository interface {
	CreateReservation(ctx context.Context,
		reservation *commands.CreateReservation) error
	DeleteReservation(ctx context.Context,
		reservation *commands.DeleteReservation) error
	UpdateReservation(ctx context.Context,
		reservation *commands.UpdateReservation) error
	UpdateState(ctx context.Context, update *commands.UpdateReservation) error
}

type QueryRepository interface {
	ListReservation(ctx context.Context,
		condition map[string]interface{}) ([]reservation.Reservation, error)
	SelectReservation(ctx context.Context,
		condition map[string]interface{}) (*reservation.Reservation, error)
}
