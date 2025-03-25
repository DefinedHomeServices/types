package customer

//customer/customer.go

import (
	"time"
)

type CustomerAddress struct {
    Line1       string `json:"line1"`
    Line2       string `json:"line2"`
}

type CustomerLocation struct {
    Address CustomerAddress     `json:"address"`
    City             string     `json:"city"`
    State            string     `json:"state"`
    Zip              string     `json:"zip"`
}

type Customer struct {
    ID               string             `json:"id"`
    FirstName        string             `json:"first_name"`
    LastName         string             `json:"last_name"`
    Phone            string             `json:"phone"`
    Email            string             `json:"email"`
    DateCreated      time.Time          `json:"date_created"`
    LastUpdated      time.Time          `json:"last_updated"`
    Location         CustomerLocation   `json:"location"`
    OptedInEmails    bool               `json:"opted_in_emails"`
    StripeCustomerID string             `json:"stripe_customer_id"`
}

type AcuityAppointmentsRelationship struct {
    CustomerID 		int `json:"customer_id"`
    AppointmentID 	int `json:"appointment_id"`
}

type StripeInvoicesRelationship struct {
	CustomerID 				int `json:"customer_id"`
	StripeInvoiceID 		string `json:"invoice_id"`
	AcuityAppointmentID 	int `json:"acuity_appointment_id"`
	StripeCustomerID 		string `json:"stripe_customer_id"`
}
