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

// checks if the ConjurExternalServerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConjurExternalServerResponse{}

// ConjurExternalServerResponse struct for ConjurExternalServerResponse
type ConjurExternalServerResponse struct {
	Schemas []EnumconjurExternalServerSchemaUrn `json:"schemas"`
	// The base URL needed to access the CyberArk Conjur server. The base URL should consist of the protocol (\"http\" or \"https\"), the server address (resolvable name or IP address), and the port number. For example, \"https://conjur.example.com:8443/\".
	ConjurServerBaseURI []string `json:"conjurServerBaseURI"`
	// The mechanism used to authenticate to the Conjur server.
	ConjurAuthenticationMethod string `json:"conjurAuthenticationMethod"`
	// The name of the account with which the desired secrets are associated.
	ConjurAccountName string `json:"conjurAccountName"`
	// The maximum length of time to wait to obtain an HTTP connection.
	HttpConnectTimeout *string `json:"httpConnectTimeout,omitempty"`
	// The maximum length of time to wait for a response to an HTTP request.
	HttpResponseTimeout *string `json:"httpResponseTimeout,omitempty"`
	// The path to a file containing the information needed to trust the certificate presented by the Conjur servers.
	TrustStoreFile *string `json:"trustStoreFile,omitempty"`
	// The PIN needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents.
	TrustStorePin *string `json:"trustStorePin,omitempty"`
	// The store type for the specified trust store file. The value should likely be one of \"JKS\", \"PKCS12\", or \"BCFKS\".
	TrustStoreType *string `json:"trustStoreType,omitempty"`
	// A description for this External Server
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the External Server
	Id string `json:"id"`
}

// NewConjurExternalServerResponse instantiates a new ConjurExternalServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConjurExternalServerResponse(schemas []EnumconjurExternalServerSchemaUrn, conjurServerBaseURI []string, conjurAuthenticationMethod string, conjurAccountName string, id string) *ConjurExternalServerResponse {
	this := ConjurExternalServerResponse{}
	this.Schemas = schemas
	this.ConjurServerBaseURI = conjurServerBaseURI
	this.ConjurAuthenticationMethod = conjurAuthenticationMethod
	this.ConjurAccountName = conjurAccountName
	this.Id = id
	return &this
}

// NewConjurExternalServerResponseWithDefaults instantiates a new ConjurExternalServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConjurExternalServerResponseWithDefaults() *ConjurExternalServerResponse {
	this := ConjurExternalServerResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ConjurExternalServerResponse) GetSchemas() []EnumconjurExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumconjurExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetSchemasOk() ([]EnumconjurExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ConjurExternalServerResponse) SetSchemas(v []EnumconjurExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetConjurServerBaseURI returns the ConjurServerBaseURI field value
func (o *ConjurExternalServerResponse) GetConjurServerBaseURI() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConjurServerBaseURI
}

// GetConjurServerBaseURIOk returns a tuple with the ConjurServerBaseURI field value
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetConjurServerBaseURIOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConjurServerBaseURI, true
}

// SetConjurServerBaseURI sets field value
func (o *ConjurExternalServerResponse) SetConjurServerBaseURI(v []string) {
	o.ConjurServerBaseURI = v
}

// GetConjurAuthenticationMethod returns the ConjurAuthenticationMethod field value
func (o *ConjurExternalServerResponse) GetConjurAuthenticationMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConjurAuthenticationMethod
}

// GetConjurAuthenticationMethodOk returns a tuple with the ConjurAuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetConjurAuthenticationMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConjurAuthenticationMethod, true
}

// SetConjurAuthenticationMethod sets field value
func (o *ConjurExternalServerResponse) SetConjurAuthenticationMethod(v string) {
	o.ConjurAuthenticationMethod = v
}

// GetConjurAccountName returns the ConjurAccountName field value
func (o *ConjurExternalServerResponse) GetConjurAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConjurAccountName
}

// GetConjurAccountNameOk returns a tuple with the ConjurAccountName field value
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetConjurAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConjurAccountName, true
}

// SetConjurAccountName sets field value
func (o *ConjurExternalServerResponse) SetConjurAccountName(v string) {
	o.ConjurAccountName = v
}

// GetHttpConnectTimeout returns the HttpConnectTimeout field value if set, zero value otherwise.
func (o *ConjurExternalServerResponse) GetHttpConnectTimeout() string {
	if o == nil || IsNil(o.HttpConnectTimeout) {
		var ret string
		return ret
	}
	return *o.HttpConnectTimeout
}

// GetHttpConnectTimeoutOk returns a tuple with the HttpConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetHttpConnectTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.HttpConnectTimeout) {
		return nil, false
	}
	return o.HttpConnectTimeout, true
}

// HasHttpConnectTimeout returns a boolean if a field has been set.
func (o *ConjurExternalServerResponse) HasHttpConnectTimeout() bool {
	if o != nil && !IsNil(o.HttpConnectTimeout) {
		return true
	}

	return false
}

// SetHttpConnectTimeout gets a reference to the given string and assigns it to the HttpConnectTimeout field.
func (o *ConjurExternalServerResponse) SetHttpConnectTimeout(v string) {
	o.HttpConnectTimeout = &v
}

// GetHttpResponseTimeout returns the HttpResponseTimeout field value if set, zero value otherwise.
func (o *ConjurExternalServerResponse) GetHttpResponseTimeout() string {
	if o == nil || IsNil(o.HttpResponseTimeout) {
		var ret string
		return ret
	}
	return *o.HttpResponseTimeout
}

