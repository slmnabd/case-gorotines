package main

type BookList struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	ImageUrl   string `json:"image_url"`
	Author     string `json:"author"`
	Price      int    `json:"price"`
}

type Book struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Title       string `json:"title"`
	ImageUrl    string `json:"image_url"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Format      string `json:"format"`
	Dimensions  string `json:"dimensions"`
	Language    string `json:"language"`
	ISBN13      string `json:"ISBN13"`
	Price       int    `json:"price"`
}

type Rating struct {
	ID        int             `json:"id"`
	ProductID int             `json:"product_id"`
	Qty       int             `json:"qty"`
	Rating    float64         `json:"rating"`
	Details   []RatingDetails `json:"details"`
}

type RatingDetails struct {
	Qty        int    `json:"qty"`
	Rating     int    `json:"rating"`
	Percentage string `json:"percentage"`
}

type ResponseBookDetail struct {
	Book       Book
	Rating     Rating
	Bestseller []BookList
}
