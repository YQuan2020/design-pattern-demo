package creational

import "testing"

func TestDoAbstractFactory(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test abstract factory"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoAbstractFactory()
		})
	}
}
