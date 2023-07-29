/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 4.0.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// LogSeverity the model 'LogSeverity'
type LogSeverity string

// List of LogSeverity
const (
	LOGSEVERITY_DEBUG LogSeverity = "DEBUG"
	LOGSEVERITY_ERROR LogSeverity = "ERROR"
	LOGSEVERITY_INFO  LogSeverity = "INFO"
	LOGSEVERITY_WARN  LogSeverity = "WARN"
)

// All allowed values of LogSeverity enum
var AllowedLogSeverityEnumValues = []LogSeverity{
	"DEBUG",
	"ERROR",
	"INFO",
	"WARN",
}

func (v *LogSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogSeverity(value)
	for _, existing := range AllowedLogSeverityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogSeverity", value)
}

// NewLogSeverityFromValue returns a pointer to a valid LogSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogSeverityFromValue(v string) (*LogSeverity, error) {
	ev := LogSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogSeverity: valid values are %v", v, AllowedLogSeverityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogSeverity) IsValid() bool {
	for _, existing := range AllowedLogSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogSeverity value
func (v LogSeverity) Ptr() *LogSeverity {
	return &v
}

type NullableLogSeverity struct {
	value *LogSeverity
	isSet bool
}

func (v NullableLogSeverity) Get() *LogSeverity {
	return v.value
}

func (v *NullableLogSeverity) Set(val *LogSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSeverity(val *LogSeverity) *NullableLogSeverity {
	return &NullableLogSeverity{value: val, isSet: true}
}

func (v NullableLogSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}