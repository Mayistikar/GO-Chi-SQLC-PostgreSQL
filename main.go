package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type User struct {
	UserID   int64  `json:"userID,omitempty"`
	UserName string `json:"userName,omitempty"`
	Wildcard string `json:"wildcard,omitempty"`
}

func (u *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	router.Get("/users/{id}/{username}/*", getUser)

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
