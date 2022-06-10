package main

import (
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

type Pokemon struct {
	Name    graphql.String `json:"name,omitempty"`
	Message graphql.String `json:"message,omitempty"`
}

func (p *Pokemon) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func main() {
	// HTTP Server
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	router.Get("/users/{id}/{username}/*", getUser)
	router.Post("/users", createUser)

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
	if err := render.Bind(r, data); err != nil {
		log.Print(err.Error())
		render.Render(w, r, data)
		return
	}

	render.Render(w, r, data)
}
