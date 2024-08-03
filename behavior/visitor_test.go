package behavior

import "testing"

func TestDoVisitor(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test visitor"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoVisitor()
		})
	}
}
