# SubjectAttributeToUserAttributeCertificateMapperShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn**](EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn.md) |  | 
**SubjectAttributeMapping** | **[]string** |  | 
**UserBaseDN** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 

## Methods

### NewSubjectAttributeToUserAttributeCertificateMapperShared

`func NewSubjectAttributeToUserAttributeCertificateMapperShared(schemas []EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn, subjectAttributeMapping []string, enabled bool, ) *SubjectAttributeToUserAttributeCertificateMapperShared`

NewSubjectAttributeToUserAttributeCertificateMapperShared instantiates a new SubjectAttributeToUserAttributeCertificateMapperShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectAttributeToUserAttributeCertificateMapperSharedWithDefaults

`func NewSubjectAttributeToUserAttributeCertificateMapperSharedWithDefaults() *SubjectAttributeToUserAttributeCertificateMapperShared`

NewSubjectAttributeToUserAttributeCertificateMapperSharedWithDefaults instantiates a new SubjectAttributeToUserAttributeCertificateMapperShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) GetSchemas() []EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) GetSchemasOk() (*[]EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) SetSchemas(v []EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSubjectAttributeMapping

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) GetSubjectAttributeMapping() []string`

GetSubjectAttributeMapping returns the SubjectAttributeMapping field if non-nil, zero value otherwise.

### GetSubjectAttributeMappingOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) GetSubjectAttributeMappingOk() (*[]string, bool)`

GetSubjectAttributeMappingOk returns a tuple with the SubjectAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttributeMapping

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) SetSubjectAttributeMapping(v []string)`

SetSubjectAttributeMapping sets SubjectAttributeMapping field to given value.


### GetUserBaseDN

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SubjectAttributeToUserAttributeCertificateMapperShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


