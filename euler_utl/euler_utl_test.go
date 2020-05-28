package euler_utl

import (
	"testing"
	"time"

	"project_euler/logging"
)

func TestP1(t *testing.T) {
	logger, _ := logging.New(time.RFC3339, logging.DefaultLevel)
	result, _ := P1(logger)
	p1_result := uint64(233168)
	if result != p1_result {
		t.Errorf("ERROR: Result of the problem 1 should be %d and not %d ", uint64(p1_result), uint64(result))
	}
}

func TestP2(t *testing.T) {
	logger, _ := logging.New(time.RFC3339, logging.DefaultLevel)
	result, _ := P2(logger)
	p1_result := 4613732
	if uint64(result) != uint64(p1_result) {
		t.Errorf("ERROR: Result of the problem 2 should be %d and not %d ", uint64(p1_result), uint64(result))
	}
}

func TestP3(t *testing.T) {
	logger, _ := logging.New(time.RFC3339, logging.DefaultLevel)
	result, _ := P3(logger)
	p1_result := 6857
	if uint64(result) != uint64(p1_result) {
		t.Errorf("ERROR: Result of the problem 3 should be %d and not %d ", uint64(p1_result), uint64(result))
	}
}
