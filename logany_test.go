package logany

import "testing"

func Test_log(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test result"},
		{name: "another result"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := tt.name
			got := log(tt.name)
			if got != tt.name {
				t.Errorf("log() = %v, want %v", got, expected)
			}
		})
	}
}
