package design

import "testing"

func TestDoMemento(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test memento"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoMemento()
		})
	}
}