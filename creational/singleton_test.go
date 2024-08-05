package creational

import "testing"

func TestDoSingleton(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test singleton"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoSingleton()
		})
	}
}
