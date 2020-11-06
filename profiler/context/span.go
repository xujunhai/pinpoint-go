package context

import (
	trace2 "pinpoint-go/common/trace"
	"pinpoint-go/common/util"
	"pinpoint-go/config"
	"pinpoint-go/proto/thrift/gen-go/trace"
)

type Span struct {
	TSpan *trace.TSpan
}

func NewSpan(traceID *trace2.TraceID, header *trace2.HttpHeader) *Span {
	span := &Span{
		TSpan: trace.NewTSpan(),
	}

	if header != nil && header.HttpType == trace2.ValidHttpHeader {
		pAppName := header.PAppName
		pAppType := int16(header.PAppType)
		span.TSpan.ParentApplicationName = &pAppName
		span.TSpan.ParentApplicationType = &pAppType
	}

	span.TSpan.TransactionId = trace2.FmtTransIDByte(traceID.AgentID, traceID.AgentStartTime, traceID.TransSeq)
	span.TSpan.AgentId = config.AgentId()
	span.TSpan.ApplicationName = config.AppName()
	span.TSpan.AgentStartTime = config.StartTime()
	span.TSpan.ServiceType = config.ServerType()
	span.TSpan.SpanId = traceID.SpanID
	span.TSpan.ParentSpanId = traceID.PSpanID
	span.TSpan.Flag = int16(traceID.Flags)

	return span
}

func (p *Span) MarkStartTime() {
	p.TSpan.StartTime = util.GetNowMSec()
}

func (p *Span) MarkEndTime() {
	p.TSpan.Elapsed = int32(util.GetNowMSec() - p.TSpan.StartTime)
}

func (p *Span) SetServiceType(serviceType int16) {
	p.TSpan.ServiceType = serviceType
}

func (p *Span) GetStartTime() int64 {
	return p.TSpan.StartTime
}

func (p *Span) GetEndTime() int64 {
	return p.TSpan.StartTime + int64(p.TSpan.Elapsed)
}

func (p *Span) SetAPIID(apiID int32) {
	p.TSpan.ApiId = &apiID
}

func (p *Span) SetEndPoint(endPoint string) {
	p.TSpan.EndPoint = &endPoint
}

func (p *Span) SetRemoteAddr(remote string) {
	p.TSpan.RemoteAddr = &remote
}

func (p *Span) SetAcceptorHost(host string) {
	p.TSpan.AcceptorHost = &host
}

func (p *Span) SetRpc(rpc string) {
	p.TSpan.RPC = &rpc
}

func (p *Span) AddSpanEvent(spanEvent *trace.TSpanEvent) {
	p.TSpan.SpanEventList = append(p.TSpan.SpanEventList, spanEvent)
}

func (p *Span) AddAnnotation(annotation *trace.TAnnotation) {
	p.TSpan.Annotations = append(p.TSpan.Annotations, annotation)
}

func (p *Span) SetExceptionInfo(errID int32, errMsg string) {
	value := trace.NewTIntStringValue()
	value.IntValue = errID
	value.StringValue = &errMsg
	p.TSpan.ExceptionInfo = value
}
