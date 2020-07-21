package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Field :nodoc:
type Field struct {
	Key   string
	Value interface{}
}

// FieldFunc :nodoc:
type FieldFunc func(key string, value interface{}) *Field

// SetField :nodoc:
func SetField(k string, v interface{}) Field {
	return Field{
		Key:   k,
		Value: v,
	}

}

// SetMessageFormat :nodoc:
func SetMessageFormat(format string, args ...interface{}) interface{} {
	return fmt.Sprintf(format, args...)
}

func extract(args ...Field) map[string]interface{} {
	if len(args) == 0 {
		return nil
	}

	data := map[string]interface{}{}
	for _, fl := range args {
		data[fl.Key] = fl.Value
	}

	return data
}

// Error :nodoc:
func Error(arg interface{}, fl ...Field) {
	logrus.WithFields(extract(fl...)).Error(arg)
}

// Info :nodoc:
func Info(arg interface{}, fl ...Field) {
	logrus.WithFields(extract(fl...)).Info(arg)
}

// Debug :nodoc:
func Debug(arg interface{}, fl ...Field) {
	logrus.WithFields(extract(fl...)).Debug(arg)
}

// Fatal :nodoc:
func Fatal(arg interface{}, fl ...Field) {
	logrus.WithFields(extract(fl...)).Fatal(arg)
}

// Warn :nodoc:
func Warn(arg interface{}, fl ...Field) {
	logrus.WithFields(extract(fl...)).Warn(arg)
}

// Trace :nodoc:
func Trace(arg interface{}, fl ...Field) {
	logrus.WithFields(extract(fl...)).Trace(arg)
}
