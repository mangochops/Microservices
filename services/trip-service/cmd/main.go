package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()

	repo := repository.NewInMemoryRepository()
	tripService := service.NewTripService(repo)

	// Use tripService for handling trip-related operations
	fare := &domain.RideFareModel{
		UserID:        "user123",
		PackageSlug:   "standard",
		TotalPriceInCents: 1000,
	}
	trip, err := tripService.CreateTrip(ctx, fare)

	if err != nil {
		log.Fatalf("Failed to create trip: %v", err)
	}
	log.Printf("Trip created successfully: %+v", trip)
	log.Printf("Ride Fare: %+v", trip.RideFare)

	for {
		time.Sleep(5 * time.Second)
	}
}
