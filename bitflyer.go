package bitflyer

import (
	"encoding/json"
	"errors"
	"fmt"
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

type PageParams struct {
	// 結果の個数を指定します。省略した場合の値は 100 です。
	Count int
	// このパラメータに指定した値より小さい id を持つデータを取得します。
	Before int64
	// このパラメータに指定した値より大きい id を持つデータを取得します。
	After int64
}

func NewPageParams() *PageParams {
	return &PageParams{-1, -1, -1}
}

func (p *PageParams) QueryParams(params url.Values) {
	if p.Count > 0 {
		params.Add("count", fmt.Sprint(p.Count))
	}
	if p.Before > 0 {
		params.Add("before", fmt.Sprint(p.Before))
	}
	if p.After > 0 {
		params.Add("after", fmt.Sprint(p.After))
	}
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
