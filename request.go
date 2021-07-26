package mundipagg

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BASEURL string = "https://api.mundipagg.com/core/v1/"

func Do(method string, data interface{}, secretKey string, uuid string, url string) ([]byte, error) {
	postData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	token := base64.StdEncoding.EncodeToString([]byte(secretKey + ":"))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Content-Type", "application/json")
	if uuid != "" {
		req.Header.Set("Idempotency-Key", uuid)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, fmt.Errorf("invalid status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
