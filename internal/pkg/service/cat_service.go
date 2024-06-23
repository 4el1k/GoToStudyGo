package service

import (
	"awesomeProject/internal/pkg/handler/dto/req"
	"awesomeProject/internal/pkg/model"
	repo "awesomeProject/internal/pkg/repository"
	"awesomeProject/internal/pkg/util/hasher"
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"log"
)

type CatService struct {
	repo *repo.CatRepository
	log  *logrus.Logger
}

func NewCatService(repo *repo.CatRepository, log *logrus.Logger) *CatService {
	return &CatService{
		repo: repo,
		log:  log,
	}
}

func (srv *CatService) Create(ctx context.Context, c *req.CatSave) error {
	srv.log.Println("start cat service, func create")
	passhash := hasher.HashPass(c.Password)
	println("passhash: ", passhash)
	err := srv.repo.Save(ctx, &model.Cat{
		Name:         c.Name,
		Age:          c.Age,
		PasswordHash: passhash,
	})
	if err != nil {
		srv.log.Printf("error in cat service, func create %w", err)
		return err
	}
	return nil
}

func (srv *CatService) FindById(ctx context.Context, id *uuid.UUID) (*model.Cat, error) {
	srv.log.Println("start cat service, findById")
	cat, err := srv.repo.FindById(ctx, id)
	if err != nil {
		srv.log.Printf("error in cat service, findById %w", err)
		return &model.Cat{}, err
	}
	return cat, nil
}

func (srv *CatService) Update(ctx context.Context, u *req.CatUpdate) error {
	srv.log.Println("start cat service, func update")
	err := srv.repo.Update(ctx, &model.Cat{
		Name: u.Name,
		Age:  u.Age,
		Id:   u.Id,
	})
	if err != nil {
		log.Printf("error in cat service %w", err)
		return err
	}
	return nil
}

func (srv *CatService) DeleteById(ctx context.Context, id *uuid.UUID) error {
	srv.log.Println("start cat service, func delete")
	err := srv.repo.Delete(ctx, id)
	if err != nil {
		srv.log.Printf("error in cat service %w", err)
		return err
	}
	return nil
}
