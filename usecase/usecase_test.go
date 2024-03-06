package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"homework/entity"
	"homework/repo"
	mockrepo "homework/repo/mocks"
)

const (
	testID     int    = 1
	testStatus        = 10
	testName   string = "task_test"
)

var (
	testTask = entity.Task{
		Name: testName, Status: entity.TaskCompleted,
	}

	defulatTasks = []entity.Task{
		{ID: 1, Name: testName, Status: entity.TaskCompleted},
		{ID: 2, Name: testName, Status: entity.TaskIncomplete},
		{ID: 3, Name: testName, Status: entity.TaskIncomplete},
	}

	ErrTest = errors.New("test error")
)

type UseCaseSuite struct {
	suite.Suite

	tasks   *mockrepo.Tasks
	useCase UseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.tasks = mockrepo.NewTasks(s.T())
	s.useCase = mustNewUseCase(s.tasks)
}

func TestUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}

func (s *UseCaseSuite) TestListTasks() {
	testCases := []struct {
		description  string
		excepetErr   error
		exceptResult []entity.Task
	}{
		{
			description:  "empty task",
			exceptResult: []entity.Task{},
		},
		{
			description:  "three tasks",
			exceptResult: defulatTasks,
		},
	}
	for _, testCase := range testCases {
		s.tasks.On("List").Return(testCase.exceptResult, testCase.excepetErr).Once()
		result, err := s.useCase.ListTasks()
		s.NoError(err)
		s.Equal(result, testCase.exceptResult, testCase.description)
	}
}

func (s *UseCaseSuite) TestCreateTask() {
	testCases := []struct {
		description  string
		task         entity.Task
		excepetErr   error
		exceptResult int
	}{
		{
			description:  "default create",
			task:         testTask,
			exceptResult: 0,
		},
		{
			description:  "default create 2",
			task:         entity.Task{Name: testName, Status: entity.TaskCompleted},
			exceptResult: 0,
		},
		{
			description: "invalid status",
			task:        entity.Task{Name: testName, Status: testStatus},
			excepetErr:  ErrInvalidParameter,
		},
		{
			description: "empty name",
			task:        entity.Task{Name: "", Status: entity.TaskCompleted},
			excepetErr:  ErrInvalidParameter,
		},
		{
			description: "with ID",
			task:        entity.Task{ID: testID, Name: testName, Status: entity.TaskCompleted},
			excepetErr:  ErrInvalidID,
		},
	}
	for _, testCase := range testCases {
		if testCase.excepetErr == nil {
			s.tasks.On("Create", testCase.task).Return(testCase.exceptResult, testCase.excepetErr).Once()
		}
		result, err := s.useCase.CreateTask(testCase.task)
		if testCase.excepetErr != nil {
			s.ErrorIs(err, testCase.excepetErr, testCase.description)
		} else {
			s.NoError(err)
			s.Equal(result, testCase.exceptResult, testCase.description)
		}
	}
}

func (s *UseCaseSuite) TestUpdateTask() {
	testCases := []struct {
		description  string
		mock         bool
		task         entity.Task
		mockResult   error
		exceptResult error
	}{
		{
			description:  "default update",
			mock:         true,
			task:         defulatTasks[1],
			exceptResult: nil,
		},
		{
			description:  "Invalid ID",
			task:         testTask,
			exceptResult: ErrInvalidID,
		},
		{
			description:  "invalid status",
			task:         entity.Task{ID: testID, Name: testName, Status: testStatus},
			exceptResult: ErrInvalidParameter,
		},
		{
			description:  "empty name",
			task:         entity.Task{ID: testID, Name: "", Status: entity.TaskCompleted},
			exceptResult: ErrInvalidParameter,
		},
		{
			description:  "ID not found",
			mock:         true,
			task:         defulatTasks[1],
			mockResult:   repo.ErrIdNotFound,
			exceptResult: ErrIdNotFound,
		},
		{
			description:  "other errors",
			mock:         true,
			task:         defulatTasks[1],
			mockResult:   ErrTest,
			exceptResult: ErrTest,
		},
	}
	for _, testCase := range testCases {
		if testCase.mock {
			s.tasks.On("Update", testCase.task).Return(testCase.mockResult).Once()
		}
		err := s.useCase.UpdateTask(testCase.task)
		if testCase.exceptResult != nil {
			s.ErrorIs(err, testCase.exceptResult, testCase.description)
		} else {
			s.NoError(err, testCase.description)
		}
	}
}

func (s *UseCaseSuite) TestDeleteTask() {
	testCases := []struct {
		description  string
		mock         bool
		taskID       int
		mockResult   error
		exceptResult error
	}{
		{
			description:  "default delete",
			mock:         true,
			taskID:       testID,
			exceptResult: nil,
		},
		{
			description:  "Invalid ID",
			taskID:       -1,
			exceptResult: ErrInvalidID,
		},
		{
			description:  "ID not found",
			mock:         true,
			taskID:       testID,
			mockResult:   repo.ErrIdNotFound,
			exceptResult: ErrIdNotFound,
		},
		{
			description:  "other errors",
			mock:         true,
			taskID:       testID,
			mockResult:   ErrTest,
			exceptResult: ErrTest,
		},
	}
	for _, testCase := range testCases {
		if testCase.mock {
			s.tasks.On("Delete", testCase.taskID).Return(testCase.mockResult).Once()
		}
		err := s.useCase.DeleteTask(testCase.taskID)
		if testCase.exceptResult != nil {
			s.ErrorIs(err, testCase.exceptResult, testCase.description)
		} else {
			s.NoError(err, testCase.description)
		}
	}
}
