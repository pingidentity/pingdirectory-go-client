# KeyPairResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Key Pair | 
**Schemas** | Pointer to [**[]EnumkeyPairSchemaUrn**](EnumkeyPairSchemaUrn.md) |  | [optional] 
**KeyAlgorithm** | [**EnumkeyPairKeyAlgorithmProp**](EnumkeyPairKeyAlgorithmProp.md) |  | 
**SelfSignedCertificateValidity** | Pointer to **string** | The validity period for a self-signed certificate. If not specified, the self-signed certificate will be valid for approximately 20 years. This is not used when importing an existing key-pair. The system will not automatically rotate expired certificates. It is up to the administrator to do that when that happens. | [optional] 
**SubjectDN** | Pointer to **string** | The DN that should be used as the subject for the self-signed certificate and certificate signing request. This is not used when importing an existing key-pair. | [optional] 
**CertificateChain** | Pointer to **string** | The PEM-encoded X.509 certificate chain. | [optional] 
**PrivateKey** | Pointer to **string** | The base64-encoded private key that is encrypted using the preferred encryption settings definition. | [optional] 

## Methods

### NewKeyPairResponse

`func NewKeyPairResponse(id string, keyAlgorithm EnumkeyPairKeyAlgorithmProp, ) *KeyPairResponse`

NewKeyPairResponse instantiates a new KeyPairResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyPairResponseWithDefaults

`func NewKeyPairResponseWithDefaults() *KeyPairResponse`

NewKeyPairResponseWithDefaults instantiates a new KeyPairResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyPairResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyPairResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyPairResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *KeyPairResponse) GetSchemas() []EnumkeyPairSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *KeyPairResponse) GetSchemasOk() (*[]EnumkeyPairSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *KeyPairResponse) SetSchemas(v []EnumkeyPairSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *KeyPairResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *KeyPairResponse) GetKeyAlgorithm() EnumkeyPairKeyAlgorithmProp`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *KeyPairResponse) GetKeyAlgorithmOk() (*EnumkeyPairKeyAlgorithmProp, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *KeyPairResponse) SetKeyAlgorithm(v EnumkeyPairKeyAlgorithmProp)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.


### GetSelfSignedCertificateValidity

`func (o *KeyPairResponse) GetSelfSignedCertificateValidity() string`

GetSelfSignedCertificateValidity returns the SelfSignedCertificateValidity field if non-nil, zero value otherwise.

### GetSelfSignedCertificateValidityOk

`func (o *KeyPairResponse) GetSelfSignedCertificateValidityOk() (*string, bool)`

GetSelfSignedCertificateValidityOk returns a tuple with the SelfSignedCertificateValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSignedCertificateValidity

`func (o *KeyPairResponse) SetSelfSignedCertificateValidity(v string)`

SetSelfSignedCertificateValidity sets SelfSignedCertificateValidity field to given value.

### HasSelfSignedCertificateValidity

`func (o *KeyPairResponse) HasSelfSignedCertificateValidity() bool`

HasSelfSignedCertificateValidity returns a boolean if a field has been set.

### GetSubjectDN

`func (o *KeyPairResponse) GetSubjectDN() string`

GetSubjectDN returns the SubjectDN field if non-nil, zero value otherwise.

### GetSubjectDNOk

`func (o *KeyPairResponse) GetSubjectDNOk() (*string, bool)`

GetSubjectDNOk returns a tuple with the SubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDN

`func (o *KeyPairResponse) SetSubjectDN(v string)`

SetSubjectDN sets SubjectDN field to given value.

### HasSubjectDN

`func (o *KeyPairResponse) HasSubjectDN() bool`

HasSubjectDN returns a boolean if a field has been set.

### GetCertificateChain

`func (o *KeyPairResponse) GetCertificateChain() string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *KeyPairResponse) GetCertificateChainOk() (*string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *KeyPairResponse) SetCertificateChain(v string)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *KeyPairResponse) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.

### GetPrivateKey

`func (o *KeyPairResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *KeyPairResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *KeyPairResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *KeyPairResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


