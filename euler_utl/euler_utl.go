package euler_utl

import (
	"fmt"
	"math"

	//"strings"

	"project_euler/euler_utl/el_math"
	"project_euler/logging"
)

/*
Template for functions
func P[Num]l *logging.Logger) (uint64, error) {
	pnum := [Num]
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Start with Problem ", pnum))
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	result := 1
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return uint64(result), nil
}
*/
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func P1(l *logging.Logger) (uint64, error) {

	pnum := 1
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Start with Problem ", pnum))

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

	return uint64(result), nil
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// Each new term in the Fibonacci sequence is generated by adding the previous two terms.
// By starting with 1 and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// By considering the terms in the Fibonacci sequence whose values do not exceed four million,
// find the sum of the even-valued terms.
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func P2(l *logging.Logger) (uint64, error) {

	pnum := 2
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Start with Problem ", pnum))

	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	const max_value = 4000000
	next_fib := 0.0
	term_1 := 1.0
	term_2 := 1.0
	result := 0.0
	// First fibonacci Number 1
	l.Log(logging.Info, fmt.Sprintf("Fibonacci Num: %d", 1))
	for next_fib < max_value {
		next_fib = term_1 + term_2
		l.Log(logging.Info, fmt.Sprintf("Fibonacci Num: %d", uint64(next_fib)))
		if math.Mod(float64(next_fib), float64(2)) == 0 {
			l.Log(logging.Debug, fmt.Sprintf("even-term: %d +%d = %d", uint(next_fib), uint(result), uint(result+next_fib)))
			result = result + next_fib

		}

		term_1 = term_2
		term_2 = next_fib
	}
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Sum of the even-valued terms in the Fibonacci sequence is ", uint64(result)))
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return uint64(result), nil
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//The prime factors of 13195 are 5, 7, 13 and 29.
//
//What is the largest prime factor of the number 600851475143 ?
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func P3(l *logging.Logger) (uint64, error) {
	pnum := 3
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Start with Problem ", pnum))

	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	var next_num_str string
	result_str := "NULL"

	results := el_math.Prime_Factorization(600851475143)
	for i := 0; i < len(results); i++ {
		next_num_str = fmt.Sprintf("%.0f", results[i])
		if result_str == "NULL" {
			result_str = next_num_str
		} else {
			result_str = result_str + "*" + next_num_str
		}
	}
	l.Log(logging.Info, fmt.Sprintf("Prime_Factorization: %d = %s", 600851475143, result_str))
	result := results[len(results)-1]
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return uint64(result), nil
}

func P4(l *logging.Logger) (uint64, error) {
	pnum := 4
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Start with Problem ", pnum))
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	var (
		max_palindrom float64 = 0
		a             float64 = 0
		b             float64 = 0
	)

	for a = 999; a >= 100; {

		for b = 999; b >= a; {
			if el_math.Is_Palindrom(a * b) {
				l.Log(logging.Debug, fmt.Sprintf("New Palindrom: %0.f * %.0f = %.0f", a, b, a*b))
				if (a * b) > max_palindrom {
					max_palindrom = a * b
					l.Log(logging.Info, fmt.Sprintf("New highest Palindrom: %0.f * %.0f = %.0f", a, b, max_palindrom))
				}
			}
			b = b - 1
		}

		a = a - 1
	}
	result := max_palindrom
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return uint64(result), nil
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
// Solution from the ProjectEuler blog/pdf
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func P5(l *logging.Logger) (uint64, error) {
	pnum := 5
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Start with Problem ", pnum))
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	N := float64(1.0)
	k := float64(20.0)
	check := true
	limit := math.Sqrt(k)
	nr := float64(0)
	var a [9]float64

	//List of primes
	p := [9]float64{2, 3, 5, 7, 11, 13, 17, 19, 23}

	for i := 0; p[i] <= k; i++ {
		a[i] = 1
		if check {
			if p[i] <= limit {
				a[i] = math.Floor(math.Log(float64(k)) / math.Log(float64(p[i])))
			} else {
				check = false
			}
		}
		nr = N * math.Pow(float64(p[i]), float64(a[i]))
		l.Log(logging.Info, fmt.Sprintf("%.0f = %.0f * %.0f ^ %.0f", nr, N, p[i], a[i]))
		N = nr
	}
	result := N
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return uint64(result), nil
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//The sum of the squares of the first ten natural numbers is,
//
//1^2+2^2+...+10^2=385
//The square of the sum of the first ten natural numbers is,
//
//(1+2+...+10)^2=55^2=3025
//Hence the difference between the sum of the squares of the first ten natural numbers
//and the square of the sum is 3025−385=2640.
//
//Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func P6(l *logging.Logger) (uint64, error) {
	pnum := 6
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Start with Problem ", pnum))
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	var (
		sum_nat      float64 = 0
		sum_sqrt_nat float64 = 0
	)
	for i := 1.0; i <= 100; i++ {
		sum_sqrt_nat = sum_sqrt_nat + math.Pow(i, 2.0)
		sum_nat = sum_nat + i
	}
	result := math.Abs(math.Pow(sum_nat, 2.0) - sum_sqrt_nat)

	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return uint64(result), nil
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//
//What is the 10001st prime number?`
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func P7(l *logging.Logger) (uint64, error) {
	pnum := 7
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Start with Problem ", pnum))
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	next_prime := 1.0
	for count := 0; count < 10001; count++ {
		next_prime = el_math.Get_Next_Prime(next_prime)
		l.Log(logging.Info, fmt.Sprintf("%.0f", next_prime))
	}
	result := next_prime
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return uint64(result), nil
}

func P8(l *logging.Logger) (uint64, error) {
	pnum := 8
	l.Log(logging.Info, fmt.Sprintf("%s %d", "Start with Problem ", pnum))
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	result := 1
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return uint64(result), nil
}
