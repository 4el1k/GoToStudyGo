package repo

import (
	"context"
	"github.com/jackc/pgtype/pgxtype"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type TokensRepository struct {
	db  pgxtype.Querier
	log *logrus.Logger
}

func NewTokensRepository(db pgxtype.Querier, log *logrus.Logger) *TokensRepository {
	return &TokensRepository{
		db:  db,
		log: log,
	}
}

const (
	createTokens         = "INSERT INTO tokens (refresh_token, jwt_token, exp_date) VALUES ($1, $2, $3)"
	updateJwtToken       = "UPDATE tokens SET jwt_token = $1 WHERE refresh_token = $2"
	deleteByRefreshToken = "DELETE FROM tokens WHERE refresh_token = $1"
)

func (r *TokensRepository) Create(ctx context.Context, jwtToken string, createDate time.Time) (*uuid.UUID, error) {
	r.log.Printf("start create tokens")
	refreshToken := uuid.NewV4()
	_, err := r.db.Exec(ctx, createTokens, refreshToken, jwtToken, createDate)
	if err != nil {
		r.log.Printf("error creating tokens: %v", err)
		return nil, err
	}
	return &refreshToken, nil
}

func (r *TokensRepository) UpdateJwtToken(ctx context.Context, jwtToken string, refreshToken *uuid.UUID) error {
	r.log.Printf("start update tokens")
	_, err := r.db.Exec(ctx, updateJwtToken, jwtToken, refreshToken)
	if err != nil {
		r.log.Printf("error updating tokens: %v", err)
		return err
	}
	return nil
}

func (r *TokensRepository) DeleteByRefreshToken(ctx context.Context, refreshToken *uuid.UUID) error {
	r.log.Printf("start delete tokens")
	_, err := r.db.Exec(ctx, deleteByRefreshToken, refreshToken)
	if err != nil {
		r.log.Printf("error deleting tokens: %v", err)
		return err
	}
	return nil
}
