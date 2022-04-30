package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// AddService 注册服务
func AddService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}

	res, err := http.Post(ServersURL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("faild to registry service, registry service "+
			"responded with code %v", res.StatusCode)
	}

	return nil
}

// DelService 服务取消
func DelService(url string) error {
	req, err := http.NewRequest(http.MethodDelete, ServersURL, bytes.NewBuffer([]byte(url)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("faild to deregister service, registry service "+
			"responded with code %v", res.StatusCode)
	}

	return nil
}
