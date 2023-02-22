# AddDefaultUncachedEntryCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Uncached Entry Criteria | 
**Schemas** | [**[]EnumdefaultUncachedEntryCriteriaSchemaUrn**](EnumdefaultUncachedEntryCriteriaSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Uncached Entry Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Entry Criteria is enabled for use in the server. | 

## Methods

### NewAddDefaultUncachedEntryCriteriaRequest

`func NewAddDefaultUncachedEntryCriteriaRequest(criteriaName string, schemas []EnumdefaultUncachedEntryCriteriaSchemaUrn, enabled bool, ) *AddDefaultUncachedEntryCriteriaRequest`

NewAddDefaultUncachedEntryCriteriaRequest instantiates a new AddDefaultUncachedEntryCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDefaultUncachedEntryCriteriaRequestWithDefaults

`func NewAddDefaultUncachedEntryCriteriaRequestWithDefaults() *AddDefaultUncachedEntryCriteriaRequest`

NewAddDefaultUncachedEntryCriteriaRequestWithDefaults instantiates a new AddDefaultUncachedEntryCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddDefaultUncachedEntryCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddDefaultUncachedEntryCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddDefaultUncachedEntryCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddDefaultUncachedEntryCriteriaRequest) GetSchemas() []EnumdefaultUncachedEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDefaultUncachedEntryCriteriaRequest) GetSchemasOk() (*[]EnumdefaultUncachedEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDefaultUncachedEntryCriteriaRequest) SetSchemas(v []EnumdefaultUncachedEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddDefaultUncachedEntryCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDefaultUncachedEntryCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDefaultUncachedEntryCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDefaultUncachedEntryCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddDefaultUncachedEntryCriteriaRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddDefaultUncachedEntryCriteriaRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddDefaultUncachedEntryCriteriaRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


