package repo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	database "homework/data_base"
	mockdb "homework/data_base/mocks"
	"homework/entity"
)

const (
	testID   int    = 0
	testName string = "task_test"
)

var (
	testTask = entity.Task{
		Name:   testName,
		Status: entity.TaskCompleted,
	}

	defulatTasks = []entity.Task{
		{ID: 1, Name: testName, Status: entity.TaskCompleted},
		{ID: 2, Name: testName, Status: entity.TaskIncomplete},
		{ID: 3, Name: testName, Status: entity.TaskIncomplete},
	}

	ErrTest = errors.New("test error")
)

type TasksSuite struct {
	suite.Suite
	dataBase *mockdb.TempDataBase
	tasks    Tasks
}

func (s *TasksSuite) SetupTest() {
	s.dataBase = mockdb.NewTempDataBase(s.T())
	s.tasks = NewTasks(s.dataBase)
}

func TestTasksSuite(t *testing.T) {
	suite.Run(t, new(TasksSuite))
}

func (s *TasksSuite) TestTasksList() {
	testCases := []struct {
		description  string
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
		s.dataBase.On("List").Return(testCase.exceptResult).Once()
		result, err := s.tasks.List()
		s.NoError(err)
		s.Equal(result, testCase.exceptResult, testCase.description)
	}
}

func (s *TasksSuite) TestTasksCreate() {
	testCases := []struct {
		description  string
		exceptResult int
	}{
		{
			description:  "First create",
			exceptResult: 1,
		},
		{
			description:  "Second create",
			exceptResult: 2,
		},
	}
	for _, testCase := range testCases {
		s.dataBase.On("Insert", testTask).Return(testCase.exceptResult).Once()
		result, err := s.tasks.Create(testTask)
		s.NoError(err)
		s.Equal(result, testCase.exceptResult, testCase.description)
	}
}

func (s *TasksSuite) TestTasksUpdate() {
	testCases := []struct {
		description  string
		mockResult   error
		exceptResult error
	}{
		{
			description:  "ID not found",
			mockResult:   database.ErrIdNotExist,
			exceptResult: ErrIdNotFound,
		},
		{
			description:  "updated successfully",
			mockResult:   nil,
			exceptResult: nil,
		},
		{
			description:  "other errors",
			mockResult:   ErrTest,
			exceptResult: ErrTest,
		},
	}
	for _, testCase := range testCases {
		s.dataBase.On("Upsert", testTask.ID, testTask).Return(testCase.mockResult).Once()
		err := s.tasks.Update(testTask)
		s.Equal(err, testCase.exceptResult, testCase.description)
	}
}

func (s *TasksSuite) TestTasksDelete() {
	testCases := []struct {
		description  string
		mockResult   error
		exceptResult error
	}{
		{
			description:  "ID not found",
			mockResult:   database.ErrIdNotExist,
			exceptResult: ErrIdNotFound,
		},
		{
			description:  "updated successfully",
			mockResult:   nil,
			exceptResult: nil,
		},
		{
			description:  "other errors",
			mockResult:   ErrTest,
			exceptResult: ErrTest,
		},
	}
	for _, testCase := range testCases {
		s.dataBase.On("Delete", testID).Return(testCase.mockResult).Once()
		err := s.tasks.Delete(testID)
		s.Equal(err, testCase.exceptResult, testCase.description)
	}
}
