package common

import "log"

const (
	SocketStatNone                    byte = 1
	SocketStatBeingConnect            byte = 2
	SocketStatConnected               byte = 3
	SocketStatConnectFailed           byte = 6
	SocketStatIgnore                  byte = 9
	SocketStatRunWithoutHandshake     byte = 10
	SocketStatRunSimplex              byte = 11
	SocketStatRunDuplex               byte = 12
	SocketStatBeingCloseByClient      byte = 20
	SocketStatClosedByClient          byte = 22
	SocketStatUnexpectedCloseByClient byte = 26
	SocketStatBeingCloseByServer      byte = 30
	SocketStatClosedByServer          byte = 32
	SocketStatUnexpectedCloseByServer byte = 36
	SocketStatErrorUnknown            byte = 40
	SocketStatIllegalStateChange      byte = 41
	SocketStatErrorSyncStateSession   byte = 42
)

var sockStatCodeOp SocketStateCodeOp

type SocketStateCodeOp struct{}

func (p *SocketStateCodeOp) canChangeState(state byte, nextState byte) bool {
	switch nextState {
	case SocketStatBeingConnect:
		return state == SocketStatNone

	case SocketStatConnected:
		return state == SocketStatNone || state == SocketStatBeingConnect ||
			state == SocketStatConnectFailed

	case SocketStatConnectFailed:
		return state == SocketStatBeingConnect || state == SocketStatRunWithoutHandshake ||
			state == SocketStatRunSimplex || state == SocketStatRunDuplex

	case SocketStatIgnore:
		return state == SocketStatConnected

	case SocketStatRunWithoutHandshake:
		return state == SocketStatConnected

	case SocketStatRunSimplex:
		return state == SocketStatRunWithoutHandshake

	case SocketStatRunDuplex:
		return state == SocketStatRunWithoutHandshake

	case SocketStatBeingCloseByClient:
		return state == SocketStatRunWithoutHandshake || state == SocketStatRunSimplex ||
			state == SocketStatRunDuplex

	case SocketStatClosedByClient:
		return state == SocketStatBeingCloseByClient

	case SocketStatUnexpectedCloseByClient:
		return state == SocketStatConnected || state == SocketStatRunWithoutHandshake ||
			state == SocketStatRunSimplex || state == SocketStatRunDuplex

	case SocketStatBeingCloseByServer:
		return state == SocketStatRunWithoutHandshake || state == SocketStatRunSimplex ||
			state == SocketStatRunDuplex

	case SocketStatClosedByServer:
		return state == SocketStatBeingCloseByServer

	case SocketStatUnexpectedCloseByServer:
		return state == SocketStatRunWithoutHandshake || state == SocketStatRunSimplex ||
			state == SocketStatRunDuplex

	default:
		log.Println("error state change: ", state, " ", nextState)
		return false
	}
}

func (p *SocketStateCodeOp) isBeforeConnected(state byte) bool {
	switch state {
	case SocketStatNone:
		return true
	case SocketStatBeingConnect:
		return true
	}

	return false
}

func (p *SocketStateCodeOp) isRun(state byte) bool {
	switch state {
	case SocketStatRunWithoutHandshake:
		return true
	case SocketStatRunSimplex:
		return true
	case SocketStatRunDuplex:
		return true
	}

	return false
}

func (p *SocketStateCodeOp) isRunDuplex(state byte) bool {
	return state == SocketStatRunDuplex
}

func (p *SocketStateCodeOp) onClose(state byte) bool {
	switch state {
	case SocketStatBeingCloseByClient:
		return true
	case SocketStatBeingCloseByServer:
		return true
	}

	return false
}

func (p *SocketStateCodeOp) isClosed(state byte) bool {
	switch state {
	case SocketStatClosedByClient:
		return true
	case SocketStatClosedByServer:
		return true
	case SocketStatUnexpectedCloseByServer:
		return true
	case SocketStatErrorUnknown:
		return true
	case SocketStatIllegalStateChange:
		return true
	case SocketStatErrorSyncStateSession:
		return true
	}

	return false
}

