package types

type CreditRequest struct {
	Amount int64 `json:"amount" validate:"numeric,min=0"`
}

type LoginRequest struct{
	User string `json:"user" validate:"required"`
}