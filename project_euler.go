package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"project_euler/euler_utl"
	"project_euler/logging"
)

const version string = "v0.0.2"
const project_name = "EulerGo"
const max_problem_num = 2

func print_result(pnum int, result uint64, text string) {
	fmt.Println(":-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) \n")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("ProjectEuler.net Problem %d\n\n", pnum)
	fmt.Println(text)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("%s %d is %d %s", "\nResult of the problem", pnum, result, "\n")
	fmt.Println("\n:-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) :-) ")
}

func print_version(name string) {
	fmt.Printf("Program: %s Version: %s \n", path.Base(name), version)
	//fmt.Printf(" %s\n", version)
	fmt.Printf("Max implemented problem number is %d\n", max_problem_num)
}

func main() {

	argsWithProg := os.Args

	numbPtr := flag.Int("problemnum", 1, "Number of the Eulerprojectproblem")
	wordPtr := flag.String("loglevel", logging.DefaultLevel, "Set loglevel [DEBUG, INFO, WARNING, ERROR]")
	versionPtr := flag.Bool("version", false, "Print version and exit the program")
	flag.Parse()

	if *versionPtr {
		print_version(argsWithProg[0])
		os.Exit(0)
	}

	logger, err := logging.New(time.RFC3339, strings.ToUpper(*wordPtr))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmdwithargs := strings.Join(argsWithProg, " ")
	logger.Log(logging.Debug, "Start : "+cmdwithargs)
	logger.Log(logging.Debug, fmt.Sprintf("Project %s %s", project_name, version))
	logger.Log(logging.Info, fmt.Sprintf("Max implemented problem number is %d", max_problem_num))

	switch *numbPtr {
	case 1:
		result, problem_text, _ := euler_utl.P1(logger)
		print_result(1, uint64(result), problem_text)
	case 2:
		result, problem_text, _ := euler_utl.P2(logger)
		print_result(2, result, problem_text)
	default:
		logger.Log(logging.Error, fmt.Sprintf("%s %d %s", "Problem ", *numbPtr, " not yet implemented"))
		logger.Log(logging.Error, fmt.Sprintf("Max implemented problem number is %d", max_problem_num))
		os.Exit(1)
	}
	logger.Log(logging.Debug, "Stop : "+cmdwithargs)
}
