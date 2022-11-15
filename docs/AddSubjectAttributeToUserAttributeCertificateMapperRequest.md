# AddSubjectAttributeToUserAttributeCertificateMapperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapperName** | **string** | Name of the new Certificate Mapper | 
**Schemas** | [**[]EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn**](EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn.md) |  | 
**SubjectAttributeMapping** | **[]string** |  | 
**UserBaseDN** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 

## Methods

### NewAddSubjectAttributeToUserAttributeCertificateMapperRequest

`func NewAddSubjectAttributeToUserAttributeCertificateMapperRequest(mapperName string, schemas []EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn, subjectAttributeMapping []string, enabled bool, ) *AddSubjectAttributeToUserAttributeCertificateMapperRequest`

NewAddSubjectAttributeToUserAttributeCertificateMapperRequest instantiates a new AddSubjectAttributeToUserAttributeCertificateMapperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSubjectAttributeToUserAttributeCertificateMapperRequestWithDefaults

`func NewAddSubjectAttributeToUserAttributeCertificateMapperRequestWithDefaults() *AddSubjectAttributeToUserAttributeCertificateMapperRequest`

NewAddSubjectAttributeToUserAttributeCertificateMapperRequestWithDefaults instantiates a new AddSubjectAttributeToUserAttributeCertificateMapperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapperName

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetMapperName() string`

GetMapperName returns the MapperName field if non-nil, zero value otherwise.

### GetMapperNameOk

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetMapperNameOk() (*string, bool)`

GetMapperNameOk returns a tuple with the MapperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperName

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) SetMapperName(v string)`

SetMapperName sets MapperName field to given value.


### GetSchemas

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetSchemas() []EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetSchemasOk() (*[]EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) SetSchemas(v []EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSubjectAttributeMapping

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetSubjectAttributeMapping() []string`

GetSubjectAttributeMapping returns the SubjectAttributeMapping field if non-nil, zero value otherwise.

### GetSubjectAttributeMappingOk

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetSubjectAttributeMappingOk() (*[]string, bool)`

GetSubjectAttributeMappingOk returns a tuple with the SubjectAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttributeMapping

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) SetSubjectAttributeMapping(v []string)`

SetSubjectAttributeMapping sets SubjectAttributeMapping field to given value.


### GetUserBaseDN

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSubjectAttributeToUserAttributeCertificateMapperRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


