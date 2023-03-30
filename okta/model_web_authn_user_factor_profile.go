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

// WebAuthnUserFactorProfile struct for WebAuthnUserFactorProfile
type WebAuthnUserFactorProfile struct {
	AuthenticatorName    *string `json:"authenticatorName,omitempty"`
	CredentialId         *string `json:"credentialId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebAuthnUserFactorProfile WebAuthnUserFactorProfile

// NewWebAuthnUserFactorProfile instantiates a new WebAuthnUserFactorProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebAuthnUserFactorProfile() *WebAuthnUserFactorProfile {
	this := WebAuthnUserFactorProfile{}
	return &this
}

// NewWebAuthnUserFactorProfileWithDefaults instantiates a new WebAuthnUserFactorProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebAuthnUserFactorProfileWithDefaults() *WebAuthnUserFactorProfile {
	this := WebAuthnUserFactorProfile{}
	return &this
}

// GetAuthenticatorName returns the AuthenticatorName field value if set, zero value otherwise.
func (o *WebAuthnUserFactorProfile) GetAuthenticatorName() string {
	if o == nil || o.AuthenticatorName == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatorName
}

// GetAuthenticatorNameOk returns a tuple with the AuthenticatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnUserFactorProfile) GetAuthenticatorNameOk() (*string, bool) {
	if o == nil || o.AuthenticatorName == nil {
		return nil, false
	}
	return o.AuthenticatorName, true
}

// HasAuthenticatorName returns a boolean if a field has been set.
func (o *WebAuthnUserFactorProfile) HasAuthenticatorName() bool {
	if o != nil && o.AuthenticatorName != nil {
		return true
	}

	return false
}

// SetAuthenticatorName gets a reference to the given string and assigns it to the AuthenticatorName field.
func (o *WebAuthnUserFactorProfile) SetAuthenticatorName(v string) {
	o.AuthenticatorName = &v
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *WebAuthnUserFactorProfile) GetCredentialId() string {
	if o == nil || o.CredentialId == nil {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnUserFactorProfile) GetCredentialIdOk() (*string, bool) {
	if o == nil || o.CredentialId == nil {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *WebAuthnUserFactorProfile) HasCredentialId() bool {
	if o != nil && o.CredentialId != nil {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *WebAuthnUserFactorProfile) SetCredentialId(v string) {
	o.CredentialId = &v
}

func (o WebAuthnUserFactorProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticatorName != nil {
		toSerialize["authenticatorName"] = o.AuthenticatorName
	}
	if o.CredentialId != nil {
		toSerialize["credentialId"] = o.CredentialId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WebAuthnUserFactorProfile) UnmarshalJSON(bytes []byte) (err error) {
	varWebAuthnUserFactorProfile := _WebAuthnUserFactorProfile{}

	err = json.Unmarshal(bytes, &varWebAuthnUserFactorProfile)
	if err == nil {
		*o = WebAuthnUserFactorProfile(varWebAuthnUserFactorProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authenticatorName")
		delete(additionalProperties, "credentialId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWebAuthnUserFactorProfile struct {
	value *WebAuthnUserFactorProfile
	isSet bool
}

func (v NullableWebAuthnUserFactorProfile) Get() *WebAuthnUserFactorProfile {
	return v.value
}

func (v *NullableWebAuthnUserFactorProfile) Set(val *WebAuthnUserFactorProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableWebAuthnUserFactorProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableWebAuthnUserFactorProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebAuthnUserFactorProfile(val *WebAuthnUserFactorProfile) *NullableWebAuthnUserFactorProfile {
	return &NullableWebAuthnUserFactorProfile{value: val, isSet: true}
}

func (v NullableWebAuthnUserFactorProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebAuthnUserFactorProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
