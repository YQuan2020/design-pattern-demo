package design

import "testing"

func TestDoObserver(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test observer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoObserver()
		})
	}
}
