package simaland

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httputil"
	"time"
)

var (
	baseURL = "https://www.sima-land.ru/api/v3/"
)

func NewClient(opts ...Options) (*Client, error) {
	cl := &Client{
		mods: []RequestModificator{},
		http: &http.Client{
			Timeout: time.Second * 10,
		},
	}

	var err error

	for i := range opts {
		cl, err = opts[i](cl)
		if err != nil {
			return cl, err
		}
	}

	return cl, nil
}

type Logger interface {
	Logf(format string, args ...any)
}

type Client struct {
	http *http.Client
	log  Logger
	mods []RequestModificator
}

func (c *Client) Do(method, path string, body, data any, modificators ...RequestModificator) error {
	buffer := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buffer).Encode(body)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, baseURL+path, buffer)
	if err != nil {
		return err
	}

	modificators = append(modificators, c.mods...)

	modificators = append(modificators,
		SetHeader("Accept", "application/json"),
		SetHeader("Content-Type", "application/json"),
	)

	for i := range modificators {
		modificators[i](req)
	}

	resp, err := c.request(req)
	if err != nil {
		return err
	}

	err = tryError(resp)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(data)
}

func (c *Client) get(path string, object any, modificators ...RequestModificator) error {
	return c.Do(http.MethodGet, path, nil, object, modificators...)
}

func (c *Client) post(path string, body, object any, modificators ...RequestModificator) error {
	return c.Do(http.MethodPost, path, body, object, modificators...)
}

func (c *Client) put(path string, body, object any, modificators ...RequestModificator) error {
	return c.Do(http.MethodPost, path, body, object)
}

func (c *Client) delete(path string, body, object any, modificators ...RequestModificator) error {
	return c.Do(http.MethodDelete, path, body, object, modificators...)
}

// request Обрабатывает http запрос
func (c *Client) request(httpReq *http.Request) (*http.Response, error) {
	if c == nil {
		return nil, errors.New("client is nil")
	}

	if c.log != nil {
		data, err := httputil.DumpRequest(httpReq, true)
		if err != nil {
			return nil, err
		}

		c.log.Logf("[DEBUG] Requset \n%v", string(data))
	}

	httpResp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if c.log != nil {
		data, err := httputil.DumpResponse(httpResp, true)
		if err != nil {
			return nil, err
		}

		c.log.Logf("[DEBUG] Response \n%v", string(data))
	}

	return httpResp, nil
}
