/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddHttpServletCrossOriginPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddHttpServletCrossOriginPolicyRequest{}

// AddHttpServletCrossOriginPolicyRequest struct for AddHttpServletCrossOriginPolicyRequest
type AddHttpServletCrossOriginPolicyRequest struct {
	Schemas []EnumhttpServletCrossOriginPolicySchemaUrn `json:"schemas,omitempty"`
	// A description for this HTTP Servlet Cross Origin Policy
	Description *string `json:"description,omitempty"`
	// A list of HTTP methods allowed for cross-origin access to resources. i.e. one or more of GET, POST, PUT, DELETE, etc.
	CorsAllowedMethods []string `json:"corsAllowedMethods,omitempty"`
	// A list of origins that are allowed to execute cross-origin requests.
	CorsAllowedOrigins []string `json:"corsAllowedOrigins,omitempty"`
	// A list of HTTP headers other than the simple response headers that browsers are allowed to access.
	CorsExposedHeaders []string `json:"corsExposedHeaders,omitempty"`
	// A list of HTTP headers that are supported by the resource and can be specified in a cross-origin request.
	CorsAllowedHeaders []string `json:"corsAllowedHeaders,omitempty"`
	// The maximum amount of time that a preflight request can be cached by a client.
	CorsPreflightMaxAge *string `json:"corsPreflightMaxAge,omitempty"`
	// Indicates whether the servlet extension allows CORS requests with username/password credentials.
	CorsAllowCredentials *bool `json:"corsAllowCredentials,omitempty"`
	// Name of the new HTTP Servlet Cross Origin Policy
	PolicyName string `json:"policyName"`
}

// NewAddHttpServletCrossOriginPolicyRequest instantiates a new AddHttpServletCrossOriginPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddHttpServletCrossOriginPolicyRequest(policyName string) *AddHttpServletCrossOriginPolicyRequest {
	this := AddHttpServletCrossOriginPolicyRequest{}
	this.PolicyName = policyName
	return &this
}

// NewAddHttpServletCrossOriginPolicyRequestWithDefaults instantiates a new AddHttpServletCrossOriginPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddHttpServletCrossOriginPolicyRequestWithDefaults() *AddHttpServletCrossOriginPolicyRequest {
	this := AddHttpServletCrossOriginPolicyRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddHttpServletCrossOriginPolicyRequest) GetSchemas() []EnumhttpServletCrossOriginPolicySchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumhttpServletCrossOriginPolicySchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) GetSchemasOk() ([]EnumhttpServletCrossOriginPolicySchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumhttpServletCrossOriginPolicySchemaUrn and assigns it to the Schemas field.
func (o *AddHttpServletCrossOriginPolicyRequest) SetSchemas(v []EnumhttpServletCrossOriginPolicySchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddHttpServletCrossOriginPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddHttpServletCrossOriginPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCorsAllowedMethods returns the CorsAllowedMethods field value if set, zero value otherwise.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedMethods() []string {
	if o == nil || IsNil(o.CorsAllowedMethods) {
		var ret []string
		return ret
	}
	return o.CorsAllowedMethods
}

// GetCorsAllowedMethodsOk returns a tuple with the CorsAllowedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.CorsAllowedMethods) {
		return nil, false
	}
	return o.CorsAllowedMethods, true
}

// HasCorsAllowedMethods returns a boolean if a field has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsAllowedMethods() bool {
	if o != nil && !IsNil(o.CorsAllowedMethods) {
		return true
	}

	return false
}

// SetCorsAllowedMethods gets a reference to the given []string and assigns it to the CorsAllowedMethods field.
func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsAllowedMethods(v []string) {
	o.CorsAllowedMethods = v
}

// GetCorsAllowedOrigins returns the CorsAllowedOrigins field value if set, zero value otherwise.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedOrigins() []string {
	if o == nil || IsNil(o.CorsAllowedOrigins) {
		var ret []string
		return ret
	}
	return o.CorsAllowedOrigins
}

// GetCorsAllowedOriginsOk returns a tuple with the CorsAllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedOriginsOk() ([]string, bool) {
	if o == nil || IsNil(o.CorsAllowedOrigins) {
		return nil, false
	}
	return o.CorsAllowedOrigins, true
}

// HasCorsAllowedOrigins returns a boolean if a field has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsAllowedOrigins() bool {
	if o != nil && !IsNil(o.CorsAllowedOrigins) {
		return true
	}

	return false
}

// SetCorsAllowedOrigins gets a reference to the given []string and assigns it to the CorsAllowedOrigins field.
func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsAllowedOrigins(v []string) {
	o.CorsAllowedOrigins = v
}

// GetCorsExposedHeaders returns the CorsExposedHeaders field value if set, zero value otherwise.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsExposedHeaders() []string {
	if o == nil || IsNil(o.CorsExposedHeaders) {
		var ret []string
		return ret
	}
	return o.CorsExposedHeaders
}

// GetCorsExposedHeadersOk returns a tuple with the CorsExposedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsExposedHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.CorsExposedHeaders) {
		return nil, false
	}
	return o.CorsExposedHeaders, true
}

// HasCorsExposedHeaders returns a boolean if a field has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsExposedHeaders() bool {
	if o != nil && !IsNil(o.CorsExposedHeaders) {
		return true
	}

	return false
}

