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

// checks if the Scim2HttpServletExtensionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Scim2HttpServletExtensionResponse{}

// Scim2HttpServletExtensionResponse struct for Scim2HttpServletExtensionResponse
type Scim2HttpServletExtensionResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []Enumscim2HttpServletExtensionSchemaUrn           `json:"schemas"`
	// Name of the HTTP Servlet Extension
	Id string `json:"id"`
	// The context path to use to access the SCIM 2.0 interface. The value must start with a forward slash and must represent a valid HTTP context path.
	BaseContextPath string `json:"baseContextPath"`
	// If specified, the Access Token Validator(s) that may be used to validate access tokens for requests submitted to this SCIM2 HTTP Servlet Extension.
	AccessTokenValidator        []string                                                 `json:"accessTokenValidator,omitempty"`
	MapAccessTokensToLocalUsers *EnumhttpServletExtensionMapAccessTokensToLocalUsersProp `json:"mapAccessTokensToLocalUsers,omitempty"`
	// Enables debug logging of the SCIM 2.0 SDK. Debug messages will be forwarded to the Directory Server debug logger with the scope of com.unboundid.directory.broker.http.scim2.extension.SCIM2HTTPServletExtension.
	DebugEnabled *bool                                   `json:"debugEnabled,omitempty"`
	DebugLevel   EnumhttpServletExtensionDebugLevelProp  `json:"debugLevel"`
	DebugType    []EnumhttpServletExtensionDebugTypeProp `json:"debugType"`
	// Indicates whether a stack trace of the thread which called the debug method should be included in debug log messages.
	IncludeStackTrace bool `json:"includeStackTrace"`
	// Indicates whether the SCIM2 HTTP Servlet Extension will generate a Swagger specification document.
	SwaggerEnabled *bool `json:"swaggerEnabled,omitempty"`
	// A description for this HTTP Servlet Extension
	Description *string `json:"description,omitempty"`
	// The cross-origin request policy to use for the HTTP Servlet Extension.
	CrossOriginPolicy *string `json:"crossOriginPolicy,omitempty"`
	// Specifies HTTP header fields and values added to response headers for all requests.
	ResponseHeader []string `json:"responseHeader,omitempty"`
	// Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \"Correlation-Id\", \"X-Amzn-Trace-Id\", and \"X-Request-Id\".
	CorrelationIDResponseHeader *string `json:"correlationIDResponseHeader,omitempty"`
}

// NewScim2HttpServletExtensionResponse instantiates a new Scim2HttpServletExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScim2HttpServletExtensionResponse(schemas []Enumscim2HttpServletExtensionSchemaUrn, id string, baseContextPath string, debugLevel EnumhttpServletExtensionDebugLevelProp, debugType []EnumhttpServletExtensionDebugTypeProp, includeStackTrace bool) *Scim2HttpServletExtensionResponse {
	this := Scim2HttpServletExtensionResponse{}
	this.Schemas = schemas
	this.Id = id
	this.BaseContextPath = baseContextPath
	this.DebugLevel = debugLevel
	this.DebugType = debugType
	this.IncludeStackTrace = includeStackTrace
	return &this
}

// NewScim2HttpServletExtensionResponseWithDefaults instantiates a new Scim2HttpServletExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScim2HttpServletExtensionResponseWithDefaults() *Scim2HttpServletExtensionResponse {
	this := Scim2HttpServletExtensionResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Scim2HttpServletExtensionResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Scim2HttpServletExtensionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *Scim2HttpServletExtensionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *Scim2HttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *Scim2HttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *Scim2HttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *Scim2HttpServletExtensionResponse) GetSchemas() []Enumscim2HttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []Enumscim2HttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetSchemasOk() ([]Enumscim2HttpServletExtensionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *Scim2HttpServletExtensionResponse) SetSchemas(v []Enumscim2HttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *Scim2HttpServletExtensionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Scim2HttpServletExtensionResponse) SetId(v string) {
	o.Id = v
}

// GetBaseContextPath returns the BaseContextPath field value
func (o *Scim2HttpServletExtensionResponse) GetBaseContextPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseContextPath
}

// GetBaseContextPathOk returns a tuple with the BaseContextPath field value
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetBaseContextPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseContextPath, true
}

// SetBaseContextPath sets field value
func (o *Scim2HttpServletExtensionResponse) SetBaseContextPath(v string) {
	o.BaseContextPath = v
}

