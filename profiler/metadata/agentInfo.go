package metadata

type AgentInfo struct {
	AgentInformation AgentInformation
	ServerMetaData   ServerMetaData
}

func (ai AgentInfo) ConvertThrift() {

}
