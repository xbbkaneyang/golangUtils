package golangUtils

import (
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	red                       = "31"
	green                     = "32"
	yellow                    = "33"
	blue                      = "34"
	magenta                   = "35"
	cyan                      = "36"
	white                     = "37"
	LogDefaultLogFormat       = "[%time%] [%lvl%] %f% - %msg%"
	LogDefaultTimestampFormat = "2006-01-02 15:04:05.000"
)

type MyFormatter struct {
	logrus.TextFormatter
	LogFormat  string
	LoggerName string
}

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	// if output == "" {
	// 	output = defaultLogFormat
	// }

	level := strings.ToUpper(entry.Level.String())
	var levelColor string
	switch level {
	case "TRACE":
		levelColor = white
	case "DEBUG":
		levelColor = cyan
	case "INFO":
		levelColor = green
	case "WARNING":
		levelColor = yellow
	case "ERROR":
		levelColor = red
	default:
		levelColor = blue
	}
	msg := entry.Message
	if !f.DisableColors {
		output = "\x1b[" + levelColor + "m" + output
		msg = "\x1b[0m" + msg
	}
	output = strings.Replace(output, "%lvl%", level, 1)
	output = strings.Replace(output, "%f%", f.LoggerName, 1)
	timestampFormat := f.TimestampFormat
	// if timestampFormat == "" {
	// 	timestampFormat = defaultTimestampFormat
	// }
	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)
	output = strings.Replace(output, "%msg%", msg, 1)

	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}
	output += "\n"
	return []byte(output), nil
}

func GetLogger(loggerName string) *logrus.Logger {
	myFormatter := new(MyFormatter)
	myFormatter.LoggerName = loggerName
	myFormatter.TimestampFormat = LogDefaultTimestampFormat
	myFormatter.LogFormat = LogDefaultLogFormat
	myFormatter.DisableColors = true

	logger := &logrus.Logger{
		Out:       os.Stdout,
		Level:     logrus.TraceLevel,
		Formatter: myFormatter,
	}
	return logger
}
