package pvpc

import (
	"context"
)

type PricesZonesRepository interface {
	// GetByID returns the prices zone with the given ID.
	GetByID(ctx context.Context, id string) (int, error)
}
