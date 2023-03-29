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

// checks if the MacSecretKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MacSecretKeyResponse{}

// MacSecretKeyResponse struct for MacSecretKeyResponse
type MacSecretKeyResponse struct {
	Schemas []EnummacSecretKeySchemaUrn `json:"schemas,omitempty"`
	// The algorithm name used to generate this MAC key, e.g. HmacMD5, HmacSHA1, HMacSHA256, etc.
	MacAlgorithmName *string `json:"macAlgorithmName,omitempty"`
	// The unique system-generated identifier for the Secret Key.
	KeyID string `json:"keyID"`
	// If the key is compromised, an administrator may set this flag to immediately trigger the creation of a new secret key. After the new key is generated, the value of this property will be reset to false.
	IsCompromised *bool `json:"isCompromised,omitempty"`
	// The symmetric key that is used for both encryption of plain text and decryption of cipher text. This stores the secret key for each server instance encrypted with that server's inter-server certificate.
	SymmetricKey []string `json:"symmetricKey,omitempty"`
	// The length of the key in bits.
	KeyLengthBits                                 int32                                              `json:"keyLengthBits"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewMacSecretKeyResponse instantiates a new MacSecretKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacSecretKeyResponse(keyID string, keyLengthBits int32) *MacSecretKeyResponse {
	this := MacSecretKeyResponse{}
	this.KeyID = keyID
	this.KeyLengthBits = keyLengthBits
	return &this
}

// NewMacSecretKeyResponseWithDefaults instantiates a new MacSecretKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacSecretKeyResponseWithDefaults() *MacSecretKeyResponse {
	this := MacSecretKeyResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *MacSecretKeyResponse) GetSchemas() []EnummacSecretKeySchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnummacSecretKeySchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacSecretKeyResponse) GetSchemasOk() ([]EnummacSecretKeySchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *MacSecretKeyResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnummacSecretKeySchemaUrn and assigns it to the Schemas field.
func (o *MacSecretKeyResponse) SetSchemas(v []EnummacSecretKeySchemaUrn) {
	o.Schemas = v
}

// GetMacAlgorithmName returns the MacAlgorithmName field value if set, zero value otherwise.
func (o *MacSecretKeyResponse) GetMacAlgorithmName() string {
	if o == nil || IsNil(o.MacAlgorithmName) {
		var ret string
		return ret
	}
	return *o.MacAlgorithmName
}

// GetMacAlgorithmNameOk returns a tuple with the MacAlgorithmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacSecretKeyResponse) GetMacAlgorithmNameOk() (*string, bool) {
	if o == nil || IsNil(o.MacAlgorithmName) {
		return nil, false
	}
	return o.MacAlgorithmName, true
}

// HasMacAlgorithmName returns a boolean if a field has been set.
func (o *MacSecretKeyResponse) HasMacAlgorithmName() bool {
	if o != nil && !IsNil(o.MacAlgorithmName) {
		return true
	}

	return false
}

// SetMacAlgorithmName gets a reference to the given string and assigns it to the MacAlgorithmName field.
func (o *MacSecretKeyResponse) SetMacAlgorithmName(v string) {
	o.MacAlgorithmName = &v
}

// GetKeyID returns the KeyID field value
func (o *MacSecretKeyResponse) GetKeyID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyID
}

// GetKeyIDOk returns a tuple with the KeyID field value
// and a boolean to check if the value has been set.
func (o *MacSecretKeyResponse) GetKeyIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyID, true
}

// SetKeyID sets field value
func (o *MacSecretKeyResponse) SetKeyID(v string) {
	o.KeyID = v
}

// GetIsCompromised returns the IsCompromised field value if set, zero value otherwise.
func (o *MacSecretKeyResponse) GetIsCompromised() bool {
	if o == nil || IsNil(o.IsCompromised) {
		var ret bool
		return ret
	}
	return *o.IsCompromised
}

// GetIsCompromisedOk returns a tuple with the IsCompromised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacSecretKeyResponse) GetIsCompromisedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompromised) {
		return nil, false
	}
	return o.IsCompromised, true
}

// HasIsCompromised returns a boolean if a field has been set.
func (o *MacSecretKeyResponse) HasIsCompromised() bool {
	if o != nil && !IsNil(o.IsCompromised) {
		return true
	}

	return false
}

// SetIsCompromised gets a reference to the given bool and assigns it to the IsCompromised field.
func (o *MacSecretKeyResponse) SetIsCompromised(v bool) {
	o.IsCompromised = &v
}

// GetSymmetricKey returns the SymmetricKey field value if set, zero value otherwise.
func (o *MacSecretKeyResponse) GetSymmetricKey() []string {
	if o == nil || IsNil(o.SymmetricKey) {
		var ret []string
		return ret
	}
	return o.SymmetricKey
}

// GetSymmetricKeyOk returns a tuple with the SymmetricKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacSecretKeyResponse) GetSymmetricKeyOk() ([]string, bool) {
	if o == nil || IsNil(o.SymmetricKey) {
		return nil, false
	}
	return o.SymmetricKey, true
}

// HasSymmetricKey returns a boolean if a field has been set.
func (o *MacSecretKeyResponse) HasSymmetricKey() bool {
	if o != nil && !IsNil(o.SymmetricKey) {
		return true
	}

	return false
}

// SetSymmetricKey gets a reference to the given []string and assigns it to the SymmetricKey field.
func (o *MacSecretKeyResponse) SetSymmetricKey(v []string) {
	o.SymmetricKey = v
}

// GetKeyLengthBits returns the KeyLengthBits field value
func (o *MacSecretKeyResponse) GetKeyLengthBits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KeyLengthBits
}

// GetKeyLengthBitsOk returns a tuple with the KeyLengthBits field value
// and a boolean to check if the value has been set.
func (o *MacSecretKeyResponse) GetKeyLengthBitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyLengthBits, true
}

// SetKeyLengthBits sets field value
func (o *MacSecretKeyResponse) SetKeyLengthBits(v int32) {
	o.KeyLengthBits = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MacSecretKeyResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacSecretKeyResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MacSecretKeyResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *MacSecretKeyResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *MacSecretKeyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacSecretKeyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *MacSecretKeyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *MacSecretKeyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o MacSecretKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MacSecretKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.MacAlgorithmName) {
		toSerialize["macAlgorithmName"] = o.MacAlgorithmName
	}
	toSerialize["keyID"] = o.KeyID
	if !IsNil(o.IsCompromised) {
		toSerialize["isCompromised"] = o.IsCompromised
	}
	if !IsNil(o.SymmetricKey) {
		toSerialize["symmetricKey"] = o.SymmetricKey
	}
	toSerialize["keyLengthBits"] = o.KeyLengthBits
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableMacSecretKeyResponse struct {
	value *MacSecretKeyResponse
	isSet bool
}

func (v NullableMacSecretKeyResponse) Get() *MacSecretKeyResponse {
	return v.value
}

func (v *NullableMacSecretKeyResponse) Set(val *MacSecretKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMacSecretKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMacSecretKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacSecretKeyResponse(val *MacSecretKeyResponse) *NullableMacSecretKeyResponse {
	return &NullableMacSecretKeyResponse{value: val, isSet: true}
}

func (v NullableMacSecretKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacSecretKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
