package services

import (
	"errors"
	"github.com/donaldcrane/my-go-todo-api/models"
	"github.com/donaldcrane/my-go-todo-api/repositories"
)

type TodoService struct {
	Repository *repositories.TodoRepository
}

func NewTodoService(
	repository *repositories.TodoRepository,
) *TodoService {

	return &TodoService{
		Repository: repository,
	}

}

func (s *TodoService) Create(todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}

	return s.Repository.Create(todo)

}

func (s *TodoService) GetAll() ([]models.Todo, error) {

	return s.Repository.FindAll()

}

func (s *TodoService) GetByID(id int) (models.Todo, error) {

	return s.Repository.FindByID(id)

}

func (s *TodoService) GetCompleted() ([]models.Todo, error) {

	return s.Repository.GetCompleted()

}

func (s *TodoService) Update(todo *models.Todo) error {

	return s.Repository.Update(todo)

}

func (s *TodoService) Delete(todo *models.Todo) error {

	return s.Repository.Delete(todo)

}