// SetCorsExposedHeaders gets a reference to the given []string and assigns it to the CorsExposedHeaders field.
func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsExposedHeaders(v []string) {
	o.CorsExposedHeaders = v
}

// GetCorsAllowedHeaders returns the CorsAllowedHeaders field value if set, zero value otherwise.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedHeaders() []string {
	if o == nil || IsNil(o.CorsAllowedHeaders) {
		var ret []string
		return ret
	}
	return o.CorsAllowedHeaders
}

// GetCorsAllowedHeadersOk returns a tuple with the CorsAllowedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.CorsAllowedHeaders) {
		return nil, false
	}
	return o.CorsAllowedHeaders, true
}

// HasCorsAllowedHeaders returns a boolean if a field has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsAllowedHeaders() bool {
	if o != nil && !IsNil(o.CorsAllowedHeaders) {
		return true
	}

	return false
}

// SetCorsAllowedHeaders gets a reference to the given []string and assigns it to the CorsAllowedHeaders field.
func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsAllowedHeaders(v []string) {
	o.CorsAllowedHeaders = v
}

// GetCorsPreflightMaxAge returns the CorsPreflightMaxAge field value if set, zero value otherwise.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsPreflightMaxAge() string {
	if o == nil || IsNil(o.CorsPreflightMaxAge) {
		var ret string
		return ret
	}
	return *o.CorsPreflightMaxAge
}

// GetCorsPreflightMaxAgeOk returns a tuple with the CorsPreflightMaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsPreflightMaxAgeOk() (*string, bool) {
	if o == nil || IsNil(o.CorsPreflightMaxAge) {
		return nil, false
	}
	return o.CorsPreflightMaxAge, true
}

// HasCorsPreflightMaxAge returns a boolean if a field has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsPreflightMaxAge() bool {
	if o != nil && !IsNil(o.CorsPreflightMaxAge) {
		return true
	}

	return false
}

// SetCorsPreflightMaxAge gets a reference to the given string and assigns it to the CorsPreflightMaxAge field.
func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsPreflightMaxAge(v string) {
	o.CorsPreflightMaxAge = &v
}

// GetCorsAllowCredentials returns the CorsAllowCredentials field value if set, zero value otherwise.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowCredentials() bool {
	if o == nil || IsNil(o.CorsAllowCredentials) {
		var ret bool
		return ret
	}
	return *o.CorsAllowCredentials
}

// GetCorsAllowCredentialsOk returns a tuple with the CorsAllowCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowCredentialsOk() (*bool, bool) {
	if o == nil || IsNil(o.CorsAllowCredentials) {
		return nil, false
	}
	return o.CorsAllowCredentials, true
}

// HasCorsAllowCredentials returns a boolean if a field has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsAllowCredentials() bool {
	if o != nil && !IsNil(o.CorsAllowCredentials) {
		return true
	}

	return false
}

// SetCorsAllowCredentials gets a reference to the given bool and assigns it to the CorsAllowCredentials field.
func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsAllowCredentials(v bool) {
	o.CorsAllowCredentials = &v
}

// GetPolicyName returns the PolicyName field value
func (o *AddHttpServletCrossOriginPolicyRequest) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *AddHttpServletCrossOriginPolicyRequest) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value
func (o *AddHttpServletCrossOriginPolicyRequest) SetPolicyName(v string) {
	o.PolicyName = v
}

func (o AddHttpServletCrossOriginPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddHttpServletCrossOriginPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CorsAllowedMethods) {
		toSerialize["corsAllowedMethods"] = o.CorsAllowedMethods
	}
	if !IsNil(o.CorsAllowedOrigins) {
		toSerialize["corsAllowedOrigins"] = o.CorsAllowedOrigins
	}
	if !IsNil(o.CorsExposedHeaders) {
		toSerialize["corsExposedHeaders"] = o.CorsExposedHeaders
	}
	if !IsNil(o.CorsAllowedHeaders) {
		toSerialize["corsAllowedHeaders"] = o.CorsAllowedHeaders
	}
	if !IsNil(o.CorsPreflightMaxAge) {
		toSerialize["corsPreflightMaxAge"] = o.CorsPreflightMaxAge
	}
	if !IsNil(o.CorsAllowCredentials) {
		toSerialize["corsAllowCredentials"] = o.CorsAllowCredentials
	}
	toSerialize["policyName"] = o.PolicyName
	return toSerialize, nil
}

type NullableAddHttpServletCrossOriginPolicyRequest struct {
	value *AddHttpServletCrossOriginPolicyRequest
	isSet bool
}

func (v NullableAddHttpServletCrossOriginPolicyRequest) Get() *AddHttpServletCrossOriginPolicyRequest {
	return v.value
}

func (v *NullableAddHttpServletCrossOriginPolicyRequest) Set(val *AddHttpServletCrossOriginPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddHttpServletCrossOriginPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddHttpServletCrossOriginPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddHttpServletCrossOriginPolicyRequest(val *AddHttpServletCrossOriginPolicyRequest) *NullableAddHttpServletCrossOriginPolicyRequest {
	return &NullableAddHttpServletCrossOriginPolicyRequest{value: val, isSet: true}
}

func (v NullableAddHttpServletCrossOriginPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddHttpServletCrossOriginPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
