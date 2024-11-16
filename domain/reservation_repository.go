package domain

import "context"

type ReservationRepository interface {
	Add(ctx context.Context, reservation *Reservation) error
	FindByUserID(ctx context.Context, userID string) ([]Reservation, error)
	RemoveByID(ctx context.Context, id string) error
}
