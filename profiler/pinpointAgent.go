package profiler

import (
	"net/http"
	trace2 "pinpoint-go/common/trace"
	"pinpoint-go/common/util"
	"pinpoint-go/config"
	"pinpoint-go/profiler/metadata"
	"pinpoint-go/profiler/monitor"
	"pinpoint-go/profiler/sender"
	"pinpoint-go/proto/thrift"
	"pinpoint-go/proto/thrift/gen-go/trace"
	"pinpoint-go/rpc/packet"
)

const (
	MysqlExec    = "mysql.Exec"
	MysqlConnect = "mysql.Connect"

	StrError     = "Error"
	StrException = "Exception"
	StrWarn      = "Warn"
)

var GAgent = NewPinpointAgent()

type PinpointAgent struct {
	AgentID         string
	ApplicationName string
	client          *sender.TcpDataSender
	statMonitor     *monitor.AgentStatMonitor
	spanClient      *sender.AgentSpanClient

	apiMetaManager *metadata.SimpleCache
	sqlMetaManager *metadata.SimpleCache
	strMetaManager *metadata.SimpleCache
	StartTime      int64
	ServerType     int16
	hostName       string
	EndPoint       string
}

func NewPinpointAgent() *PinpointAgent {
	return &PinpointAgent{
		client:         sender.NewTcpDataSender(),
		statMonitor:    monitor.NewAgentStatMonitor(config.StatAddr()),
		spanClient:     sender.NewAgentSpanClient(config.SpanAddr()),
		StartTime:      util.GetNowMSec(),
		ServerType:     trace2.StandAlone,
		apiMetaManager: metadata.NewSimpleCache(),
		sqlMetaManager: metadata.NewSimpleCache(),
		strMetaManager: metadata.NewSimpleCache(),
		hostName:       util.GetHostName(),
		EndPoint:       util.GetHostIP(),
	}
}

func InitCommonAPI() {
	GAgent.AddAPIMeta("http."+http.MethodHead, -1, -1)
	GAgent.AddAPIMeta("http."+http.MethodGet, -1, -1)
	GAgent.AddAPIMeta("http."+http.MethodDelete, -1, -1)
	GAgent.AddAPIMeta("http."+http.MethodConnect, -1, -1)
	GAgent.AddAPIMeta("http."+http.MethodOptions, -1, -1)
	GAgent.AddAPIMeta("http."+http.MethodPatch, -1, -1)
	GAgent.AddAPIMeta("http."+http.MethodPost, -1, -1)
	GAgent.AddAPIMeta("http."+http.MethodTrace, -1, -1)
	GAgent.AddAPIMeta("http."+http.MethodPut, -1, -1)

	GAgent.AddAPIMeta(MysqlExec, -1, -1)
	GAgent.AddAPIMeta(MysqlConnect, -1, -1)

	GAgent.AddStrMeta(StrError)
	GAgent.AddStrMeta(StrException)
	GAgent.AddStrMeta(StrWarn)

}

func (p *PinpointAgent) Start() {
	p.AgentID = config.AgentId()
	p.ApplicationName = config.AppName()
	p.statMonitor = monitor.NewAgentStatMonitor(config.StatAddr())
	p.spanClient = sender.NewAgentSpanClient(config.SpanAddr())

	quitC := make(chan struct{})

	go p.client.Init(quitC)
	go p.statMonitor.Start()
	go p.spanClient.Start()

	//初始化agentInfo
	agentInfoSender := NewAgentInfoSender(p.client)
	go agentInfoSender.SendAgentInfo(quitC) //agent的信息采集打点

	InitCommonAPI()

	//待优化
	GAgent.SendAllMetaData()
}

func (p *PinpointAgent) AddAPIMeta(api string, line, apiType int32) int32 {
	apiID, isNew := p.apiMetaManager.GetOrCreateID(api)

	//非新的说明之前已经发过了
	if !isNew {
		return apiID
	}

	apiMetaObject := &metadata.ApiMetaData{
		ApiInfo: api,
		ApiId:   apiID,
		Line:    line,
		Type:    apiType,
	}

	p.apiMetaManager.SetMetaObject(apiID, apiMetaObject)

	p.sendApiMeta(api, apiID, line, apiType)

	return apiID
}

