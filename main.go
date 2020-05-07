package main

import (
	"log"
	"net/http"
	"truco-backend/routers"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", routers.BuildRouter()))
}
