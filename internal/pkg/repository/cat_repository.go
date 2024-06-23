package repo

import (
	"awesomeProject/internal/pkg/model"
	"awesomeProject/internal/pkg/util/hasher"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

const (
	createCat  = "INSERT INTO cat (name, age, passwordhash) VALUES ($1, $2, $3)"
	findById   = "SELECT name, age FROM cat WHERE id = $1"
	findByName = "SELECT id, age, passwordhash FROM cat WHERE name = $1"
	updateCat  = "UPDATE cat SET name = $1, age = $2 WHERE id = $3"
	delete     = "DELETE FROM cat WHERE id = $1"
)

type CatRepository struct {
	db  pgxtype.Querier
	log *logrus.Logger
}

// в zuzu возращалась ссылка *
func NewCatRepository(db pgxtype.Querier, log *logrus.Logger) *CatRepository {
	return &CatRepository{
		db:  db,
		log: log,
	}
}

func (r *CatRepository) Save(ctx context.Context, c *model.Cat) error {
	r.log.Println("Start save cat")
	_, err := r.db.Exec(ctx, createCat, c.Name, c.Age, c.PasswordHash)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec, func save: %w", err)
		return err
	}
	return nil
}

func (r *CatRepository) FindById(ctx context.Context, id *uuid.UUID) (*model.Cat, error) {
	r.log.Println("start findById cat repo")
	row := r.db.QueryRow(ctx, findById, id)
	result := &model.Cat{Id: *id}
	err := row.Scan(&result.Name, &result.Age, &result.PasswordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("error happened in db.QueryRow, func findById: %w", err)
			r.log.Errorf("error in cat repo", err.Error())
			return &model.Cat{}, err
		}
	}
	r.log.Println("end findById cat repo, result: {id: %s, name: %s, age: %s}",
		result.Id, result.Name, result.Age)
	return result, nil
}

func (r *CatRepository) FindByName(ctx context.Context, name string) (*model.Cat, error) {
	r.log.Println("start findByName cat repo")
	row := r.db.QueryRow(ctx, findByName, name)
	result := &model.Cat{Name: name}
	err := row.Scan(&result.Id, &result.Age, &result.PasswordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("error happened in db.QueryRow, func findById: %w", err)
			r.log.Errorf("error in cat repo", err.Error())
			return &model.Cat{}, err
		}
	}
	r.log.Println("end findById cat repo, result: {id: %s, name: %s, age: %s}",
		result.Id, result.Name, result.Age)
	fmt.Printf("from repo: ")
	println(&result.PasswordHash)
	println()
	fmt.Printf("hash of \"123\" : %s\n")
	println(hasher.HashPass("123"))
	println()
	return result, nil
}

func (r *CatRepository) Update(ctx context.Context, cat *model.Cat) error {
	r.log.Println("start update cat repo")
	_, err := r.db.Exec(ctx, updateCat, cat.Name, cat.Age, cat.Id)
	if err != nil {
		r.log.Errorf("error happened in db.Exec, func update: %w", err)
		return err
	}
	return nil
}

func (r *CatRepository) Delete(ctx context.Context, id *uuid.UUID) error {
	r.log.Println("start delete cat repo")
	_, err := r.db.Exec(ctx, delete, id)
	if err != nil {
		r.log.Errorf("error happened in db.Exec, func delete: %w", err)
		return err
	}
	return nil
}
