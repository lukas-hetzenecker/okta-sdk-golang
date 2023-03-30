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

// GroupRuleStatus the model 'GroupRuleStatus'
type GroupRuleStatus string

// List of GroupRuleStatus
const (
	GROUPRULESTATUS_ACTIVE   GroupRuleStatus = "ACTIVE"
	GROUPRULESTATUS_INACTIVE GroupRuleStatus = "INACTIVE"
	GROUPRULESTATUS_INVALID  GroupRuleStatus = "INVALID"
)

// All allowed values of GroupRuleStatus enum
var AllowedGroupRuleStatusEnumValues = []GroupRuleStatus{
	"ACTIVE",
	"INACTIVE",
	"INVALID",
}

func (v *GroupRuleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupRuleStatus(value)
	for _, existing := range AllowedGroupRuleStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupRuleStatus", value)
}

// NewGroupRuleStatusFromValue returns a pointer to a valid GroupRuleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupRuleStatusFromValue(v string) (*GroupRuleStatus, error) {
	ev := GroupRuleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupRuleStatus: valid values are %v", v, AllowedGroupRuleStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupRuleStatus) IsValid() bool {
	for _, existing := range AllowedGroupRuleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupRuleStatus value
func (v GroupRuleStatus) Ptr() *GroupRuleStatus {
	return &v
}

type NullableGroupRuleStatus struct {
	value *GroupRuleStatus
	isSet bool
}

func (v NullableGroupRuleStatus) Get() *GroupRuleStatus {
	return v.value
}

func (v *NullableGroupRuleStatus) Set(val *GroupRuleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRuleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRuleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRuleStatus(val *GroupRuleStatus) *NullableGroupRuleStatus {
	return &NullableGroupRuleStatus{value: val, isSet: true}
}

func (v NullableGroupRuleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRuleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
