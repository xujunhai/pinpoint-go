package context

import (
	"pinpoint-go/common/util"
	"pinpoint-go/proto/thrift/gen-go/trace"
)

type SpanEvent struct {
	TSpanEvent *trace.TSpanEvent
	Span       *Span
}

func NewSpanEvent(span *Span, spanId int64) *SpanEvent {
	if spanId == 0 || span == nil {
		return nil
	}

	spanEvent := &SpanEvent{
		TSpanEvent: trace.NewTSpanEvent(),
		Span:       span,
	}
	spanEvent.Span = span
	spanEvent.TSpanEvent.SpanId = &spanId

	return spanEvent
}

func (p *SpanEvent) MarkStartTime() {
	elapsed := int32(util.GetNowMSec() - p.Span.GetStartTime())
	p.TSpanEvent.StartElapsed = elapsed
}

func (p *SpanEvent) MarkEndTime() {
	elapsed := int32(util.GetNowMSec() - p.Span.GetStartTime())
	p.TSpanEvent.EndElapsed = elapsed
}

func (p *SpanEvent) SetApiID(apiID int32) {
	p.TSpanEvent.ApiId = &apiID
}

func (p *SpanEvent) SetServiceType(serviceType int16) {
	p.TSpanEvent.ServiceType = serviceType
}

func (p *SpanEvent) AddAnnotation(annotation *trace.TAnnotation) {
	p.TSpanEvent.Annotations = append(p.TSpanEvent.Annotations, annotation)
}

func (p *SpanEvent) SetEndPoint(endPoint string) {
	p.TSpanEvent.EndPoint = &endPoint
}

func (p *SpanEvent) SetRPC(rpc string) {
	p.TSpanEvent.RPC = &rpc
}

func (p *SpanEvent) SetExceptionInfo(errID int32, errMsg string) {
	value := trace.NewTIntStringValue()
	value.IntValue = errID
	value.StringValue = &errMsg
	p.TSpanEvent.ExceptionInfo = value
}

func (p *SpanEvent) SetNextSpanID(nextSpanID int64) {
	p.TSpanEvent.NextSpanId = nextSpanID
}

func (p *SpanEvent) SetDestinationID(desID string) {
	p.TSpanEvent.DestinationId = &desID
}

func (p *SpanEvent) Finish() {
	p.MarkEndTime()
}

func (p *SpanEvent) SetSequence(seq int16) {
	p.TSpanEvent.Sequence = seq
}

func (p *SpanEvent) SetDepth(depth int32) {
	p.TSpanEvent.Depth = depth
}
