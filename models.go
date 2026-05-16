package main

type ShortenRequest struct {
	Url string `json:"Url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
