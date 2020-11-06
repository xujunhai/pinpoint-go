package trace

type ServiceTypeProperty int

const (
	TERMINAL ServiceTypeProperty = iota
	QUEUE
	RecordStatistics
	IncludeDestinationId
	ALIAS
)
