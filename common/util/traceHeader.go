package util

type TraceHeader struct {
	id           string
	spanId       string
	parentSpanId string
	sampling     string
	flag         string
}
