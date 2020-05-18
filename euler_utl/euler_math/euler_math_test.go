// euler_math_test project euler_math_test.go
package euler_math

import (
	"testing"
	"time"

	"project_euler/logging"
)

func TestSumdivisibleby(t *testing.T) {
	logger, _ := logging.New(time.RFC3339, logging.DefaultLevel)
	sum3_below10 := Sumdivisibleby(3, 9, logger)
	if sum3_below10 != 18 {
		t.Errorf("Error: euler_math.Sumdivisibleby(%d,%d) %s", 3, 9, " is 18")
	}
}
