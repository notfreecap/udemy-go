package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/notfreecap/bookstore/pkg/models"
	"github.com/notfreecap/bookstore/pkg/utils"
)

func GetAllBooks(w http.ResponseWriter, _ *http.Request) {
	b := models.GetAllBooks()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "Application/json")

	if bookDetails.ID != 0 {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "book not found"})
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	b := &models.Book{}
	utils.ParseBody(r, b)
	res, _ := json.Marshal(b.CreateBook())
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		er := models.ErrorResponse{
			Error: fmt.Sprintf(
				"Error while parsing the id: %v", mux.Vars(r)["id"],
			),
		}
		er.ReturnError(w)
		return
	}

	book := &models.Book{}
	utils.ParseBody(r, book)

	dbBook, db := models.GetBookById(id)

	dbBook.Name = book.Author
	dbBook.Author = book.Author
	dbBook.Publication = book.Publication

	res, _ := json.Marshal(db.Update(dbBook).Value)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		er := models.ErrorResponse{
			Error: fmt.Sprintf(
				"Error while parsing the id: %v", mux.Vars(r)["id"],
			),
		}
		er.ReturnError(w)
		return
	}
	res, _ := json.Marshal(models.DeleteBook(id))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
