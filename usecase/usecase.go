package usecase

import (
	"errors"
	"fmt"
	database "homework/data_base"
	"homework/entity"
	"homework/repo"
)

type actionType string

const (
	createAction  actionType = "create"
	updatedAction actionType = "update"
)

var (
	ErrInvalidParameter = errors.New("invaild parameter error")
	ErrInvalidID        = errors.New("invalid ID")
	ErrIdNotFound       = errors.New("id is not found in data base")
)

type UseCase interface {
	ListTasks() ([]entity.Task, error)
	CreateTask(task entity.Task) (int, error)
	UpdateTask(updatedTask entity.Task) error
	DeleteTask(taskID int) error
}

type UseCaseImpl struct {
	tasks repo.Tasks
}

func NewUseCase(Db database.TempDataBase) UseCase {
	return mustNewUseCase(repo.NewTasks(Db))
}

func mustNewUseCase(tasks repo.Tasks) UseCase {
	return UseCaseImpl{tasks: tasks}
}

func (t UseCaseImpl) checkTaskInValid(task entity.Task, action actionType) error {
	if !(task.Status == entity.TaskCompleted || task.Status == entity.TaskIncomplete) {
		return fmt.Errorf("%w, error parameter is %s", ErrInvalidParameter, entity.StatusJsonName)
	}
	if task.Name == "" {
		return fmt.Errorf("%w, error parameter is %s", ErrInvalidParameter, entity.NameJsonName)
	}
	if action == createAction && task.ID != 0 {
		return fmt.Errorf("%w, ID should not be given while creating task", ErrInvalidID)
	}
	if action == updatedAction && task.ID <= 0 {
		return fmt.Errorf("%w, ID should be positive", ErrInvalidID)
	}
	return nil
}

func (t UseCaseImpl) ListTasks() ([]entity.Task, error) {
	return t.tasks.List()
}

func (t UseCaseImpl) CreateTask(task entity.Task) (int, error) {
	if err := t.checkTaskInValid(task, createAction); err != nil {
		return 0, err
	}

	return t.tasks.Create(task)
}

func (t UseCaseImpl) UpdateTask(updatedTask entity.Task) error {
	if err := t.checkTaskInValid(updatedTask, updatedAction); err != nil {
		return err
	}

	err := t.tasks.Update(updatedTask)
	if errors.Is(err, repo.ErrIdNotFound) {
		return ErrIdNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (t UseCaseImpl) DeleteTask(taskID int) error {
	if taskID <= 0 {
		return fmt.Errorf("%w, ID should be positive", ErrInvalidID)
	}

	err := t.tasks.Delete(taskID)
	if errors.Is(err, repo.ErrIdNotFound) {
		return ErrIdNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
