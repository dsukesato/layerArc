package usecase

import (
	"github.com/dsukesato/layerArc/domain/model"
	"github.com/dsukesato/layerArc/domain/repository"
	"github.com/dsukesato/layerArc/infrastructure/persistence"
)

type BookUsecase struct {}

func (book BookUsecase) GetAll() ([]*model.Book, error) {
	var books []*model.Book
	var err error

	books, err = repository.BookRepository(persistence.BookPersistence{}).GetAll()

	if err != nil {
		return nil, err
	}

	return books, nil
}
