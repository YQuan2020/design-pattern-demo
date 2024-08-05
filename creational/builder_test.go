package creational

import "testing"

func TestDoBuilder(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test builder"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoBuilder()
		})
	}
}
