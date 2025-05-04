package api

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ahawi/diploma/internal/entity"
)

func (a *API) PostFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		log.Printf("failed to parse userID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read body from request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var object entity.Form
	err = json.Unmarshal(body, &object)
	if err != nil {
		log.Printf("failed to parse form: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), a.timeout)
	defer cancel()

	err = a.userService.AddForm(ctx, userID, &object)
	if err != nil {
		log.Printf("userService.AddForm: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) PostAddInteractionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var object entity.Interaction
	err = json.Unmarshal(body, &object)
	if err != nil {
		log.Printf("failed to parse interaction: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), a.timeout)
	defer cancel()

	err = a.userService.AddInteraction(ctx, &object)
	if err != nil {
		log.Printf("userService.AddInteraction: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) GetProfileHandler(w http.ResponseWriter, r *http.Request) {

}

func (a *API) PostRegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var object struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal(body, &object)
	if err != nil {
		log.Printf("failed to parse name: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), a.timeout)
	defer cancel()

	err = a.userService.RegisterUser(ctx, object.Name)
	if err != nil {
		log.Printf("userService.RegisterUser: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
