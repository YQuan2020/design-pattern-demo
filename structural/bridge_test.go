package structural

import "testing"

func TestDoBridge(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test bridge"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoBridge()
		})
	}
}
