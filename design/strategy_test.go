package design

import "testing"

func TestDoStrategy(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test strategy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoStrategy()
		})
	}
}
