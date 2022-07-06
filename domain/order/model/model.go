package model

type Order struct {
	Id          string
	ProductName string
}

type Product struct {
	Url  string
	Name string
}

type DetailProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Url         string `json:"url"`
}
