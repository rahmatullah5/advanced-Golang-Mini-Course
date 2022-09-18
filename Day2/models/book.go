package models

type BooksResponse []struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	Writer    string `json:"writer"`
}

type BookResponse struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	Writer    string `json:"writer"`
}

type CreateBookResponse struct {
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
	ClientId     uint   `json:"client_id"`
	ClientKey    string `json:"client_key"`
	ClientSecret string `json:"client_secret"`
	Status       bool   `json:"status"`
}

type DeleteBookResponse struct {
	Message string `json:"message"`
}
