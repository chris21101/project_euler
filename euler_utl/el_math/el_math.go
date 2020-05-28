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

/*
   FUNCTION prime_factorization_array(
       p_num IN   NUMBER
   )RETURN t_number_ar AS

       v_list         t_number_ar;
       v_start_prim   NUMBER := 2;
       v_index        INTEGER := 0;
       v_kgt          NUMBER := 0;
       v_rest         NUMBER := 0;
   BEGIN
       v_rest := p_num;
       LOOP
           v_kgt := kgt(v_rest);

           -- v_rest is Prim Number
           IF v_kgt = 1 THEN
               v_list(v_index):= v_rest;
               RETURN v_list;
           ELSE
               v_list(v_index):= v_kgt;
               v_rest := v_rest / v_kgt;
           END IF;

           v_index := v_index + 1;
       END LOOP;

       RETURN v_list;
   END prime_factorization_array;
*/
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
