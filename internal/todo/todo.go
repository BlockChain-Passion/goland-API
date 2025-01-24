package todo

import (
	"awesomeProject/internal/db"
	"context"
	"fmt"
	"strings"
)

type Item struct {
	Task   string
	Status string
}

type Manager interface {
	InsertItem(ctx context.Context, item db.Item) error
	GetAllItems(ctx context.Context) ([]db.Item, error)
}

type Service struct {
	db Manager
}

func NewService(db Manager) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetAll() ([]Item, error) {
	var results []Item

	items, err := s.db.GetAllItems(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error in fetching all records from db: %w", err)
	}

	for _, item := range items {
		results = append(results, Item{
			Task:   item.Task,
			Status: item.Status,
		})
	}
	return results, nil
}

func (s *Service) AddTodo(str string) error {

	items, err := s.GetAll()

	if err != nil {
		return fmt.Errorf("error in fetching all records from db: %w", err)
	}

	for _, i := range items {
		if i.Task == str && i.Task != "" {
			return fmt.Errorf("%s is already present", str)
		}
	}

	if err := s.db.InsertItem(context.Background(), db.Item{
		Task:   str,
		Status: "TO_BE_STARTED",
	}); err != nil {
		return fmt.Errorf("error in inserting item into db: %w", err)
	}

	return nil
}

func (s *Service) Search(str string) ([]string, error) {
	items, err := s.GetAll()

	if err != nil {
		return nil, fmt.Errorf("error in fetching all records from db: %w", err)
	}

	var result []string
	for _, i := range items {
		if strings.Contains(i.Task, str) {
			result = append(result, i.Task)
		}
	}
	return result, nil
}
