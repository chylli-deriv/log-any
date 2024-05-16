package util

import (
	"testing"
)

func TestLogLevels(t *testing.T) {
	expected := 0 // As defined by iota in the original file
	got := LOG_LEVELS["EMERGENCY"]

	if got != expected {
		t.Errorf("LOG_LEVELS[\"EMERGENCY\"] = %d; want %d", got, expected)
	}
}
