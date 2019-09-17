package handler

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/dsukesato/layerArc/domain/model"
	"github.com/dsukesato/layerArc/usecase"
)

type BookHandler struct {}

func (book BookHandler)BookIndex(w http.ResponseWriter, r *http.Request) {
	var books []*model.Book
	var err error

	books, err = usecase.BookUsecase{}.GetAll()

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	fmt.Println(http.StatusOK, r.URL.Path)
}

