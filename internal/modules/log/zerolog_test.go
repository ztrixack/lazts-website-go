package log

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() (*bytes.Buffer, *module) {
	buffer := new(bytes.Buffer)
	c := &config{Level: 0, Writer: buffer}
	logger := New(c)
	return buffer, logger
}

func validate(lv string, msg string, log string) bool {
	format := fmt.Sprintf(`.* %s (\w+\.go:\d+) > %s`, lv, msg)
	re := regexp.MustCompile(format)
	return re.FindStringSubmatch(log) != nil
}

func TestZerologLogger(t *testing.T) {
	buffer, logger := setup()

	tests := []struct {
		lv  string
		msg string
		fn  func(format string, args ...interface{})
	}{
		{"DBG", "debug message", logger.D},
		{"INF", "info message", logger.I},
		{"WRN", "warn message", logger.W},
		{"ERR", "error message", logger.E},
	}

	for _, tt := range tests {
		buffer.Reset()
		tt.fn(tt.msg)

		if !validate(tt.lv, tt.msg, buffer.String()) {
			t.Errorf("expected '%s' and '%s' included but got '%s'", tt.lv, tt.msg, buffer.String())
		}
	}

}

func TestFalal(t *testing.T) {
	_, logger := setup()

	if os.Getenv("TEST_FALAL_LOG") == "1" {
		logger.C("fatal message")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFalal")
	cmd.Env = append(os.Environ(), "TEST_FALAL_LOG=1")
	err := cmd.Run()
	e, ok := err.(*exec.ExitError)
	if ok && e.ExitCode() == 1 {
		return
	}

	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestFields(t *testing.T) {
	buffer, logger := setup()
	fields := []any{"user", "john_doe", "id", 12345}
	expectedUserField := "user=john_doe"
	expectedIdField := "id=12345"

	extendedLogger := logger.Fields(fields...).(*module)
	extendedLogger.logger.Info().Msg("testing fields")

	output := buffer.String()
	assert.Contains(t, output, expectedUserField)
	assert.Contains(t, output, expectedIdField)
}

func TestErr(t *testing.T) {
	buffer, logger := setup()
	testError := errors.New("test error")
	expectedError := fmt.Sprintf("error=\"%s\"", testError.Error())

	errorLogger := logger.Err(testError).(*module)
	errorLogger.logger.Error().Msg("error occurred")

	output := buffer.String()
	assert.Contains(t, output, expectedError)
}
