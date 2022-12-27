package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func HTTPMethod(t *testing.T, method string, endpoint string, accessToken string, body io.Reader, expectedStatusCode int) []byte {
	req, err := http.NewRequest(method, endpoint, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	Assert(t, err == nil, "HTTP client creation should not fail: ", err)

	if accessToken != "" {
		req.Header.Add("Authorization", "Bearer "+accessToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	Assert(t, err == nil, "Request to "+endpoint+" should not fail: ", err)

	respbody := ReadResponseBody(t, resp)
	Assert(t, resp.StatusCode == expectedStatusCode, fmt.Sprintf("HTTP response code should be %d but got %d instead for "+endpoint, expectedStatusCode, resp.StatusCode)+"\n"+string(respbody))

	return respbody
}

func ReadResponseBody(t *testing.T, resp *http.Response) []byte {
	Assert(t, resp.Body != nil, "Response should not be nil")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	Assert(t, err == nil, "Response reading should not fail: ", err)

	return body
}

func UnmarshalJSON(t *testing.T, body []byte, value interface{}) {
	err := json.Unmarshal(body, value)
	bodystr := string(body)
	Assert(t, err == nil, "Response data should be a valid JSON: ", bodystr, err)
}

func MarshalJSON(t *testing.T, value interface{}) []byte {
	body, err := json.Marshal(value)
	Assert(t, err == nil, "Failed to marshal to JSON: ", err)
	return body
}

func Assert(t *testing.T, cond bool, msg ...interface{}) {
	if !cond {
		t.Error(msg...)
	}
}
