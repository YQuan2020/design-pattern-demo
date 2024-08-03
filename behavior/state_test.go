package behavior

import "testing"

func TestDoState(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "do state"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoState()
		})
	}
}
