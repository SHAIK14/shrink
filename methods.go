package main

type ShortenRequest struct {
	Url string
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
