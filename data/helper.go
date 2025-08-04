package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FromJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func MakeRequest(url string, v interface{}) (interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("the HTTP request failed with error %s", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("the HTTP request failed with error %s", err)
	}

	if err := json.Unmarshal(body, v); err != nil {
		return nil, fmt.Errorf("the HTTP request failed with error %s", err)
	}
	return v, nil
}