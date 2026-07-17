package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/donaldcrane/my-go-todo-api/config"
	"github.com/donaldcrane/my-go-todo-api/models"
	"github.com/donaldcrane/my-go-todo-api/services"
)
import "errors"
import "gorm.io/gorm"
import "strconv"

type TodoHandler struct {
	Service *services.TodoService
}

func NewTodoHandler(
	service *services.TodoService,
) *TodoHandler {

	return &TodoHandler{
		Service: service,
	}

}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := h.Service.Create(&todo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, todo)

}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	var todos []models.Todo
	todos, err := h.Service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch todos",
		})

		return
	}
	c.JSON(http.StatusOK, todos)

}

func (h *TodoHandler) GetTodo(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})

		return
	}

	todo, err := h.Service.GetByID(id)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Todo not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	c.JSON(http.StatusOK, todo)

}

func (h *TodoHandler) GetCompleted(c *gin.Context) {

	todos, err := h.Service.GetCompleted()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	c.JSON(http.StatusOK, todos)

}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})

		return
	}

	var todo models.Todo

	todo, todoExist := h.Service.GetByID(id)

	if todoExist != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})

		return
	}

	var input models.Todo

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	todo.Title = input.Title
	todo.Description = input.Description
	todo.Completed = input.Completed

	h.Service.Update(&todo)

	c.JSON(http.StatusOK, todo)

}

func (h *TodoHandler) CompleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})

		return
	}

	var todo models.Todo

	todo, todoExist := h.Service.GetByID(id)

	if todoExist != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})

		return
	}
	todo.Completed = true
	h.Service.Update(&todo)
	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})

		return
	}

	var todo models.Todo

	todo, todoExist := h.Service.GetByID(id)

	if todoExist != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})

		return
	}

	err := h.Service.Delete(todo)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete todo",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully",
	})
}
