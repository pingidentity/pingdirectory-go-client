# SubjectDnToUserAttributeCertificateMapperShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn**](EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn.md) |  | 
**SubjectAttribute** | **string** | Specifies the name or OID of the attribute whose value should exactly match the certificate subject DN. | 
**UserBaseDN** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 

## Methods

### NewSubjectDnToUserAttributeCertificateMapperShared

`func NewSubjectDnToUserAttributeCertificateMapperShared(schemas []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, subjectAttribute string, enabled bool, ) *SubjectDnToUserAttributeCertificateMapperShared`

NewSubjectDnToUserAttributeCertificateMapperShared instantiates a new SubjectDnToUserAttributeCertificateMapperShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectDnToUserAttributeCertificateMapperSharedWithDefaults

`func NewSubjectDnToUserAttributeCertificateMapperSharedWithDefaults() *SubjectDnToUserAttributeCertificateMapperShared`

NewSubjectDnToUserAttributeCertificateMapperSharedWithDefaults instantiates a new SubjectDnToUserAttributeCertificateMapperShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SubjectDnToUserAttributeCertificateMapperShared) GetSchemas() []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SubjectDnToUserAttributeCertificateMapperShared) GetSchemasOk() (*[]EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SubjectDnToUserAttributeCertificateMapperShared) SetSchemas(v []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSubjectAttribute

`func (o *SubjectDnToUserAttributeCertificateMapperShared) GetSubjectAttribute() string`

GetSubjectAttribute returns the SubjectAttribute field if non-nil, zero value otherwise.

### GetSubjectAttributeOk

`func (o *SubjectDnToUserAttributeCertificateMapperShared) GetSubjectAttributeOk() (*string, bool)`

GetSubjectAttributeOk returns a tuple with the SubjectAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttribute

`func (o *SubjectDnToUserAttributeCertificateMapperShared) SetSubjectAttribute(v string)`

SetSubjectAttribute sets SubjectAttribute field to given value.


### GetUserBaseDN

`func (o *SubjectDnToUserAttributeCertificateMapperShared) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *SubjectDnToUserAttributeCertificateMapperShared) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *SubjectDnToUserAttributeCertificateMapperShared) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *SubjectDnToUserAttributeCertificateMapperShared) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *SubjectDnToUserAttributeCertificateMapperShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubjectDnToUserAttributeCertificateMapperShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubjectDnToUserAttributeCertificateMapperShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubjectDnToUserAttributeCertificateMapperShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SubjectDnToUserAttributeCertificateMapperShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SubjectDnToUserAttributeCertificateMapperShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SubjectDnToUserAttributeCertificateMapperShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


