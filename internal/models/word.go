package models

import "time"

type Word struct {
    ID          int       `json:"id"`
    Word        string    `json:"word"`
    Translation string    `json:"translation"`
    Example     string    `json:"example"`
    CreatedAt   time.Time `json:"created_at"`
}

type PaginatedResponse struct {
    Items      interface{} `json:"items"`
    Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
    CurrentPage   int `json:"current_page"`
    TotalPages   int `json:"total_pages"`
    TotalItems   int `json:"total_items"`
    ItemsPerPage int `json:"items_per_page"`
} 