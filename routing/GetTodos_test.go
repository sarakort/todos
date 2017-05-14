package routing

import (
	"testing"
)

func TestGetAllTodosSizeShouldNotBeEmpty(t *testing.T) {
	result := GetAllTodos()
	if result == nil {
		t.Error("GetAlTodos return nil")
	}

}
