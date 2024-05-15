package dummydbstore

import (
	"context"
	"cqrs-demo/modules/reservation/model/commands"
	"cqrs-demo/modules/reservation/model/reservation"
)

type store struct {
	db *DummyDb
}

func (s *store) CreateReservation(ctx context.Context,
	reservation *commands.CreateReservation) error {
	return s.db.Create(ctx, reservation)
}

func (s *store) DeleteReservation(ctx context.Context,
	reservation *commands.DeleteReservation) error {
	return s.db.Delete(ctx, reservation.Id)
}
func (s *store) UpdateState(ctx context.Context, update *commands.UpdateReservation) error {
	err := s.db.Update(ctx, update)
	if err != nil {
		return err
	}
	return nil
}
func (s *store) UpdateReservation(ctx context.Context,
	reservation *commands.UpdateReservation) error {
	return s.db.Update(ctx, reservation)
}

func (s *store) SelectReservation(ctx context.Context,
	condition map[string]interface{}) (*reservation.Reservation, error) {
	res, err := s.db.Select(ctx, condition)
	if err != nil {
		return nil, err
	}
	return res.(*reservation.Reservation), err
}

func (s *store) ListReservation(ctx context.Context,
	condition map[string]interface{}) ([]reservation.Reservation, error) {
	res, err := s.db.Select(ctx, condition)
	if err != nil {
		return nil, err
	}
	return res.([]reservation.Reservation), err
}
