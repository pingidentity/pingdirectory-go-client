# AddFingerprintCertificateMapperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapperName** | **string** | Name of the new Certificate Mapper | 
**Schemas** | [**[]EnumfingerprintCertificateMapperSchemaUrn**](EnumfingerprintCertificateMapperSchemaUrn.md) |  | 
**FingerprintAttribute** | Pointer to **string** | Specifies the attribute in which to look for the fingerprint. | [optional] 
**FingerprintAlgorithm** | [**EnumcertificateMapperFingerprintAlgorithmProp**](EnumcertificateMapperFingerprintAlgorithmProp.md) |  | 
**UserBaseDN** | Pointer to **[]string** | Specifies the set of base DNs below which to search for users. | [optional] 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 

## Methods

### NewAddFingerprintCertificateMapperRequest

`func NewAddFingerprintCertificateMapperRequest(mapperName string, schemas []EnumfingerprintCertificateMapperSchemaUrn, fingerprintAlgorithm EnumcertificateMapperFingerprintAlgorithmProp, enabled bool, ) *AddFingerprintCertificateMapperRequest`

NewAddFingerprintCertificateMapperRequest instantiates a new AddFingerprintCertificateMapperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFingerprintCertificateMapperRequestWithDefaults

`func NewAddFingerprintCertificateMapperRequestWithDefaults() *AddFingerprintCertificateMapperRequest`

NewAddFingerprintCertificateMapperRequestWithDefaults instantiates a new AddFingerprintCertificateMapperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapperName

`func (o *AddFingerprintCertificateMapperRequest) GetMapperName() string`

GetMapperName returns the MapperName field if non-nil, zero value otherwise.

### GetMapperNameOk

`func (o *AddFingerprintCertificateMapperRequest) GetMapperNameOk() (*string, bool)`

GetMapperNameOk returns a tuple with the MapperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperName

`func (o *AddFingerprintCertificateMapperRequest) SetMapperName(v string)`

SetMapperName sets MapperName field to given value.


### GetSchemas

`func (o *AddFingerprintCertificateMapperRequest) GetSchemas() []EnumfingerprintCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddFingerprintCertificateMapperRequest) GetSchemasOk() (*[]EnumfingerprintCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddFingerprintCertificateMapperRequest) SetSchemas(v []EnumfingerprintCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFingerprintAttribute

`func (o *AddFingerprintCertificateMapperRequest) GetFingerprintAttribute() string`

GetFingerprintAttribute returns the FingerprintAttribute field if non-nil, zero value otherwise.

### GetFingerprintAttributeOk

`func (o *AddFingerprintCertificateMapperRequest) GetFingerprintAttributeOk() (*string, bool)`

GetFingerprintAttributeOk returns a tuple with the FingerprintAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintAttribute

`func (o *AddFingerprintCertificateMapperRequest) SetFingerprintAttribute(v string)`

SetFingerprintAttribute sets FingerprintAttribute field to given value.

### HasFingerprintAttribute

`func (o *AddFingerprintCertificateMapperRequest) HasFingerprintAttribute() bool`

HasFingerprintAttribute returns a boolean if a field has been set.

### GetFingerprintAlgorithm

`func (o *AddFingerprintCertificateMapperRequest) GetFingerprintAlgorithm() EnumcertificateMapperFingerprintAlgorithmProp`

GetFingerprintAlgorithm returns the FingerprintAlgorithm field if non-nil, zero value otherwise.

### GetFingerprintAlgorithmOk

`func (o *AddFingerprintCertificateMapperRequest) GetFingerprintAlgorithmOk() (*EnumcertificateMapperFingerprintAlgorithmProp, bool)`

GetFingerprintAlgorithmOk returns a tuple with the FingerprintAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintAlgorithm

`func (o *AddFingerprintCertificateMapperRequest) SetFingerprintAlgorithm(v EnumcertificateMapperFingerprintAlgorithmProp)`

SetFingerprintAlgorithm sets FingerprintAlgorithm field to given value.


### GetUserBaseDN

`func (o *AddFingerprintCertificateMapperRequest) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *AddFingerprintCertificateMapperRequest) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *AddFingerprintCertificateMapperRequest) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *AddFingerprintCertificateMapperRequest) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *AddFingerprintCertificateMapperRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddFingerprintCertificateMapperRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddFingerprintCertificateMapperRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddFingerprintCertificateMapperRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddFingerprintCertificateMapperRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddFingerprintCertificateMapperRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddFingerprintCertificateMapperRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


