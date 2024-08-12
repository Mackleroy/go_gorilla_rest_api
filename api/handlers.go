package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, marshalErr := json.Marshal(ProjectStorageInstance.categories)
	if marshalErr != nil {
		http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
		return
	}
	_, writeErr := w.Write(response)
	if writeErr != nil {
		http.Error(w, writeErr.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	body, readerErr := io.ReadAll(r.Body)
	if readerErr != nil {
		http.Error(w, readerErr.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var jsonBody = CreateCategoryRequest{}
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := ProjectStorageInstance.AddCategory(jsonBody.CategoryName); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func CreateGoodHandler(w http.ResponseWriter, r *http.Request) {
	body, readerErr := io.ReadAll(r.Body)
	if readerErr != nil {
		http.Error(w, readerErr.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var jsonBody = CreateGoodRequest{}
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := jsonBody.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := ProjectStorageInstance.AddGood(jsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
