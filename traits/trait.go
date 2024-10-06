package traits

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type Response struct {
	Status      string `json:"status,omitempty"`
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

func SendNotification(url string, data interface{}) error {

	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}

	urls := strings.Split(url, ",")

	for _, url = range urls {

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshal))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("error send webhook, url: %s, code: %d, response: %s", url, resp.StatusCode, string(body))
			}
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func GetLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

func GetFile() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}
