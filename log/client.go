package log

import (
	"bytes"
	"distributed/registry"
	"fmt"
	stLog "log"
	"net/http"
)

func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stLog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stLog.SetFlags(0)
	stLog.SetOutput(&clientLogger{url: serviceURL})
}

type clientLogger struct {
	url string
}

func (c clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer([]byte(data))
	res, err := http.Post(c.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to send log message, service responded error")
	}

	return len(data), nil
}
