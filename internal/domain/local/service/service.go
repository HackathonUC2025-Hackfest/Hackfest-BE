package service

import (
	"context"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local"
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local/repository"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type localService struct {
	repository repository.RepositoryItf
	snap       snap.Client
	coreapi    coreapi.Client
}

type LocalServiceItf interface {
	GetAllLocalsWithCity(ctx context.Context, request local.QueryParamRequestGetLocals) ([]local.ResponseGetLocalBusinesses, error)
	GetSpecificLocalBusiness(ctx context.Context, localBusinessID uuid.UUID) (local.ResponseGetLocalBusinesses, error)
	GetTourGuides(ctx context.Context, city string) ([]local.ReseponseGetTourGuide, error)
	GetSpecificTourGuide(ctx context.Context, touristAttractionsID uuid.UUID) (local.ReseponseGetTourGuide, error)
	GenerateSnapPayment(ctx context.Context, request local.RequestGenerateSnapLink) (local.ResponseGenerateSnapLink, error)
	GetFullBook(ctx context.Context, taID string) ([]string, error)
}

func New(repository repository.RepositoryItf, snap snap.Client, coreapi coreapi.Client) LocalServiceItf {

	return &localService{
		repository: repository,
		snap:       snap,
		coreapi:    coreapi,
	}
}
