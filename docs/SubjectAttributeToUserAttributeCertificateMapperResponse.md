# SubjectAttributeToUserAttributeCertificateMapperResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Certificate Mapper | 
**Schemas** | [**[]EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn**](EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn.md) |  | 
**SubjectAttributeMapping** | **[]string** | Specifies a mapping between certificate attributes and user attributes. | 
**UserBaseDN** | Pointer to **[]string** | Specifies the base DNs that should be used when performing searches to map the client certificate to a user entry. | [optional] 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewSubjectAttributeToUserAttributeCertificateMapperResponse

`func NewSubjectAttributeToUserAttributeCertificateMapperResponse(id string, schemas []EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn, subjectAttributeMapping []string, enabled bool, ) *SubjectAttributeToUserAttributeCertificateMapperResponse`

NewSubjectAttributeToUserAttributeCertificateMapperResponse instantiates a new SubjectAttributeToUserAttributeCertificateMapperResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectAttributeToUserAttributeCertificateMapperResponseWithDefaults

`func NewSubjectAttributeToUserAttributeCertificateMapperResponseWithDefaults() *SubjectAttributeToUserAttributeCertificateMapperResponse`

NewSubjectAttributeToUserAttributeCertificateMapperResponseWithDefaults instantiates a new SubjectAttributeToUserAttributeCertificateMapperResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetSchemas() []EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetSchemasOk() (*[]EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) SetSchemas(v []EnumsubjectAttributeToUserAttributeCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSubjectAttributeMapping

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetSubjectAttributeMapping() []string`

GetSubjectAttributeMapping returns the SubjectAttributeMapping field if non-nil, zero value otherwise.

### GetSubjectAttributeMappingOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetSubjectAttributeMappingOk() (*[]string, bool)`

GetSubjectAttributeMappingOk returns a tuple with the SubjectAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttributeMapping

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) SetSubjectAttributeMapping(v []string)`

SetSubjectAttributeMapping sets SubjectAttributeMapping field to given value.


### GetUserBaseDN

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SubjectAttributeToUserAttributeCertificateMapperResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


