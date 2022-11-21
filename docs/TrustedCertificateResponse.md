# TrustedCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Trusted Certificate | 
**Schemas** | Pointer to [**[]EnumtrustedCertificateSchemaUrn**](EnumtrustedCertificateSchemaUrn.md) |  | [optional] 
**Certificate** | **string** | The PEM-encoded X.509v3 certificate. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewTrustedCertificateResponse

`func NewTrustedCertificateResponse(id string, certificate string, ) *TrustedCertificateResponse`

NewTrustedCertificateResponse instantiates a new TrustedCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedCertificateResponseWithDefaults

`func NewTrustedCertificateResponseWithDefaults() *TrustedCertificateResponse`

NewTrustedCertificateResponseWithDefaults instantiates a new TrustedCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrustedCertificateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustedCertificateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustedCertificateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *TrustedCertificateResponse) GetSchemas() []EnumtrustedCertificateSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TrustedCertificateResponse) GetSchemasOk() (*[]EnumtrustedCertificateSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TrustedCertificateResponse) SetSchemas(v []EnumtrustedCertificateSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *TrustedCertificateResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetCertificate

`func (o *TrustedCertificateResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TrustedCertificateResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TrustedCertificateResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetMeta

`func (o *TrustedCertificateResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TrustedCertificateResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TrustedCertificateResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TrustedCertificateResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


