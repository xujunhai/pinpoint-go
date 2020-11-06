package trace

import (
	"net/http"
	"strconv"
)

const (
	SamplingRateFalse = "s0"
	SamplingRateTrue  = "s1"

	NullHttpHeader    = 0
	InvalidHttpHeader = 1
	ValidHttpHeader   = 2

	HttpTraceId               = "Pinpoint-TraceID"
	HttpSpanId                = "Pinpoint-SpanID"
	HttpParentSpanId          = "Pinpoint-pSpanID"
	HttpSampled               = "Pinpoint-Sampled"
	HttpFlags                 = "Pinpoint-Flags"
	HttpParentApplicationName = "Pinpoint-pAppName"
	HttpParentApplicationType = "Pinpoint-pAppType"
	HttpHost                  = "Pinpoint-Host"
)

type HttpHeader struct {
	HttpType      int
	TransactionID string
	SpanID        int64
	PSpanID       int64
	Sampled       bool
	Flag          int64
	PAppName      string
	PAppType      int64
	PHostName     string
}

func NewTraceHttpHeader(header *http.Header) *HttpHeader {
	var traceHeader = new(HttpHeader)
	traceHeader.Sampled = true
	traceHeader.HttpType = ValidHttpHeader
	if header == nil {
		traceHeader.HttpType = NullHttpHeader
		return traceHeader
	}

	sampled := header.Get(HttpSampled)
	if sampled == SamplingRateFalse {
		traceHeader.Sampled = false
		return traceHeader
	}

	traceHeader.TransactionID = header.Get(HttpTraceId)
	if traceHeader.TransactionID == "" {
		traceHeader.HttpType = NullHttpHeader
		return traceHeader
	}

	spanID, err := strconv.ParseInt(header.Get(HttpSpanId), 10, 64)
	traceHeader.SpanID = spanID
	if err != nil {
		traceHeader.HttpType = InvalidHttpHeader
		return traceHeader
	}

	traceHeader.PSpanID, err = strconv.ParseInt(header.Get(HttpParentSpanId), 10, 64)
	if err != nil {
		traceHeader.HttpType = InvalidHttpHeader
		return traceHeader
	}

	traceHeader.Flag, err = strconv.ParseInt(header.Get(HttpFlags), 10, 64)
	traceHeader.PAppName = header.Get(HttpParentApplicationName)
	traceHeader.PAppType, err = strconv.ParseInt(header.Get(HttpParentApplicationType), 10, 64)
	traceHeader.PHostName = header.Get(HttpHost)
	return traceHeader
}
