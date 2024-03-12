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

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// ProvisioningConnectionProfileToken The app provisioning connection profile used to configure the method of authentication and the credentials. Currently, token-based and OAuth 2.0-based authentication are supported.
type ProvisioningConnectionProfileToken struct {
	// Defines the method of authentication
	AuthScheme string `json:"authScheme"`
	// Token used to authenticate with the app
	Token                string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionProfileToken ProvisioningConnectionProfileToken

// NewProvisioningConnectionProfileToken instantiates a new ProvisioningConnectionProfileToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionProfileToken(authScheme string, token string) *ProvisioningConnectionProfileToken {
	this := ProvisioningConnectionProfileToken{}
	this.AuthScheme = authScheme
	this.Token = token
	return &this
}

// NewProvisioningConnectionProfileTokenWithDefaults instantiates a new ProvisioningConnectionProfileToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionProfileTokenWithDefaults() *ProvisioningConnectionProfileToken {
	this := ProvisioningConnectionProfileToken{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value
func (o *ProvisioningConnectionProfileToken) GetAuthScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionProfileToken) GetAuthSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthScheme, true
}

// SetAuthScheme sets field value
func (o *ProvisioningConnectionProfileToken) SetAuthScheme(v string) {
	o.AuthScheme = v
}

// GetToken returns the Token field value
func (o *ProvisioningConnectionProfileToken) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionProfileToken) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ProvisioningConnectionProfileToken) SetToken(v string) {
	o.Token = v
}

func (o ProvisioningConnectionProfileToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authScheme"] = o.AuthScheme
	}
	if true {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConnectionProfileToken) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConnectionProfileToken := _ProvisioningConnectionProfileToken{}

	err = json.Unmarshal(bytes, &varProvisioningConnectionProfileToken)
	if err == nil {
		*o = ProvisioningConnectionProfileToken(varProvisioningConnectionProfileToken)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProvisioningConnectionProfileToken struct {
	value *ProvisioningConnectionProfileToken
	isSet bool
}

func (v NullableProvisioningConnectionProfileToken) Get() *ProvisioningConnectionProfileToken {
	return v.value
}

func (v *NullableProvisioningConnectionProfileToken) Set(val *ProvisioningConnectionProfileToken) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionProfileToken) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionProfileToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionProfileToken(val *ProvisioningConnectionProfileToken) *NullableProvisioningConnectionProfileToken {
	return &NullableProvisioningConnectionProfileToken{value: val, isSet: true}
}

func (v NullableProvisioningConnectionProfileToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionProfileToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}