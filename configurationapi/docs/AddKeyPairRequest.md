# AddKeyPairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PairName** | **string** | Name of the new Key Pair | 
**Schemas** | Pointer to [**[]EnumkeyPairSchemaUrn**](EnumkeyPairSchemaUrn.md) |  | [optional] 
**KeyAlgorithm** | [**EnumkeyPairKeyAlgorithmProp**](EnumkeyPairKeyAlgorithmProp.md) |  | 
**SelfSignedCertificateValidity** | Pointer to **string** | The validity period for a self-signed certificate. If not specified, the self-signed certificate will be valid for approximately 20 years. This is not used when importing an existing key-pair. The system will not automatically rotate expired certificates. It is up to the administrator to do that when that happens. | [optional] 
**SubjectDN** | Pointer to **string** | The DN that should be used as the subject for the self-signed certificate and certificate signing request. This is not used when importing an existing key-pair. | [optional] 
**CertificateChain** | Pointer to **string** | The PEM-encoded X.509 certificate chain. | [optional] 
**PrivateKey** | Pointer to **string** | The base64-encoded private key that is encrypted using the preferred encryption settings definition. | [optional] 

## Methods

### NewAddKeyPairRequest

`func NewAddKeyPairRequest(pairName string, keyAlgorithm EnumkeyPairKeyAlgorithmProp, ) *AddKeyPairRequest`

NewAddKeyPairRequest instantiates a new AddKeyPairRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddKeyPairRequestWithDefaults

`func NewAddKeyPairRequestWithDefaults() *AddKeyPairRequest`

NewAddKeyPairRequestWithDefaults instantiates a new AddKeyPairRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPairName

`func (o *AddKeyPairRequest) GetPairName() string`

GetPairName returns the PairName field if non-nil, zero value otherwise.

### GetPairNameOk

`func (o *AddKeyPairRequest) GetPairNameOk() (*string, bool)`

GetPairNameOk returns a tuple with the PairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairName

`func (o *AddKeyPairRequest) SetPairName(v string)`

SetPairName sets PairName field to given value.


### GetSchemas

`func (o *AddKeyPairRequest) GetSchemas() []EnumkeyPairSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddKeyPairRequest) GetSchemasOk() (*[]EnumkeyPairSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddKeyPairRequest) SetSchemas(v []EnumkeyPairSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddKeyPairRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *AddKeyPairRequest) GetKeyAlgorithm() EnumkeyPairKeyAlgorithmProp`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *AddKeyPairRequest) GetKeyAlgorithmOk() (*EnumkeyPairKeyAlgorithmProp, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *AddKeyPairRequest) SetKeyAlgorithm(v EnumkeyPairKeyAlgorithmProp)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.


### GetSelfSignedCertificateValidity

`func (o *AddKeyPairRequest) GetSelfSignedCertificateValidity() string`

GetSelfSignedCertificateValidity returns the SelfSignedCertificateValidity field if non-nil, zero value otherwise.

### GetSelfSignedCertificateValidityOk

`func (o *AddKeyPairRequest) GetSelfSignedCertificateValidityOk() (*string, bool)`

GetSelfSignedCertificateValidityOk returns a tuple with the SelfSignedCertificateValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSignedCertificateValidity

`func (o *AddKeyPairRequest) SetSelfSignedCertificateValidity(v string)`

SetSelfSignedCertificateValidity sets SelfSignedCertificateValidity field to given value.

### HasSelfSignedCertificateValidity

`func (o *AddKeyPairRequest) HasSelfSignedCertificateValidity() bool`

HasSelfSignedCertificateValidity returns a boolean if a field has been set.

### GetSubjectDN

`func (o *AddKeyPairRequest) GetSubjectDN() string`

GetSubjectDN returns the SubjectDN field if non-nil, zero value otherwise.

### GetSubjectDNOk

`func (o *AddKeyPairRequest) GetSubjectDNOk() (*string, bool)`

GetSubjectDNOk returns a tuple with the SubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDN

`func (o *AddKeyPairRequest) SetSubjectDN(v string)`

SetSubjectDN sets SubjectDN field to given value.

### HasSubjectDN

`func (o *AddKeyPairRequest) HasSubjectDN() bool`

HasSubjectDN returns a boolean if a field has been set.

### GetCertificateChain

`func (o *AddKeyPairRequest) GetCertificateChain() string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *AddKeyPairRequest) GetCertificateChainOk() (*string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *AddKeyPairRequest) SetCertificateChain(v string)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *AddKeyPairRequest) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.

### GetPrivateKey

`func (o *AddKeyPairRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *AddKeyPairRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *AddKeyPairRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *AddKeyPairRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


