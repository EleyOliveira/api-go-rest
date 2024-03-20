package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EleyOliveira/api-go-rest/routes"
)

func main() {
	fmt.Println("iniciando o servidor")
	routes.HandleRequest()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
