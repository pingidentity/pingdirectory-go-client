# AddSubjectDnToUserAttributeCertificateMapperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapperName** | **string** | Name of the new Certificate Mapper | 
**Schemas** | [**[]EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn**](EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn.md) |  | 
**SubjectAttribute** | Pointer to **string** | Specifies the name or OID of the attribute whose value should exactly match the certificate subject DN. | [optional] 
**UserBaseDN** | Pointer to **[]string** | Specifies the base DNs that should be used when performing searches to map the client certificate to a user entry. | [optional] 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 

## Methods

### NewAddSubjectDnToUserAttributeCertificateMapperRequest

`func NewAddSubjectDnToUserAttributeCertificateMapperRequest(mapperName string, schemas []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, enabled bool, ) *AddSubjectDnToUserAttributeCertificateMapperRequest`

NewAddSubjectDnToUserAttributeCertificateMapperRequest instantiates a new AddSubjectDnToUserAttributeCertificateMapperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSubjectDnToUserAttributeCertificateMapperRequestWithDefaults

`func NewAddSubjectDnToUserAttributeCertificateMapperRequestWithDefaults() *AddSubjectDnToUserAttributeCertificateMapperRequest`

NewAddSubjectDnToUserAttributeCertificateMapperRequestWithDefaults instantiates a new AddSubjectDnToUserAttributeCertificateMapperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapperName

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetMapperName() string`

GetMapperName returns the MapperName field if non-nil, zero value otherwise.

### GetMapperNameOk

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetMapperNameOk() (*string, bool)`

GetMapperNameOk returns a tuple with the MapperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperName

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetMapperName(v string)`

SetMapperName sets MapperName field to given value.


### GetSchemas

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetSchemas() []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetSchemasOk() (*[]EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetSchemas(v []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSubjectAttribute

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetSubjectAttribute() string`

GetSubjectAttribute returns the SubjectAttribute field if non-nil, zero value otherwise.

### GetSubjectAttributeOk

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetSubjectAttributeOk() (*string, bool)`

GetSubjectAttributeOk returns a tuple with the SubjectAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttribute

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetSubjectAttribute(v string)`

SetSubjectAttribute sets SubjectAttribute field to given value.

### HasSubjectAttribute

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) HasSubjectAttribute() bool`

HasSubjectAttribute returns a boolean if a field has been set.

### GetUserBaseDN

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


