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

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// Function Kgt get the smallest common divisor                                +
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Kgt(nat_num float64) float64 {
	var next_div float64 = 2
	sqrt_nat_num := math.Sqrt(nat_num)
	if nat_num == 0 {
		return 0
	} else {
		for next_div <= sqrt_nat_num {
			if math.Mod(nat_num, next_div) == 0 {
				return next_div
			} else {
				next_div = next_div + 1
			}
		}
	}
	return 1
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//func is_prime_number: Function calculates whether a number is a prime number +
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Is_Prime_Number(num float64) bool {
	if num < 2 {
		return false
	}
	if num == 2 || num == 3 {
		return true
	}
	if Kgt(num) == 1 {
		return true
	}
	return false
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//func Prime_Factorization: Get a slice with the prime factores
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Prime_Factorization(num float64) []float64 {
	rest := num
	v_kgt := float64(0)
	facts := make([]float64, 0)
	for i := 0; i == 0; {

		v_kgt = Kgt(rest)
		if v_kgt == 1 {
			facts = append(facts, rest)
			i = 1
		} else {
			facts = append(facts, v_kgt)
			rest = rest / v_kgt
		}
	}

	return facts
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//func Get_Reverse_Numn: Get the reverse number 1234=> 4321q
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Get_Reverse_Num(num float64) float64 {
	var (
		n       float64 = 0.0
		reverse float64 = 0.0
	)

	for n = num; n > 0; {
		reverse = 10.0*reverse + math.Mod(n, 10.0)
		n = math.Floor(n / 10)
	}

	return reverse
}

func Is_Palindrom(num float64) bool {
	return num == Get_Reverse_Num(num)
}

func Get_Next_Prime(start_num float64) float64 {
	var (
		next_num float64 = 0.0
		result   float64 = 0.0
		exit     bool    = false
	)
	if math.Mod(start_num, 2) == 0 || start_num == 1 {
		next_num = start_num + 1
	} else {
		next_num = start_num + 2
	}

	for i := next_num; exit == false; {
		if Is_Prime_Number(i) == true {
			result = i
			exit = true
		}
		result = i
		i = i + 2
	}

	return result
}
