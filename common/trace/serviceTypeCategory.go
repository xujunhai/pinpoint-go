package trace

var (
	UndefinedCategory     = ServiceTypeCategory{minCode: -1, maxCode: -1}
	PinpointInternal      = ServiceTypeCategory{minCode: 0, maxCode: 999}
	SERVER                = ServiceTypeCategory{minCode: 1000, maxCode: 1999}
	DATABASE              = ServiceTypeCategory{minCode: 2000, maxCode: 2999}
	LIBRARY               = ServiceTypeCategory{minCode: 5000, maxCode: 7999}
	CacheLibrary          = ServiceTypeCategory{minCode: 8000, maxCode: 8999}
	RPC                   = ServiceTypeCategory{minCode: 9000, maxCode: 9999}
	ServiceTypeCategories = [...]*ServiceTypeCategory{
		&UndefinedCategory,
		&PinpointInternal,
		&SERVER,
		&DATABASE,
		&CacheLibrary,
		&RPC,
	}
)

type ServiceTypeCategory struct {
	minCode int16
	maxCode int16
}

func findCategory(code int16) *ServiceTypeCategory {
	for _, v := range ServiceTypeCategories {
		if v.contains(code) {
			return v
		}
	}
	return nil
}

func (stc ServiceTypeCategory) contains(code int16) bool {
	return stc.minCode <= code && code <= stc.maxCode
}
