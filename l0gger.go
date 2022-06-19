package l0gger

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

// CallerPrettyfier for Log Formater
func prettyfier(f *runtime.Frame) (string, string) {
	_, filename := path.Split(f.File)
	filename = fmt.Sprintf("%s:%d", filename, f.Line)
	return "", filename
}

var (
	// Log Formaters
	JsonFormater = &logrus.JSONFormatter{CallerPrettyfier: prettyfier}
	TextFormater = &logrus.TextFormatter{CallerPrettyfier: prettyfier}
)

// Log Levels
const (
	DebugLvl   = logrus.DebugLevel
	InfoLvl    = logrus.InfoLevel
	WarningLvl = logrus.WarnLevel
	ErrorLvl   = logrus.ErrorLevel
	FatalLvl   = logrus.FatalLevel
)

var СonsoleOut = os.Stdout

// New Logger
func New(outType interface{}, formater logrus.Formatter, lvl logrus.Level) (*logrus.Logger, error) {
	var log = logrus.New()

	switch out := outType.(type) {
	case string:
		file, err := os.OpenFile(out, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return nil, err
		}
		log.SetOutput(file)
	case nil:
		log.SetOutput(os.Stdout)
	default:
		log.SetOutput(out.(io.Writer))
	}

	// Set Formater
	log.SetFormatter(formater)

	log.SetReportCaller(true)

	// Set Log Levels
	log.SetLevel(lvl)

	return log, nil
}
