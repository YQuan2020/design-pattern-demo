package structural

import "testing"

func TestDoFacade(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test facade"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoFacade()
		})
	}
}
