package routes

import (
	"net/http"

	"github.com/EleyOliveira/api-go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
}
