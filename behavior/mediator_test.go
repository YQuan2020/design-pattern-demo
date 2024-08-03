package behavior

import "testing"

func TestDoMediator(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test mediator"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoMediator()
		})
	}
}
