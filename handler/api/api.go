package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mustafakocatepe/go-api-example/handler/api/user"
	"github.com/mustafakocatepe/go-api-example/model"
)

type Server struct {
	Users model.User
}

func New(users model.User) Server {
	return Server{
		Users: users,
	}
}

func (s Server) Handler() http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/", user.HandleCreateUser(s.Users))
		r.Get("/", user.HandleUsers(s.Users))
		r.Get("/{username}", user.HandleUser(s.Users))
		r.Patch("/{id}", user.HandleUpdateUserName(s.Users))
		r.Put("/{id}", user.HandleUpdate(s.Users))
		r.Delete("/{id}", user.HandleDelete(s.Users))
	})

	return r
}
