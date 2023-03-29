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

// checks if the ApiKeyConjurAuthenticationMethodResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyConjurAuthenticationMethodResponse{}

// ApiKeyConjurAuthenticationMethodResponse struct for ApiKeyConjurAuthenticationMethodResponse
type ApiKeyConjurAuthenticationMethodResponse struct {
	// Name of the Conjur Authentication Method
	Id      string                                          `json:"id"`
	Schemas []EnumapiKeyConjurAuthenticationMethodSchemaUrn `json:"schemas,omitempty"`
	// The username for the user to authenticate.
	Username string `json:"username"`
	// The password for the user to authenticate. This will be used to obtain an API key for the target user.
	Password *string `json:"password,omitempty"`
	// The API key for the user to authenticate.
	ApiKey *string `json:"apiKey,omitempty"`
	// A description for this Conjur Authentication Method
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewApiKeyConjurAuthenticationMethodResponse instantiates a new ApiKeyConjurAuthenticationMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyConjurAuthenticationMethodResponse(id string, username string) *ApiKeyConjurAuthenticationMethodResponse {
	this := ApiKeyConjurAuthenticationMethodResponse{}
	this.Id = id
	this.Username = username
	return &this
}

// NewApiKeyConjurAuthenticationMethodResponseWithDefaults instantiates a new ApiKeyConjurAuthenticationMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyConjurAuthenticationMethodResponseWithDefaults() *ApiKeyConjurAuthenticationMethodResponse {
	this := ApiKeyConjurAuthenticationMethodResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ApiKeyConjurAuthenticationMethodResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiKeyConjurAuthenticationMethodResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetSchemas() []EnumapiKeyConjurAuthenticationMethodSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumapiKeyConjurAuthenticationMethodSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetSchemasOk() ([]EnumapiKeyConjurAuthenticationMethodSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumapiKeyConjurAuthenticationMethodSchemaUrn and assigns it to the Schemas field.
func (o *ApiKeyConjurAuthenticationMethodResponse) SetSchemas(v []EnumapiKeyConjurAuthenticationMethodSchemaUrn) {
	o.Schemas = v
}

// GetUsername returns the Username field value
func (o *ApiKeyConjurAuthenticationMethodResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ApiKeyConjurAuthenticationMethodResponse) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiKeyConjurAuthenticationMethodResponse) SetPassword(v string) {
	o.Password = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *ApiKeyConjurAuthenticationMethodResponse) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiKeyConjurAuthenticationMethodResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ApiKeyConjurAuthenticationMethodResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ApiKeyConjurAuthenticationMethodResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ApiKeyConjurAuthenticationMethodResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o ApiKeyConjurAuthenticationMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyConjurAuthenticationMethodResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["username"] = o.Username
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableApiKeyConjurAuthenticationMethodResponse struct {
	value *ApiKeyConjurAuthenticationMethodResponse
	isSet bool
}

func (v NullableApiKeyConjurAuthenticationMethodResponse) Get() *ApiKeyConjurAuthenticationMethodResponse {
	return v.value
}

func (v *NullableApiKeyConjurAuthenticationMethodResponse) Set(val *ApiKeyConjurAuthenticationMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyConjurAuthenticationMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyConjurAuthenticationMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyConjurAuthenticationMethodResponse(val *ApiKeyConjurAuthenticationMethodResponse) *NullableApiKeyConjurAuthenticationMethodResponse {
	return &NullableApiKeyConjurAuthenticationMethodResponse{value: val, isSet: true}
}

func (v NullableApiKeyConjurAuthenticationMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyConjurAuthenticationMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
