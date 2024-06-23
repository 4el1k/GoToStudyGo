package service

import (
	"awesomeProject/internal/pkg/handler/dto/req"
	repo "awesomeProject/internal/pkg/repository"
	"awesomeProject/internal/pkg/util/hasher"
	"awesomeProject/internal/pkg/util/jwter"
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type JwtService struct {
	catRepo    *repo.CatRepository
	tokensRepo *repo.TokensRepository
	log        *logrus.Logger
}

func NewJwtService(catRepo *repo.CatRepository, tokensRepo *repo.TokensRepository, log *logrus.Logger) *JwtService {
	return &JwtService{
		catRepo:    catRepo,
		tokensRepo: tokensRepo,
		log:        log,
	}
}

var (
	ErrPassMismatch = errors.New("password does not match")
)

func (s *JwtService) Auth(ctx context.Context, in *req.SignIn) (*req.CoupleToken, error) {
	s.log.Println("start auth service")
	c, err := s.catRepo.FindByName(ctx, in.Username)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	if !hasher.CheckPass(c.PasswordHash, in.Password) {
		s.log.Errorf(ErrPassMismatch.Error())
		return nil, ErrPassMismatch
	}
	return s.createCoupleToken(ctx, in.Username)
}

func (s *JwtService) RefreshJwtToken(ctx context.Context,
	currentJwtToken string, refreshToken *uuid.UUID) (string, error) {
	s.log.Printf("start refresh token service")
	username, err := jwter.DecodeToken(currentJwtToken)
	if err != nil {
		s.log.Error("error0 %w", err)
		return "", err
	}
	newJwtToken, _, err := jwter.EncodeToken(username)
	if err != nil {
		s.log.Error("error1 %w", err)
		return "", err
	}
	err = s.tokensRepo.UpdateJwtToken(ctx, newJwtToken, refreshToken)
	if err != nil {
		s.log.Error("error2 %w", err)
		return "", err
	}
	return newJwtToken, nil
}

func (s *JwtService) Logout(ctx context.Context, refreshToken *uuid.UUID) error {
	s.log.Printf("start logout service")
	err := s.tokensRepo.DeleteByRefreshToken(ctx, refreshToken)
	if err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func (s *JwtService) createCoupleToken(ctx context.Context, username string) (*req.CoupleToken, error) {
	token, exp, err := jwter.EncodeToken(username)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.tokensRepo.Create(ctx, token, exp.Add(time.Hour*24*365))
	if err != nil {
		return nil, err
	}
	return &req.CoupleToken{AccessToken: token, RefreshToken: refreshToken.String()}, nil
}
