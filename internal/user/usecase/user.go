package usecase

import (
	"context"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/entity"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/repository"
	"github.com/mert-yigittop/cxp-api-starter/pkg/jwt"
	"net/http"
	"time"
)

type Usecase interface {
	Register(ctx context.Context, payload dto.RegisterRequest) (dto.RegisterResponse, int, error)
	Login(ctx context.Context, payload dto.LoginRequest) (dto.LoginResponse, int, error)
}

type usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Register(ctx context.Context, payload dto.RegisterRequest) (dto.RegisterResponse, int, error) {
	user := entity.New(payload)
	response, err := uc.repo.Register(ctx, user)
	if err != nil {
		return dto.RegisterResponse{}, http.StatusBadRequest, err
	}

	return response, http.StatusOK, nil
}

func (uc *usecase) Login(ctx context.Context, payload dto.LoginRequest) (dto.LoginResponse, int, error) {
	userId, err := uc.repo.Login(ctx, payload)
	if err != nil {
		return dto.LoginResponse{}, http.StatusBadRequest, err
	}

	accessToken, err := jwt.Sign(userId, time.Hour*8)
	if err != nil {
		return dto.LoginResponse{}, http.StatusUnauthorized, err
	}

	return dto.LoginResponse{
		ID:          userId,
		AccessToken: accessToken,
	}, http.StatusOK, nil
}
