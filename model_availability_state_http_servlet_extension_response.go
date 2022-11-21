/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AvailabilityStateHttpServletExtensionResponse struct for AvailabilityStateHttpServletExtensionResponse
type AvailabilityStateHttpServletExtensionResponse struct {
	// Name of the HTTP Servlet Extension
	Id string `json:"id"`
	Schemas []EnumavailabilityStateHttpServletExtensionSchemaUrn `json:"schemas"`
	// Specifies the base context path that HTTP clients should use to access this servlet. The value must start with a forward slash and must represent a valid HTTP context path.
	BaseContextPath string `json:"baseContextPath"`
	// Specifies the HTTP status code that the servlet should return if the server considers itself to be available.
	AvailableStatusCode int32 `json:"availableStatusCode"`
	// Specifies the HTTP status code that the servlet should return if the server considers itself to be degraded.
	DegradedStatusCode int32 `json:"degradedStatusCode"`
	// Specifies the HTTP status code that the servlet should return if the server considers itself to be unavailable.
	UnavailableStatusCode int32 `json:"unavailableStatusCode"`
	// Specifies a HTTP status code that the servlet should always return, regardless of the server's availability. If this value is defined, it will override the availability-based return codes.
	OverrideStatusCode *int32 `json:"overrideStatusCode,omitempty"`
	// Indicates whether the response should include a body that is a JSON object.
	IncludeResponseBody *bool `json:"includeResponseBody,omitempty"`
	// A JSON-formatted string containing additional fields to be returned in the response body. For example, an additional-response-contents value of '{ \"key\": \"value\" }' would add the key and value to the root of the JSON response body.
	AdditionalResponseContents *string `json:"additionalResponseContents,omitempty"`
	// A description for this HTTP Servlet Extension
	Description *string `json:"description,omitempty"`
	// The cross-origin request policy to use for the HTTP Servlet Extension.
	CrossOriginPolicy *string `json:"crossOriginPolicy,omitempty"`
	// Specifies HTTP header fields and values added to response headers for all requests.
	ResponseHeader []string `json:"responseHeader,omitempty"`
	// Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \"Correlation-Id\", \"X-Amzn-Trace-Id\", and \"X-Request-Id\".
	CorrelationIDResponseHeader *string `json:"correlationIDResponseHeader,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewAvailabilityStateHttpServletExtensionResponse instantiates a new AvailabilityStateHttpServletExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityStateHttpServletExtensionResponse(id string, schemas []EnumavailabilityStateHttpServletExtensionSchemaUrn, baseContextPath string, availableStatusCode int32, degradedStatusCode int32, unavailableStatusCode int32) *AvailabilityStateHttpServletExtensionResponse {
	this := AvailabilityStateHttpServletExtensionResponse{}
	this.Id = id
	this.Schemas = schemas
	this.BaseContextPath = baseContextPath
	this.AvailableStatusCode = availableStatusCode
	this.DegradedStatusCode = degradedStatusCode
	this.UnavailableStatusCode = unavailableStatusCode
	return &this
}

// NewAvailabilityStateHttpServletExtensionResponseWithDefaults instantiates a new AvailabilityStateHttpServletExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityStateHttpServletExtensionResponseWithDefaults() *AvailabilityStateHttpServletExtensionResponse {
	this := AvailabilityStateHttpServletExtensionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AvailabilityStateHttpServletExtensionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AvailabilityStateHttpServletExtensionResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *AvailabilityStateHttpServletExtensionResponse) GetSchemas() []EnumavailabilityStateHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumavailabilityStateHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetSchemasOk() ([]EnumavailabilityStateHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AvailabilityStateHttpServletExtensionResponse) SetSchemas(v []EnumavailabilityStateHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetBaseContextPath returns the BaseContextPath field value
func (o *AvailabilityStateHttpServletExtensionResponse) GetBaseContextPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseContextPath
}

// GetBaseContextPathOk returns a tuple with the BaseContextPath field value
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetBaseContextPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BaseContextPath, true
}

// SetBaseContextPath sets field value
func (o *AvailabilityStateHttpServletExtensionResponse) SetBaseContextPath(v string) {
	o.BaseContextPath = v
}

// GetAvailableStatusCode returns the AvailableStatusCode field value
func (o *AvailabilityStateHttpServletExtensionResponse) GetAvailableStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AvailableStatusCode
}

// GetAvailableStatusCodeOk returns a tuple with the AvailableStatusCode field value
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetAvailableStatusCodeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AvailableStatusCode, true
}

// SetAvailableStatusCode sets field value
func (o *AvailabilityStateHttpServletExtensionResponse) SetAvailableStatusCode(v int32) {
	o.AvailableStatusCode = v
}

// GetDegradedStatusCode returns the DegradedStatusCode field value
func (o *AvailabilityStateHttpServletExtensionResponse) GetDegradedStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DegradedStatusCode
}

// GetDegradedStatusCodeOk returns a tuple with the DegradedStatusCode field value
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetDegradedStatusCodeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DegradedStatusCode, true
}

// SetDegradedStatusCode sets field value
func (o *AvailabilityStateHttpServletExtensionResponse) SetDegradedStatusCode(v int32) {
	o.DegradedStatusCode = v
}

// GetUnavailableStatusCode returns the UnavailableStatusCode field value
func (o *AvailabilityStateHttpServletExtensionResponse) GetUnavailableStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnavailableStatusCode
}

// GetUnavailableStatusCodeOk returns a tuple with the UnavailableStatusCode field value
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetUnavailableStatusCodeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UnavailableStatusCode, true
}

// SetUnavailableStatusCode sets field value
func (o *AvailabilityStateHttpServletExtensionResponse) SetUnavailableStatusCode(v int32) {
	o.UnavailableStatusCode = v
}

// GetOverrideStatusCode returns the OverrideStatusCode field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetOverrideStatusCode() int32 {
	if o == nil || isNil(o.OverrideStatusCode) {
		var ret int32
		return ret
	}
	return *o.OverrideStatusCode
}

// GetOverrideStatusCodeOk returns a tuple with the OverrideStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetOverrideStatusCodeOk() (*int32, bool) {
	if o == nil || isNil(o.OverrideStatusCode) {
    return nil, false
	}
	return o.OverrideStatusCode, true
}

// HasOverrideStatusCode returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasOverrideStatusCode() bool {
	if o != nil && !isNil(o.OverrideStatusCode) {
		return true
	}

	return false
}

// SetOverrideStatusCode gets a reference to the given int32 and assigns it to the OverrideStatusCode field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetOverrideStatusCode(v int32) {
	o.OverrideStatusCode = &v
}

// GetIncludeResponseBody returns the IncludeResponseBody field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetIncludeResponseBody() bool {
	if o == nil || isNil(o.IncludeResponseBody) {
		var ret bool
		return ret
	}
	return *o.IncludeResponseBody
}

// GetIncludeResponseBodyOk returns a tuple with the IncludeResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetIncludeResponseBodyOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeResponseBody) {
    return nil, false
	}
	return o.IncludeResponseBody, true
}

// HasIncludeResponseBody returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasIncludeResponseBody() bool {
	if o != nil && !isNil(o.IncludeResponseBody) {
		return true
	}

	return false
}

// SetIncludeResponseBody gets a reference to the given bool and assigns it to the IncludeResponseBody field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetIncludeResponseBody(v bool) {
	o.IncludeResponseBody = &v
}

// GetAdditionalResponseContents returns the AdditionalResponseContents field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetAdditionalResponseContents() string {
	if o == nil || isNil(o.AdditionalResponseContents) {
		var ret string
		return ret
	}
	return *o.AdditionalResponseContents
}

// GetAdditionalResponseContentsOk returns a tuple with the AdditionalResponseContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetAdditionalResponseContentsOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalResponseContents) {
    return nil, false
	}
	return o.AdditionalResponseContents, true
}

// HasAdditionalResponseContents returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasAdditionalResponseContents() bool {
	if o != nil && !isNil(o.AdditionalResponseContents) {
		return true
	}

	return false
}

// SetAdditionalResponseContents gets a reference to the given string and assigns it to the AdditionalResponseContents field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetAdditionalResponseContents(v string) {
	o.AdditionalResponseContents = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetCrossOriginPolicy() string {
	if o == nil || isNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || isNil(o.CrossOriginPolicy) {
    return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasCrossOriginPolicy() bool {
	if o != nil && !isNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetResponseHeader() []string {
	if o == nil || isNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || isNil(o.ResponseHeader) {
    return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasResponseHeader() bool {
	if o != nil && !isNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
    return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool {
	if o != nil && !isNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o AvailabilityStateHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["baseContextPath"] = o.BaseContextPath
	}
	if true {
		toSerialize["availableStatusCode"] = o.AvailableStatusCode
	}
	if true {
		toSerialize["degradedStatusCode"] = o.DegradedStatusCode
	}
	if true {
		toSerialize["unavailableStatusCode"] = o.UnavailableStatusCode
	}
	if !isNil(o.OverrideStatusCode) {
		toSerialize["overrideStatusCode"] = o.OverrideStatusCode
	}
	if !isNil(o.IncludeResponseBody) {
		toSerialize["includeResponseBody"] = o.IncludeResponseBody
	}
	if !isNil(o.AdditionalResponseContents) {
		toSerialize["additionalResponseContents"] = o.AdditionalResponseContents
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.CrossOriginPolicy) {
		toSerialize["crossOriginPolicy"] = o.CrossOriginPolicy
	}
	if !isNil(o.ResponseHeader) {
		toSerialize["responseHeader"] = o.ResponseHeader
	}
	if !isNil(o.CorrelationIDResponseHeader) {
		toSerialize["correlationIDResponseHeader"] = o.CorrelationIDResponseHeader
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableAvailabilityStateHttpServletExtensionResponse struct {
	value *AvailabilityStateHttpServletExtensionResponse
	isSet bool
}

func (v NullableAvailabilityStateHttpServletExtensionResponse) Get() *AvailabilityStateHttpServletExtensionResponse {
	return v.value
}

func (v *NullableAvailabilityStateHttpServletExtensionResponse) Set(val *AvailabilityStateHttpServletExtensionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityStateHttpServletExtensionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityStateHttpServletExtensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityStateHttpServletExtensionResponse(val *AvailabilityStateHttpServletExtensionResponse) *NullableAvailabilityStateHttpServletExtensionResponse {
	return &NullableAvailabilityStateHttpServletExtensionResponse{value: val, isSet: true}
}

func (v NullableAvailabilityStateHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityStateHttpServletExtensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


