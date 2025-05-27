package repository

import (
	"attRest/db/model"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo Repository) Create(task *model.Task) (*model.Task, error) {
	query := `INSERT INTO list_tasks (Title, Done) VALUES ($1, $2) RETURNING id`

	err := repo.db.QueryRow(query, task.Title, task.Done).Scan(&task.ID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (repo Repository) Delete(id int) error {
	_, err := repo.db.Exec("DELETE FROM list_tasks WHERE Id = $1", id)
	return err
}

func (repo Repository) Get() ([]*model.Task, error) {
	rows, err := repo.db.Query("SELECT * FROM list_tasks")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		task := &model.Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (repo Repository) MarkTaskAsDone(id int) error {
	_, err := repo.db.Exec("UPDATE list_tasks SET Done = TRUE WHERE Id = $1", id)
	return err
}