func (p *SocketStateCodeOp) isError(state byte) bool {
	switch state {
	case SocketStatIllegalStateChange:
		return true
	case SocketStatErrorUnknown:
		return true
	}

	return false
}

type SocketStateChangeResult struct {
	changed           bool
	beforeState       byte
	currentState      byte
	updateWantedState byte
}

func NewSocketStateChange(changed bool, beforeState, currentState, updateWantedState byte) *SocketStateChangeResult {
	return &SocketStateChangeResult{
		changed:           changed,
		beforeState:       beforeState,
		currentState:      currentState,
		updateWantedState: updateWantedState,
	}
}

func (p *SocketStateChangeResult) isChange() bool {
	return p.changed
}

func (p *SocketStateChangeResult) getBeforeState() byte {
	return p.beforeState
}

func (p *SocketStateChangeResult) getCurrentState() byte {
	return p.currentState
}

func (p *SocketStateChangeResult) getUpdateWantedState() byte {
	return p.updateWantedState
}

type SocketState struct {
	beforeState  byte
	currentState byte
}

func NewSocketState() *SocketState {
	return &SocketState{
		beforeState:  SocketStatNone,
		currentState: SocketStatNone,
	}
}

func (p *SocketState) to(nextState byte) SocketStateChangeResult {
	enable := sockStatCodeOp.canChangeState(p.currentState, nextState)
	changed := false
	if enable {
		p.beforeState = p.currentState
		p.currentState = nextState
		changed = true
	}

	return SocketStateChangeResult{
		changed:           changed,
		beforeState:       p.beforeState,
		currentState:      p.currentState,
		updateWantedState: nextState,
	}
}

func (p *SocketState) toBeingConnect() SocketStateChangeResult {
	return p.to(SocketStatBeingConnect)
}

func (p *SocketState) ToConnected() SocketStateChangeResult {
	return p.to(SocketStatConnected)
}

func (p *SocketState) toConnectFailed() SocketStateChangeResult {
	return p.to(SocketStatConnectFailed)
}

func (p *SocketState) toIgnore() SocketStateChangeResult {
	return p.to(SocketStatIgnore)
}

func (p *SocketState) ToRunWithoutHandShake() SocketStateChangeResult {
	return p.to(SocketStatRunWithoutHandshake)
}

func (p *SocketState) ToRunSimplex() SocketStateChangeResult {
	return p.to(SocketStatRunSimplex)
}

func (p *SocketState) ToRunDuplex() SocketStateChangeResult {
	return p.to(SocketStatRunDuplex)
}

func (p *SocketState) toBeingCloseByClient() SocketStateChangeResult {
	return p.to(SocketStatBeingCloseByClient)
}

func (p *SocketState) ToClosedByClient() SocketStateChangeResult {
	return p.to(SocketStatClosedByClient)
}

func (p *SocketState) toUnexpectedCloseByClient() SocketStateChangeResult {
	return p.to(SocketStatUnexpectedCloseByClient)
}

func (p *SocketState) toBeingCloseByServer() SocketStateChangeResult {
	return p.to(SocketStatBeingCloseByServer)
}

func (p *SocketState) toClosedByServer() SocketStateChangeResult {
	return p.to(SocketStatClosedByServer)
}

func (p *SocketState) toUnexpectedCloseByServer() SocketStateChangeResult {
	return p.to(SocketStatUnexpectedCloseByServer)
}

func (p *SocketState) toUnknownError() SocketStateChangeResult {
	return p.to(SocketStatErrorUnknown)
}

func (p *SocketState) toSyncStateSessionError() SocketStateChangeResult {
	return p.to(SocketStatErrorSyncStateSession)
}

func (p *SocketState) GetCurrentState() byte {
	return p.currentState
}
