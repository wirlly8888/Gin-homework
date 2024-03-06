package handler

import (
	"errors"
	database "homework/data_base"
	"homework/entity"
	"homework/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidParameterType = errors.New("invaild parameter type")
	ErrInvalidParameter     = errors.New("invaild parameter")
	ErrInvalidID            = errors.New("invalid task ID")
	ErrIdNotFound           = errors.New("id not found")
	ErrInternal             = errors.New("internal service error")
)

type ServerHandler interface {
	ListTasks(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type HandlerImpl struct {
	Usecase usecase.UseCase
}

func SetupRoute(r *gin.Engine, Db database.TempDataBase) *gin.Engine {
	h := newHandler(Db)
	r.GET("/tasks", h.ListTasks)
	r.POST("/tasks", h.CreateTask)
	r.PUT("/tasks/:id", h.UpdateTask)
	r.DELETE("/tasks/:id", h.DeleteTask)
	return r
}

func newHandler(Db database.TempDataBase) ServerHandler {
	return HandlerImpl{
		// Usecase: usecase.NewUseCase(Db),
		Usecase: usecase.NewUseCase(Db),
	}
}

func (h HandlerImpl) ListTasks(c *gin.Context) {
	tasks, err := h.Usecase.ListTasks()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInternal.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h HandlerImpl) CreateTask(c *gin.Context) {
	var newTask entity.Task
	if err := c.BindJSON(&newTask); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidParameterType.Error()})
		return
	}

	taskID, err := h.Usecase.CreateTask(newTask)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidParameter.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"task id": taskID})
}

func (h HandlerImpl) UpdateTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidID.Error()})
		return
	}

	var updatedTask entity.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidParameterType.Error()})
		return
	}

	updatedTask.ID = taskID
	if err := h.Usecase.UpdateTask(updatedTask); err != nil {
		log.Println(err.Error())
		switch {
		case errors.Is(err, usecase.ErrIdNotFound):
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrIdNotFound.Error()})
		case errors.Is(err, usecase.ErrInvalidParameter):
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidParameter.Error()})
		case errors.Is(err, usecase.ErrInvalidID):
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidParameter.Error()})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInternal.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"updated task ID": taskID})
}

func (h HandlerImpl) DeleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidID.Error()})
		return
	}

	if err := h.Usecase.DeleteTask(taskID); err != nil {
		log.Println(err.Error())
		switch {
		case errors.Is(err, usecase.ErrIdNotFound):
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrIdNotFound.Error()})
		case errors.Is(err, usecase.ErrInvalidID):
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidID.Error()})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInternal.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted task ID": taskID})
}
