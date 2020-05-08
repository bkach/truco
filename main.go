package main

import (
	"github.com/bkach/truco/routers"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", routers.BuildRouter()))
}
