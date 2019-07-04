package repsitory

import (
	"context"

	"PatientData/restAPI/models"
)

type PostRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.Post, error)
}
