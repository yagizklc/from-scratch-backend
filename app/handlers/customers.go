package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/yagizklc/from-scratch-server/app/internal/customers"
)

type CustomersHandler struct {
	repo *customers.Repository
}

func (h CustomersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.GetCustomerByEmail(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func NewCustomersHandler(repo *customers.Repository) *CustomersHandler {
	return &CustomersHandler{repo: repo}
}
func NewDefaultCustomersHandler() *CustomersHandler {
	connStr := "host=localhost port=5432 user=myuser password=mypassword dbname=myapp sslmode=disable"
	repo, err := customers.NewRepository(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	return &CustomersHandler{repo: repo}
}

func (h CustomersHandler) GetCustomerByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	customer, err := h.repo.GetCustomerByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(customer)
}
