package tools

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"
)

func loggerInit() {
	//初始化日志
	log.SetReportCaller(true) //需要设置这个为true
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:03:04",
		ForceColors:     true,
		FullTimestamp:   true,
		DisableQuote:    true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//处理函数名
			fs := strings.Split(frame.Function, ".")
			fun := ""
			if len(fs) > 0 {
				fun = fs[len(fs)-1]
			}
			fileName := path.Base(frame.File)
			return fmt.Sprintf("[\033[1;34m%s\033[0m]", fun), fmt.Sprintf("[%s:%d]", fileName, frame.Line)
		},
	})
}
