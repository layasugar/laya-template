// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package edb

import (
	"net/http"

	"github.com/layasugar/laya/store/cm"
	"go.opentelemetry.io/otel/attribute"
)

const (
	tSpanName = "elasticsearch"
)

// Transport for tracing Elastic operations.
type Transport struct {
	rt http.RoundTripper
}

// Option signature for specifying options, e.g. WithRoundTripper.
type Option func(t *Transport)

// WithRoundTripper specifies the http.RoundTripper to call
// next after this transport. If it is nil (default), the
// transport will use http.DefaultTransport.
func WithRoundTripper(rt http.RoundTripper) Option {
	return func(t *Transport) {
		t.rt = rt
	}
}

// NewTransport specifies a transport that will trace Elastic
// and report back via OpenTracing.
func NewTransport(opts ...Option) *Transport {
	t := &Transport{}
	for _, o := range opts {
		o(t)
	}
	return t
}

// RoundTrip captures the request and starts an OpenTracing span
// for Elastic PerformRequest operation.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	_, span := cm.ParseSpanByCtx(req.Context(), tSpanName)
	if nil != span {
		span.SetAttributes(attribute.String("component", "go-elasticsearch/v7"))
		span.SetAttributes(attribute.String("http.url", req.URL.String()))
		span.SetAttributes(attribute.String("http.method", req.Method))
		span.SetAttributes(attribute.String("peer.hostname", req.URL.Hostname()))
		span.SetAttributes(attribute.String("peer.port", req.URL.Port()))
		defer span.End()
	}
	var (
		resp *http.Response
		err  error
	)
	if t.rt != nil {
		resp, err = t.rt.RoundTrip(req)
	} else {
		resp, err = http.DefaultTransport.RoundTrip(req)
	}
	if err != nil {
		if nil != span {
			span.RecordError(err)
		}
	}
	if resp != nil {
		if nil != span {
			span.SetAttributes(attribute.String("http.status_code", resp.Status))
		}
	}

	return resp, err
}
