package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ride-sharing/services/trip-service/internal/domain"
)

type TripService struct {
	repo domain.TripRepository
}

func NewTripService(repo domain.TripRepository) *TripService {
	return &TripService{
		repo: repo,
	}
}

func (s *TripService) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	trip := &domain.TripModel{
		ID:        primitive.NewObjectID(),
		UserID:    fare.UserID,
		Status:    "created",
		RideFare:  fare,
	}

	return s.repo.CreateTrip(ctx, trip)
}