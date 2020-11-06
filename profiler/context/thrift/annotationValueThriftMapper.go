package thrift

import "pinpoint-go/proto/thrift/gen-go/trace"

func buildTAnnotationValue(value interface{}) trace.TAnnotationValue {
	switch value.(type) {
	case string:

	}
}
