package main

import (
	"attRest/db/internal/config"
	"attRest/db/internal/handler"
	repository "attRest/db/internal/repo"
	"attRest/db/internal/service"
	"attRest/db/pkg/db"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	db, err := db.NewDB(cfg)
	if err != nil {
		log.Println(err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandlerTask(service)

	http.HandleFunc("POST /create", handler.CreateTask())
	http.HandleFunc("GET /list", handler.GetTasksHandler())
	http.HandleFunc("DELETE /delete/{id}", handler.DeleteTaskHandler())
	http.HandleFunc("PATCH /done/{id}", handler.MarkTaskAsDoneHandler())

	log.Fatal(http.ListenAndServe(":8082", nil))
}
