package logger

import "github.com/stretchr/testify/mock"

var _ Logger = (*MockLog)(nil)

type MockLog struct {
	mock.Mock
}

func NewMockLogger() *MockLog {
	return &MockLog{}
}

func (l *MockLog) Config() config {
	return config{}
}

func (l *MockLog) D(format string, args ...interface{}) {
	// Do nothing
}

func (l *MockLog) I(format string, args ...interface{}) {
	// Do nothing
}

func (l *MockLog) W(format string, args ...interface{}) {
	// Do nothing
}

func (l *MockLog) E(format string, args ...interface{}) {
	// Do nothing
}

func (l *MockLog) C(format string, args ...interface{}) {
	// Do nothing
}

func (l *MockLog) Fields(fields ...interface{}) Logger {
	return &MockLog{}
}

func (l *MockLog) Err(err error) Logger {
	return &MockLog{}
}
