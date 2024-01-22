# AddSubjectEqualsDnCertificateMapperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsubjectEqualsDnCertificateMapperSchemaUrn**](EnumsubjectEqualsDnCertificateMapperSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 
**MapperName** | **string** | Name of the new Certificate Mapper | 

## Methods

### NewAddSubjectEqualsDnCertificateMapperRequest

`func NewAddSubjectEqualsDnCertificateMapperRequest(schemas []EnumsubjectEqualsDnCertificateMapperSchemaUrn, enabled bool, mapperName string, ) *AddSubjectEqualsDnCertificateMapperRequest`

NewAddSubjectEqualsDnCertificateMapperRequest instantiates a new AddSubjectEqualsDnCertificateMapperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSubjectEqualsDnCertificateMapperRequestWithDefaults

`func NewAddSubjectEqualsDnCertificateMapperRequestWithDefaults() *AddSubjectEqualsDnCertificateMapperRequest`

NewAddSubjectEqualsDnCertificateMapperRequestWithDefaults instantiates a new AddSubjectEqualsDnCertificateMapperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddSubjectEqualsDnCertificateMapperRequest) GetSchemas() []EnumsubjectEqualsDnCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSubjectEqualsDnCertificateMapperRequest) GetSchemasOk() (*[]EnumsubjectEqualsDnCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSubjectEqualsDnCertificateMapperRequest) SetSchemas(v []EnumsubjectEqualsDnCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddSubjectEqualsDnCertificateMapperRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSubjectEqualsDnCertificateMapperRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSubjectEqualsDnCertificateMapperRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSubjectEqualsDnCertificateMapperRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSubjectEqualsDnCertificateMapperRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSubjectEqualsDnCertificateMapperRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSubjectEqualsDnCertificateMapperRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMapperName

`func (o *AddSubjectEqualsDnCertificateMapperRequest) GetMapperName() string`

GetMapperName returns the MapperName field if non-nil, zero value otherwise.

### GetMapperNameOk

`func (o *AddSubjectEqualsDnCertificateMapperRequest) GetMapperNameOk() (*string, bool)`

GetMapperNameOk returns a tuple with the MapperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperName

`func (o *AddSubjectEqualsDnCertificateMapperRequest) SetMapperName(v string)`

SetMapperName sets MapperName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


