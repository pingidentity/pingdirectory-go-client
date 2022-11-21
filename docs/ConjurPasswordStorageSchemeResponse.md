# ConjurPasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Password Storage Scheme | 
**Schemas** | [**[]EnumconjurPasswordStorageSchemeSchemaUrn**](EnumconjurPasswordStorageSchemeSchemaUrn.md) |  | 
**ConjurExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Conjur instance containing user passwords. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewConjurPasswordStorageSchemeResponse

`func NewConjurPasswordStorageSchemeResponse(id string, schemas []EnumconjurPasswordStorageSchemeSchemaUrn, conjurExternalServer string, enabled bool, ) *ConjurPasswordStorageSchemeResponse`

NewConjurPasswordStorageSchemeResponse instantiates a new ConjurPasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConjurPasswordStorageSchemeResponseWithDefaults

`func NewConjurPasswordStorageSchemeResponseWithDefaults() *ConjurPasswordStorageSchemeResponse`

NewConjurPasswordStorageSchemeResponseWithDefaults instantiates a new ConjurPasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConjurPasswordStorageSchemeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConjurPasswordStorageSchemeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConjurPasswordStorageSchemeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ConjurPasswordStorageSchemeResponse) GetSchemas() []EnumconjurPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConjurPasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumconjurPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConjurPasswordStorageSchemeResponse) SetSchemas(v []EnumconjurPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetConjurExternalServer

`func (o *ConjurPasswordStorageSchemeResponse) GetConjurExternalServer() string`

GetConjurExternalServer returns the ConjurExternalServer field if non-nil, zero value otherwise.

### GetConjurExternalServerOk

`func (o *ConjurPasswordStorageSchemeResponse) GetConjurExternalServerOk() (*string, bool)`

GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurExternalServer

`func (o *ConjurPasswordStorageSchemeResponse) SetConjurExternalServer(v string)`

SetConjurExternalServer sets ConjurExternalServer field to given value.


### GetDescription

`func (o *ConjurPasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConjurPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConjurPasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConjurPasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ConjurPasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConjurPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConjurPasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *ConjurPasswordStorageSchemeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConjurPasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConjurPasswordStorageSchemeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConjurPasswordStorageSchemeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


