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

// DeviceAssuranceAndroidPlatformAllOf struct for DeviceAssuranceAndroidPlatformAllOf
type DeviceAssuranceAndroidPlatformAllOf struct {
	DiskEncryptionType    *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType `json:"diskEncryptionType,omitempty"`
	Jailbreak             *bool                                                  `json:"jailbreak,omitempty"`
	OsVersion             *OSVersion                                             `json:"osVersion,omitempty"`
	ScreenLockType        *DeviceAssuranceAndroidPlatformAllOfScreenLockType     `json:"screenLockType,omitempty"`
	SecureHardwarePresent *bool                                                  `json:"secureHardwarePresent,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _DeviceAssuranceAndroidPlatformAllOf DeviceAssuranceAndroidPlatformAllOf

// NewDeviceAssuranceAndroidPlatformAllOf instantiates a new DeviceAssuranceAndroidPlatformAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceAndroidPlatformAllOf() *DeviceAssuranceAndroidPlatformAllOf {
	this := DeviceAssuranceAndroidPlatformAllOf{}
	return &this
}

// NewDeviceAssuranceAndroidPlatformAllOfWithDefaults instantiates a new DeviceAssuranceAndroidPlatformAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceAndroidPlatformAllOfWithDefaults() *DeviceAssuranceAndroidPlatformAllOf {
	this := DeviceAssuranceAndroidPlatformAllOf{}
	return &this
}

// GetDiskEncryptionType returns the DiskEncryptionType field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatformAllOf) GetDiskEncryptionType() DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType {
	if o == nil || o.DiskEncryptionType == nil {
		var ret DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType
		return ret
	}
	return *o.DiskEncryptionType
}

// GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatformAllOf) GetDiskEncryptionTypeOk() (*DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType, bool) {
	if o == nil || o.DiskEncryptionType == nil {
		return nil, false
	}
	return o.DiskEncryptionType, true
}

// HasDiskEncryptionType returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatformAllOf) HasDiskEncryptionType() bool {
	if o != nil && o.DiskEncryptionType != nil {
		return true
	}

	return false
}

// SetDiskEncryptionType gets a reference to the given DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType and assigns it to the DiskEncryptionType field.
func (o *DeviceAssuranceAndroidPlatformAllOf) SetDiskEncryptionType(v DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) {
	o.DiskEncryptionType = &v
}

// GetJailbreak returns the Jailbreak field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatformAllOf) GetJailbreak() bool {
	if o == nil || o.Jailbreak == nil {
		var ret bool
		return ret
	}
	return *o.Jailbreak
}

// GetJailbreakOk returns a tuple with the Jailbreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatformAllOf) GetJailbreakOk() (*bool, bool) {
	if o == nil || o.Jailbreak == nil {
		return nil, false
	}
	return o.Jailbreak, true
}

// HasJailbreak returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatformAllOf) HasJailbreak() bool {
	if o != nil && o.Jailbreak != nil {
		return true
	}

	return false
}

// SetJailbreak gets a reference to the given bool and assigns it to the Jailbreak field.
func (o *DeviceAssuranceAndroidPlatformAllOf) SetJailbreak(v bool) {
	o.Jailbreak = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatformAllOf) GetOsVersion() OSVersion {
	if o == nil || o.OsVersion == nil {
		var ret OSVersion
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatformAllOf) GetOsVersionOk() (*OSVersion, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatformAllOf) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given OSVersion and assigns it to the OsVersion field.
func (o *DeviceAssuranceAndroidPlatformAllOf) SetOsVersion(v OSVersion) {
	o.OsVersion = &v
}

// GetScreenLockType returns the ScreenLockType field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatformAllOf) GetScreenLockType() DeviceAssuranceAndroidPlatformAllOfScreenLockType {
	if o == nil || o.ScreenLockType == nil {
		var ret DeviceAssuranceAndroidPlatformAllOfScreenLockType
		return ret
	}
	return *o.ScreenLockType
}

// GetScreenLockTypeOk returns a tuple with the ScreenLockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatformAllOf) GetScreenLockTypeOk() (*DeviceAssuranceAndroidPlatformAllOfScreenLockType, bool) {
	if o == nil || o.ScreenLockType == nil {
		return nil, false
	}
	return o.ScreenLockType, true
}

// HasScreenLockType returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatformAllOf) HasScreenLockType() bool {
	if o != nil && o.ScreenLockType != nil {
		return true
	}

	return false
}

// SetScreenLockType gets a reference to the given DeviceAssuranceAndroidPlatformAllOfScreenLockType and assigns it to the ScreenLockType field.
func (o *DeviceAssuranceAndroidPlatformAllOf) SetScreenLockType(v DeviceAssuranceAndroidPlatformAllOfScreenLockType) {
	o.ScreenLockType = &v
}

// GetSecureHardwarePresent returns the SecureHardwarePresent field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatformAllOf) GetSecureHardwarePresent() bool {
	if o == nil || o.SecureHardwarePresent == nil {
		var ret bool
		return ret
	}
	return *o.SecureHardwarePresent
}

// GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatformAllOf) GetSecureHardwarePresentOk() (*bool, bool) {
	if o == nil || o.SecureHardwarePresent == nil {
		return nil, false
	}
	return o.SecureHardwarePresent, true
}

// HasSecureHardwarePresent returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatformAllOf) HasSecureHardwarePresent() bool {
	if o != nil && o.SecureHardwarePresent != nil {
		return true
	}

	return false
}

// SetSecureHardwarePresent gets a reference to the given bool and assigns it to the SecureHardwarePresent field.
func (o *DeviceAssuranceAndroidPlatformAllOf) SetSecureHardwarePresent(v bool) {
	o.SecureHardwarePresent = &v
}

func (o DeviceAssuranceAndroidPlatformAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiskEncryptionType != nil {
		toSerialize["diskEncryptionType"] = o.DiskEncryptionType
	}
	if o.Jailbreak != nil {
		toSerialize["jailbreak"] = o.Jailbreak
	}
	if o.OsVersion != nil {
		toSerialize["osVersion"] = o.OsVersion
	}
	if o.ScreenLockType != nil {
		toSerialize["screenLockType"] = o.ScreenLockType
	}
	if o.SecureHardwarePresent != nil {
		toSerialize["secureHardwarePresent"] = o.SecureHardwarePresent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAssuranceAndroidPlatformAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceAssuranceAndroidPlatformAllOf := _DeviceAssuranceAndroidPlatformAllOf{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceAndroidPlatformAllOf)
	if err == nil {
		*o = DeviceAssuranceAndroidPlatformAllOf(varDeviceAssuranceAndroidPlatformAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "diskEncryptionType")
		delete(additionalProperties, "jailbreak")
		delete(additionalProperties, "osVersion")
		delete(additionalProperties, "screenLockType")
		delete(additionalProperties, "secureHardwarePresent")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceAssuranceAndroidPlatformAllOf struct {
	value *DeviceAssuranceAndroidPlatformAllOf
	isSet bool
}

func (v NullableDeviceAssuranceAndroidPlatformAllOf) Get() *DeviceAssuranceAndroidPlatformAllOf {
	return v.value
}

func (v *NullableDeviceAssuranceAndroidPlatformAllOf) Set(val *DeviceAssuranceAndroidPlatformAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceAndroidPlatformAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceAndroidPlatformAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceAndroidPlatformAllOf(val *DeviceAssuranceAndroidPlatformAllOf) *NullableDeviceAssuranceAndroidPlatformAllOf {
	return &NullableDeviceAssuranceAndroidPlatformAllOf{value: val, isSet: true}
}

func (v NullableDeviceAssuranceAndroidPlatformAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceAndroidPlatformAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}