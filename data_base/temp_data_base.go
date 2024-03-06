package database

import (
	"errors"
	"homework/entity"
)

var (
	ErrIdNotExist = errors.New("id not exist")
)

type TempDataBase interface {
	List() []entity.Task
	Insert(datum entity.Task) int
	Upsert(ID int, datum entity.Task) error
	Delete(ID int) error
}

type TempDbImpl struct {
	Db *[]entity.Task
}

func NewTempDataBase() TempDataBase {
	var tasks []entity.Task
	return TempDbImpl{
		Db: &tasks,
	}
}

func (t TempDbImpl) Insert(datum entity.Task) int {
	datum.ID = len(*t.Db) + 1
	*t.Db = append(*t.Db, datum)
	return datum.ID
}

func (t TempDbImpl) List() []entity.Task {
	return *t.Db
}

func (t TempDbImpl) Upsert(ID int, datum entity.Task) error {
	for i, task := range *t.Db {
		if task.ID == datum.ID {
			(*t.Db)[i] = datum
			return nil
		}
	}
	return ErrIdNotExist
}

func (t TempDbImpl) Delete(ID int) error {
	for i, task := range *t.Db {
		if task.ID == ID {
			*t.Db = append((*t.Db)[:i], (*t.Db)[i+1:]...)
			return nil
		}
	}
	return ErrIdNotExist
}
