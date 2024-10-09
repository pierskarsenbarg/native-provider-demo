// Package code provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package code

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

const (
	BearerScopes = "Bearer.Scopes"
)

// Organisation defines model for Organisation.
type Organisation struct {
	Id   float32 `json:"id"`
	Name string  `json:"name"`
}

// PostOrganisationJSONBody defines parameters for PostOrganisation.
type PostOrganisationJSONBody struct {
	Name string `json:"name"`
}

// PostOrganisationJSONRequestBody defines body for PostOrganisation for application/json ContentType.
type PostOrganisationJSONRequestBody PostOrganisationJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetOrganisation request
	GetOrganisation(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostOrganisationWithBody request with any body
	PostOrganisationWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostOrganisation(ctx context.Context, body PostOrganisationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetOrganisationId request
	GetOrganisationId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetOrganisation(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetOrganisationRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostOrganisationWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostOrganisationRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostOrganisation(ctx context.Context, body PostOrganisationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostOrganisationRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetOrganisationId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetOrganisationIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetOrganisationRequest generates requests for GetOrganisation
func NewGetOrganisationRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/organisation")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostOrganisationRequest calls the generic PostOrganisation builder with application/json body
func NewPostOrganisationRequest(server string, body PostOrganisationJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostOrganisationRequestWithBody(server, "application/json", bodyReader)
}

// NewPostOrganisationRequestWithBody generates requests for PostOrganisation with any type of body
func NewPostOrganisationRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/organisation")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetOrganisationIdRequest generates requests for GetOrganisationId
func NewGetOrganisationIdRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/organisation/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetOrganisationWithResponse request
	GetOrganisationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOrganisationResponse, error)

	// PostOrganisationWithBodyWithResponse request with any body
	PostOrganisationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostOrganisationResponse, error)

	PostOrganisationWithResponse(ctx context.Context, body PostOrganisationJSONRequestBody, reqEditors ...RequestEditorFn) (*PostOrganisationResponse, error)

	// GetOrganisationIdWithResponse request
	GetOrganisationIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetOrganisationIdResponse, error)
}

type GetOrganisationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Result []Organisation `json:"result"`
	}
}

// Status returns HTTPResponse.Status
func (r GetOrganisationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOrganisationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostOrganisationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Result Organisation `json:"result"`
	}
}

// Status returns HTTPResponse.Status
func (r PostOrganisationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostOrganisationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetOrganisationIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Result Organisation `json:"result"`
	}
	JSON400 *struct {
		Error string `json:"error"`
	}
	JSON404 *struct {
		Error string `json:"error"`
	}
}

// Status returns HTTPResponse.Status
func (r GetOrganisationIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOrganisationIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetOrganisationWithResponse request returning *GetOrganisationResponse
func (c *ClientWithResponses) GetOrganisationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOrganisationResponse, error) {
	rsp, err := c.GetOrganisation(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetOrganisationResponse(rsp)
}

// PostOrganisationWithBodyWithResponse request with arbitrary body returning *PostOrganisationResponse
func (c *ClientWithResponses) PostOrganisationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostOrganisationResponse, error) {
	rsp, err := c.PostOrganisationWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostOrganisationResponse(rsp)
}

func (c *ClientWithResponses) PostOrganisationWithResponse(ctx context.Context, body PostOrganisationJSONRequestBody, reqEditors ...RequestEditorFn) (*PostOrganisationResponse, error) {
	rsp, err := c.PostOrganisation(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostOrganisationResponse(rsp)
}

// GetOrganisationIdWithResponse request returning *GetOrganisationIdResponse
func (c *ClientWithResponses) GetOrganisationIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetOrganisationIdResponse, error) {
	rsp, err := c.GetOrganisationId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetOrganisationIdResponse(rsp)
}

// ParseGetOrganisationResponse parses an HTTP response from a GetOrganisationWithResponse call
func ParseGetOrganisationResponse(rsp *http.Response) (*GetOrganisationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetOrganisationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Result []Organisation `json:"result"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostOrganisationResponse parses an HTTP response from a PostOrganisationWithResponse call
func ParsePostOrganisationResponse(rsp *http.Response) (*PostOrganisationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostOrganisationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Result Organisation `json:"result"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseGetOrganisationIdResponse parses an HTTP response from a GetOrganisationIdWithResponse call
func ParseGetOrganisationIdResponse(rsp *http.Response) (*GetOrganisationIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetOrganisationIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Result Organisation `json:"result"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			Error string `json:"error"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Error string `json:"error"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /organisation)
	GetOrganisation(ctx echo.Context) error

	// (POST /organisation)
	PostOrganisation(ctx echo.Context) error

	// (GET /organisation/{id})
	GetOrganisationId(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetOrganisation converts echo context to params.
func (w *ServerInterfaceWrapper) GetOrganisation(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetOrganisation(ctx)
	return err
}

// PostOrganisation converts echo context to params.
func (w *ServerInterfaceWrapper) PostOrganisation(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostOrganisation(ctx)
	return err
}

// GetOrganisationId converts echo context to params.
func (w *ServerInterfaceWrapper) GetOrganisationId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetOrganisationId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/organisation", wrapper.GetOrganisation)
	router.POST(baseURL+"/organisation", wrapper.PostOrganisation)
	router.GET(baseURL+"/organisation/:id", wrapper.GetOrganisationId)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xUTW/bMAz9Kwa3o1CnW0++LZchh23FdgxyUCwmURFLKkkPMAL990FSvtx0aLBuQy+W",
	"TZHU03t83kHru+AdOmFodsDtBjudX7/RWjvLWqx36TuQD0hiMe9ak54yBIQGXN8tkSAqcLrDsw0Wsm4N",
	"MSogfOwtoYFmnor3qQt1SPXLB2wlpzK2PVkZfiQw5bgpakI6IkwFyxI6NtiIBIip3rqVzyCsbNPOl6H6",
	"dD8DBT+RON8Gbm8mN5ME2Ad0Olho4GMOKQhaNvnM2j9hYI2SlkRDjs0MNPAZZcRUuikH77gA/zCZpKX1",
	"TtDlch3C1rY5uX7g0rnQfskyIffbXGUFuxx6T7iCBt7VJ+HqvWr1CEg8MqOJ9HAhwr73MwpEBQa5JRvK",
	"zeE7Sk+OK11tLUvlV9U5NTwSDZr5Sa75Ii6iguD5GebuPV9S99gjy9Sb4RWsXTeEv5u/kjeS8PavSHi9",
	"cn+u1AtSRDWe6npnTbx2tGcmu4N0h4LEubtNhybHHAzdFHOf4Av1qM7YeSrK4h8Z5n+wffCFJ1N+f3ev",
	"Ao9Enl4e3JJ2Db6pNtXeUgXd3VtC99VLtfK9M9l0Mf4KAAD//1h/Zr6MBgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
