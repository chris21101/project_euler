package euler_utl

import (
	"fmt"

	"project_euler/euler_utl/el_math"

	"project_euler/logging"
)

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func P1(l *logging.Logger) (int, string, error) {

	pnum := 1
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Start with Problem ", pnum))
	problem_text := `If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. 
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.`

	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	sum3 := el_math.Sumdivisibleby(3, 999, l)
	sum5 := el_math.Sumdivisibleby(5, 999, l)
	// 15 is the first number that is a multiple of 3 and 5.
	// The sum of multiples of 15 must therefore be subtracted from the sum of 3 and 5
	// Otherwise they would count twice
	sum15 := el_math.Sumdivisibleby(15, 999, l)
	result := sum3 + sum5 - sum15
	l.Log(logging.Info, fmt.Sprintf("%s %d and % d lower %d = %d", "Sum of multiples from ", 3, 5, 1000, result))
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return result, problem_text, nil
}
