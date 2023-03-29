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

// checks if the DelegatedAdminHttpServletExtensionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelegatedAdminHttpServletExtensionResponse{}

// DelegatedAdminHttpServletExtensionResponse struct for DelegatedAdminHttpServletExtensionResponse
type DelegatedAdminHttpServletExtensionResponse struct {
	Schemas []EnumdelegatedAdminHttpServletExtensionSchemaUrn `json:"schemas"`
	// Name of the HTTP Servlet Extension
	Id string `json:"id"`
	// Enables HTTP Basic authentication, using a username and password. The Identity Mapper specified by the identity-mapper property will be used to map the username to a DN.
	BasicAuthEnabled *bool `json:"basicAuthEnabled,omitempty"`
	// Specifies the Identity Mapper that is to be used for associating user entries with basic authentication user names.
	IdentityMapper *string `json:"identityMapper,omitempty"`
	// If specified, the Access Token Validator(s) that may be used to validate access tokens for requests submitted to this Delegated Admin HTTP Servlet Extension.
	AccessTokenValidator []string `json:"accessTokenValidator,omitempty"`
	// The name of a scope that must be present in an access token accepted by the Delegated Admin HTTP Servlet Extension.
	AccessTokenScope *string `json:"accessTokenScope,omitempty"`
	// A string or URI that identifies the Delegated Admin HTTP Servlet Extension in the context of OAuth2 authorization.
	Audience *string `json:"audience,omitempty"`
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
}

// NewDelegatedAdminHttpServletExtensionResponse instantiates a new DelegatedAdminHttpServletExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegatedAdminHttpServletExtensionResponse(schemas []EnumdelegatedAdminHttpServletExtensionSchemaUrn, id string) *DelegatedAdminHttpServletExtensionResponse {
	this := DelegatedAdminHttpServletExtensionResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewDelegatedAdminHttpServletExtensionResponseWithDefaults instantiates a new DelegatedAdminHttpServletExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegatedAdminHttpServletExtensionResponseWithDefaults() *DelegatedAdminHttpServletExtensionResponse {
	this := DelegatedAdminHttpServletExtensionResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *DelegatedAdminHttpServletExtensionResponse) GetSchemas() []EnumdelegatedAdminHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumdelegatedAdminHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetSchemasOk() ([]EnumdelegatedAdminHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DelegatedAdminHttpServletExtensionResponse) SetSchemas(v []EnumdelegatedAdminHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *DelegatedAdminHttpServletExtensionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DelegatedAdminHttpServletExtensionResponse) SetId(v string) {
	o.Id = v
}

// GetBasicAuthEnabled returns the BasicAuthEnabled field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetBasicAuthEnabled() bool {
	if o == nil || IsNil(o.BasicAuthEnabled) {
		var ret bool
		return ret
	}
	return *o.BasicAuthEnabled
}

// GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetBasicAuthEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BasicAuthEnabled) {
		return nil, false
	}
	return o.BasicAuthEnabled, true
}

// HasBasicAuthEnabled returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasBasicAuthEnabled() bool {
	if o != nil && !IsNil(o.BasicAuthEnabled) {
		return true
	}

	return false
}

// SetBasicAuthEnabled gets a reference to the given bool and assigns it to the BasicAuthEnabled field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetBasicAuthEnabled(v bool) {
	o.BasicAuthEnabled = &v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetIdentityMapper() string {
	if o == nil || IsNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityMapper) {
		return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasIdentityMapper() bool {
	if o != nil && !IsNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetAccessTokenValidator returns the AccessTokenValidator field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetAccessTokenValidator() []string {
	if o == nil || IsNil(o.AccessTokenValidator) {
		var ret []string
		return ret
	}
	return o.AccessTokenValidator
}

// GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetAccessTokenValidatorOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessTokenValidator) {
		return nil, false
	}
	return o.AccessTokenValidator, true
}

// HasAccessTokenValidator returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasAccessTokenValidator() bool {
	if o != nil && !IsNil(o.AccessTokenValidator) {
		return true
	}

	return false
}

// SetAccessTokenValidator gets a reference to the given []string and assigns it to the AccessTokenValidator field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetAccessTokenValidator(v []string) {
	o.AccessTokenValidator = v
}

// GetAccessTokenScope returns the AccessTokenScope field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetAccessTokenScope() string {
	if o == nil || IsNil(o.AccessTokenScope) {
		var ret string
		return ret
	}
	return *o.AccessTokenScope
}

// GetAccessTokenScopeOk returns a tuple with the AccessTokenScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetAccessTokenScopeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessTokenScope) {
		return nil, false
	}
	return o.AccessTokenScope, true
}

// HasAccessTokenScope returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasAccessTokenScope() bool {
	if o != nil && !IsNil(o.AccessTokenScope) {
		return true
	}

	return false
}

// SetAccessTokenScope gets a reference to the given string and assigns it to the AccessTokenScope field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetAccessTokenScope(v string) {
	o.AccessTokenScope = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetAudience() string {
	if o == nil || IsNil(o.Audience) {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetAudienceOk() (*string, bool) {
	if o == nil || IsNil(o.Audience) {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasAudience() bool {
	if o != nil && !IsNil(o.Audience) {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetAudience(v string) {
	o.Audience = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetCrossOriginPolicy() string {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasCrossOriginPolicy() bool {
	if o != nil && !IsNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetResponseHeader() []string {
	if o == nil || IsNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasResponseHeader() bool {
	if o != nil && !IsNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool {
	if o != nil && !IsNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DelegatedAdminHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DelegatedAdminHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DelegatedAdminHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o DelegatedAdminHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelegatedAdminHttpServletExtensionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.BasicAuthEnabled) {
		toSerialize["basicAuthEnabled"] = o.BasicAuthEnabled
	}
	if !IsNil(o.IdentityMapper) {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !IsNil(o.AccessTokenValidator) {
		toSerialize["accessTokenValidator"] = o.AccessTokenValidator
	}
	if !IsNil(o.AccessTokenScope) {
		toSerialize["accessTokenScope"] = o.AccessTokenScope
	}
	if !IsNil(o.Audience) {
		toSerialize["audience"] = o.Audience
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
	return toSerialize, nil
}

type NullableDelegatedAdminHttpServletExtensionResponse struct {
	value *DelegatedAdminHttpServletExtensionResponse
	isSet bool
}

func (v NullableDelegatedAdminHttpServletExtensionResponse) Get() *DelegatedAdminHttpServletExtensionResponse {
	return v.value
}

func (v *NullableDelegatedAdminHttpServletExtensionResponse) Set(val *DelegatedAdminHttpServletExtensionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegatedAdminHttpServletExtensionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegatedAdminHttpServletExtensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegatedAdminHttpServletExtensionResponse(val *DelegatedAdminHttpServletExtensionResponse) *NullableDelegatedAdminHttpServletExtensionResponse {
	return &NullableDelegatedAdminHttpServletExtensionResponse{value: val, isSet: true}
}

func (v NullableDelegatedAdminHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegatedAdminHttpServletExtensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
