package repo

import (
	"errors"
	database "homework/data_base"
	"homework/entity"
)

var ErrIdNotFound = errors.New("error not found")

type Tasks interface {
	List() ([]entity.Task, error)
	Create(task entity.Task) (int, error)
	Update(task entity.Task) error
	Delete(taskID int) error
}

type TasksTempDbImpl struct {
	Db database.TempDataBase
}

func NewTasks(Db database.TempDataBase) Tasks {
	return TasksTempDbImpl{
		Db: Db,
	}
}

func (t TasksTempDbImpl) List() ([]entity.Task, error) {
	return t.Db.List(), nil
}

func (t TasksTempDbImpl) Create(task entity.Task) (int, error) {
	return t.Db.Insert(task), nil
}

func (t TasksTempDbImpl) Update(task entity.Task) error {
	err := t.Db.Upsert(task.ID, task)
	if errors.Is(err, database.ErrIdNotExist) {
		return ErrIdNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (t TasksTempDbImpl) Delete(taskID int) error {
	err := t.Db.Delete(taskID)
	if errors.Is(err, database.ErrIdNotExist) {
		return ErrIdNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
