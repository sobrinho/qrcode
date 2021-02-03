package main

import (
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/qpliu/qrencode-go/qrencode"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Path[1:]

	if len(data) == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	grid, err := qrencode.Encode(data, qrencode.ECLevelQ)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	png.Encode(w, grid.Image(8))
}

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		log.Fatal("Missing PORT")
	}

	http.HandleFunc("/", handler)
	log.Printf("Listening on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
