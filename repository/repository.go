package repsitory

import (
	"context"

	"github.com/jinwo-o/PatientData/models"
)

type PostRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.Post, error)
}
