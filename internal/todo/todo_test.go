package todo_test

import (
	"awesomeProject/internal/db"
	"awesomeProject/internal/todo"
	"context"
	"reflect"
	"testing"
)

type MockDB struct {
	items []db.Item
}

func (m *MockDB) InsertItem(ctx context.Context, item db.Item) error {
	m.items = append(m.items, item)
	return nil
}

func (m *MockDB) GetAllItems(ctx context.Context) ([]db.Item, error) {
	return m.items, nil
}

func TestService_Search(t *testing.T) {

	tests := []struct {
		name           string
		toDosToAdd     []string
		query          string
		expectedResult []string
	}{
		// TODO: Add test cases.
		{name: "Todo test test case1", toDosToAdd: []string{"shop"}, query: "sh", expectedResult: []string{"shop"}},
		{name: "Todo test case2", toDosToAdd: []string{"shop1"}, query: "sh", expectedResult: []string{"shop1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDB{}
			svc := todo.NewService(m)
			for _, toAdd := range tt.toDosToAdd {
				err := svc.AddTodo(toAdd)
				if err != nil {
					t.Errorf("%v", err)
				}
			}
			got, err := svc.Search(tt.query)
			if err != nil {
				t.Errorf("%v", err)
			}
			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
