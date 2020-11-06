package context

import (
	"pinpoint-go/common/trace"
	"pinpoint-go/config"
	"strconv"
	"sync/atomic"
)

type Context struct {
	Span          *Span
	TraceHeader   *trace.HttpHeader
	SpanEventList []*SpanEvent
	TraceID       *trace.TraceID
	Depth         int32
	Sequence      int32
}

func NewTraceContext(traceHeader *trace.HttpHeader, traceID *trace.TraceID, span *Span) *Context {
	return &Context{
		Span:        span,
		TraceHeader: traceHeader,
		TraceID:     traceID,
		Sequence:    -1,
		Depth:       0, //从1开始
	}
}

func (p *Context) StartTraceSpanEvent() *SpanEvent {
	spanEvent := NewSpanEvent(p.Span, p.TraceID.SpanID)
	spanEvent.SetSequence(int16(atomic.AddInt32(&p.Sequence, 1)))
	spanEvent.SetDepth(atomic.AddInt32(&p.Depth, 1))
	spanEvent.MarkStartTime()
	p.SpanEventList = append(p.SpanEventList, spanEvent)
	return spanEvent
}

func (p *Context) GetNextSpanInfo() map[string]string {
	m := make(map[string]string)
	nextSpanID := trace.NextSpanID(p.TraceID.SpanID, p.TraceID.PSpanID)

	m[trace.HttpTraceId] = trace.FmtTransIDString(p.TraceID.AgentID, p.TraceID.AgentStartTime, p.TraceID.TransSeq)
	m[trace.HttpParentSpanId] = strconv.FormatInt(p.TraceID.SpanID, 10)
	m[trace.HttpSpanId] = strconv.FormatInt(nextSpanID, 10)
	m[trace.HttpFlags] = "0"

	sampled := trace.SamplingRateTrue
	if !p.TraceHeader.Sampled {
		sampled = trace.SamplingRateFalse
	}

	m[trace.HttpSampled] = sampled
	m[trace.HttpParentApplicationName] = config.AppName()
	m[trace.HttpParentApplicationType] = strconv.Itoa(int(config.ServerType()))

	return m
}

func (p *Context) Finish() {
	p.Span.MarkEndTime()
	count := len(p.SpanEventList)
	for i := 0; i < count; i++ {
		p.Span.AddSpanEvent(p.SpanEventList[i].TSpanEvent)
	}
}
