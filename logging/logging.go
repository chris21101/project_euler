package logging

import (
	"errors"
	"fmt"

	//"strconv"
	"strings"
	"time"
)

type Logger struct {
	timeFormat string
	//{"DEBUG": 1, "INFO": 2, "WARNING": 3, "ERROR": 4}
	loglevel    string
	loglevelnum int
}

const Debug = "DEBUG"
const Info = "INFO"
const Warning = "WARNING"
const Error = "ERROR"
const DefaultLevel = "ERROR"

func get_loglevelnum(loglevel string) (int, error) {
	loglevel_up := strings.ToUpper(loglevel)

	loglevelnums := map[string]int{"DEBUG": 1, "INFO": 2, "WARNING": 3, "ERROR": 4}

	loglevelnum, ok := loglevelnums[loglevel_up]

	if !ok {
		return -1, errors.New("Loglevel " + loglevel_up + " not exists")
	} else {
		return loglevelnum, nil
	}

}

func New(timeFormat string, loglevel string) (*Logger, error) {
	loglevelnum, err := get_loglevelnum(loglevel)
	//fmt.Println("L.LOGLEVEL: " + loglevel)
	return &Logger{
		timeFormat:  timeFormat,
		loglevel:    loglevel,
		loglevelnum: loglevelnum}, err
}

func (l *Logger) Log(ll string, s string) error {
	mesg_levelnum, err := get_loglevelnum(ll)

	if mesg_levelnum < l.loglevelnum {

		if err != nil {
			fmt.Printf("%s %s\n", time.Now().Format(l.timeFormat), err)
			return err
		}
		return err
	} else {
		fmt.Printf("%s [%s] %s\n", time.Now().Format(l.timeFormat), ll, s)
	}
	return err
}
