package main

import (
	"io"
	"net/http"

	"github.com/savaki/swaggering"
)

func echo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World")
}

type Owner struct {
	Name string
}

type Pet struct {
	Name  string
	Owner Owner
}

func main() {
	api := &swaggering.Api{
		BasePath: "/api",
		CORS:     true,
		Endpoints: []swaggering.Endpoint{
			{
				Method:      "get",
				Path:        "/pet",
				Summary:     "Add a New Pet",
				Description: "PetDescription",
				HandlerFunc: echo,
				Parameter: &swaggering.Parameter{
					Description: "Thingie!",
					Schema:      Owner{},
				},
				Responses: map[int]swaggering.Response{
					http.StatusOK: {
						Description: "Woo hoo!",
						Schema:      []Pet{},
					},
				},
			},
		},
	}

	http.Handle("/swagger", api)
	http.ListenAndServe(":8080", nil)
}