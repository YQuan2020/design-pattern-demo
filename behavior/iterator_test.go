package behavior

import "testing"

func TestDoIterator(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test iterator"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoIterator()
		})
	}
}