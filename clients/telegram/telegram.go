package telegram

import (
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: NewBasePath(token),
		client:   http.Client{},
	}
}

func NewBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = e.Wrap("can't do request", err) }()

	u := url.URL{
		Scheme: "https", Host: c.host, Path: path.Join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, e.Wrap(errMessage, err)
	}

	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, e.Wrap(errMessage, err)
	}

	return resp, nil
}

func (c *Client) SendMessage() {

}
