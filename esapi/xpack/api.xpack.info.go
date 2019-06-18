// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func NewXPackInfoFunc(t Transport) XPackInfo {
	return func(o ...func(*XPackInfoRequest)) (*Response, error) {
		var r = XPackInfoRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/info-api.html.
//
type XPackInfo func(o ...func(*XPackInfoRequest)) (*Response, error)

// XPackInfoRequest configures the Xpack Info API request.
//
type XPackInfoRequest struct {
	Categories []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r XPackInfoRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_xpack"))
	path.WriteString("/_xpack")

	params = make(map[string]string)

	if len(r.Categories) > 0 {
		params["categories"] = strings.Join(r.Categories, ",")
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f XPackInfo) WithContext(v context.Context) func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.ctx = v
	}
}

// WithCategories - comma-separated list of info categories. can be any of: build, license, features.
//
func (f XPackInfo) WithCategories(v ...string) func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.Categories = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f XPackInfo) WithPretty() func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f XPackInfo) WithHuman() func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f XPackInfo) WithErrorTrace() func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f XPackInfo) WithFilterPath(v ...string) func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.FilterPath = v
	}
}
