package db

import (
	"github.com/ilyas-muhammad/bookstore_oauth_api/src/domain/access_token"
	"github.com/ilyas-muhammad/bookstore_oauth_api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.InternalServerError("database connection not implemented yet")
	return nil, nil
}
