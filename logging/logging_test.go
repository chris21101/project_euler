// logging_test
package logging

import (
	"strings"
	"testing"
	"time"
)

func TestNew1(t *testing.T) {

	loglevel := "debug"
	_, err := New(time.RFC3339, strings.ToUpper(loglevel))

	if err != nil {
		t.Errorf("Error: New <%s> [%s]", loglevel, err)
	}

}

// Unknown loglevel creates an error
func TestNew2(t *testing.T) {
	loglevel := "unknown_level"
	_, err := New(time.RFC3339, strings.ToUpper(loglevel))

	if err == nil {
		t.Errorf("Error: New <%s> is not used [%s]", loglevel, err)
	}
}
