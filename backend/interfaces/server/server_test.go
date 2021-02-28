package server

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/tetsuzawa/web-spat/infrastructure/persistence_mock"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tetsuzawa/web-spat/interfaces/server/handler"
	"github.com/tetsuzawa/web-spat/interfaces/server/openapi"
	"github.com/tetsuzawa/web-spat/usecase"
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
		{
			name:    "listExperimentsMDDActive",
			args:    args{method: http.MethodGet, url: "/experiment/mdd/active", body: nil},
			wantErr: false,
		},
		{
			name:    "listExperimentsMDDInactive",
			args:    args{method: http.MethodGet, url: "/experiment/mdd/inactive", body: nil},
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
	e := echo.New()
	e.Debug = true

	// routing
	h := handler.NewIntegratedHandler(
		*handler.NewExperimentsHandler(usecase.NewExperimentUseCase(persistence_mock.NewExperimentRepository())),
		*handler.NewUtilHandler(),
	)

	openapi.RegisterHandlersWithBaseURL(e, h, "/v1")

	// Middleware
	//e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatalf("Error loading swagger spec\n: %s", err)
	}

	e.Use(oapimiddleware.OapiRequestValidator(swagger))
	// routing

	openapi.RegisterHandlersWithBaseURL(e, h, "/v1")

	ts := httptest.NewServer(e)
	defer ts.Close()

	openAPIRouter := openapi3filter.NewRouter().WithSwaggerFromFile("../../../docs/openapi.yaml")
	route, pathParams, err := openAPIRouter.FindRoute(request.Method, request.URL)
	if err != nil {
		t.Logf("method:%v, request.URL:%v\n", request.Method, request.URL)
		return fmt.Errorf("FindRoute() -> %w", err)
	}

	u, err := url.Parse(ts.URL)
	if err != nil {
		return fmt.Errorf("url.Parse() -> %w", err)
	}
	request.URL.Scheme = u.Scheme
	request.URL.Host = u.Host
	t.Logf("%+v",request)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("http.DefaultClient.Do() -> %w", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	t.Log(string(body))
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

	if err := openapi3filter.ValidateResponse(context.TODO(), responseValidationInput); err != nil {
		t.Logf("request = %+v", request)
		t.Logf("response = %+v", response)
		return err
	}

	return nil
}
