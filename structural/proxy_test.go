package structural

import "testing"

func TestDoProxy(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test proxy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoProxy()
		})
	}
}
