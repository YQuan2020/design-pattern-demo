package creational

import "testing"

func TestDoFactory(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test factory"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoFactory()
		})
	}
}
