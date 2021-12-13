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

	/*"github.com/goware/cors"
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	r.Use(cors.Handler)*/

	//log := logrus.New()
	//r.Use(logger.Logger("router", log))

	// r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.NoCache)

	/*
		r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

			// documentation for developers
			opts := sw.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
			sh := sw.SwaggerUI(opts, nil)
			r.Handle("/docs", sh)

			opts1 := sw.RedocOpts{SpecURL: "/swagger.yaml", Path: "docs1"}
			sh1 := sw.Redoc(opts1, nil)
			r.Handle("/docs1", sh1)*/

	// r.Route("/api/users", func(r chi.Router) {
	// 	r.Post("/login", user.HandleAuthentication(s.Users))
	// 	r.Post("/", user.HandleRegistration(s.Users))
	// })

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/", user.HandleCreateUser(s.Users))
		r.Get("/", user.HandleUsers(s.Users))
		r.Get("/{username}", user.HandleUser(s.Users))
		r.Put("/", user.HandleUpdate(s.Users))
		r.Delete("/{id}", user.HandleDelete(s.Users))
	})

	// r.Route("/api/profiles", func(r chi.Router) {
	// 	r.Get("/{username}", user.HandleProfile(s.Users))
	// 	r.With(middleware1.ValidateJWT).Post("/{username}/follow", user.HandleFollowUser(s.Users))
	// })
 
	// r.Route("/api/articles", func(r chi.Router) {
	// 	r.Get("/", article.HandleArticleList(s.Articles))
	// })

	return r
}
