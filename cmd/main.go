package main

import (
	"github.com/go-chi/chi"

	_ "./docs"

	"log"
	"net/http"

	"./api/controller"
	"github.com/swaggo/http-swagger"
)

func InitializeRouter() chi.Router {

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8001/swagger/doc.json"), //The url pointing to API definition"
	))

	r.Get("/sensor", controller.SensorIndex)
	r.Get("/sensorAverage", controller.SensorIndexAvg)

	return r
}

// @title API GoMeteo MQTT
// @version 1.0
// @description Project Go - IMTA.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /
func main() {
	// set routes
	router := InitializeRouter()
	log.Fatal(http.ListenAndServe(":8001", router))
}
