package main

type ListProductsResponse struct {
	Products []*Product `json:"products"`
}

type Product struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	PriceUsd    *Money   `json:"price_usd"`
	Categories  []string `json:"categories"`
}

type Money struct {
	CurrencyCode string `json:"currency_code"`
	Units        int64  `json:"units"`
	Nanos        int32  `json:"nanos"`
}
