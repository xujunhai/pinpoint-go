package profiler

import (
	"net/http"
	"pinpoint-go/common/trace"
	"pinpoint-go/config"
	"pinpoint-go/profiler/context"
)

//param httpHeader is nil means its is root trace
func StartTrace(httpHeader *http.Header) *context.Context {

	traceHeader := trace.NewTraceHttpHeader(httpHeader)

	if traceHeader.HttpType == trace.InvalidHttpHeader || !traceHeader.Sampled {
		return nil
	}

	var traceID *trace.TraceID

	if traceHeader.HttpType == trace.ValidHttpHeader {
		//transID := ParseTransID([]byte(traceHeader.TransactionID))
		//traceID = NewTraceID(transID, traceHeader.PSpanID, traceHeader.SpanID, traceHeader.Flag)
		traceID = trace.NewTraceID(traceHeader.TransactionID, traceHeader.PSpanID, traceHeader.SpanID, traceHeader.Flag)
		if traceID == nil {
			return nil
		}
	} else {
		transSeq := trace.NewTransIDSeq()
		transID := trace.FmtTransIDString(config.AgentId(), GAgent.StartTime, transSeq)
		traceID = &trace.TraceID{
			AgentID:        config.AgentId(),
			AgentStartTime: GAgent.StartTime,
			PSpanID:        -1,
			SpanID:         trace.NewSpanID(),
			TransSeq:       transSeq,
			Flags:          0,
			TransID:        transID,
		}
	}

	span := context.NewSpan(traceID, traceHeader)

	if !traceID.IsRoot() {
		span.SetAcceptorHost(traceHeader.PHostName)
	}
	span.MarkStartTime()
	span.SetEndPoint(GAgent.EndPoint)

	return context.NewTraceContext(traceHeader, traceID, span)
}

func FinishTrace(traceContext *context.Context) {
	if traceContext == nil {
		return
	}
	traceContext.Finish()
	GAgent.SendSpan(traceContext.Span.TSpan)
}
