package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
)

func main() {
	userHandler := InitUserHandler()

	reflector := openapi31.NewReflector()
	svc := web.NewService(reflector)

	svc.OpenAPISchema().SetTitle("User API")
	svc.OpenAPISchema().SetDescription("Service for managing users")
	svc.OpenAPISchema().SetVersion("v1.0.0")

	svc.Route("/", func(r chi.Router) {
		userHandler.RegisterRoutes(r)
	})

	svc.Docs("/docs", swgui.New)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", svc))
}
