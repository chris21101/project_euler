package euler_utl

import (
	"testing"
	"time"

	"project_euler/logging"
)

func TestP1(t *testing.T) {
	logger, _ := logging.New(time.RFC3339, logging.DefaultLevel)
	result, _, _ := P1(logger)
	p1_result := 233168
	if result != p1_result {
		t.Errorf("ERROR: Result of the problem 1 should be %d and not %d ", p1_result, result)
	}
}
