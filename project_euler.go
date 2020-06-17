package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"project_euler/euler_utl"
	"project_euler/euler_utl/el_localize"
	"project_euler/euler_utl/el_output"
	"project_euler/logging"
)

const version string = "v0.1.8"
const project_name = "EulerGo"
const max_problem_num = 8

func main() {

	argsWithProg := os.Args

	numbPtr := flag.Int("problemnum", max_problem_num, "Number of the Eulerprojectproblem")
	wordPtr := flag.String("loglevel", logging.DefaultLevel, "Set loglevel [DEBUG, INFO, WARNING, ERROR]")
	langPtr := flag.String("lang", "en", "Set language en|de")

	versionPtr := flag.Bool("version", false, "Print version and exit the program")
	flag.Parse()

	if *versionPtr {
		el_output.Print_txt_version(version, max_problem_num, argsWithProg[0])
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

	localizer, err := el_localize.New(*numbPtr, *langPtr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch *numbPtr {
	case 1:
		result, _ := euler_utl.P1(logger)
		el_output.Print_Loc_Text(localizer, *numbPtr, result)
	case 2:
		result, _ := euler_utl.P2(logger)
		el_output.Print_Loc_Text(localizer, *numbPtr, result)
	case 3:
		result, _ := euler_utl.P3(logger)
		el_output.Print_Loc_Text(localizer, *numbPtr, result)
	case 4:
		result, _ := euler_utl.P4(logger)
		el_output.Print_Loc_Text(localizer, *numbPtr, result)
	case 5:
		result, _ := euler_utl.P5(logger)
		el_output.Print_Loc_Text(localizer, *numbPtr, result)
	case 6:
		result, _ := euler_utl.P6(logger)
		el_output.Print_Loc_Text(localizer, *numbPtr, result)
	case 7:
		result, _ := euler_utl.P7(logger)
		el_output.Print_Loc_Text(localizer, *numbPtr, result)
	case 8:
		result, _ := euler_utl.P8(logger)
		el_output.Print_Loc_Text(localizer, *numbPtr, result)
	default:
		logger.Log(logging.Error, fmt.Sprintf("%s %d %s", "Problem ", *numbPtr, " not yet implemented"))
		logger.Log(logging.Error, fmt.Sprintf("Max implemented problem number is %d", max_problem_num))
		os.Exit(1)
	}
	logger.Log(logging.Debug, "Stop : "+cmdwithargs)
}
