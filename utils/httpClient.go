package utils

import (
	"net/http"
	"bytes"
	"encoding/json"

	"github.com/spf13/viper"
)

type HttpClient struct {
	WithToken bool
}

func (httpClient *HttpClient) Call(method string, url string, param interface{}) (resp *(http.Response), err error) {
	var params []byte
	if param != nil {
		params, _ = json.Marshal(param)
	}
	req, err := http.NewRequest(method, "http://localhost:8080" + url, bytes.NewBuffer(params))
	if err != nil {
		return resp, err
	}
	if httpClient.WithToken {
		token := viper.Get("token").(string)
		req.Header.Set("X-Termworld-Token", token)
	}
	client := new(http.Client)
	resp, err = client.Do(req)
	return resp, err
}
