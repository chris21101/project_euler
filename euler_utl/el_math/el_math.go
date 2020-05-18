package el_math

import (
	"fmt"
	"math"

	"project_euler/logging"
)

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// Function Sumdivisibleby calculate the sum of all multiples                  +
// from nat_num <= nat_taget_num                                               +
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Sumdivisibleby(nat_num int, nat_taget_num int, l *logging.Logger) int {
	p := math.Floor(float64(nat_taget_num / nat_num))
	result := nat_num * (int(p) * (int(p) + 1)) / 2
	l.Log(logging.Debug, fmt.Sprintf("Sum of all the multiples of %d <= %d is %d", nat_num, nat_taget_num, result))
	return result
}
