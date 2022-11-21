# SubjectDnToUserAttributeCertificateMapperResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Certificate Mapper | 
**Schemas** | [**[]EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn**](EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn.md) |  | 
**SubjectAttribute** | **string** | Specifies the name or OID of the attribute whose value should exactly match the certificate subject DN. | 
**UserBaseDN** | Pointer to **[]string** | Specifies the base DNs that should be used when performing searches to map the client certificate to a user entry. | [optional] 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewSubjectDnToUserAttributeCertificateMapperResponse

`func NewSubjectDnToUserAttributeCertificateMapperResponse(id string, schemas []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, subjectAttribute string, enabled bool, ) *SubjectDnToUserAttributeCertificateMapperResponse`

NewSubjectDnToUserAttributeCertificateMapperResponse instantiates a new SubjectDnToUserAttributeCertificateMapperResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectDnToUserAttributeCertificateMapperResponseWithDefaults

`func NewSubjectDnToUserAttributeCertificateMapperResponseWithDefaults() *SubjectDnToUserAttributeCertificateMapperResponse`

NewSubjectDnToUserAttributeCertificateMapperResponseWithDefaults instantiates a new SubjectDnToUserAttributeCertificateMapperResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetSchemas() []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetSchemasOk() (*[]EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetSchemas(v []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSubjectAttribute

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetSubjectAttribute() string`

GetSubjectAttribute returns the SubjectAttribute field if non-nil, zero value otherwise.

### GetSubjectAttributeOk

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetSubjectAttributeOk() (*string, bool)`

GetSubjectAttributeOk returns a tuple with the SubjectAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttribute

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetSubjectAttribute(v string)`

SetSubjectAttribute sets SubjectAttribute field to given value.


### GetUserBaseDN

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SubjectDnToUserAttributeCertificateMapperResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


