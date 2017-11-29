package bitflyer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type BitFlyer struct {
	baseURL string
}

const baseURL = "https://api.bitflyer.com/v1/"

func New() *BitFlyer {
	return &BitFlyer{baseURL}
}

func (b *BitFlyer) get(path string, params url.Values, entity interface{}) error {
	u := b.baseURL + path
	if params != nil {
		u = u + "?" + params.Encode()
	}
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, entity)
}
