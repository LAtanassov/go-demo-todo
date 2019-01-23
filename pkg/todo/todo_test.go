package todo_test

import (
	"testing"

	"github.com/LAtanassov/go-demo-todo/pkg/todo"
)

func TestService_Create(t *testing.T) {
	type args struct {
		t todo.Item
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"should create a new item", args{t: todo.Item{Description: "description"}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := todo.NewService()
			if err := s.Create(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Service.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
