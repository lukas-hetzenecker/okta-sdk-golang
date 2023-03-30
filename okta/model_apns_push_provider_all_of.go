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

// APNSPushProviderAllOf struct for APNSPushProviderAllOf
type APNSPushProviderAllOf struct {
	Configuration        *APNSConfiguration `json:"configuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _APNSPushProviderAllOf APNSPushProviderAllOf

// NewAPNSPushProviderAllOf instantiates a new APNSPushProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPNSPushProviderAllOf() *APNSPushProviderAllOf {
	this := APNSPushProviderAllOf{}
	return &this
}

// NewAPNSPushProviderAllOfWithDefaults instantiates a new APNSPushProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPNSPushProviderAllOfWithDefaults() *APNSPushProviderAllOf {
	this := APNSPushProviderAllOf{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *APNSPushProviderAllOf) GetConfiguration() APNSConfiguration {
	if o == nil || o.Configuration == nil {
		var ret APNSConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APNSPushProviderAllOf) GetConfigurationOk() (*APNSConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *APNSPushProviderAllOf) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given APNSConfiguration and assigns it to the Configuration field.
func (o *APNSPushProviderAllOf) SetConfiguration(v APNSConfiguration) {
	o.Configuration = &v
}

func (o APNSPushProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *APNSPushProviderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAPNSPushProviderAllOf := _APNSPushProviderAllOf{}

	err = json.Unmarshal(bytes, &varAPNSPushProviderAllOf)
	if err == nil {
		*o = APNSPushProviderAllOf(varAPNSPushProviderAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "configuration")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAPNSPushProviderAllOf struct {
	value *APNSPushProviderAllOf
	isSet bool
}

func (v NullableAPNSPushProviderAllOf) Get() *APNSPushProviderAllOf {
	return v.value
}

func (v *NullableAPNSPushProviderAllOf) Set(val *APNSPushProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAPNSPushProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAPNSPushProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPNSPushProviderAllOf(val *APNSPushProviderAllOf) *NullableAPNSPushProviderAllOf {
	return &NullableAPNSPushProviderAllOf{value: val, isSet: true}
}

func (v NullableAPNSPushProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPNSPushProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
