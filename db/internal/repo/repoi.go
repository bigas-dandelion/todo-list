package repository

import "attRest/db/model"

type IRepository interface {
	Create(task *model.Task) (*model.Task, error)
	Delete(id int) error
	Get() ([]*model.Task, error)
	MarkTaskAsDone(id int) error
}