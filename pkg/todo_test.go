package todo

import (
	"testing"
)

func TestService_Create(t *testing.T) {
	type args struct {
		t Item
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"should create a new item", args{t: Item{Description: "description"}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if err := s.Create(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Service.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
