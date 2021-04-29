package main

import (
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	req.Write(os.Stdout)
}
