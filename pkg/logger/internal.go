package logger

type Logger interface {
	Config() config
	D(format string, args ...interface{})
	I(format string, args ...interface{})
	W(format string, args ...interface{})
	E(format string, args ...interface{})
	C(format string, args ...interface{})
	Fields(kv ...interface{}) Logger
	Err(err error) Logger
}