// GetHttpResponseTimeoutOk returns a tuple with the HttpResponseTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetHttpResponseTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.HttpResponseTimeout) {
		return nil, false
	}
	return o.HttpResponseTimeout, true
}

// HasHttpResponseTimeout returns a boolean if a field has been set.
func (o *ConjurExternalServerResponse) HasHttpResponseTimeout() bool {
	if o != nil && !IsNil(o.HttpResponseTimeout) {
		return true
	}

	return false
}

// SetHttpResponseTimeout gets a reference to the given string and assigns it to the HttpResponseTimeout field.
func (o *ConjurExternalServerResponse) SetHttpResponseTimeout(v string) {
	o.HttpResponseTimeout = &v
}

// GetTrustStoreFile returns the TrustStoreFile field value if set, zero value otherwise.
func (o *ConjurExternalServerResponse) GetTrustStoreFile() string {
	if o == nil || IsNil(o.TrustStoreFile) {
		var ret string
		return ret
	}
	return *o.TrustStoreFile
}

// GetTrustStoreFileOk returns a tuple with the TrustStoreFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetTrustStoreFileOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStoreFile) {
		return nil, false
	}
	return o.TrustStoreFile, true
}

// HasTrustStoreFile returns a boolean if a field has been set.
func (o *ConjurExternalServerResponse) HasTrustStoreFile() bool {
	if o != nil && !IsNil(o.TrustStoreFile) {
		return true
	}

	return false
}

// SetTrustStoreFile gets a reference to the given string and assigns it to the TrustStoreFile field.
func (o *ConjurExternalServerResponse) SetTrustStoreFile(v string) {
	o.TrustStoreFile = &v
}

// GetTrustStorePin returns the TrustStorePin field value if set, zero value otherwise.
func (o *ConjurExternalServerResponse) GetTrustStorePin() string {
	if o == nil || IsNil(o.TrustStorePin) {
		var ret string
		return ret
	}
	return *o.TrustStorePin
}

// GetTrustStorePinOk returns a tuple with the TrustStorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetTrustStorePinOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStorePin) {
		return nil, false
	}
	return o.TrustStorePin, true
}

// HasTrustStorePin returns a boolean if a field has been set.
func (o *ConjurExternalServerResponse) HasTrustStorePin() bool {
	if o != nil && !IsNil(o.TrustStorePin) {
		return true
	}

	return false
}

// SetTrustStorePin gets a reference to the given string and assigns it to the TrustStorePin field.
func (o *ConjurExternalServerResponse) SetTrustStorePin(v string) {
	o.TrustStorePin = &v
}

// GetTrustStoreType returns the TrustStoreType field value if set, zero value otherwise.
func (o *ConjurExternalServerResponse) GetTrustStoreType() string {
	if o == nil || IsNil(o.TrustStoreType) {
		var ret string
		return ret
	}
	return *o.TrustStoreType
}

// GetTrustStoreTypeOk returns a tuple with the TrustStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetTrustStoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStoreType) {
		return nil, false
	}
	return o.TrustStoreType, true
}

// HasTrustStoreType returns a boolean if a field has been set.
func (o *ConjurExternalServerResponse) HasTrustStoreType() bool {
	if o != nil && !IsNil(o.TrustStoreType) {
		return true
	}

	return false
}

// SetTrustStoreType gets a reference to the given string and assigns it to the TrustStoreType field.
func (o *ConjurExternalServerResponse) SetTrustStoreType(v string) {
	o.TrustStoreType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConjurExternalServerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConjurExternalServerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConjurExternalServerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ConjurExternalServerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ConjurExternalServerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ConjurExternalServerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ConjurExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ConjurExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ConjurExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ConjurExternalServerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConjurExternalServerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConjurExternalServerResponse) SetId(v string) {
	o.Id = v
}

func (o ConjurExternalServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConjurExternalServerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["conjurServerBaseURI"] = o.ConjurServerBaseURI
	toSerialize["conjurAuthenticationMethod"] = o.ConjurAuthenticationMethod
	toSerialize["conjurAccountName"] = o.ConjurAccountName
	if !IsNil(o.HttpConnectTimeout) {
		toSerialize["httpConnectTimeout"] = o.HttpConnectTimeout
	}
	if !IsNil(o.HttpResponseTimeout) {
		toSerialize["httpResponseTimeout"] = o.HttpResponseTimeout
	}
	if !IsNil(o.TrustStoreFile) {
		toSerialize["trustStoreFile"] = o.TrustStoreFile
	}
	if !IsNil(o.TrustStorePin) {
		toSerialize["trustStorePin"] = o.TrustStorePin
	}
	if !IsNil(o.TrustStoreType) {
		toSerialize["trustStoreType"] = o.TrustStoreType
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
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableConjurExternalServerResponse struct {
	value *ConjurExternalServerResponse
	isSet bool
}

func (v NullableConjurExternalServerResponse) Get() *ConjurExternalServerResponse {
	return v.value
}

func (v *NullableConjurExternalServerResponse) Set(val *ConjurExternalServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConjurExternalServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConjurExternalServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConjurExternalServerResponse(val *ConjurExternalServerResponse) *NullableConjurExternalServerResponse {
	return &NullableConjurExternalServerResponse{value: val, isSet: true}
}

func (v NullableConjurExternalServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConjurExternalServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
