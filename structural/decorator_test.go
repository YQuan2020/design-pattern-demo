package structural

import "testing"

func TestDoDecorator(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test decorator"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoDecorator()
		})
	}
}
