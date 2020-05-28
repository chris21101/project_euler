// euler_math_test project euler_math_test.go
package el_math

import (
	"fmt"
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

func TestKgt(t *testing.T) {

	kgt := Kgt(0)
	if kgt != 0 {
		t.Errorf("Error: euler_math.Kgt(%d) = %d and not (%f)", 0, 0, kgt)
	}

	kgt = Kgt(5)
	if kgt != 1 {
		t.Errorf("Error: euler_math.Kgt(%d) = %d and not (%f)", 5, 1, kgt)
	}

	kgt = Kgt(10)
	if kgt != 2 {
		t.Errorf("Error: euler_math.Kgt(%d) = %d and not (%f)", 10, 2, kgt)
	}
	kgt = Kgt(15)
	if kgt != 3 {
		t.Errorf("Error: euler_math.Kgt(%d) = %d and not (%f)", 15, 3, kgt)
	}

}

func TestIs_Prime_number(t *testing.T) {
	test_num := 19.0
	is_prime := Is_Prime_Number(test_num)
	if is_prime == false {
		t.Errorf("Error: euler_math.Is_Prime_Number(%f) is a prime_number", test_num)
	}
	test_num = 222.0
	is_prime = Is_Prime_Number(test_num)
	if is_prime == true {
		t.Errorf("Error: euler_math.Is_Prime_Number(%f) is not a prime_number", test_num)
	}
	test_num = 550.0
	is_prime = Is_Prime_Number(test_num)
	if is_prime == true {
		t.Errorf("Error: euler_math.Is_Prime_Number(%f) is not a prime_number", test_num)
	}
	test_num = 10000000097.0
	is_prime = Is_Prime_Number(test_num)
	if is_prime == false {
		t.Errorf("Error: euler_math.Is_Prime_Number(%f) is a prime_number", test_num)
	}
}

func TestPrime_Factorization(t *testing.T) {

	var next_num_str string

	// Test 10 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	result_str := "NULL"
	test_num := 10.0
	test_string := "2,5"

	facts := Prime_Factorization(10)

	for i := 0; i < len(facts); i++ {
		next_num_str = fmt.Sprintf("%.0f", facts[i])
		if result_str == "NULL" {
			result_str = next_num_str
		} else {
			result_str = result_str + "," + next_num_str
		}
	}

	if result_str != test_string {
		t.Errorf("Error: euler_math.Prime_Factorization(%.0f) = %s and not %s", test_num, test_string, result_str)
	}

	// Test 25 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	test_num = 25.0
	test_string = "5,5"
	result_str = "NULL"
	facts = Prime_Factorization(test_num)

	for i := 0; i < len(facts); i++ {
		next_num_str = fmt.Sprintf("%.0f", facts[i])
		if result_str == "NULL" {
			result_str = next_num_str
		} else {
			result_str = result_str + "," + next_num_str
		}
	}

	if result_str != test_string {
		t.Errorf("Error: euler_math.Prime_Factorization(%.0f) = %s and not %s", test_num, test_string, result_str)
	}

	// Test 13195 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	test_num = 13195.0
	test_string = "5,7,13,29"
	result_str = "NULL"
	facts = Prime_Factorization(test_num)

	for i := 0; i < len(facts); i++ {
		next_num_str = fmt.Sprintf("%.0f", facts[i])
		if result_str == "NULL" {
			result_str = next_num_str
		} else {
			result_str = result_str + "," + next_num_str
		}
	}

	if result_str != test_string {
		t.Errorf("Error: euler_math.Prime_Factorization(%.0f) = %s and not %s", test_num, test_string, result_str)
	}

	// Test 25999999990  ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	test_num = 25999999990.0
	test_string = "2,5,109,1049,22739"
	result_str = "NULL"
	facts = Prime_Factorization(test_num)

	for i := 0; i < len(facts); i++ {
		next_num_str = fmt.Sprintf("%.0f", facts[i])
		if result_str == "NULL" {
			result_str = next_num_str
		} else {
			result_str = result_str + "," + next_num_str
		}
	}

	if result_str != test_string {
		t.Errorf("Error: euler_math.Prime_Factorization(%.0f) = %s and not %s", test_num, test_string, result_str)
	}
}

func TestGet_Reverse_Num(t *testing.T) {
	test_num := 123456789.0
	result := Get_Reverse_Num(test_num)

	if result != 987654321.0 {
		t.Errorf("Error: euler_math.Get_Reverse_Num(%.0f) = %.0f and not %0.f", test_num, 987654321.0, result)
	}

	test_num = 4567.0
	result = Get_Reverse_Num(test_num)

	if result != 7654.0 {
		t.Errorf("Error: euler_math.Get_Reverse_Num(%.0f) = %.0f and not %0.f", test_num, 7654.0, result)
	}

}

func TestIs_Pallindrom(t *testing.T) {
	if !Is_Palindrom(9009.0) {
		t.Errorf("Error: euler_math.Is_Palindrom(9009) is a pallindrom")
	}
}
