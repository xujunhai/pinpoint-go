package trace

func AnnotationKeyFactory(code int, name string, properties ...AnnotationKeyProperty) AnnotationKey {
	viewInRecordSet := false
	errorApiMetadata := false

	for _, v := range properties {
		switch v {
		case ViewInRecordSet:
			viewInRecordSet = true
		case ErrorApiMetadata:
			errorApiMetadata = true
		}
	}
	return AnnotationKey{
		code:             code,
		name:             name,
		viewInRecordSet:  viewInRecordSet,
		errorApiMetadata: errorApiMetadata,
	}
}
