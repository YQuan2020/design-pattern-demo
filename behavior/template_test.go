package behavior

import "testing"

func TestDoTemplate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test template"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoTemplate()
		})
	}
}
