package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}

type Client struct {
	APIKey string
}

type Request struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

func (c *Client) GetResponse(model, prompt string) (*Response, error) {
	request := &Request{
		Model: model,
		Input: prompt,
	}
	b, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/responses", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	r := &Response{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		panic(err)
	}
	return r, nil
}
