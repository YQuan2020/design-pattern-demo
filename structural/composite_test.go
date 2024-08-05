package structural

import "testing"

func TestDoComposite(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test composite"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoComposite()
		})
	}
}
