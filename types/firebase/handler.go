package customer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CustomerService interface defines the methods needed for customer operations
type CustomerService interface {
    CreateCustomer(customer map[string]interface{}) (CustomerIds, error)
}

// CustomerAPIClient implementation of the CustomerService interface
type CustomerAPIClient struct{}

// CreateCustomerHandler handles customer creation operations
type CreateCustomerHandler struct {
    client CustomerService
}

// CustomerHandler handles general customer operations with Firebase
type CustomerHandler struct {
    client FirebaseAPIClient
}

// CustomerIds holds IDs for a customer in different systems
type CustomerIds struct {
	FirebaseId string `json:"firebase_id"`
	StripeId string `json:"stripe_customer_id"`
}

// NewCreateCustomerHandler creates a new handler for customer creation
func NewCreateCustomerHandler(client CustomerService) *CreateCustomerHandler {
    return &CreateCustomerHandler{client: client}
}

// NewCustomerHandler creates a new handler for general customer operations
func NewCustomerHandler(client FirebaseAPIClient) *CustomerHandler {
    return &CustomerHandler{client: client}
}

// HandleGetCustomer handles HTTP requests to get customer information
func (h *CustomerHandler) HandleGetCustomer(w http.ResponseWriter, r *http.Request) {
    email := r.URL.Query().Get("email")
    if email == "" {
        http.Error(w, "Missing email parameter", http.StatusBadRequest)
        return
    }
    customer, err := h.client.GetCustomer(r.Context(), email)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get customer: %v", err), http.StatusInternalServerError)
        return
    }
    if customer == nil {
        http.Error(w, "Customer not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(customer)
    return
}

// HandleCreateCustomer handles HTTP requests to create a new customer
func (h *CustomerHandler) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {
    var customer map[string]interface{}
    
    err := json.NewDecoder(r.Body).Decode(&customer)
    
    fmt.Printf("Customer Decoded: %v", customer)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    firebaseId, err := h.CreateCustomer(customer)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to create customer: %v", err), http.StatusInternalServerError)
        return
    }

    response := map[string]string{
        "firebase_id": firebaseId,
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)    
}

// CreateCustomer handles the logic for creating a customer
func (h *CustomerHandler) CreateCustomer(customer map[string]interface{}) (string, error) {
    ctx := context.Background()

    fmt.Printf("Creating customer in Firebase %v", customer)

    firebaseId, err := h.client.AddCustomerToDatabase(ctx, customer)

    if err != nil {
        return "", err
    }

    return firebaseId, nil
}

// HandleCreateCustomer handles HTTP requests to create a customer in multiple systems
func (h *CreateCustomerHandler) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	client_ids, err := h.client.CreateCustomer(customer)	

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client_ids)
}