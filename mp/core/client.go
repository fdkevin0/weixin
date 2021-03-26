package core

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"net/url"
)

type Client struct {
	AccessToken string
}

func (clt Client) GetDataCommon(method string, path string, query url.Values, postBody interface{}, data interface{}) (errData Error, err error) {
	requestUrl := url.URL{
		Scheme: "https",
		Host:   "api.weixin.qq.com",
	}
	requestUrl.Path = path
	query.Add("access_token", clt.AccessToken)
	requestUrl.RawQuery = query.Encode()

	request := gorequest.New()

	switch method {
	case http.MethodGet:
		request.Get(requestUrl.String())
	case http.MethodPost:
		request.Post(requestUrl.String()).Send(postBody)
	default:
		request.Get(requestUrl.String())
	}

	resp, body, errs := request.EndStruct(errData)
	if errData.ErrorCode != 0 || resp.StatusCode != http.StatusOK || errs != nil {
		err = errors.New("")
	}
	if data != nil {
		err = json.Unmarshal(body, data)
		if err != nil {
			return
		}
	}
	return
}

func (clt Client) GetJson(path string, query url.Values, data interface{}) (errData Error, err error) {
	return clt.GetDataCommon(http.MethodGet, path, query, nil, data)
}

func (clt Client) PostJson(path string, query url.Values, postBody interface{}, data interface{}) (errData Error, err error) {
	return clt.GetDataCommon(http.MethodPost, path, query, postBody, data)
}
