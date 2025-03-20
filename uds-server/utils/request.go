package utils

type RequestID struct {
	ID int `json:"id" validate:"required"`
}

type RequestPaginate struct {
	Limit int `json:"limit" validate:"required"`
	Page  int `json:"page" validate:"required"`
}
