# ConjurPasswordStorageSchemeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumconjurPasswordStorageSchemeSchemaUrn**](EnumconjurPasswordStorageSchemeSchemaUrn.md) |  | 
**ConjurExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Conjur instance containing user passwords. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewConjurPasswordStorageSchemeShared

`func NewConjurPasswordStorageSchemeShared(schemas []EnumconjurPasswordStorageSchemeSchemaUrn, conjurExternalServer string, enabled bool, ) *ConjurPasswordStorageSchemeShared`

NewConjurPasswordStorageSchemeShared instantiates a new ConjurPasswordStorageSchemeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConjurPasswordStorageSchemeSharedWithDefaults

`func NewConjurPasswordStorageSchemeSharedWithDefaults() *ConjurPasswordStorageSchemeShared`

NewConjurPasswordStorageSchemeSharedWithDefaults instantiates a new ConjurPasswordStorageSchemeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ConjurPasswordStorageSchemeShared) GetSchemas() []EnumconjurPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConjurPasswordStorageSchemeShared) GetSchemasOk() (*[]EnumconjurPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConjurPasswordStorageSchemeShared) SetSchemas(v []EnumconjurPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetConjurExternalServer

`func (o *ConjurPasswordStorageSchemeShared) GetConjurExternalServer() string`

GetConjurExternalServer returns the ConjurExternalServer field if non-nil, zero value otherwise.

### GetConjurExternalServerOk

`func (o *ConjurPasswordStorageSchemeShared) GetConjurExternalServerOk() (*string, bool)`

GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurExternalServer

`func (o *ConjurPasswordStorageSchemeShared) SetConjurExternalServer(v string)`

SetConjurExternalServer sets ConjurExternalServer field to given value.


### GetDescription

`func (o *ConjurPasswordStorageSchemeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConjurPasswordStorageSchemeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConjurPasswordStorageSchemeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConjurPasswordStorageSchemeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ConjurPasswordStorageSchemeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConjurPasswordStorageSchemeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConjurPasswordStorageSchemeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