func (p *PinpointAgent) sendApiMeta(api string, apiID, line, apiType int32) {
	apiMetaData := trace.NewTApiMetaData()
	apiMetaData.AgentId = GAgent.AgentID
	apiMetaData.AgentStartTime = GAgent.StartTime
	apiMetaData.ApiId = apiID
	apiMetaData.ApiInfo = api

	if line != -1 {
		apiMetaData.Line = &line
	}
	if apiType != -1 {
		apiMetaData.Type = &apiType
	}

	payload := thrift.Serialize(apiMetaData)
	request := packet.NewRequestPacket()
	request.Payload = payload
	p.client.SendPacket(request)
}

func (p *PinpointAgent) GetApiID(metaInfo string) int32 {
	return p.apiMetaManager.GetID(metaInfo)
}

func (p *PinpointAgent) AddSQLMeta(sql string) int32 {
	sqlID, isNew := p.sqlMetaManager.GetOrCreateID(sql)

	if !isNew {
		return sqlID
	}

	sqlMetaObject := &metadata.SqlMetaData{
		SqlId: sqlID,
		Sql:   sql,
	}

	p.sqlMetaManager.SetMetaObject(sqlID, sqlMetaObject)

	p.sendSqlMeta(sqlID, sql)
	return sqlID
}

func (p *PinpointAgent) sendSqlMeta(sqlID int32, sql string) {
	sqlMetaData := trace.NewTSqlMetaData()
	sqlMetaData.AgentId = GAgent.AgentID
	sqlMetaData.AgentStartTime = GAgent.StartTime
	sqlMetaData.Sql = sql
	sqlMetaData.SqlId = sqlID

	payload := thrift.Serialize(sqlMetaData)
	request := packet.NewRequestPacket()
	request.Payload = payload
	p.client.SendPacket(request)
}

func (p *PinpointAgent) AddStrMeta(metaInfo string) int32 {
	strID, isNew := p.strMetaManager.GetOrCreateID(metaInfo)
	if !isNew {
		return strID
	}

	strMetaObject := &metadata.StringMetaData{
		StringId:    strID,
		StringValue: metaInfo,
	}

	p.strMetaManager.SetMetaObject(strID, strMetaObject)
	p.sendStrMeta(strID, metaInfo)
	return strID
}

func (p *PinpointAgent) sendStrMeta(strID int32, str string) {
	strMetaData := trace.NewTStringMetaData()
	strMetaData.AgentId = GAgent.AgentID
	strMetaData.AgentStartTime = GAgent.StartTime
	strMetaData.StringId = strID
	strMetaData.StringValue = str

	payload := thrift.Serialize(strMetaData)
	request := packet.NewRequestPacket()
	request.Payload = payload
	p.client.SendPacket(request)
}

func (p *PinpointAgent) GetStrID(metaInfo string) int32 {
	return p.strMetaManager.GetID(metaInfo)
}

func (p *PinpointAgent) SendAllMetaData() {
	sqlMetaMap := p.sqlMetaManager.ID2MetaObject

	sqlMetaMap.Range(func(key, metaObject interface{}) bool {
		sqlMeta, ok := metaObject.(*metadata.SqlMetaData)
		if !ok {
			return true
		}
		p.sendSqlMeta(sqlMeta.SqlId, sqlMeta.Sql)
		return true
	})

	apiMetaMap := p.apiMetaManager.ID2MetaObject

	apiMetaMap.Range(func(key, metaObject interface{}) bool {
		apiMeta, ok := metaObject.(*metadata.ApiMetaData)
		if !ok {
			return true
		}
		p.sendApiMeta(apiMeta.ApiInfo, apiMeta.ApiId, apiMeta.Line, apiMeta.Type)
		return true
	})

	strMetaMap := p.strMetaManager.ID2MetaObject
	strMetaMap.Range(func(key, metaObject interface{}) bool {
		strMeta, ok := metaObject.(*metadata.StringMetaData)
		if !ok {
			return true
		}
		p.sendStrMeta(strMeta.StringId, strMeta.StringValue)
		return true
	})
}

func (p *PinpointAgent) GetSQLID(metaInfo string) int32 {
	return p.sqlMetaManager.GetID(metaInfo)
}

func (p *PinpointAgent) SendSpan(tspan *trace.TSpan) {
	p.spanClient.SendTrace(tspan)
}
