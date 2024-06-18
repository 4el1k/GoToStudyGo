package service

import (
	"awesomeProject/internal/pkg/handler/dto/req"
	"awesomeProject/internal/pkg/model"
	repo "awesomeProject/internal/pkg/repository"
	"context"
	"fmt"
)

type CatService struct {
	repo repo.CatRepository
}

func NewCatService(repo repo.CatRepository) *CatService {
	return &CatService{
		repo: repo,
	}
}

func (srv *CatService) Create(ctx context.Context, c req.CatReq) error {
	err := srv.repo.Save(ctx, model.Cat{
		Name: c.Name,
		Age:  c.Age,
	})
	if err != nil {
		fmt.Printf("error in cat service %w", err)
		return err
	}
	return nil
}