// GetAccessTokenValidator returns the AccessTokenValidator field value if set, zero value otherwise.
func (o *Scim2HttpServletExtensionResponse) GetAccessTokenValidator() []string {
	if o == nil || IsNil(o.AccessTokenValidator) {
		var ret []string
		return ret
	}
	return o.AccessTokenValidator
}

// GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetAccessTokenValidatorOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessTokenValidator) {
		return nil, false
	}
	return o.AccessTokenValidator, true
}

// HasAccessTokenValidator returns a boolean if a field has been set.
func (o *Scim2HttpServletExtensionResponse) HasAccessTokenValidator() bool {
	if o != nil && !IsNil(o.AccessTokenValidator) {
		return true
	}

	return false
}

// SetAccessTokenValidator gets a reference to the given []string and assigns it to the AccessTokenValidator field.
func (o *Scim2HttpServletExtensionResponse) SetAccessTokenValidator(v []string) {
	o.AccessTokenValidator = v
}

// GetMapAccessTokensToLocalUsers returns the MapAccessTokensToLocalUsers field value if set, zero value otherwise.
func (o *Scim2HttpServletExtensionResponse) GetMapAccessTokensToLocalUsers() EnumhttpServletExtensionMapAccessTokensToLocalUsersProp {
	if o == nil || IsNil(o.MapAccessTokensToLocalUsers) {
		var ret EnumhttpServletExtensionMapAccessTokensToLocalUsersProp
		return ret
	}
	return *o.MapAccessTokensToLocalUsers
}

// GetMapAccessTokensToLocalUsersOk returns a tuple with the MapAccessTokensToLocalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetMapAccessTokensToLocalUsersOk() (*EnumhttpServletExtensionMapAccessTokensToLocalUsersProp, bool) {
	if o == nil || IsNil(o.MapAccessTokensToLocalUsers) {
		return nil, false
	}
	return o.MapAccessTokensToLocalUsers, true
}

// HasMapAccessTokensToLocalUsers returns a boolean if a field has been set.
func (o *Scim2HttpServletExtensionResponse) HasMapAccessTokensToLocalUsers() bool {
	if o != nil && !IsNil(o.MapAccessTokensToLocalUsers) {
		return true
	}

	return false
}

// SetMapAccessTokensToLocalUsers gets a reference to the given EnumhttpServletExtensionMapAccessTokensToLocalUsersProp and assigns it to the MapAccessTokensToLocalUsers field.
func (o *Scim2HttpServletExtensionResponse) SetMapAccessTokensToLocalUsers(v EnumhttpServletExtensionMapAccessTokensToLocalUsersProp) {
	o.MapAccessTokensToLocalUsers = &v
}

// GetDebugEnabled returns the DebugEnabled field value if set, zero value otherwise.
func (o *Scim2HttpServletExtensionResponse) GetDebugEnabled() bool {
	if o == nil || IsNil(o.DebugEnabled) {
		var ret bool
		return ret
	}
	return *o.DebugEnabled
}

// GetDebugEnabledOk returns a tuple with the DebugEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetDebugEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DebugEnabled) {
		return nil, false
	}
	return o.DebugEnabled, true
}

// HasDebugEnabled returns a boolean if a field has been set.
func (o *Scim2HttpServletExtensionResponse) HasDebugEnabled() bool {
	if o != nil && !IsNil(o.DebugEnabled) {
		return true
	}

	return false
}

// SetDebugEnabled gets a reference to the given bool and assigns it to the DebugEnabled field.
func (o *Scim2HttpServletExtensionResponse) SetDebugEnabled(v bool) {
	o.DebugEnabled = &v
}

// GetDebugLevel returns the DebugLevel field value
func (o *Scim2HttpServletExtensionResponse) GetDebugLevel() EnumhttpServletExtensionDebugLevelProp {
	if o == nil {
		var ret EnumhttpServletExtensionDebugLevelProp
		return ret
	}

	return o.DebugLevel
}

// GetDebugLevelOk returns a tuple with the DebugLevel field value
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetDebugLevelOk() (*EnumhttpServletExtensionDebugLevelProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DebugLevel, true
}

// SetDebugLevel sets field value
func (o *Scim2HttpServletExtensionResponse) SetDebugLevel(v EnumhttpServletExtensionDebugLevelProp) {
	o.DebugLevel = v
}

// GetDebugType returns the DebugType field value
func (o *Scim2HttpServletExtensionResponse) GetDebugType() []EnumhttpServletExtensionDebugTypeProp {
	if o == nil {
		var ret []EnumhttpServletExtensionDebugTypeProp
		return ret
	}

	return o.DebugType
}

