package tools

import (
	"bytes"
	logrus "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"runtime"
)

var logger = &lumberjack.Logger{
	Filename:   "../../log/log.txt",
	MaxSize:    10, // megabytes
	MaxBackups: 3,
	MaxAge:     28, //days
}

func LoggerInit() {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:03:04",

		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//处理文件名
			fileName := path.Base(frame.File)
			return frame.Function, fileName
		},
	})
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, logger)
	logrus.SetOutput(fileAndStdoutWriter)
	//设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)
}

type OutputSplitter struct{}

func (splitter *OutputSplitter) Write(p []byte) (n int, err error) {
	// your logs filter logic here. For ex:
	if bytes.Contains(p, []byte("level=error")) {
		return os.Stderr.Write(p)
	}
	return os.Stdout.Write(p)
}
