/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fi

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func IsNilOrEmpty(s *string) bool {
	if s == nil {
		return true
	}
	return *s == ""
}

// String is a helper that builds a *string from a string value
// This is similar to aws.String, except that we use it for non-AWS values
func String(s string) *string {
	return &s
}

func Bool(v bool) *bool {
	return &v
}

func BoolValue(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func Int32(v int32) *int32 {
	return &v
}

func Int32Value(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

// Int64 is a helper that builds a *int64 from an int64 value
// This is similar to aws.Int64, except that we use it for non-AWS values
func Int64(v int64) *int64 {
	return &v
}

func Int64Value(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

func Uint64Value(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

func DebugPrint(o interface{}) string {
	if o == nil {
		return "<nil>"
	}
	if rh, ok := o.(*ResourceHolder); ok {
		if rh == nil {
			// Avoid go nil vs interface problems
			return "<nil>"
		}

		if rh.Resource == nil {
			return fmt.Sprintf("unknown resource %q", rh.Name)
		}
	}
	if resource, ok := o.(Resource); ok {
		if resource == nil {
			// Avoid go nil vs interface problems
			return "<nil>"
		}

		s, err := ResourceAsString(resource)
		if err != nil {
			return fmt.Sprintf("error converting resource to string: %v", err)
		}
		if len(s) >= 256 {
			s = s[:256] + "... (truncated)"
		}
		return s
	}

	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return "<nil>"
		}
		v = v.Elem()
	}
	if !v.IsValid() {
		return "<?>"
	}
	o = v.Interface()
	if stringer, ok := o.(fmt.Stringer); ok {
		if stringer == nil {
			// Avoid go nil vs interface problems
			return "<nil>"
		}
		return stringer.String()
	}

	return fmt.Sprint(o)
}

func DebugAsJsonString(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("error marshalling: %v", err)
	}
	return string(data)
}

func DebugAsJsonStringIndent(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("error marshalling: %v", err)
	}
	return string(data)
}
