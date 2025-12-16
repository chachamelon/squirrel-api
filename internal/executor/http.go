package executor

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func ExecuteRequest(method, url string, headers map[string]string, body interface{}) (*http.Response, []byte, error) {
	var buf *bytes.Buffer

	if body != nil {
		b, _ := json.Marshal(body)
		buf = bytes.NewBuffer(b)
	} else {
		buf = bytes.NewBuffer(nil)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody, nil
}
