# AddTrustedCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumtrustedCertificateSchemaUrn**](EnumtrustedCertificateSchemaUrn.md) |  | [optional] 
**Certificate** | **string** | The PEM-encoded X.509v3 certificate. | 
**CertificateName** | **string** | Name of the new Trusted Certificate | 

## Methods

### NewAddTrustedCertificateRequest

`func NewAddTrustedCertificateRequest(certificate string, certificateName string, ) *AddTrustedCertificateRequest`

NewAddTrustedCertificateRequest instantiates a new AddTrustedCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTrustedCertificateRequestWithDefaults

`func NewAddTrustedCertificateRequestWithDefaults() *AddTrustedCertificateRequest`

NewAddTrustedCertificateRequestWithDefaults instantiates a new AddTrustedCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddTrustedCertificateRequest) GetSchemas() []EnumtrustedCertificateSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTrustedCertificateRequest) GetSchemasOk() (*[]EnumtrustedCertificateSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTrustedCertificateRequest) SetSchemas(v []EnumtrustedCertificateSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddTrustedCertificateRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetCertificate

`func (o *AddTrustedCertificateRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *AddTrustedCertificateRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *AddTrustedCertificateRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetCertificateName

`func (o *AddTrustedCertificateRequest) GetCertificateName() string`

GetCertificateName returns the CertificateName field if non-nil, zero value otherwise.

### GetCertificateNameOk

`func (o *AddTrustedCertificateRequest) GetCertificateNameOk() (*string, bool)`

GetCertificateNameOk returns a tuple with the CertificateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateName

`func (o *AddTrustedCertificateRequest) SetCertificateName(v string)`

SetCertificateName sets CertificateName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


