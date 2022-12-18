package httpclient

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type HttpResponsePayload struct {
	StatusCode int         `json:"status_code"`
	Body       string      `json:"body,omitempty"`
	Header     http.Header `json:"header,omitempty"`
}

type HttpRequestPayload struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	Body   string `json:"body,omitempty"`
}

type HttpDoer interface {
	Do(req *http.Request) ([]byte, int, http.Header, error)
	SetTimeout(duration string) error
}

type ProxiedHttpDoer interface {
	HttpDoer
	WithProxyAuthHeader(value string) (ProxiedHttpDoer, error)
}

type httpDoer struct {
	client *http.Client
}

type proxiedHttpDoer struct {
	httpDoer
}

func NewDoer() HttpDoer {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	return newDoer(client)
}

func newDoer(client *http.Client) HttpDoer {
	return &httpDoer{client: client}
}

// SetTimeout sets timeout tolerance for the next http client call via Do()
func (d *httpDoer) SetTimeout(duration string) error {
	timeoutDuration, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}
	d.client.Timeout = timeoutDuration
	return nil
}

// The main logic of *httpDoer.Do()
func (d *httpDoer) do(req *http.Request) ([]byte, int, error, HttpRequestPayload, HttpResponsePayload) {
	reqObj := HttpRequestPayload{
		Method: req.Method,
		URL:    req.URL.String(),
	}
	resObj := HttpResponsePayload{}

	reqBodyBytes := copyBodyBytes(req)
	resp, err := d.client.Do(req)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() { // https://stackoverflow.com/questions/23494950/specifically-check-for-timeout-error
			return nil, 500, err, reqObj, resObj
		}

		return nil, 0, err, reqObj, resObj
	}
	defer resp.Body.Close()

	resObj.StatusCode = resp.StatusCode
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err, reqObj, resObj
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes))
	resObj.Header = resp.Header

	return respBytes, resp.StatusCode, err, reqObj, resObj
}

func (d *httpDoer) Do(req *http.Request) ([]byte, int, http.Header, error) {
	respBytes, statusCode, err, _, response := d.do(req)

	// sendHTTPLogToKafka is used to log request and response to Kafka,
	// so that api health can be inferred.
	// It is still called even though the `err` value from d.do() function is NOT nil.
	// This is because we want to log the case where the response is 500 (server error),
	// in this case, the `err` is not nil, and we want to monitor this case.
	// produce message to kafka in a separate goroutine so function can return immediately without waiting for producing kafka message to finish
	// go sendHTTPLogToKafka(ctx, reqObj, resObj)

	return respBytes, statusCode, response.Header, err
}

func copyBodyBytes(req *http.Request) []byte {
	if req.Body == nil {
		return nil
	}
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}
