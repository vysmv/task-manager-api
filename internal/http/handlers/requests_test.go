package handlers

import "testing"

func TestCreateTaskRequest_Validate(t *testing.T) {
	tests := []struct {
		name string
		req  CreateTaskRequest
		want bool // есть ли ошибки
	}{
		{
			name: "valid",
			req:  CreateTaskRequest{Title: "learn go"},
			want: false,
		},
		{
			name: "empty title",
			req:  CreateTaskRequest{Title: ""},
			want: true,
		},
		{
			name: "spaces only",
			req:  CreateTaskRequest{Title: "   "},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := tt.req.Validate()

			if tt.want && len(errs) == 0 {
				t.Fatalf("expected validation error, got none")
			}

			if !tt.want && len(errs) != 0 {
				t.Fatalf("unexpected validation errors: %v", errs)
			}
		})
	}
}