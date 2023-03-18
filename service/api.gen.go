// Package service provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"github.com/arrodriguez/selling-partner-api-sdk/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

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
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
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

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetServiceJobs request
	GetServiceJobs(ctx context.Context, params *GetServiceJobsParams) (*http.Response, error)

	// GetServiceJobByServiceJobId request
	GetServiceJobByServiceJobId(ctx context.Context, serviceJobId string) (*http.Response, error)

	// AddAppointmentForServiceJobByServiceJobId request  with any body
	AddAppointmentForServiceJobByServiceJobIdWithBody(ctx context.Context, serviceJobId string, contentType string, body io.Reader) (*http.Response, error)

	AddAppointmentForServiceJobByServiceJobId(ctx context.Context, serviceJobId string, body AddAppointmentForServiceJobByServiceJobIdJSONRequestBody) (*http.Response, error)

	// RescheduleAppointmentForServiceJobByServiceJobId request  with any body
	RescheduleAppointmentForServiceJobByServiceJobIdWithBody(ctx context.Context, serviceJobId string, appointmentId string, contentType string, body io.Reader) (*http.Response, error)

	RescheduleAppointmentForServiceJobByServiceJobId(ctx context.Context, serviceJobId string, appointmentId string, body RescheduleAppointmentForServiceJobByServiceJobIdJSONRequestBody) (*http.Response, error)

	// CancelServiceJobByServiceJobId request
	CancelServiceJobByServiceJobId(ctx context.Context, serviceJobId string, params *CancelServiceJobByServiceJobIdParams) (*http.Response, error)

	// CompleteServiceJobByServiceJobId request
	CompleteServiceJobByServiceJobId(ctx context.Context, serviceJobId string) (*http.Response, error)
}

