package tranquility

import (
	"io/ioutil"
	"net/http"
	"strings"
)

//Action provides an abstraction for Http Request basic parameters
type Action struct {
	Method     string
	URL        string
	Headers    map[string]string
	Parameters map[string]string
	Body       string
}

//Response provides an abstraction for a parsed HTTP response
type Response struct {
	Status        string
	StatusCode    int
	Header        http.Header
	Body          string
	ContentLength int64
	Request       *http.Request
}

//Run perfoms a HTTP response based on the parameters of the action and parses the result into a Response
func (action Action) Run() (*Response, error) {
	req, err := http.NewRequest(action.Method, action.URL, strings.NewReader(action.Body))
	if err != nil {
		return nil, err
	}
	if action.Headers != nil {
		for k, v := range action.Headers {
			req.Header.Add(k, v)
		}
	}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	response := &Response{Status: res.Status, StatusCode: res.StatusCode, Header: res.Header, Body: string(body),
		ContentLength: res.ContentLength, Request: res.Request}
	return response, nil

}
