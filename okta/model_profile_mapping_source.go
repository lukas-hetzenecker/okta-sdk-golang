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

// ProfileMappingSource struct for ProfileMappingSource
type ProfileMappingSource struct {
	Id                   *string                           `json:"id,omitempty"`
	Name                 *string                           `json:"name,omitempty"`
	Type                 *string                           `json:"type,omitempty"`
	Links                map[string]map[string]interface{} `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileMappingSource ProfileMappingSource

// NewProfileMappingSource instantiates a new ProfileMappingSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileMappingSource() *ProfileMappingSource {
	this := ProfileMappingSource{}
	return &this
}

// NewProfileMappingSourceWithDefaults instantiates a new ProfileMappingSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileMappingSourceWithDefaults() *ProfileMappingSource {
	this := ProfileMappingSource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProfileMappingSource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMappingSource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProfileMappingSource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProfileMappingSource) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProfileMappingSource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMappingSource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProfileMappingSource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProfileMappingSource) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProfileMappingSource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMappingSource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProfileMappingSource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProfileMappingSource) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProfileMappingSource) GetLinks() map[string]map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMappingSource) GetLinksOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProfileMappingSource) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]map[string]interface{} and assigns it to the Links field.
func (o *ProfileMappingSource) SetLinks(v map[string]map[string]interface{}) {
	o.Links = v
}

func (o ProfileMappingSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProfileMappingSource) UnmarshalJSON(bytes []byte) (err error) {
	varProfileMappingSource := _ProfileMappingSource{}

	err = json.Unmarshal(bytes, &varProfileMappingSource)
	if err == nil {
		*o = ProfileMappingSource(varProfileMappingSource)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProfileMappingSource struct {
	value *ProfileMappingSource
	isSet bool
}

func (v NullableProfileMappingSource) Get() *ProfileMappingSource {
	return v.value
}

func (v *NullableProfileMappingSource) Set(val *ProfileMappingSource) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileMappingSource) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileMappingSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileMappingSource(val *ProfileMappingSource) *NullableProfileMappingSource {
	return &NullableProfileMappingSource{value: val, isSet: true}
}

func (v NullableProfileMappingSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileMappingSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}