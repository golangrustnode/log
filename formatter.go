package log

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"runtime"
	"strconv"
	"strings"
)

type CustomFormatter struct {
	logrus.JSONFormatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Adjusting the depth to get the caller's caller
	data := make(logrus.Fields, len(entry.Data)+5)

	// 拷贝原始的 entry.Data 字段
	for k, v := range entry.Data {
		data[k] = v
	}

	// 处理Caller
	if entry.HasCaller() {
		/*
			data["file"] = trimPath(entry.Caller.File) + ":" + getLine(entry)
			data["func"] = trimFunc(entry.Caller.Function)
		*/
		if pc, file, line, ok := runtime.Caller(9); ok {
			stackInfo := trimPath(file) + ":" + strconv.Itoa(line)

			if fn := runtime.FuncForPC(pc); fn != nil {
				stackInfo = stackInfo + "  " + trimFunc(fn.Name())
			}
			data["caller_file"] = stackInfo

		}
	}
	// 设置基础字段
	data["time"] = entry.Time.Format(f.TimestampFormat)
	data["level"] = entry.Level.String()
	data["msg"] = entry.Message
	// Marshal为JSON
	var serialized []byte
	var err error
	if f.PrettyPrint {
		serialized, err = json.MarshalIndent(data, "", "  ")
	} else {
		serialized, err = json.Marshal(data)
	}
	return append(serialized, '\n'), err
}

func trimPath(fullPath string) string {
	const delimiter = "golangrustnode/"
	if idx := strings.Index(fullPath, delimiter); idx != -1 {
		return fullPath[idx+len(delimiter):]
	}
	return fullPath // 找不到就只取文件名
}

// trimFunc 裁剪函数路径，去掉 github.com 之前的部分
func trimFunc(fullFunc string) string {
	const delimiter = "golangrustnode/"
	if idx := strings.Index(fullFunc, delimiter); idx != -1 {
		return fullFunc[idx+len(delimiter):]
	}
	return fullFunc // 找不到就只取函数名
}

func getLine(entry *logrus.Entry) string {
	return strconv.Itoa(entry.Caller.Line)
}
