package trace

import (
	"log"
	"math/rand"
	"pinpoint-go/common/util"
	"strconv"
	"strings"
	"sync/atomic"
)

const TVERSION byte = 0

var transIDSeq int64

type TraceID struct {
	AgentID        string
	TransID        string
	PSpanID        int64
	SpanID         int64
	TransSeq       int64
	AgentStartTime int64
	Flags          int64
}

/**
 * @author emeroad
 */
type ITraceId interface {
	getNextTraceId() ITraceId
	getSpanId() int64
	getTransactionId() string
	getAgentId() string
	getAgentStartTime() int64
	getTransactionSequence() int64
	getParentSpanId() int64
	getFlags() int16
	isRoot() bool
}

func NewTransIDSeq() int64 {
	return atomic.AddInt64(&transIDSeq, 1)
}

func FmtTransIDByte(agentID string, agentStartTime int64, transSeq int64) []byte {
	byteUtil := util.NewByteUtil("")
	byteUtil.PutByte(TVERSION)
	byteUtil.PutPrefixedString(agentID)
	byteUtil.PutVar64(agentStartTime)
	byteUtil.PutVar64(transSeq)
	return byteUtil.GetBytes()
}

func FmtTransIDString(agentID string, agentStartTime int64, transSeq int64) string {
	return agentID + "^" +
		strconv.FormatInt(agentStartTime, 10) + "^" +
		strconv.FormatInt(transSeq, 10)
}

func ParseTransID(bys []byte) string {
	byteUtil := util.NewByteUtil(string(bys))
	version := byteUtil.ReadByte()
	if version != TVERSION {
		return ""
	}
	agentID := byteUtil.ReadPrefixedString()
	agentStartTime := byteUtil.ReadVar64()
	transSeq := byteUtil.ReadVar64()
	return FmtTransIDString(agentID, agentStartTime, transSeq)
}

func createSpanID() int64 {
	spanID := rand.Int63()
	for spanID == -1 {
		spanID = rand.Int63()
	}
	return spanID
}

func NewSpanID() int64 {
	return createSpanID()
}

func NextSpanID(spanID, pSpanID int64) int64 {
	nextSpanID := createSpanID()
	for spanID == nextSpanID || nextSpanID == pSpanID {
		nextSpanID = createSpanID()
	}
	return nextSpanID
}

func NewTraceID(transID string, pSpanID, spanID, flags int64) *TraceID {
	transInfo := strings.Split(transID, "^")
	if len(transInfo) != 3 {
		log.Println("New TraceID Parse Error")
		return nil
	}

	agentID := transInfo[0]
	agentStartTime, err := strconv.ParseInt(transInfo[1], 10, 64)
	if err != nil {
		log.Println("New TraceID Parse AgentStartTime error")
		return nil
	}

	transSeq, err := strconv.ParseInt(transInfo[2], 10, 64)
	if err != nil {
		log.Println("New TraceID Parse TransSeq error")
		return nil
	}

	return &TraceID{
		AgentID:        agentID,
		AgentStartTime: agentStartTime,
		TransSeq:       transSeq,
		PSpanID:        pSpanID,
		SpanID:         spanID,
		Flags:          flags,
	}
}

func (p *TraceID) IsRoot() bool {
	return p.PSpanID == -1
}

func (p *TraceID) GetTransID() string {
	return p.TransID
}
