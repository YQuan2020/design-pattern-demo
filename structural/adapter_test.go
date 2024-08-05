package structural

import "testing"

func TestDoAdapter(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test adapter"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoAdapter()
		})
	}
}
