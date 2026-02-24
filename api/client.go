package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	accessToken string
	httpClient  http.Client
}

func NewClient(token string) *Client {
	c := Client{accessToken: token, httpClient: http.Client{}}
	return &c
}

func (c *Client) UnrestrictLink(link string) (UnrestrictLinkResponse, error) {
	v := url.Values{}
	response := UnrestrictLinkResponse{}

	v.Set("link", link)

	req, err := http.NewRequest("POST", RD_URL+UNRESTRICT_PATH, strings.NewReader(v.Encode()))

	if err != nil {
		return response, err
	}

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)

	if err != nil {
		return response, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, nil
}
