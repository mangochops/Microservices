package main

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {
	ctx := context.Background()

	repo := repository.NewInMemoryRepository()
	tripService := service.NewTripService(repo)

	// Use tripService for handling trip-related operations
	tripService.CreateTrip(ctx, &domain.RideFareModel{
		UserID:         "user123",
		PackageSlug:    "standard",
		TotalPriceInCents: 1000,
	})

	
}
