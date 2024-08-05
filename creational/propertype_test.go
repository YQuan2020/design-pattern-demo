package creational

import "testing"

func TestDoPrototype(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test prototype"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoPrototype()
		})
	}
}