func (c *Client) GetServiceJobs(ctx context.Context, params *GetServiceJobsParams) (*http.Response, error) {
	req, err := NewGetServiceJobsRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetServiceJobByServiceJobId(ctx context.Context, serviceJobId string) (*http.Response, error) {
	req, err := NewGetServiceJobByServiceJobIdRequest(c.Endpoint, serviceJobId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) AddAppointmentForServiceJobByServiceJobIdWithBody(ctx context.Context, serviceJobId string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewAddAppointmentForServiceJobByServiceJobIdRequestWithBody(c.Endpoint, serviceJobId, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) AddAppointmentForServiceJobByServiceJobId(ctx context.Context, serviceJobId string, body AddAppointmentForServiceJobByServiceJobIdJSONRequestBody) (*http.Response, error) {
	req, err := NewAddAppointmentForServiceJobByServiceJobIdRequest(c.Endpoint, serviceJobId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) RescheduleAppointmentForServiceJobByServiceJobIdWithBody(ctx context.Context, serviceJobId string, appointmentId string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewRescheduleAppointmentForServiceJobByServiceJobIdRequestWithBody(c.Endpoint, serviceJobId, appointmentId, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) RescheduleAppointmentForServiceJobByServiceJobId(ctx context.Context, serviceJobId string, appointmentId string, body RescheduleAppointmentForServiceJobByServiceJobIdJSONRequestBody) (*http.Response, error) {
	req, err := NewRescheduleAppointmentForServiceJobByServiceJobIdRequest(c.Endpoint, serviceJobId, appointmentId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) CancelServiceJobByServiceJobId(ctx context.Context, serviceJobId string, params *CancelServiceJobByServiceJobIdParams) (*http.Response, error) {
	req, err := NewCancelServiceJobByServiceJobIdRequest(c.Endpoint, serviceJobId, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CompleteServiceJobByServiceJobId(ctx context.Context, serviceJobId string) (*http.Response, error) {
	req, err := NewCompleteServiceJobByServiceJobIdRequest(c.Endpoint, serviceJobId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewGetServiceJobsRequest generates requests for GetServiceJobs
func NewGetServiceJobsRequest(endpoint string, params *GetServiceJobsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/service/v1/serviceJobs")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.ServiceOrderIds != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "serviceOrderIds", *params.ServiceOrderIds); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ServiceJobStatus != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "serviceJobStatus", *params.ServiceJobStatus); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageToken != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "pageToken", *params.PageToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageSize != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "pageSize", *params.PageSize); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.SortField != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "sortField", *params.SortField); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.SortOrder != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "sortOrder", *params.SortOrder); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.CreatedAfter != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "createdAfter", *params.CreatedAfter); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.CreatedBefore != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "createdBefore", *params.CreatedBefore); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.LastUpdatedAfter != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "lastUpdatedAfter", *params.LastUpdatedAfter); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.LastUpdatedBefore != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "lastUpdatedBefore", *params.LastUpdatedBefore); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ScheduleStartDate != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "scheduleStartDate", *params.ScheduleStartDate); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ScheduleEndDate != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "scheduleEndDate", *params.ScheduleEndDate); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if queryFrag, err := runtime.StyleParam("form", true, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetServiceJobByServiceJobIdRequest generates requests for GetServiceJobByServiceJobId
func NewGetServiceJobByServiceJobIdRequest(endpoint string, serviceJobId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "serviceJobId", serviceJobId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/service/v1/serviceJobs/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAddAppointmentForServiceJobByServiceJobIdRequest calls the generic AddAppointmentForServiceJobByServiceJobId builder with application/json body
func NewAddAppointmentForServiceJobByServiceJobIdRequest(endpoint string, serviceJobId string, body AddAppointmentForServiceJobByServiceJobIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAddAppointmentForServiceJobByServiceJobIdRequestWithBody(endpoint, serviceJobId, "application/json", bodyReader)
}

// NewAddAppointmentForServiceJobByServiceJobIdRequestWithBody generates requests for AddAppointmentForServiceJobByServiceJobId with any type of body
func NewAddAppointmentForServiceJobByServiceJobIdRequestWithBody(endpoint string, serviceJobId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "serviceJobId", serviceJobId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/service/v1/serviceJobs/%s/appointments", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewRescheduleAppointmentForServiceJobByServiceJobIdRequest calls the generic RescheduleAppointmentForServiceJobByServiceJobId builder with application/json body
func NewRescheduleAppointmentForServiceJobByServiceJobIdRequest(endpoint string, serviceJobId string, appointmentId string, body RescheduleAppointmentForServiceJobByServiceJobIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewRescheduleAppointmentForServiceJobByServiceJobIdRequestWithBody(endpoint, serviceJobId, appointmentId, "application/json", bodyReader)
}

// NewRescheduleAppointmentForServiceJobByServiceJobIdRequestWithBody generates requests for RescheduleAppointmentForServiceJobByServiceJobId with any type of body
func NewRescheduleAppointmentForServiceJobByServiceJobIdRequestWithBody(endpoint string, serviceJobId string, appointmentId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "serviceJobId", serviceJobId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParam("simple", false, "appointmentId", appointmentId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/service/v1/serviceJobs/%s/appointments/%s", pathParam0, pathParam1)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewCancelServiceJobByServiceJobIdRequest generates requests for CancelServiceJobByServiceJobId
func NewCancelServiceJobByServiceJobIdRequest(endpoint string, serviceJobId string, params *CancelServiceJobByServiceJobIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "serviceJobId", serviceJobId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/service/v1/serviceJobs/%s/cancellations", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "cancellationReasonCode", params.CancellationReasonCode); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("PUT", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCompleteServiceJobByServiceJobIdRequest generates requests for CompleteServiceJobByServiceJobId
func NewCompleteServiceJobByServiceJobIdRequest(endpoint string, serviceJobId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "serviceJobId", serviceJobId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/service/v1/serviceJobs/%s/completions", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
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
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetServiceJobs request
	GetServiceJobsWithResponse(ctx context.Context, params *GetServiceJobsParams) (*GetServiceJobsResp, error)

	// GetServiceJobByServiceJobId request
	GetServiceJobByServiceJobIdWithResponse(ctx context.Context, serviceJobId string) (*GetServiceJobByServiceJobIdResp, error)

	// AddAppointmentForServiceJobByServiceJobId request  with any body
	AddAppointmentForServiceJobByServiceJobIdWithBodyWithResponse(ctx context.Context, serviceJobId string, contentType string, body io.Reader) (*AddAppointmentForServiceJobByServiceJobIdResp, error)

	AddAppointmentForServiceJobByServiceJobIdWithResponse(ctx context.Context, serviceJobId string, body AddAppointmentForServiceJobByServiceJobIdJSONRequestBody) (*AddAppointmentForServiceJobByServiceJobIdResp, error)

	// RescheduleAppointmentForServiceJobByServiceJobId request  with any body
	RescheduleAppointmentForServiceJobByServiceJobIdWithBodyWithResponse(ctx context.Context, serviceJobId string, appointmentId string, contentType string, body io.Reader) (*RescheduleAppointmentForServiceJobByServiceJobIdResp, error)

	RescheduleAppointmentForServiceJobByServiceJobIdWithResponse(ctx context.Context, serviceJobId string, appointmentId string, body RescheduleAppointmentForServiceJobByServiceJobIdJSONRequestBody) (*RescheduleAppointmentForServiceJobByServiceJobIdResp, error)

	// CancelServiceJobByServiceJobId request
	CancelServiceJobByServiceJobIdWithResponse(ctx context.Context, serviceJobId string, params *CancelServiceJobByServiceJobIdParams) (*CancelServiceJobByServiceJobIdResp, error)

	// CompleteServiceJobByServiceJobId request
	CompleteServiceJobByServiceJobIdWithResponse(ctx context.Context, serviceJobId string) (*CompleteServiceJobByServiceJobIdResp, error)
}

type GetServiceJobsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetServiceJobsResponse
}

// Status returns HTTPResponse.Status
func (r GetServiceJobsResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetServiceJobsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetServiceJobByServiceJobIdResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetServiceJobByServiceJobIdResponse
}

// Status returns HTTPResponse.Status
func (r GetServiceJobByServiceJobIdResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetServiceJobByServiceJobIdResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AddAppointmentForServiceJobByServiceJobIdResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *SetAppointmentResponse
}

// Status returns HTTPResponse.Status
func (r AddAppointmentForServiceJobByServiceJobIdResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AddAppointmentForServiceJobByServiceJobIdResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RescheduleAppointmentForServiceJobByServiceJobIdResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *SetAppointmentResponse
}

// Status returns HTTPResponse.Status
func (r RescheduleAppointmentForServiceJobByServiceJobIdResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RescheduleAppointmentForServiceJobByServiceJobIdResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CancelServiceJobByServiceJobIdResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CancelServiceJobByServiceJobIdResponse
}

// Status returns HTTPResponse.Status
func (r CancelServiceJobByServiceJobIdResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CancelServiceJobByServiceJobIdResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CompleteServiceJobByServiceJobIdResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CompleteServiceJobByServiceJobIdResponse
}

// Status returns HTTPResponse.Status
func (r CompleteServiceJobByServiceJobIdResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CompleteServiceJobByServiceJobIdResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetServiceJobsWithResponse request returning *GetServiceJobsResponse
func (c *ClientWithResponses) GetServiceJobsWithResponse(ctx context.Context, params *GetServiceJobsParams) (*GetServiceJobsResp, error) {
	rsp, err := c.GetServiceJobs(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetServiceJobsResp(rsp)
}

// GetServiceJobByServiceJobIdWithResponse request returning *GetServiceJobByServiceJobIdResponse
func (c *ClientWithResponses) GetServiceJobByServiceJobIdWithResponse(ctx context.Context, serviceJobId string) (*GetServiceJobByServiceJobIdResp, error) {
	rsp, err := c.GetServiceJobByServiceJobId(ctx, serviceJobId)
	if err != nil {
		return nil, err
	}
	return ParseGetServiceJobByServiceJobIdResp(rsp)
}

// AddAppointmentForServiceJobByServiceJobIdWithBodyWithResponse request with arbitrary body returning *AddAppointmentForServiceJobByServiceJobIdResponse
func (c *ClientWithResponses) AddAppointmentForServiceJobByServiceJobIdWithBodyWithResponse(ctx context.Context, serviceJobId string, contentType string, body io.Reader) (*AddAppointmentForServiceJobByServiceJobIdResp, error) {
	rsp, err := c.AddAppointmentForServiceJobByServiceJobIdWithBody(ctx, serviceJobId, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseAddAppointmentForServiceJobByServiceJobIdResp(rsp)
}

func (c *ClientWithResponses) AddAppointmentForServiceJobByServiceJobIdWithResponse(ctx context.Context, serviceJobId string, body AddAppointmentForServiceJobByServiceJobIdJSONRequestBody) (*AddAppointmentForServiceJobByServiceJobIdResp, error) {
	rsp, err := c.AddAppointmentForServiceJobByServiceJobId(ctx, serviceJobId, body)
	if err != nil {
		return nil, err
	}
	return ParseAddAppointmentForServiceJobByServiceJobIdResp(rsp)
}

// RescheduleAppointmentForServiceJobByServiceJobIdWithBodyWithResponse request with arbitrary body returning *RescheduleAppointmentForServiceJobByServiceJobIdResponse
func (c *ClientWithResponses) RescheduleAppointmentForServiceJobByServiceJobIdWithBodyWithResponse(ctx context.Context, serviceJobId string, appointmentId string, contentType string, body io.Reader) (*RescheduleAppointmentForServiceJobByServiceJobIdResp, error) {
	rsp, err := c.RescheduleAppointmentForServiceJobByServiceJobIdWithBody(ctx, serviceJobId, appointmentId, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseRescheduleAppointmentForServiceJobByServiceJobIdResp(rsp)
}

func (c *ClientWithResponses) RescheduleAppointmentForServiceJobByServiceJobIdWithResponse(ctx context.Context, serviceJobId string, appointmentId string, body RescheduleAppointmentForServiceJobByServiceJobIdJSONRequestBody) (*RescheduleAppointmentForServiceJobByServiceJobIdResp, error) {
	rsp, err := c.RescheduleAppointmentForServiceJobByServiceJobId(ctx, serviceJobId, appointmentId, body)
	if err != nil {
		return nil, err
	}
	return ParseRescheduleAppointmentForServiceJobByServiceJobIdResp(rsp)
}

// CancelServiceJobByServiceJobIdWithResponse request returning *CancelServiceJobByServiceJobIdResponse
func (c *ClientWithResponses) CancelServiceJobByServiceJobIdWithResponse(ctx context.Context, serviceJobId string, params *CancelServiceJobByServiceJobIdParams) (*CancelServiceJobByServiceJobIdResp, error) {
	rsp, err := c.CancelServiceJobByServiceJobId(ctx, serviceJobId, params)
	if err != nil {
		return nil, err
	}
	return ParseCancelServiceJobByServiceJobIdResp(rsp)
}

// CompleteServiceJobByServiceJobIdWithResponse request returning *CompleteServiceJobByServiceJobIdResponse
func (c *ClientWithResponses) CompleteServiceJobByServiceJobIdWithResponse(ctx context.Context, serviceJobId string) (*CompleteServiceJobByServiceJobIdResp, error) {
	rsp, err := c.CompleteServiceJobByServiceJobId(ctx, serviceJobId)
	if err != nil {
		return nil, err
	}
	return ParseCompleteServiceJobByServiceJobIdResp(rsp)
}

// ParseGetServiceJobsResp parses an HTTP response from a GetServiceJobsWithResponse call
func ParseGetServiceJobsResp(rsp *http.Response) (*GetServiceJobsResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetServiceJobsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetServiceJobsResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetServiceJobByServiceJobIdResp parses an HTTP response from a GetServiceJobByServiceJobIdWithResponse call
func ParseGetServiceJobByServiceJobIdResp(rsp *http.Response) (*GetServiceJobByServiceJobIdResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetServiceJobByServiceJobIdResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetServiceJobByServiceJobIdResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseAddAppointmentForServiceJobByServiceJobIdResp parses an HTTP response from a AddAppointmentForServiceJobByServiceJobIdWithResponse call
func ParseAddAppointmentForServiceJobByServiceJobIdResp(rsp *http.Response) (*AddAppointmentForServiceJobByServiceJobIdResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &AddAppointmentForServiceJobByServiceJobIdResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest SetAppointmentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseRescheduleAppointmentForServiceJobByServiceJobIdResp parses an HTTP response from a RescheduleAppointmentForServiceJobByServiceJobIdWithResponse call
func ParseRescheduleAppointmentForServiceJobByServiceJobIdResp(rsp *http.Response) (*RescheduleAppointmentForServiceJobByServiceJobIdResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &RescheduleAppointmentForServiceJobByServiceJobIdResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest SetAppointmentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCancelServiceJobByServiceJobIdResp parses an HTTP response from a CancelServiceJobByServiceJobIdWithResponse call
func ParseCancelServiceJobByServiceJobIdResp(rsp *http.Response) (*CancelServiceJobByServiceJobIdResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CancelServiceJobByServiceJobIdResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CancelServiceJobByServiceJobIdResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCompleteServiceJobByServiceJobIdResp parses an HTTP response from a CompleteServiceJobByServiceJobIdWithResponse call
func ParseCompleteServiceJobByServiceJobIdResp(rsp *http.Response) (*CompleteServiceJobByServiceJobIdResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CompleteServiceJobByServiceJobIdResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CompleteServiceJobByServiceJobIdResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}
