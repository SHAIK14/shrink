package main

import (
	"encoding/json"
	"net/http"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	key := generateCode()
	urlStore[key] = req.Url
	json.NewEncoder(w).Encode(ShortenResponse{
		ShortURL: "http://localhost:8080/" + key,
	})

}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	reqU := r.URL.Path
	prefix := reqU[1:]
	resp, ok := urlStore[prefix]
	if !ok {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, resp, http.StatusFound)

}

func main() {
	initDB()
	http.HandleFunc("POST /url", ShortenURL)
	http.HandleFunc("GET /", RedirectURL)

	http.ListenAndServe(":8080", nil)

}
