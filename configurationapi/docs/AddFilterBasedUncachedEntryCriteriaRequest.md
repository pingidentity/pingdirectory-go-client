# AddFilterBasedUncachedEntryCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Uncached Entry Criteria | 
**Schemas** | [**[]EnumfilterBasedUncachedEntryCriteriaSchemaUrn**](EnumfilterBasedUncachedEntryCriteriaSchemaUrn.md) |  | 
**Filter** | **string** | Specifies the search filter that should be used to differentiate entries into cached and uncached sets. | 
**FilterIdentifiesUncachedEntries** | **bool** | Indicates whether the associated filter identifies those entries which should be stored in the uncached-id2entry database (if true) or entries which should be stored in the id2entry database (if false). | 
**Description** | Pointer to **string** | A description for this Uncached Entry Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Entry Criteria is enabled for use in the server. | 

## Methods

### NewAddFilterBasedUncachedEntryCriteriaRequest

`func NewAddFilterBasedUncachedEntryCriteriaRequest(criteriaName string, schemas []EnumfilterBasedUncachedEntryCriteriaSchemaUrn, filter string, filterIdentifiesUncachedEntries bool, enabled bool, ) *AddFilterBasedUncachedEntryCriteriaRequest`

NewAddFilterBasedUncachedEntryCriteriaRequest instantiates a new AddFilterBasedUncachedEntryCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFilterBasedUncachedEntryCriteriaRequestWithDefaults

`func NewAddFilterBasedUncachedEntryCriteriaRequestWithDefaults() *AddFilterBasedUncachedEntryCriteriaRequest`

NewAddFilterBasedUncachedEntryCriteriaRequestWithDefaults instantiates a new AddFilterBasedUncachedEntryCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetSchemas() []EnumfilterBasedUncachedEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetSchemasOk() (*[]EnumfilterBasedUncachedEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) SetSchemas(v []EnumfilterBasedUncachedEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFilter

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetFilterIdentifiesUncachedEntries

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetFilterIdentifiesUncachedEntries() bool`

GetFilterIdentifiesUncachedEntries returns the FilterIdentifiesUncachedEntries field if non-nil, zero value otherwise.

### GetFilterIdentifiesUncachedEntriesOk

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetFilterIdentifiesUncachedEntriesOk() (*bool, bool)`

GetFilterIdentifiesUncachedEntriesOk returns a tuple with the FilterIdentifiesUncachedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIdentifiesUncachedEntries

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) SetFilterIdentifiesUncachedEntries(v bool)`

SetFilterIdentifiesUncachedEntries sets FilterIdentifiesUncachedEntries field to given value.


### GetDescription

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddFilterBasedUncachedEntryCriteriaRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


