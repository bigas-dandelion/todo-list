package handler

import (
	"attRest/db/internal/service"
	"attRest/db/model"
	"encoding/json"
	"net/http"
	"strconv"
)

type HandlerTask struct {
	service *service.Service
}

func NewHandlerTask(service *service.Service) *HandlerTask {
	return &HandlerTask{
		service: service,
	}
}

func (h *HandlerTask) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task model.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdTask, err := h.service.Create(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdTask)
	}
}

func (h *HandlerTask) GetTasksHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := h.service.Get()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

func (h *HandlerTask) DeleteTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.service.Delete(int(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("deleted"))
	}
}

func (h *HandlerTask) MarkTaskAsDoneHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.service.MarkTaskAsDone(int(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("marked as done"))
	}
}
