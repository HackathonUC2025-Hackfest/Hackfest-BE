package service

import (
	"context"
	"fmt"
	"time"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (s *localService) GetTourGuides(ctx context.Context, city string) ([]local.ReseponseGetTourGuide, error) {
	localRepository, err := s.repository.NewClient(false)
	if err != nil {
		return []local.ReseponseGetTourGuide{}, err
	}

	touristAttractions := new([]local.TouristAttractions)

	err = localRepository.GetAllTourGuides(ctx, city, touristAttractions)
	if err != nil {
		return []local.ReseponseGetTourGuide{}, err
	}

	tourGuides := make([]local.ReseponseGetTourGuide, len(*touristAttractions))
	for i, d := range *touristAttractions {
		tourGuides[i] = local.ReseponseGetTourGuide{
			ID:                          d.ID,
			Name:                        d.Name,
			Description:                 d.Description,
			Address:                     d.Address,
			City:                        d.City,
			Province:                    d.Province,
			Longitude:                   d.Longitude,
			Latitude:                    d.Latitude,
			PhotoUrl:                    d.PhotoURL,
			TourGuidePrice:              d.TourGuidePrice,
			TourGuideCount:              d.TourGuideCount,
			TourGuideDiscountPercentage: d.TourGuideDiscountPercentage,
			Price:                       d.Price,
			DiscountPercentage:          d.DiscountPercentage,
			CreatedAt:                   d.CreatedAt,
		}
	}

	return tourGuides, nil
}

func (s *localService) GetSpecificTourGuide(ctx context.Context, touristAttractionsID uuid.UUID) (local.ReseponseGetTourGuide, error) {
	taRepository, err := s.repository.NewClient(false)
	if err != nil {
		return local.ReseponseGetTourGuide{}, err
	}

	touristAttractions := &local.TouristAttractions{
		ID: touristAttractionsID,
	}
	err = taRepository.GetTAByID(ctx, touristAttractions)
	if err != nil {
		return local.ReseponseGetTourGuide{}, err
	}

	err = taRepository.GetBookingsByTAID(ctx, touristAttractionsID.String(), &touristAttractions.Bookings)
	if err != nil {
		return local.ReseponseGetTourGuide{}, err
	}

	reviews := make([]local.ResponseReviews, len(touristAttractions.Bookings))
	for i, d := range touristAttractions.Bookings {
		reviews[i] = local.ResponseReviews{
			ID:        d.ID,
			Star:      d.Star,
			Content:   d.Content,
			CreatedAt: d.CreatedAt,
			PhotoURL:  d.PhotoURL,
		}
	}

	return local.ReseponseGetTourGuide{
		ID:                          touristAttractions.ID,
		Name:                        touristAttractions.Name,
		Description:                 touristAttractions.Description,
		Address:                     touristAttractions.Address,
		City:                        touristAttractions.City,
		Province:                    touristAttractions.Province,
		Longitude:                   touristAttractions.Longitude,
		Latitude:                    touristAttractions.Latitude,
		PhotoUrl:                    touristAttractions.PhotoURL,
		TourGuidePrice:              touristAttractions.TourGuidePrice,
		TourGuideCount:              touristAttractions.TourGuideCount,
		TourGuideDiscountPercentage: touristAttractions.TourGuideDiscountPercentage,
		Price:                       touristAttractions.Price,
		DiscountPercentage:          touristAttractions.DiscountPercentage,
		CreatedAt:                   touristAttractions.CreatedAt,
		Reviews:                     reviews,
	}, nil
}

func (s *localService) GetFullBook(ctx context.Context, taID string) ([]string, error) {
	taRepository, err := s.repository.NewClient(false)
	if err != nil {
		return []string{}, err
	}

	days := new([]string)
	err = taRepository.GetFullBook(ctx, taID, days)
	if err != nil {
		return []string{}, err
	}

	return *days, nil
}

func (s *localService) GenerateSnapPayment(ctx context.Context, request local.RequestGenerateSnapLink) (local.ResponseGenerateSnapLink, error) {
	taRepository, err := s.repository.NewClient(false)
	if err != nil {
		return local.ResponseGenerateSnapLink{}, err
	}

	taIDUUID, err := uuid.Parse(request.TAID)
	if err != nil {
		return local.ResponseGenerateSnapLink{}, err
	}

	touristAttractions := &local.TouristAttractions{
		ID: taIDUUID,
	}
	err = taRepository.GetTAByID(ctx, touristAttractions)
	if err != nil {
		return local.ResponseGenerateSnapLink{}, err
	}

	transactionID, err := uuid.NewV7()
	if err != nil {
		return local.ResponseGenerateSnapLink{}, err
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionID.String(),
			GrossAmt: touristAttractions.TourGuidePrice,
		},
		Callbacks:       nil,
		EnabledPayments: snap.AllSnapPaymentType,
		Expiry: &snap.ExpiryDetails{
			Duration: 5,
			Unit:     "hours",
		},
	}

	snapR, aw := s.snap.CreateTransactionToken(req)
	if aw != nil {
		return local.ResponseGenerateSnapLink{}, err
	}

	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, request.BookedAt)
	if err != nil {
		return local.ResponseGenerateSnapLink{}, err
	}

	parsedTime = time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 0, 0, 0, 0, time.UTC)

	userIDUUID, err := uuid.Parse(request.UserID)
	if err != nil {
		return local.ResponseGenerateSnapLink{}, err
	}

	booking := local.TourGuideBookings{
		ID:                   transactionID,
		PaymentURL:           fmt.Sprintf("https://app.sandbox.midtrans.com/snap/v4/redirection/%s", snapR),
		BookedAt:             parsedTime,
		Status:               "unpaid",
		UserID:               userIDUUID,
		TouristAttractionsID: touristAttractions.ID,
	}

	err = taRepository.CreateBooking(ctx, booking)
	if err != nil {
		return local.ResponseGenerateSnapLink{}, err
	}

	return local.ResponseGenerateSnapLink{
		TAID:       touristAttractions.ID.String(),
		PaymentUrl: snapR,
	}, nil
}
