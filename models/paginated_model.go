package models

import "github.com/go-flow/template-api/pkg/paging"

// PaginatedModel model
type PaginatedModel struct {
	Results   interface{}       `json:"results"`
	Paginator *paging.Paginator `json:"paginator"`
}
