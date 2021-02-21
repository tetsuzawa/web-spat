package server

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/getkin/kin-openapi/openapi3filter"
)

func TestServer_Run(t *testing.T) {
	const serverPrefix = "http://localhost:1991/v1"

	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "ping",
			args:    args{method: http.MethodGet, url: "/ping", body: nil},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, serverPrefix+tt.args.url, tt.args.body)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
			if err := HelperTestRequest(t, req); (err != nil) != tt.wantErr {
				t.Errorf("HelperTestRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func HelperTestRequest(t *testing.T, request *http.Request) error {
	s := NewServer()
	if err := s.Init(); err != nil {
		return fmt.Errorf("failed to init server -> %w", err)
	}
	ts := httptest.NewServer(s.e)
	defer ts.Close()

	openAPIRouter := openapi3filter.NewRouter().WithSwaggerFromFile("../../../docs/openapi.yaml")
	route, pathParams, err := openAPIRouter.FindRoute(request.Method, request.URL)
	if err != nil {
		return fmt.Errorf("FindRoute() -> %w", err)
	}

	u, err := url.Parse(ts.URL)
	if err != nil {
		return fmt.Errorf("url.Parse() -> %w", err)
	}
	request.URL.Scheme = u.Scheme
	request.URL.Host = u.Host
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("http.DefaultClient.Do() -> %w", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    request,
		PathParams: pathParams,
		Route:      route,
	}

	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 response.StatusCode,
		Header:                 response.Header,
		Body:                   response.Body,
	}
	responseValidationInput.SetBodyBytes(body)

	t.Logf("request = %+v", request)
	t.Logf("response = %+v", response)

	return openapi3filter.ValidateResponse(context.TODO(), responseValidationInput)
}
