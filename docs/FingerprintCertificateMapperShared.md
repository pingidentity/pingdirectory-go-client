# FingerprintCertificateMapperShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfingerprintCertificateMapperSchemaUrn**](EnumfingerprintCertificateMapperSchemaUrn.md) |  | 
**FingerprintAttribute** | **string** | Specifies the attribute in which to look for the fingerprint. | 
**FingerprintAlgorithm** | [**EnumcertificateMapperFingerprintAlgorithmProp**](EnumcertificateMapperFingerprintAlgorithmProp.md) |  | 
**UserBaseDN** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 

## Methods

### NewFingerprintCertificateMapperShared

`func NewFingerprintCertificateMapperShared(schemas []EnumfingerprintCertificateMapperSchemaUrn, fingerprintAttribute string, fingerprintAlgorithm EnumcertificateMapperFingerprintAlgorithmProp, enabled bool, ) *FingerprintCertificateMapperShared`

NewFingerprintCertificateMapperShared instantiates a new FingerprintCertificateMapperShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFingerprintCertificateMapperSharedWithDefaults

`func NewFingerprintCertificateMapperSharedWithDefaults() *FingerprintCertificateMapperShared`

NewFingerprintCertificateMapperSharedWithDefaults instantiates a new FingerprintCertificateMapperShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FingerprintCertificateMapperShared) GetSchemas() []EnumfingerprintCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FingerprintCertificateMapperShared) GetSchemasOk() (*[]EnumfingerprintCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FingerprintCertificateMapperShared) SetSchemas(v []EnumfingerprintCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFingerprintAttribute

`func (o *FingerprintCertificateMapperShared) GetFingerprintAttribute() string`

GetFingerprintAttribute returns the FingerprintAttribute field if non-nil, zero value otherwise.

### GetFingerprintAttributeOk

`func (o *FingerprintCertificateMapperShared) GetFingerprintAttributeOk() (*string, bool)`

GetFingerprintAttributeOk returns a tuple with the FingerprintAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintAttribute

`func (o *FingerprintCertificateMapperShared) SetFingerprintAttribute(v string)`

SetFingerprintAttribute sets FingerprintAttribute field to given value.


### GetFingerprintAlgorithm

`func (o *FingerprintCertificateMapperShared) GetFingerprintAlgorithm() EnumcertificateMapperFingerprintAlgorithmProp`

GetFingerprintAlgorithm returns the FingerprintAlgorithm field if non-nil, zero value otherwise.

### GetFingerprintAlgorithmOk

`func (o *FingerprintCertificateMapperShared) GetFingerprintAlgorithmOk() (*EnumcertificateMapperFingerprintAlgorithmProp, bool)`

GetFingerprintAlgorithmOk returns a tuple with the FingerprintAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintAlgorithm

`func (o *FingerprintCertificateMapperShared) SetFingerprintAlgorithm(v EnumcertificateMapperFingerprintAlgorithmProp)`

SetFingerprintAlgorithm sets FingerprintAlgorithm field to given value.


### GetUserBaseDN

`func (o *FingerprintCertificateMapperShared) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *FingerprintCertificateMapperShared) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *FingerprintCertificateMapperShared) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *FingerprintCertificateMapperShared) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *FingerprintCertificateMapperShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FingerprintCertificateMapperShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FingerprintCertificateMapperShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FingerprintCertificateMapperShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FingerprintCertificateMapperShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FingerprintCertificateMapperShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FingerprintCertificateMapperShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


