package trace

import "fmt"

func ServiceTypeFactory(code int, name string, properties ...ServiceTypeProperty) ServiceType {
	var st ServiceType
	st.code = int16(code)
	st.name = name
	st.desc = name

	st.category = findCategory(int16(code))

	terminal := false
	queue := false
	recordStatistics := false
	includeDestinationId := false
	alias := false

	for _, v := range properties {
		switch v {
		case TERMINAL:
			terminal = true
			break

		case QUEUE:
			queue = true
			break

		case RecordStatistics:
			recordStatistics = true
			break

		case IncludeDestinationId:
			includeDestinationId = true
			break

		case ALIAS:
			alias = true
			break

		default:
			panic(fmt.Sprint("Unknown ServiceTypeProperty:", v))
		}
	}

	st.terminal = terminal
	st.queue = queue
	st.recordStatistics = recordStatistics
	st.includeDestinationId = includeDestinationId
	st.alias = alias
	return st
}
