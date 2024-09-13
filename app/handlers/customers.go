package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/yagizklc/from-scratch-server/app/internal/customers"
	"github.com/yagizklc/from-scratch-server/app/pkg"
)

type PokemonHandler struct {
	mux  *http.ServeMux
	repo *customers.Repository
}

func NewPokemonHandler(config *pkg.Config) *PokemonHandler {

	repo, err := customers.NewRepository(context.Background(), config.GetConnectionString())
	if err != nil {
		panic(err)
	}
	return &PokemonHandler{repo: repo, mux: http.NewServeMux()}
}

func (h *PokemonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.HandleFunc("GET /pokemon", h.GetPokemonByName)
	h.mux.ServeHTTP(w, r)
}

func (h *PokemonHandler) GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	customer, err := h.repo.GetCustomerByEmail(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(customer)
}
