package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	file, err := os.Create("a.txt.gz")
	if err != nil {
		panic(err)
	}
	gzipWriter := gzip.NewWriter(file)
	gzipWriter.Header.Name = "a.txt"

	writer := io.MultiWriter(gzipWriter, os.Stdout)

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	encoder.Encode(source)

	gzipWriter.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
