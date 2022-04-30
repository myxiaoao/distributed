package log

import (
	"io/ioutil"
	stLog "log"
	"net/http"
	"os"
)

var log *stLog.Logger

type fileLog string

// 日志写入文件
func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)
	return f.Write(data)
}

// Run 自定义日志
func Run(destination string) {
	log = stLog.New(fileLog(destination), "go: ", stLog.LstdFlags) // 带 go 前缀和日期的日志
}

// RegisterHandlers 注册日志服务
func RegisterHandlers() {
	http.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		// 只处理 post 请求，写入日志
		case http.MethodPost:
			// 写入 body 信息
			msg, err := ioutil.ReadAll(request.Body)
			if err != nil || len(msg) == 0 {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(massage string) {
	log.Printf("%v\n", massage)
}
