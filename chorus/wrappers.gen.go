// This is a generated file, please do not edit by hand
package chorus

import (
	"github.com/golang/protobuf/ptypes/wrappers"
)

func convertToBoolPtr(in *wrappers.BoolValue) *bool {
	if in == nil {
		return nil
	}
	v := in.GetValue()
	return &v
}

func convertFromBoolPtr(in *bool) *wrappers.BoolValue {
	if in == nil {
		return nil
	}
	return &wrappers.BoolValue{Value: *in}
}

func convertToStringPtr(in *wrappers.StringValue) *string {
	if in == nil {
		return nil
	}
	v := in.GetValue()
	return &v
}

func convertFromStringPtr(in *string) *wrappers.StringValue {
	if in == nil {
		return nil
	}
	return &wrappers.StringValue{Value: *in}
}

func convertToInt64Ptr(in *wrappers.Int64Value) *int64 {
	if in == nil {
		return nil
	}
	v := in.GetValue()
	return &v
}

func convertFromInt64Ptr(in *int64) *wrappers.Int64Value {
	if in == nil {
		return nil
	}
	return &wrappers.Int64Value{Value: *in}
}
