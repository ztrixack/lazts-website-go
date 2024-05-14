package markdown

import (
	"io"

	"github.com/stretchr/testify/mock"
)

var _ Moduler = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

func (m *Mock) Convert(source []byte, wr io.Writer) error {
	args := m.Called(source, wr)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (m *Mock) Metadata(source []byte, result interface{}) error {
	args := m.Called(source, result)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (m *Mock) ReadFile(domain string, name string) ([]byte, error) {
	args := m.Called(domain, name)
	return args.Get(0).([]byte), args.Error(1)
}
