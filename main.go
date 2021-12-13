package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mustafakocatepe/go-api-example/handler/api"
	"github.com/mustafakocatepe/go-api-example/model"
)

func main() {

	model.UserArray = []model.User{
		{UserId: 1, UserName: "John", Surname: "Doe", IsActive: true},
		{UserId: 2, UserName: "Jane", Surname: "Doe", IsActive: true},
		{UserId: 3, UserName: "Mustafa", Surname: "Kocatepe", IsActive: true},
	}

	r := api.New(model.User{})
	h := r.Handler()

	s := http.Server{
		Addr:           ":8080",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //1mb
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Println("application failed to start")
		panic(err)
	}
}
