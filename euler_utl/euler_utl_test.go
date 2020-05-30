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

func TestP4(t *testing.T) {
	logger, _ := logging.New(time.RFC3339, logging.DefaultLevel)
	result, _ := P4(logger)
	p1_result := 906609
	if uint64(result) != uint64(p1_result) {
		t.Errorf("ERROR: Result of the problem 4 should be %d and not %d ", uint64(p1_result), uint64(result))
	}
}

func TestP5(t *testing.T) {
	p_num := 5
	logger, _ := logging.New(time.RFC3339, logging.DefaultLevel)
	result, _ := P5(logger)
	p1_result := 232792560
	if uint64(result) != uint64(p1_result) {
		t.Errorf("ERROR: Result of the problem %d should be %d and not %d ", p_num, uint64(p1_result), uint64(result))
	}
}

func TestP6(t *testing.T) {
	p_num := 6
	logger, _ := logging.New(time.RFC3339, logging.DefaultLevel)
	result, _ := P6(logger)
	p1_result := 25164150
	if uint64(result) != uint64(p1_result) {
		t.Errorf("ERROR: Result of the problem %d should be %d and not %d ", p_num, uint64(p1_result), uint64(result))
	}
}