// GetDebugTypeOk returns a tuple with the DebugType field value
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetDebugTypeOk() ([]EnumhttpServletExtensionDebugTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebugType, true
}

// SetDebugType sets field value
func (o *Scim2HttpServletExtensionResponse) SetDebugType(v []EnumhttpServletExtensionDebugTypeProp) {
	o.DebugType = v
}

// GetIncludeStackTrace returns the IncludeStackTrace field value
func (o *Scim2HttpServletExtensionResponse) GetIncludeStackTrace() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeStackTrace
}

// GetIncludeStackTraceOk returns a tuple with the IncludeStackTrace field value
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetIncludeStackTraceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeStackTrace, true
}

// SetIncludeStackTrace sets field value
func (o *Scim2HttpServletExtensionResponse) SetIncludeStackTrace(v bool) {
	o.IncludeStackTrace = v
}

// GetSwaggerEnabled returns the SwaggerEnabled field value if set, zero value otherwise.
func (o *Scim2HttpServletExtensionResponse) GetSwaggerEnabled() bool {
	if o == nil || IsNil(o.SwaggerEnabled) {
		var ret bool
		return ret
	}
	return *o.SwaggerEnabled
}

// GetSwaggerEnabledOk returns a tuple with the SwaggerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetSwaggerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SwaggerEnabled) {
		return nil, false
	}
	return o.SwaggerEnabled, true
}

// HasSwaggerEnabled returns a boolean if a field has been set.
func (o *Scim2HttpServletExtensionResponse) HasSwaggerEnabled() bool {
	if o != nil && !IsNil(o.SwaggerEnabled) {
		return true
	}

	return false
}

// SetSwaggerEnabled gets a reference to the given bool and assigns it to the SwaggerEnabled field.
func (o *Scim2HttpServletExtensionResponse) SetSwaggerEnabled(v bool) {
	o.SwaggerEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Scim2HttpServletExtensionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Scim2HttpServletExtensionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Scim2HttpServletExtensionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *Scim2HttpServletExtensionResponse) GetCrossOriginPolicy() string {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *Scim2HttpServletExtensionResponse) HasCrossOriginPolicy() bool {
	if o != nil && !IsNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *Scim2HttpServletExtensionResponse) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *Scim2HttpServletExtensionResponse) GetResponseHeader() []string {
	if o == nil || IsNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *Scim2HttpServletExtensionResponse) HasResponseHeader() bool {
	if o != nil && !IsNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *Scim2HttpServletExtensionResponse) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *Scim2HttpServletExtensionResponse) GetCorrelationIDResponseHeader() string {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scim2HttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *Scim2HttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool {
	if o != nil && !IsNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *Scim2HttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

func (o Scim2HttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Scim2HttpServletExtensionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["baseContextPath"] = o.BaseContextPath
	if !IsNil(o.AccessTokenValidator) {
		toSerialize["accessTokenValidator"] = o.AccessTokenValidator
	}
	if !IsNil(o.MapAccessTokensToLocalUsers) {
		toSerialize["mapAccessTokensToLocalUsers"] = o.MapAccessTokensToLocalUsers
	}
	if !IsNil(o.DebugEnabled) {
		toSerialize["debugEnabled"] = o.DebugEnabled
	}
	toSerialize["debugLevel"] = o.DebugLevel
	toSerialize["debugType"] = o.DebugType
	toSerialize["includeStackTrace"] = o.IncludeStackTrace
	if !IsNil(o.SwaggerEnabled) {
		toSerialize["swaggerEnabled"] = o.SwaggerEnabled
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
	return toSerialize, nil
}

type NullableScim2HttpServletExtensionResponse struct {
	value *Scim2HttpServletExtensionResponse
	isSet bool
}

func (v NullableScim2HttpServletExtensionResponse) Get() *Scim2HttpServletExtensionResponse {
	return v.value
}

func (v *NullableScim2HttpServletExtensionResponse) Set(val *Scim2HttpServletExtensionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScim2HttpServletExtensionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScim2HttpServletExtensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScim2HttpServletExtensionResponse(val *Scim2HttpServletExtensionResponse) *NullableScim2HttpServletExtensionResponse {
	return &NullableScim2HttpServletExtensionResponse{value: val, isSet: true}
}

func (v NullableScim2HttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScim2HttpServletExtensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
