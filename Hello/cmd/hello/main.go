package main

import (
  "net/http"
  "/internal"
)

func main() {
	handler := http.NewHandler()
  http.ListenAndServe(":8080", handler)
}
