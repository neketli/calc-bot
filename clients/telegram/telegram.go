package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	methodGetUpdates = "getUpdates"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host, token string) Client {
	return Client{
		host:     host,
		basePath: ClientPath(token),
		client:   http.Client{},
	}
}

func (c *Client) Updates(offset, limit int) ([]Update, error) {
	query := url.Values{}
	query.Add("offset", strconv.Itoa(offset))
	query.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(methodGetUpdates, query)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res.Result, nil

}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Can't do request: %w", err)
	}
	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Can't do request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Can't do request: %w", err)
	}
	return body, nil
}

func (c *Client) SendMessage() {

}

func ClientPath(token string) string {
	return "bot" + token
}
