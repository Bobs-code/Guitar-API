package repository

import (
	"context"

	"github.com/Bobs-code/Guitar-API/models"
)

type GuitarRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.Guitar, error)
}
