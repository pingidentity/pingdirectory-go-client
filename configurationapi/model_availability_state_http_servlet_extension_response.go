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

// checks if the AvailabilityStateHttpServletExtensionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailabilityStateHttpServletExtensionResponse{}

// AvailabilityStateHttpServletExtensionResponse struct for AvailabilityStateHttpServletExtensionResponse
type AvailabilityStateHttpServletExtensionResponse struct {
	Schemas []EnumavailabilityStateHttpServletExtensionSchemaUrn `json:"schemas"`
	// Specifies the base context path that HTTP clients should use to access this servlet. The value must start with a forward slash and must represent a valid HTTP context path.
	BaseContextPath string `json:"baseContextPath"`
	// Specifies the HTTP status code that the servlet should return if the server considers itself to be available.
	AvailableStatusCode int64 `json:"availableStatusCode"`
	// Specifies the HTTP status code that the servlet should return if the server considers itself to be degraded.
	DegradedStatusCode int64 `json:"degradedStatusCode"`
	// Specifies the HTTP status code that the servlet should return if the server considers itself to be unavailable.
	UnavailableStatusCode int64 `json:"unavailableStatusCode"`
	// Specifies a HTTP status code that the servlet should always return, regardless of the server's availability. If this value is defined, it will override the availability-based return codes.
	OverrideStatusCode *int64 `json:"overrideStatusCode,omitempty"`
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
	CorrelationIDResponseHeader                   *string                                            `json:"correlationIDResponseHeader,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the HTTP Servlet Extension
	Id string `json:"id"`
}

// NewAvailabilityStateHttpServletExtensionResponse instantiates a new AvailabilityStateHttpServletExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityStateHttpServletExtensionResponse(schemas []EnumavailabilityStateHttpServletExtensionSchemaUrn, baseContextPath string, availableStatusCode int64, degradedStatusCode int64, unavailableStatusCode int64, id string) *AvailabilityStateHttpServletExtensionResponse {
	this := AvailabilityStateHttpServletExtensionResponse{}
	this.Schemas = schemas
	this.BaseContextPath = baseContextPath
	this.AvailableStatusCode = availableStatusCode
	this.DegradedStatusCode = degradedStatusCode
	this.UnavailableStatusCode = unavailableStatusCode
	this.Id = id
	return &this
}

// NewAvailabilityStateHttpServletExtensionResponseWithDefaults instantiates a new AvailabilityStateHttpServletExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityStateHttpServletExtensionResponseWithDefaults() *AvailabilityStateHttpServletExtensionResponse {
	this := AvailabilityStateHttpServletExtensionResponse{}
	return &this
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
func (o *AvailabilityStateHttpServletExtensionResponse) GetAvailableStatusCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AvailableStatusCode
}

// GetAvailableStatusCodeOk returns a tuple with the AvailableStatusCode field value
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetAvailableStatusCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableStatusCode, true
}

// SetAvailableStatusCode sets field value
func (o *AvailabilityStateHttpServletExtensionResponse) SetAvailableStatusCode(v int64) {
	o.AvailableStatusCode = v
}

// GetDegradedStatusCode returns the DegradedStatusCode field value
func (o *AvailabilityStateHttpServletExtensionResponse) GetDegradedStatusCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DegradedStatusCode
}

// GetDegradedStatusCodeOk returns a tuple with the DegradedStatusCode field value
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetDegradedStatusCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DegradedStatusCode, true
}

// SetDegradedStatusCode sets field value
func (o *AvailabilityStateHttpServletExtensionResponse) SetDegradedStatusCode(v int64) {
	o.DegradedStatusCode = v
}

// GetUnavailableStatusCode returns the UnavailableStatusCode field value
func (o *AvailabilityStateHttpServletExtensionResponse) GetUnavailableStatusCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnavailableStatusCode
}

// GetUnavailableStatusCodeOk returns a tuple with the UnavailableStatusCode field value
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetUnavailableStatusCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnavailableStatusCode, true
}

// SetUnavailableStatusCode sets field value
func (o *AvailabilityStateHttpServletExtensionResponse) SetUnavailableStatusCode(v int64) {
	o.UnavailableStatusCode = v
}

// GetOverrideStatusCode returns the OverrideStatusCode field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetOverrideStatusCode() int64 {
	if o == nil || IsNil(o.OverrideStatusCode) {
		var ret int64
		return ret
	}
	return *o.OverrideStatusCode
}

// GetOverrideStatusCodeOk returns a tuple with the OverrideStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetOverrideStatusCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.OverrideStatusCode) {
		return nil, false
	}
	return o.OverrideStatusCode, true
}

// HasOverrideStatusCode returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasOverrideStatusCode() bool {
	if o != nil && !IsNil(o.OverrideStatusCode) {
		return true
	}

	return false
}

// SetOverrideStatusCode gets a reference to the given int64 and assigns it to the OverrideStatusCode field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetOverrideStatusCode(v int64) {
	o.OverrideStatusCode = &v
}

// GetIncludeResponseBody returns the IncludeResponseBody field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetIncludeResponseBody() bool {
	if o == nil || IsNil(o.IncludeResponseBody) {
		var ret bool
		return ret
	}
	return *o.IncludeResponseBody
}

// GetIncludeResponseBodyOk returns a tuple with the IncludeResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetIncludeResponseBodyOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeResponseBody) {
		return nil, false
	}
	return o.IncludeResponseBody, true
}

// HasIncludeResponseBody returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasIncludeResponseBody() bool {
	if o != nil && !IsNil(o.IncludeResponseBody) {
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
	if o == nil || IsNil(o.AdditionalResponseContents) {
		var ret string
		return ret
	}
	return *o.AdditionalResponseContents
}

// GetAdditionalResponseContentsOk returns a tuple with the AdditionalResponseContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetAdditionalResponseContentsOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalResponseContents) {
		return nil, false
	}
	return o.AdditionalResponseContents, true
}

// HasAdditionalResponseContents returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasAdditionalResponseContents() bool {
	if o != nil && !IsNil(o.AdditionalResponseContents) {
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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
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
	if o == nil || IsNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasCrossOriginPolicy() bool {
	if o != nil && !IsNil(o.CrossOriginPolicy) {
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
	if o == nil || IsNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasResponseHeader() bool {
	if o != nil && !IsNil(o.ResponseHeader) {
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
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool {
	if o != nil && !IsNil(o.CorrelationIDResponseHeader) {
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
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AvailabilityStateHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AvailabilityStateHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AvailabilityStateHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
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

func (o AvailabilityStateHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailabilityStateHttpServletExtensionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["baseContextPath"] = o.BaseContextPath
	toSerialize["availableStatusCode"] = o.AvailableStatusCode
	toSerialize["degradedStatusCode"] = o.DegradedStatusCode
	toSerialize["unavailableStatusCode"] = o.UnavailableStatusCode
	if !IsNil(o.OverrideStatusCode) {
		toSerialize["overrideStatusCode"] = o.OverrideStatusCode
	}
	if !IsNil(o.IncludeResponseBody) {
		toSerialize["includeResponseBody"] = o.IncludeResponseBody
	}
	if !IsNil(o.AdditionalResponseContents) {
		toSerialize["additionalResponseContents"] = o.AdditionalResponseContents
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CrossOriginPolicy) {
		toSerialize["crossOriginPolicy"] = o.CrossOriginPolicy
	}
	if !IsNil(o.ResponseHeader) {
		toSerialize["responseHeader"] = o.ResponseHeader
	}
	if !IsNil(o.CorrelationIDResponseHeader) {
		toSerialize["correlationIDResponseHeader"] = o.CorrelationIDResponseHeader
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
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
