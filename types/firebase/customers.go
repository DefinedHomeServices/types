package customer

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

// FirebaseAPIClient defines a method for creating a customer in the database
type FirebaseAPIClient interface {
    AddCustomerToDatabase(ctx context.Context, customer map[string]interface{}) (string, error)
    GetCustomer(ctx context.Context, email string) (map[string]interface{}, error)
}

// FirebaseDBClient holds a Firestore client to interact with the Firebase database
type FirebaseDBClient struct {
    DB  *firestore.Client
}

func (h *FirebaseDBClient) GetCustomer(ctx context.Context, email string) (map[string]interface{}, error) {
    fmt.Println("Getting customer from email:", email)
    docRef := h.DB.Collection("customers").Where("email", "==", email).Documents(ctx)
    docs, err := docRef.GetAll()
    
    if (len(docs) == 0) {
        return nil, nil
    }

    if err != nil {
        fmt.Println("Error getting customer:", err)
        return nil, err
    }
    customer := docs[0].Data()
    
    return customer, nil
}

func (h *FirebaseDBClient) AddCustomerToDatabase(ctx context.Context, customer map[string]interface{}) (string, error) {
    fmt.Println("Creating customer in Firebase")
    docRef, _, err := h.DB.Collection("customers").Add(ctx, customer)
    if err != nil {
        fmt.Println("Error creating customer:", err)
        return "", err
    }
    fmt.Println("Customer created with ID:", docRef.ID)
    return docRef.ID, nil
}