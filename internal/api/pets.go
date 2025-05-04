package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ahawi/diploma/internal/entity"
	"github.com/ahawi/diploma/internal/repository"
)

func (a *API) GetPetsBatchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(r.Context(), a.timeout)
	defer cancel()

	lastID, err := strconv.ParseInt(r.URL.Query().Get("last_id"), 10, 64)
	if err != nil {
		log.Printf("failed to parse lastID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if err != nil {
		log.Printf("failed to parse offset: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pets, err := a.petService.PetsBatch(ctx, lastID, offset)
	if err != nil {
		log.Printf("petService.PetsBatch: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(pets)
	if err != nil {
		log.Printf("failed to write response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
}

func (a *API) GetRecommendationPetsBatchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(r.Context(), a.timeout)
	defer cancel()

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		log.Printf("failed to parse page: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		log.Printf("failed to parse pageSize: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		log.Printf("failed to parse userID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pets, err := a.recommendationService.GetRecommendations(ctx, userID, page, pageSize)
	if err != nil && errors.Is(err, entity.ErrBadRequest) {
		log.Printf("recommendationService.GetRecommendations: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Printf("recommendationService.GetRecommendations: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(pets)
	if err != nil {
		log.Printf("failed to write response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) GetPetByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(r.Context(), a.timeout)
	defer cancel()

	petID, err := strconv.ParseInt(chi.URLParam(r, "petID"), 10, 64)
	if err != nil {
		log.Printf("failed to parse petID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pet, err := a.petService.PetByID(ctx, petID)
	if err != nil {
		log.Printf("petService.PetByID: %v", err)
	}
	if err != nil && errors.Is(err, entity.ErrBadRequest) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(pet)
	if err != nil {
		log.Printf("failed to write response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
