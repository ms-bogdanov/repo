package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"repo/config"
	"repo/internal/controller"
	"repo/internal/repository"
	"repo/internal/service"
)

func main() {
	r := chi.NewRouter()

	cfg := config.NewConfig()

	rep := repository.NewUserStorage(cfg)
	svc := service.NewService(rep)
	trans := controller.NewTransport(svc)

	r.Post("/user/create", trans.Create)
	r.Get("/user/get", trans.Get)
	r.Put("/user/update", trans.Update)
	r.Delete("/user/delete", trans.Delete)
	r.Get("/user", trans.List)

	http.ListenAndServe(":8080", r)
}
