package controllers

import (
	"encoding/json"
	"net/http"
	"nex/http"

	"github.com/ogu89/9-go-bookstore/pkg/models"
)

var NewBook models.Book


func GetBook(w http.ResponseWriter, r *http.Request){
	 newBooks:=models.GetAllBooks()
	 res, _ := json.Marshal(newBooks)
	 w.Header().Set("Content-Type", "pkglication/json")
	 w.WriteHeader(http.StatusOK)
	 w.Write(res)
}

