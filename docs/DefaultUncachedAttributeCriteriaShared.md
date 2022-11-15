# DefaultUncachedAttributeCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdefaultUncachedAttributeCriteriaSchemaUrn**](EnumdefaultUncachedAttributeCriteriaSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Uncached Attribute Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Attribute Criteria is enabled for use in the server. | 

## Methods

### NewDefaultUncachedAttributeCriteriaShared

`func NewDefaultUncachedAttributeCriteriaShared(schemas []EnumdefaultUncachedAttributeCriteriaSchemaUrn, enabled bool, ) *DefaultUncachedAttributeCriteriaShared`

NewDefaultUncachedAttributeCriteriaShared instantiates a new DefaultUncachedAttributeCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultUncachedAttributeCriteriaSharedWithDefaults

`func NewDefaultUncachedAttributeCriteriaSharedWithDefaults() *DefaultUncachedAttributeCriteriaShared`

NewDefaultUncachedAttributeCriteriaSharedWithDefaults instantiates a new DefaultUncachedAttributeCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DefaultUncachedAttributeCriteriaShared) GetSchemas() []EnumdefaultUncachedAttributeCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DefaultUncachedAttributeCriteriaShared) GetSchemasOk() (*[]EnumdefaultUncachedAttributeCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DefaultUncachedAttributeCriteriaShared) SetSchemas(v []EnumdefaultUncachedAttributeCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *DefaultUncachedAttributeCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DefaultUncachedAttributeCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DefaultUncachedAttributeCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DefaultUncachedAttributeCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DefaultUncachedAttributeCriteriaShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DefaultUncachedAttributeCriteriaShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DefaultUncachedAttributeCriteriaShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


