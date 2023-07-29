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
)

// LogStreamSettingsAws struct for LogStreamSettingsAws
type LogStreamSettingsAws struct {
	// Your AWS account ID
	AccountId *string `json:"accountId,omitempty"`
	// An alphanumeric name (no spaces) to identify this event source in AWS EventBridge
	EventSourceName      *string    `json:"eventSourceName,omitempty"`
	Region               *AwsRegion `json:"region,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamSettingsAws LogStreamSettingsAws

// NewLogStreamSettingsAws instantiates a new LogStreamSettingsAws object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamSettingsAws() *LogStreamSettingsAws {
	this := LogStreamSettingsAws{}
	return &this
}

// NewLogStreamSettingsAwsWithDefaults instantiates a new LogStreamSettingsAws object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamSettingsAwsWithDefaults() *LogStreamSettingsAws {
	this := LogStreamSettingsAws{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *LogStreamSettingsAws) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsAws) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *LogStreamSettingsAws) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *LogStreamSettingsAws) SetAccountId(v string) {
	o.AccountId = &v
}

// GetEventSourceName returns the EventSourceName field value if set, zero value otherwise.
func (o *LogStreamSettingsAws) GetEventSourceName() string {
	if o == nil || o.EventSourceName == nil {
		var ret string
		return ret
	}
	return *o.EventSourceName
}

// GetEventSourceNameOk returns a tuple with the EventSourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsAws) GetEventSourceNameOk() (*string, bool) {
	if o == nil || o.EventSourceName == nil {
		return nil, false
	}
	return o.EventSourceName, true
}

// HasEventSourceName returns a boolean if a field has been set.
func (o *LogStreamSettingsAws) HasEventSourceName() bool {
	if o != nil && o.EventSourceName != nil {
		return true
	}

	return false
}

// SetEventSourceName gets a reference to the given string and assigns it to the EventSourceName field.
func (o *LogStreamSettingsAws) SetEventSourceName(v string) {
	o.EventSourceName = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *LogStreamSettingsAws) GetRegion() AwsRegion {
	if o == nil || o.Region == nil {
		var ret AwsRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsAws) GetRegionOk() (*AwsRegion, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *LogStreamSettingsAws) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given AwsRegion and assigns it to the Region field.
func (o *LogStreamSettingsAws) SetRegion(v AwsRegion) {
	o.Region = &v
}

func (o LogStreamSettingsAws) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.EventSourceName != nil {
		toSerialize["eventSourceName"] = o.EventSourceName
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogStreamSettingsAws) UnmarshalJSON(bytes []byte) (err error) {
	varLogStreamSettingsAws := _LogStreamSettingsAws{}

	err = json.Unmarshal(bytes, &varLogStreamSettingsAws)
	if err == nil {
		*o = LogStreamSettingsAws(varLogStreamSettingsAws)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "eventSourceName")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogStreamSettingsAws struct {
	value *LogStreamSettingsAws
	isSet bool
}

func (v NullableLogStreamSettingsAws) Get() *LogStreamSettingsAws {
	return v.value
}

func (v *NullableLogStreamSettingsAws) Set(val *LogStreamSettingsAws) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamSettingsAws) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamSettingsAws) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamSettingsAws(val *LogStreamSettingsAws) *NullableLogStreamSettingsAws {
	return &NullableLogStreamSettingsAws{value: val, isSet: true}
}

func (v NullableLogStreamSettingsAws) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamSettingsAws) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}