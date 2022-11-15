# AddLastAccessTimeUncachedEntryCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Uncached Entry Criteria | 
**Schemas** | [**[]EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn**](EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn.md) |  | 
**AccessTimeThreshold** | **string** | Specifies the maximum length of time that has passed since an entry was last accessed that it should still be included in the id2entry database. Entries that have not been accessed in more than this length of time may be written into the uncached-id2entry database. | 
**Description** | Pointer to **string** | A description for this Uncached Entry Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Entry Criteria is enabled for use in the server. | 

## Methods

### NewAddLastAccessTimeUncachedEntryCriteriaRequest

`func NewAddLastAccessTimeUncachedEntryCriteriaRequest(criteriaName string, schemas []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn, accessTimeThreshold string, enabled bool, ) *AddLastAccessTimeUncachedEntryCriteriaRequest`

NewAddLastAccessTimeUncachedEntryCriteriaRequest instantiates a new AddLastAccessTimeUncachedEntryCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLastAccessTimeUncachedEntryCriteriaRequestWithDefaults

`func NewAddLastAccessTimeUncachedEntryCriteriaRequestWithDefaults() *AddLastAccessTimeUncachedEntryCriteriaRequest`

NewAddLastAccessTimeUncachedEntryCriteriaRequestWithDefaults instantiates a new AddLastAccessTimeUncachedEntryCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetSchemas() []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetSchemasOk() (*[]EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) SetSchemas(v []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccessTimeThreshold

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetAccessTimeThreshold() string`

GetAccessTimeThreshold returns the AccessTimeThreshold field if non-nil, zero value otherwise.

### GetAccessTimeThresholdOk

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetAccessTimeThresholdOk() (*string, bool)`

GetAccessTimeThresholdOk returns a tuple with the AccessTimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTimeThreshold

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) SetAccessTimeThreshold(v string)`

SetAccessTimeThreshold sets AccessTimeThreshold field to given value.


### GetDescription

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddLastAccessTimeUncachedEntryCriteriaRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


