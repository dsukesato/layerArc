package repository

import (
	"github.com/dsukesato/layerArc/domain/model"
)

type BookRepository interface {
	GetAll() ([]*model.Book, error)
}
