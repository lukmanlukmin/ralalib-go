package logger

import (
	"github.com/sirupsen/logrus"
)

// Setup :nodoc:
func Setup(cfg Config, h logrus.Hook) {
	logrus.SetFormatter(&Formatter{
		ChildFormatter: &logrus.JSONFormatter{},
		Line:           true,
		Package:        false,
		File:           true,
		BaseNameOnly:   false,
	})

	if h != nil {
		logrus.AddHook(h)
	}

	//logrus.SetOutput(os.Stdout)

	if cfg.Debug {
		logrus.SetLevel(logrus.TraceLevel)
		return
	}

	logrus.SetLevel(logrus.InfoLevel)

}
