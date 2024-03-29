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

// checks if the AddKeyPairRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddKeyPairRequest{}

// AddKeyPairRequest struct for AddKeyPairRequest
type AddKeyPairRequest struct {
	Schemas      []EnumkeyPairSchemaUrn       `json:"schemas,omitempty"`
	KeyAlgorithm *EnumkeyPairKeyAlgorithmProp `json:"keyAlgorithm,omitempty"`
	// The validity period for a self-signed certificate. If not specified, the self-signed certificate will be valid for approximately 20 years. This is not used when importing an existing key-pair. The system will not automatically rotate expired certificates. It is up to the administrator to do that when that happens.
	SelfSignedCertificateValidity *string `json:"selfSignedCertificateValidity,omitempty"`
	// The DN that should be used as the subject for the self-signed certificate and certificate signing request. This is not used when importing an existing key-pair.
	SubjectDN *string `json:"subjectDN,omitempty"`
	// The PEM-encoded X.509 certificate chain.
	CertificateChain *string `json:"certificateChain,omitempty"`
	// The base64-encoded private key that is encrypted using the preferred encryption settings definition.
	PrivateKey *string `json:"privateKey,omitempty"`
	// Name of the new Key Pair
	PairName string `json:"pairName"`
}

// NewAddKeyPairRequest instantiates a new AddKeyPairRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddKeyPairRequest(pairName string) *AddKeyPairRequest {
	this := AddKeyPairRequest{}
	this.PairName = pairName
	return &this
}

// NewAddKeyPairRequestWithDefaults instantiates a new AddKeyPairRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddKeyPairRequestWithDefaults() *AddKeyPairRequest {
	this := AddKeyPairRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddKeyPairRequest) GetSchemas() []EnumkeyPairSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumkeyPairSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddKeyPairRequest) GetSchemasOk() ([]EnumkeyPairSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddKeyPairRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumkeyPairSchemaUrn and assigns it to the Schemas field.
func (o *AddKeyPairRequest) SetSchemas(v []EnumkeyPairSchemaUrn) {
	o.Schemas = v
}

// GetKeyAlgorithm returns the KeyAlgorithm field value if set, zero value otherwise.
func (o *AddKeyPairRequest) GetKeyAlgorithm() EnumkeyPairKeyAlgorithmProp {
	if o == nil || IsNil(o.KeyAlgorithm) {
		var ret EnumkeyPairKeyAlgorithmProp
		return ret
	}
	return *o.KeyAlgorithm
}

// GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddKeyPairRequest) GetKeyAlgorithmOk() (*EnumkeyPairKeyAlgorithmProp, bool) {
	if o == nil || IsNil(o.KeyAlgorithm) {
		return nil, false
	}
	return o.KeyAlgorithm, true
}

// HasKeyAlgorithm returns a boolean if a field has been set.
func (o *AddKeyPairRequest) HasKeyAlgorithm() bool {
	if o != nil && !IsNil(o.KeyAlgorithm) {
		return true
	}

	return false
}

// SetKeyAlgorithm gets a reference to the given EnumkeyPairKeyAlgorithmProp and assigns it to the KeyAlgorithm field.
func (o *AddKeyPairRequest) SetKeyAlgorithm(v EnumkeyPairKeyAlgorithmProp) {
	o.KeyAlgorithm = &v
}

// GetSelfSignedCertificateValidity returns the SelfSignedCertificateValidity field value if set, zero value otherwise.
func (o *AddKeyPairRequest) GetSelfSignedCertificateValidity() string {
	if o == nil || IsNil(o.SelfSignedCertificateValidity) {
		var ret string
		return ret
	}
	return *o.SelfSignedCertificateValidity
}

// GetSelfSignedCertificateValidityOk returns a tuple with the SelfSignedCertificateValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddKeyPairRequest) GetSelfSignedCertificateValidityOk() (*string, bool) {
	if o == nil || IsNil(o.SelfSignedCertificateValidity) {
		return nil, false
	}
	return o.SelfSignedCertificateValidity, true
}

// HasSelfSignedCertificateValidity returns a boolean if a field has been set.
func (o *AddKeyPairRequest) HasSelfSignedCertificateValidity() bool {
	if o != nil && !IsNil(o.SelfSignedCertificateValidity) {
		return true
	}

	return false
}

// SetSelfSignedCertificateValidity gets a reference to the given string and assigns it to the SelfSignedCertificateValidity field.
func (o *AddKeyPairRequest) SetSelfSignedCertificateValidity(v string) {
	o.SelfSignedCertificateValidity = &v
}

// GetSubjectDN returns the SubjectDN field value if set, zero value otherwise.
func (o *AddKeyPairRequest) GetSubjectDN() string {
	if o == nil || IsNil(o.SubjectDN) {
		var ret string
		return ret
	}
	return *o.SubjectDN
}

// GetSubjectDNOk returns a tuple with the SubjectDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddKeyPairRequest) GetSubjectDNOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectDN) {
		return nil, false
	}
	return o.SubjectDN, true
}

// HasSubjectDN returns a boolean if a field has been set.
func (o *AddKeyPairRequest) HasSubjectDN() bool {
	if o != nil && !IsNil(o.SubjectDN) {
		return true
	}

	return false
}

// SetSubjectDN gets a reference to the given string and assigns it to the SubjectDN field.
func (o *AddKeyPairRequest) SetSubjectDN(v string) {
	o.SubjectDN = &v
}

// GetCertificateChain returns the CertificateChain field value if set, zero value otherwise.
func (o *AddKeyPairRequest) GetCertificateChain() string {
	if o == nil || IsNil(o.CertificateChain) {
		var ret string
		return ret
	}
	return *o.CertificateChain
}

// GetCertificateChainOk returns a tuple with the CertificateChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddKeyPairRequest) GetCertificateChainOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateChain) {
		return nil, false
	}
	return o.CertificateChain, true
}

// HasCertificateChain returns a boolean if a field has been set.
func (o *AddKeyPairRequest) HasCertificateChain() bool {
	if o != nil && !IsNil(o.CertificateChain) {
		return true
	}

	return false
}

// SetCertificateChain gets a reference to the given string and assigns it to the CertificateChain field.
func (o *AddKeyPairRequest) SetCertificateChain(v string) {
	o.CertificateChain = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *AddKeyPairRequest) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddKeyPairRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *AddKeyPairRequest) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *AddKeyPairRequest) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPairName returns the PairName field value
func (o *AddKeyPairRequest) GetPairName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PairName
}

// GetPairNameOk returns a tuple with the PairName field value
// and a boolean to check if the value has been set.
func (o *AddKeyPairRequest) GetPairNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PairName, true
}

// SetPairName sets field value
func (o *AddKeyPairRequest) SetPairName(v string) {
	o.PairName = v
}

func (o AddKeyPairRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddKeyPairRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.KeyAlgorithm) {
		toSerialize["keyAlgorithm"] = o.KeyAlgorithm
	}
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
	toSerialize["pairName"] = o.PairName
	return toSerialize, nil
}

type NullableAddKeyPairRequest struct {
	value *AddKeyPairRequest
	isSet bool
}

func (v NullableAddKeyPairRequest) Get() *AddKeyPairRequest {
	return v.value
}

func (v *NullableAddKeyPairRequest) Set(val *AddKeyPairRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddKeyPairRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddKeyPairRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddKeyPairRequest(val *AddKeyPairRequest) *NullableAddKeyPairRequest {
	return &NullableAddKeyPairRequest{value: val, isSet: true}
}

func (v NullableAddKeyPairRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddKeyPairRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
