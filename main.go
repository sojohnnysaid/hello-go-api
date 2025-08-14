package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func main() {
	// Create a new router & API
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

	// Add more info to the API
	api.OpenAPI().Info.Description = "A bare-bones Go API with automatic OpenAPI/Swagger generation"

	// Register GET /greeting/{name}
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	// Start the server!
	fmt.Println("API server starting on http://127.0.0.1:8888")
	fmt.Println("Documentation available at:")
	fmt.Println("  - Interactive docs: http://127.0.0.1:8888/docs")
	fmt.Println("  - OpenAPI JSON: http://127.0.0.1:8888/openapi.json")
	fmt.Println("  - OpenAPI YAML: http://127.0.0.1:8888/openapi.yaml")
	
	http.ListenAndServe("127.0.0.1:8888", router)
}