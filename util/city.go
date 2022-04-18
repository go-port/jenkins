package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetLocalCity 获取当前网络对应的城市
func GetLocalCity() (string, error) {
	response, err := http.Get("https://restapi.amap.com/v3/ip?key=3279da073706b4846e9e90abd7523c0a")
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	rsp := struct {
		City string `json:"city"`
	}{}
	err = json.Unmarshal(data, &rsp)
	if err != nil {
		return "", err
	}
	rsp.City = strings.ReplaceAll(rsp.City, "市", "")
	if !strings.Contains(rsp.City, "新区") {
		rsp.City = strings.ReplaceAll(rsp.City, "区", "")
	}
	return rsp.City, nil
}
