package el_output

import (
	"fmt"
	"path"
)

func Print_txt_result(pnum int, result uint64, text string) {
	fmt.Println(":-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) \n")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("ProjectEuler.net Problem %d\n\n", pnum)
	fmt.Println(text)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("%s %d is %d %s", "\nResult of the problem", pnum, result, "\n")
	fmt.Println("\n:-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) ")
}

func Print_txt_version(version string, max_problem_num int, name string) {
	fmt.Printf("Program: %s Version: %s \n", path.Base(name), version)
	//fmt.Printf(" %s\n", version)
	fmt.Printf("Max implemented problem number is %d\n", max_problem_num)
}
