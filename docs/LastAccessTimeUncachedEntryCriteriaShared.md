# LastAccessTimeUncachedEntryCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn**](EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn.md) |  | 
**AccessTimeThreshold** | **string** | Specifies the maximum length of time that has passed since an entry was last accessed that it should still be included in the id2entry database. Entries that have not been accessed in more than this length of time may be written into the uncached-id2entry database. | 
**Description** | Pointer to **string** | A description for this Uncached Entry Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Entry Criteria is enabled for use in the server. | 

## Methods

### NewLastAccessTimeUncachedEntryCriteriaShared

`func NewLastAccessTimeUncachedEntryCriteriaShared(schemas []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn, accessTimeThreshold string, enabled bool, ) *LastAccessTimeUncachedEntryCriteriaShared`

NewLastAccessTimeUncachedEntryCriteriaShared instantiates a new LastAccessTimeUncachedEntryCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastAccessTimeUncachedEntryCriteriaSharedWithDefaults

`func NewLastAccessTimeUncachedEntryCriteriaSharedWithDefaults() *LastAccessTimeUncachedEntryCriteriaShared`

NewLastAccessTimeUncachedEntryCriteriaSharedWithDefaults instantiates a new LastAccessTimeUncachedEntryCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LastAccessTimeUncachedEntryCriteriaShared) GetSchemas() []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LastAccessTimeUncachedEntryCriteriaShared) GetSchemasOk() (*[]EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LastAccessTimeUncachedEntryCriteriaShared) SetSchemas(v []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccessTimeThreshold

`func (o *LastAccessTimeUncachedEntryCriteriaShared) GetAccessTimeThreshold() string`

GetAccessTimeThreshold returns the AccessTimeThreshold field if non-nil, zero value otherwise.

### GetAccessTimeThresholdOk

`func (o *LastAccessTimeUncachedEntryCriteriaShared) GetAccessTimeThresholdOk() (*string, bool)`

GetAccessTimeThresholdOk returns a tuple with the AccessTimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTimeThreshold

`func (o *LastAccessTimeUncachedEntryCriteriaShared) SetAccessTimeThreshold(v string)`

SetAccessTimeThreshold sets AccessTimeThreshold field to given value.


### GetDescription

`func (o *LastAccessTimeUncachedEntryCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LastAccessTimeUncachedEntryCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LastAccessTimeUncachedEntryCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LastAccessTimeUncachedEntryCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LastAccessTimeUncachedEntryCriteriaShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LastAccessTimeUncachedEntryCriteriaShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LastAccessTimeUncachedEntryCriteriaShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


