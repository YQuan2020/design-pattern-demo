package creational

import "testing"

func TestDoAnotherSingle(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test another singleton"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoAnotherSingle()
		})
	}
}
