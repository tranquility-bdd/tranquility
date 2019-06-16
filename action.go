package tranquility

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"text/template"
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

var parser = template.New("env-parser")

//parseString is an auxiliary function to replace templates in a given string
func parseString(toParsed string) string {
	var parsed bytes.Buffer
	tmpl, err := parser.Parse(toParsed)
	if err != nil {
		panic(err)
	}
	tmpl.Option("missingkey=error")
	err = tmpl.Execute(&parsed, Env)
	if err != nil {
		panic(err)
	}
	return parsed.String()
}

//setup replaces all template variables in URL, Headers and Body preparing the HTTP request to be executed
func (action Action) setup() (*http.Request, error) {
	baseURL, err := url.Parse(parseString(action.URL))
	if err != nil {
		return nil, err
	}
	if action.Parameters != nil {
		params := url.Values{}
		for k, v := range action.Parameters {
			params.Add(k, parseString(v))
		}
		baseURL.RawQuery = params.Encode()
	}
	req, err := http.NewRequest(action.Method, baseURL.String(), strings.NewReader(parseString(action.Body)))
	if err != nil {
		return nil, err
	}
	if action.Headers != nil {
		for k, v := range action.Headers {
			req.Header.Add(k, parseString(v))
		}
	}
	return req, nil
}

//Run perfoms a HTTP response based on the parameters of the action and parses the result into a Response
func (action Action) Run() (*Response, error) {
	req, err := action.setup()
	if err != nil {
		return nil, err
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
