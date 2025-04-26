package service

import (
	"context"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local"
	"github.com/google/uuid"
)

func (s *localService) GetAllLocalsWithCity(ctx context.Context, request local.QueryParamRequestGetLocals) ([]local.ResponseGetLocalBusinesses, error) {
	localRepository, err := s.repository.NewClient(false)
	if err != nil {
		return []local.ResponseGetLocalBusinesses{}, err
	}

	localBusinesses := new([]local.Locals)

	err = localRepository.GetAllLocalBusinesses(ctx, request, localBusinesses)
	if err != nil {
		return []local.ResponseGetLocalBusinesses{}, err
	}

	types := make([]local.ResponseGetLocalBusinesses, len(*localBusinesses))
	for i, d := range *localBusinesses {
		types[i] = local.ResponseGetLocalBusinesses{
			ID:          d.ID,
			Name:        d.Name,
			Description: d.Description,
			Address:     d.Address,
			City:        d.City,
			Province:    d.Province,
			Longitude:   d.Longitude,
			Latitude:    d.Latitude,
			Label:       d.Label,
			OpenedTime:  d.OpenedTime,
			PhotoUrl:    d.PhotoUrl,
			IsBusiness:  d.IsBusiness,
			CreatedAt:   d.CreatedAt,
		}
	}

	return types, nil
}

func (s *localService) GetSpecificLocalBusiness(ctx context.Context, localBusinessID uuid.UUID) (local.ResponseGetLocalBusinesses, error) {
	localRepository, err := s.repository.NewClient(false)
	if err != nil {
		return local.ResponseGetLocalBusinesses{}, err
	}

	localBusiness := &local.Locals{
		ID: localBusinessID,
	}
	err = localRepository.GetLocalBusinessByID(ctx, localBusiness)
	if err != nil {
		return local.ResponseGetLocalBusinesses{}, err
	}

	err = localRepository.GetReviewsByLocalsID(ctx, localBusinessID.String(), &localBusiness.Reviews)
	if err != nil {
		return local.ResponseGetLocalBusinesses{}, err
	}

	reviews := make([]local.ResponseReviews, len(localBusiness.Reviews))
	for i, d := range localBusiness.Reviews {
		reviews[i] = local.ResponseReviews{
			ID:        d.ID,
			Star:      d.Star,
			Content:   d.Content,
			CreatedAt: d.CreatedAt,
			PhotoURL:  d.PhotoURL,
		}
	}

	return local.ResponseGetLocalBusinesses{
		ID:          localBusiness.ID,
		Name:        localBusiness.Name,
		Description: localBusiness.Description,
		Address:     localBusiness.Address,
		City:        localBusiness.City,
		Province:    localBusiness.Province,
		Longitude:   localBusiness.Longitude,
		Latitude:    localBusiness.Latitude,
		Label:       localBusiness.Label,
		OpenedTime:  localBusiness.OpenedTime,
		PhotoUrl:    localBusiness.PhotoUrl,
		IsBusiness:  localBusiness.IsBusiness,
		CreatedAt:   localBusiness.CreatedAt,
		Reviews:     reviews,
	}, nil
}
