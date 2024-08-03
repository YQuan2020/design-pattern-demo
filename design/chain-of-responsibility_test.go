package design

import "testing"

func TestDoChainOfResponsibility(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test chain of responsibility"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoChainOfResponsibility()
		})
	}
}
