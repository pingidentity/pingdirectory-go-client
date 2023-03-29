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

// checks if the KeyPairResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyPairResponse{}

// KeyPairResponse struct for KeyPairResponse
type KeyPairResponse struct {
	// Name of the Key Pair
	Id           string                      `json:"id"`
	Schemas      []EnumkeyPairSchemaUrn      `json:"schemas,omitempty"`
	KeyAlgorithm EnumkeyPairKeyAlgorithmProp `json:"keyAlgorithm"`
	// The validity period for a self-signed certificate. If not specified, the self-signed certificate will be valid for approximately 20 years. This is not used when importing an existing key-pair. The system will not automatically rotate expired certificates. It is up to the administrator to do that when that happens.
	SelfSignedCertificateValidity *string `json:"selfSignedCertificateValidity,omitempty"`
	// The DN that should be used as the subject for the self-signed certificate and certificate signing request. This is not used when importing an existing key-pair.
	SubjectDN *string `json:"subjectDN,omitempty"`
	// The PEM-encoded X.509 certificate chain.
	CertificateChain *string `json:"certificateChain,omitempty"`
	// The base64-encoded private key that is encrypted using the preferred encryption settings definition.
	PrivateKey                                    *string                                            `json:"privateKey,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewKeyPairResponse instantiates a new KeyPairResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyPairResponse(id string, keyAlgorithm EnumkeyPairKeyAlgorithmProp) *KeyPairResponse {
	this := KeyPairResponse{}
	this.Id = id
	this.KeyAlgorithm = keyAlgorithm
	return &this
}

// NewKeyPairResponseWithDefaults instantiates a new KeyPairResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyPairResponseWithDefaults() *KeyPairResponse {
	this := KeyPairResponse{}
	return &this
}

// GetId returns the Id field value
func (o *KeyPairResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KeyPairResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KeyPairResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *KeyPairResponse) GetSchemas() []EnumkeyPairSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumkeyPairSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairResponse) GetSchemasOk() ([]EnumkeyPairSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *KeyPairResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumkeyPairSchemaUrn and assigns it to the Schemas field.
func (o *KeyPairResponse) SetSchemas(v []EnumkeyPairSchemaUrn) {
	o.Schemas = v
}

// GetKeyAlgorithm returns the KeyAlgorithm field value
func (o *KeyPairResponse) GetKeyAlgorithm() EnumkeyPairKeyAlgorithmProp {
	if o == nil {
		var ret EnumkeyPairKeyAlgorithmProp
		return ret
	}

	return o.KeyAlgorithm
}

// GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field value
// and a boolean to check if the value has been set.
func (o *KeyPairResponse) GetKeyAlgorithmOk() (*EnumkeyPairKeyAlgorithmProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyAlgorithm, true
}

// SetKeyAlgorithm sets field value
func (o *KeyPairResponse) SetKeyAlgorithm(v EnumkeyPairKeyAlgorithmProp) {
	o.KeyAlgorithm = v
}

// GetSelfSignedCertificateValidity returns the SelfSignedCertificateValidity field value if set, zero value otherwise.
func (o *KeyPairResponse) GetSelfSignedCertificateValidity() string {
	if o == nil || IsNil(o.SelfSignedCertificateValidity) {
		var ret string
		return ret
	}
	return *o.SelfSignedCertificateValidity
}

// GetSelfSignedCertificateValidityOk returns a tuple with the SelfSignedCertificateValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairResponse) GetSelfSignedCertificateValidityOk() (*string, bool) {
	if o == nil || IsNil(o.SelfSignedCertificateValidity) {
		return nil, false
	}
	return o.SelfSignedCertificateValidity, true
}

// HasSelfSignedCertificateValidity returns a boolean if a field has been set.
func (o *KeyPairResponse) HasSelfSignedCertificateValidity() bool {
	if o != nil && !IsNil(o.SelfSignedCertificateValidity) {
		return true
	}

	return false
}

// SetSelfSignedCertificateValidity gets a reference to the given string and assigns it to the SelfSignedCertificateValidity field.
func (o *KeyPairResponse) SetSelfSignedCertificateValidity(v string) {
	o.SelfSignedCertificateValidity = &v
}

// GetSubjectDN returns the SubjectDN field value if set, zero value otherwise.
func (o *KeyPairResponse) GetSubjectDN() string {
	if o == nil || IsNil(o.SubjectDN) {
		var ret string
		return ret
	}
	return *o.SubjectDN
}

// GetSubjectDNOk returns a tuple with the SubjectDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairResponse) GetSubjectDNOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectDN) {
		return nil, false
	}
	return o.SubjectDN, true
}

// HasSubjectDN returns a boolean if a field has been set.
func (o *KeyPairResponse) HasSubjectDN() bool {
	if o != nil && !IsNil(o.SubjectDN) {
		return true
	}

	return false
}

// SetSubjectDN gets a reference to the given string and assigns it to the SubjectDN field.
func (o *KeyPairResponse) SetSubjectDN(v string) {
	o.SubjectDN = &v
}

// GetCertificateChain returns the CertificateChain field value if set, zero value otherwise.
func (o *KeyPairResponse) GetCertificateChain() string {
	if o == nil || IsNil(o.CertificateChain) {
		var ret string
		return ret
	}
	return *o.CertificateChain
}

// GetCertificateChainOk returns a tuple with the CertificateChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairResponse) GetCertificateChainOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateChain) {
		return nil, false
	}
	return o.CertificateChain, true
}

// HasCertificateChain returns a boolean if a field has been set.
func (o *KeyPairResponse) HasCertificateChain() bool {
	if o != nil && !IsNil(o.CertificateChain) {
		return true
	}

	return false
}

// SetCertificateChain gets a reference to the given string and assigns it to the CertificateChain field.
func (o *KeyPairResponse) SetCertificateChain(v string) {
	o.CertificateChain = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *KeyPairResponse) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairResponse) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *KeyPairResponse) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *KeyPairResponse) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *KeyPairResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *KeyPairResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *KeyPairResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *KeyPairResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *KeyPairResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *KeyPairResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o KeyPairResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyPairResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["keyAlgorithm"] = o.KeyAlgorithm
	if !IsNil(o.SelfSignedCertificateValidity) {
		toSerialize["selfSignedCertificateValidity"] = o.SelfSignedCertificateValidity
	}
	if !IsNil(o.SubjectDN) {
		toSerialize["subjectDN"] = o.SubjectDN
	}
	if !IsNil(o.CertificateChain) {
		toSerialize["certificateChain"] = o.CertificateChain
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableKeyPairResponse struct {
	value *KeyPairResponse
	isSet bool
}

func (v NullableKeyPairResponse) Get() *KeyPairResponse {
	return v.value
}

func (v *NullableKeyPairResponse) Set(val *KeyPairResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyPairResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyPairResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyPairResponse(val *KeyPairResponse) *NullableKeyPairResponse {
	return &NullableKeyPairResponse{value: val, isSet: true}
}

func (v NullableKeyPairResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyPairResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
