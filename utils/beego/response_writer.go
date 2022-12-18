package beego

import (
	"bytes"
	"fmt"
	"net/http"

	beegoContext "github.com/astaxie/beego/context"
)

// ResponseWriter is a custom wrapper for beego.context.Response with added functionality of storing the http response for logging
type ResponseWriter struct {
	beegoContext.Response
	responseBody bytes.Buffer // buffer for storing the http response body
	proto        string       // http protocol info
}

// NewResponseWriter returns a http.ResponseWriter wrapping the given beego.context.Response from param
func NewResponseWriter(beegoResponse beegoContext.Response, proto string) *ResponseWriter {
	return &ResponseWriter{
		Response: beegoResponse,
		proto:    proto,
	}
}

// Write wraps the original Write method of beego.context.Response with the addition of writing the written bytes to buffer
func (r *ResponseWriter) Write(p []byte) (int, error) {
	_, err := r.Response.Write(p)
	if err != nil {
		return 0, err
	}
	// also write to the buffer
	return r.responseBody.Write(p)
}

// GetResponseLog returns the http response as string (including the http status code and headers) suitable for logging
func (r *ResponseWriter) GetResponseLog() (string, error) {
	headerBytes := &bytes.Buffer{}

	// compose http status in response log
	headerBytes.Write([]byte(fmt.Sprintf("%v %v %v\n", r.proto, r.Status, http.StatusText(r.Status))))

	// compose http header in response log
	err := r.Header().Write(headerBytes)
	if err != nil {
		return "", err
	}
	// return header and body as string
	return fmt.Sprintf("%v\n%v\n", headerBytes.String(), r.responseBody.String()), nil
}

// GetBody returns the http body as string
func (r *ResponseWriter) GetBody() string {
	return r.responseBody.String()
}
