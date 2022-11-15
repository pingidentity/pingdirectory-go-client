# FilterBasedUncachedEntryCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfilterBasedUncachedEntryCriteriaSchemaUrn**](EnumfilterBasedUncachedEntryCriteriaSchemaUrn.md) |  | 
**Filter** | **string** | Specifies the search filter that should be used to differentiate entries into cached and uncached sets. | 
**FilterIdentifiesUncachedEntries** | **bool** | Indicates whether the associated filter identifies those entries which should be stored in the uncached-id2entry database (if true) or entries which should be stored in the id2entry database (if false). | 
**Description** | Pointer to **string** | A description for this Uncached Entry Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Entry Criteria is enabled for use in the server. | 

## Methods

### NewFilterBasedUncachedEntryCriteriaShared

`func NewFilterBasedUncachedEntryCriteriaShared(schemas []EnumfilterBasedUncachedEntryCriteriaSchemaUrn, filter string, filterIdentifiesUncachedEntries bool, enabled bool, ) *FilterBasedUncachedEntryCriteriaShared`

NewFilterBasedUncachedEntryCriteriaShared instantiates a new FilterBasedUncachedEntryCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterBasedUncachedEntryCriteriaSharedWithDefaults

`func NewFilterBasedUncachedEntryCriteriaSharedWithDefaults() *FilterBasedUncachedEntryCriteriaShared`

NewFilterBasedUncachedEntryCriteriaSharedWithDefaults instantiates a new FilterBasedUncachedEntryCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FilterBasedUncachedEntryCriteriaShared) GetSchemas() []EnumfilterBasedUncachedEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FilterBasedUncachedEntryCriteriaShared) GetSchemasOk() (*[]EnumfilterBasedUncachedEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FilterBasedUncachedEntryCriteriaShared) SetSchemas(v []EnumfilterBasedUncachedEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFilter

`func (o *FilterBasedUncachedEntryCriteriaShared) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *FilterBasedUncachedEntryCriteriaShared) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *FilterBasedUncachedEntryCriteriaShared) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetFilterIdentifiesUncachedEntries

`func (o *FilterBasedUncachedEntryCriteriaShared) GetFilterIdentifiesUncachedEntries() bool`

GetFilterIdentifiesUncachedEntries returns the FilterIdentifiesUncachedEntries field if non-nil, zero value otherwise.

### GetFilterIdentifiesUncachedEntriesOk

`func (o *FilterBasedUncachedEntryCriteriaShared) GetFilterIdentifiesUncachedEntriesOk() (*bool, bool)`

GetFilterIdentifiesUncachedEntriesOk returns a tuple with the FilterIdentifiesUncachedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIdentifiesUncachedEntries

`func (o *FilterBasedUncachedEntryCriteriaShared) SetFilterIdentifiesUncachedEntries(v bool)`

SetFilterIdentifiesUncachedEntries sets FilterIdentifiesUncachedEntries field to given value.


### GetDescription

`func (o *FilterBasedUncachedEntryCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FilterBasedUncachedEntryCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FilterBasedUncachedEntryCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FilterBasedUncachedEntryCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FilterBasedUncachedEntryCriteriaShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FilterBasedUncachedEntryCriteriaShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FilterBasedUncachedEntryCriteriaShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


