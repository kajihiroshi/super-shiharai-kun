package models

type Client struct {
	ID             int
	CompanyID      int
	Name           string
	Representative string
	Phone          string
	PostalCode     string
	Address        string
}
