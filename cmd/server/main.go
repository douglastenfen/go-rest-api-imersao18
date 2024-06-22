package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/douglastenfen/go-rest-api/internal/router"
)

func main() {
	r := router.NewRouter()

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
