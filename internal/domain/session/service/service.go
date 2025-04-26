package service

import (
	"context"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/session"
	sessionRepository "github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/session/repository"
	userRepository "github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/user/repository"
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/pkg/jwt"
)

type authService struct {
	repository        userRepository.RepositoryItf
	sessionRepository sessionRepository.RepositoryItf
	jwt               *jwt.JWTStruct
}

type AuthServiceItf interface {
	Register(ctx context.Context, request session.RegisterRequest) (session.LoginResponse, error)
	Login(ctx context.Context, request session.LoginRequest) (session.LoginResponse, error)
}

func New(repository userRepository.RepositoryItf, sessionRepository sessionRepository.RepositoryItf, jwt *jwt.JWTStruct) AuthServiceItf {

	return &authService{
		repository:        repository,
		sessionRepository: sessionRepository,
		jwt:               jwt,
	}
}
