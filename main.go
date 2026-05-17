package main

import (
	"encoding/json"
	"net/http"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {

	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	key := generateCode()

	_, err = db.Exec(`
		INSERT INTO urls(
		short_code, original_url
		)
		values(
		?,?
		)
		`, key, req.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ShortenResponse{
		ShortURL: "http://localhost:8080/" + key,
	})

}

func RedirectURL(w http.ResponseWriter, r *http.Request) {

	reqU := r.URL.Path
	prefix := reqU[1:]

	row := db.QueryRow(`
		SELECT original_url FROM urls
		WHERE short_code = ?
		`, prefix)
	var originalURL string
	err := row.Scan(&originalURL)
	if err != nil {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)

}

func main() {
	initDB()
	http.HandleFunc("POST /url", ShortenURL)
	http.HandleFunc("GET /", RedirectURL)

	http.ListenAndServe(":8080", nil)

}
