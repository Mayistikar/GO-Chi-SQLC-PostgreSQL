package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/shurcooL/graphql"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID       int64  `json:"userID,omitempty" validate:"omitempty"`
	UserName string `json:"userName,omitempty" validate:"alphanum"`
	Wildcard string `json:"wildcard,omitempty" validate:"required"`
}

func (u *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (u *User) Bind(r *http.Request) error {
	return nil
}

type Character struct {
	ID     graphql.ID     `json:"id,omitempty"`
	Name   graphql.String `json:"name,omitempty"`
	Gender graphql.String `json:"gender,omitempty"`
}

func (p *Character) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func main() {
	// HTTP Server
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// User Routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	router.Get("/users/{id}/{username}/*", getUser)
	router.Post("/users", createUser)

	// Character Routes
	router.Get("/rick-morty/{id}", getCharacter)

	http.ListenAndServe(":3000", router)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
	userName := chi.URLParam(r, "username")
	wildcard := chi.URLParam(r, "*")

	user := &User{userID, userName, wildcard}

	render.Status(r, http.StatusOK)
	render.Render(w, r, user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	data := &User{}
	/*
		if err := render.Bind(r, data); err != nil {
			log.Print(err.Error())
			render.Render(w, r, data)
			return
		}
	*/

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	render.Render(w, r, data)
}

func getCharacter(w http.ResponseWriter, r *http.Request) {
	// GraphQL Client
	client := graphql.NewClient("https://rickandmortyapi.com/graphql", nil)

	var query struct {
		Character Character `graphql:"character(id: $id)"`
	}

	variables := map[string]interface{}{
		"id": chi.URLParam(r, "id"),
	}

	err := client.Query(r.Context(), &query, variables)
	if err != nil {
		log.Printf("error: %v", err)
	}

	render.Render(w, r, &query.Character)
}
