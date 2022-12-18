package modules

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"test-golang/utils/httpclient"
)

type Client interface {
	GetListBookByGenre(itemsPerPage int, page int, genre string) (*GetListBookByGenreResponse, int, error)
}

type client struct {
	httpDoer httpclient.HttpDoer
}

func NewClient(httpDoer httpclient.HttpDoer) Client {
	return &client{
		httpDoer: httpDoer,
	}
}

func (c *client) GetListBookByGenre(itemsPerPage int, page int, genre string) (*GetListBookByGenreResponse, int, error) {
	params := map[string]string{
		"details": "false",
		"ebooks":  "false",
		"limit":   strconv.Itoa(itemsPerPage),
		"offset":  strconv.Itoa(page),
	}

	responseByte, httpCode, err := c.callEndpoint("GET", strings.ToLower(genre)+".json", params, nil)
	if err != nil {
		return nil, httpCode, err
	}

	var resp GetListBookByGenreResponse

	err = json.Unmarshal(responseByte, &resp)
	if err != nil {
		return nil, httpCode, err
	}

	return &resp, httpCode, nil
}

func (c *client) callEndpoint(method, path string, params map[string]string, body interface{}) ([]byte, int, error) {
	var bodyReader io.Reader

	req, err := http.NewRequest(method, "http://openlibrary.org/subjects/"+path, bodyReader)
	if err != nil {
		return nil, 0, err
	}

	client := c.httpDoer

	req.Close = true

	respBytes, httpCode, _, err := client.Do(req)

	return respBytes, httpCode, nil
}
