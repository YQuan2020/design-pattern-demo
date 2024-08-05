package structural

import "testing"

func TestDoFlyweight(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test flyweight"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoFlyweight()
		})
	}
}
