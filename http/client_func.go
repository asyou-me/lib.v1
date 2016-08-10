package http

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func DoGet(urlStr string, params *url.Values) ([]byte, *http.Response, error) {
	if params != nil {
		urlStr += "?" + params.Encode()
	}
	resp, err := DefaultClient.Get(urlStr)
	if err != nil {
		return nil, resp, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp, err
	}
	return body, resp, err
}

func DoPostForm(urlStr string, params *url.Values) ([]byte, *http.Response, error) {
	resp, err := DefaultClient.PostForm(urlStr, *params)
	if err != nil {
		return nil, resp, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp, err
	}
	return body, resp, err
}

func DoPostJson(urlStr string, v interface{}) ([]byte, *http.Response, error) {
	var postReader io.Reader = nil

	data, err := json.Marshal(v)
	if err != nil {
		return nil, nil, err
	}
	postReader = bytes.NewReader(data)

	resp, err := DefaultClient.Post(urlStr, "application/json;charset=utf-8", postReader)
	if err != nil {
		return nil, nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp, err
	}
	return body, resp, err
}

func Do(method, urlStr string, v interface{}) ([]byte, *http.Response, error) {
	var postReader io.Reader = nil
	data, err := json.Marshal(v)
	if err != nil {
		return nil, nil, err
	}
	postReader = bytes.NewReader(data)

	req, err := http.NewRequest(method, urlStr, postReader)
	if err != nil {
		return nil, nil, err
	}

	resp, err := DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp, err
	}
	return body, resp, err
}

func DoPostXml(urlStr string, v interface{}) ([]byte, *http.Response, error) {
	var postReader io.Reader = nil

	data, err := xml.Marshal(v)
	if err != nil {
		return nil, nil, err
	}
	postReader = bytes.NewReader(data)

	resp, err := DefaultClient.Post(urlStr, "application/x-www-form-urlencoded", postReader)
	if err != nil {
		return nil, resp, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp, err
	}
	return body, resp, err
}
