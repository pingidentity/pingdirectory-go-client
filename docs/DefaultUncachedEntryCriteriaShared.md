# DefaultUncachedEntryCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdefaultUncachedEntryCriteriaSchemaUrn**](EnumdefaultUncachedEntryCriteriaSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Uncached Entry Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Entry Criteria is enabled for use in the server. | 

## Methods

### NewDefaultUncachedEntryCriteriaShared

`func NewDefaultUncachedEntryCriteriaShared(schemas []EnumdefaultUncachedEntryCriteriaSchemaUrn, enabled bool, ) *DefaultUncachedEntryCriteriaShared`

NewDefaultUncachedEntryCriteriaShared instantiates a new DefaultUncachedEntryCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultUncachedEntryCriteriaSharedWithDefaults

`func NewDefaultUncachedEntryCriteriaSharedWithDefaults() *DefaultUncachedEntryCriteriaShared`

NewDefaultUncachedEntryCriteriaSharedWithDefaults instantiates a new DefaultUncachedEntryCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DefaultUncachedEntryCriteriaShared) GetSchemas() []EnumdefaultUncachedEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DefaultUncachedEntryCriteriaShared) GetSchemasOk() (*[]EnumdefaultUncachedEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DefaultUncachedEntryCriteriaShared) SetSchemas(v []EnumdefaultUncachedEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *DefaultUncachedEntryCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DefaultUncachedEntryCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DefaultUncachedEntryCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DefaultUncachedEntryCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DefaultUncachedEntryCriteriaShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DefaultUncachedEntryCriteriaShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DefaultUncachedEntryCriteriaShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


