package design

import "testing"

func TestDoCommand(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test command design pattern"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoCommand()
		})
	}
}
