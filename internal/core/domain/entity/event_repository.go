package entity

import "context"

type EventRepository interface {
	Get(ctx context.Context, eventID string) (*Event, error)
}
