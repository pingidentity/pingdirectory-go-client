# AddEntryCounterPluginCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumentryCounterPluginCriteriaSchemaUrn**](EnumentryCounterPluginCriteriaSchemaUrn.md) |  | [optional] 
**BaseDN** | Pointer to **[]string** | The base DNs for entries that are eligible to match this criteria. If no base DNs are specified, then the set of public naming contexts will be used. | [optional] 
**Filter** | **string** | The filter to use to identify entries that are eligible to match this criteria. | 
**NamedSubCategoryFilter** | Pointer to **[]string** | An optional set of named filters that can be used to divide entries that match the base DN and filter criteria into additional sub-categories. | [optional] 
**WarningThresholdMinimumCount** | Pointer to **int64** | The minimum number of entries in the server that must match this criteria before the server will raise a warning alarm. | [optional] 
**ErrorThresholdMinimumCount** | Pointer to **int64** | The minimum number of entries in the server that must match this criteria before the server will raise an error alarm. | [optional] 
**TrackMatchingEntrySize** | Pointer to **bool** | Indicates whether to track information about the size of each matching entry. | [optional] 
**CriteriaName** | **string** | Name of the new Entry Counter Plugin Criteria | 

## Methods

### NewAddEntryCounterPluginCriteriaRequest

`func NewAddEntryCounterPluginCriteriaRequest(filter string, criteriaName string, ) *AddEntryCounterPluginCriteriaRequest`

NewAddEntryCounterPluginCriteriaRequest instantiates a new AddEntryCounterPluginCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEntryCounterPluginCriteriaRequestWithDefaults

`func NewAddEntryCounterPluginCriteriaRequestWithDefaults() *AddEntryCounterPluginCriteriaRequest`

NewAddEntryCounterPluginCriteriaRequestWithDefaults instantiates a new AddEntryCounterPluginCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddEntryCounterPluginCriteriaRequest) GetSchemas() []EnumentryCounterPluginCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddEntryCounterPluginCriteriaRequest) GetSchemasOk() (*[]EnumentryCounterPluginCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddEntryCounterPluginCriteriaRequest) SetSchemas(v []EnumentryCounterPluginCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddEntryCounterPluginCriteriaRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetBaseDN

`func (o *AddEntryCounterPluginCriteriaRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddEntryCounterPluginCriteriaRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddEntryCounterPluginCriteriaRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddEntryCounterPluginCriteriaRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetFilter

`func (o *AddEntryCounterPluginCriteriaRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddEntryCounterPluginCriteriaRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddEntryCounterPluginCriteriaRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetNamedSubCategoryFilter

`func (o *AddEntryCounterPluginCriteriaRequest) GetNamedSubCategoryFilter() []string`

GetNamedSubCategoryFilter returns the NamedSubCategoryFilter field if non-nil, zero value otherwise.

### GetNamedSubCategoryFilterOk

`func (o *AddEntryCounterPluginCriteriaRequest) GetNamedSubCategoryFilterOk() (*[]string, bool)`

GetNamedSubCategoryFilterOk returns a tuple with the NamedSubCategoryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedSubCategoryFilter

`func (o *AddEntryCounterPluginCriteriaRequest) SetNamedSubCategoryFilter(v []string)`

SetNamedSubCategoryFilter sets NamedSubCategoryFilter field to given value.

### HasNamedSubCategoryFilter

`func (o *AddEntryCounterPluginCriteriaRequest) HasNamedSubCategoryFilter() bool`

HasNamedSubCategoryFilter returns a boolean if a field has been set.

### GetWarningThresholdMinimumCount

`func (o *AddEntryCounterPluginCriteriaRequest) GetWarningThresholdMinimumCount() int64`

GetWarningThresholdMinimumCount returns the WarningThresholdMinimumCount field if non-nil, zero value otherwise.

### GetWarningThresholdMinimumCountOk

`func (o *AddEntryCounterPluginCriteriaRequest) GetWarningThresholdMinimumCountOk() (*int64, bool)`

GetWarningThresholdMinimumCountOk returns a tuple with the WarningThresholdMinimumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdMinimumCount

`func (o *AddEntryCounterPluginCriteriaRequest) SetWarningThresholdMinimumCount(v int64)`

SetWarningThresholdMinimumCount sets WarningThresholdMinimumCount field to given value.

### HasWarningThresholdMinimumCount

`func (o *AddEntryCounterPluginCriteriaRequest) HasWarningThresholdMinimumCount() bool`

HasWarningThresholdMinimumCount returns a boolean if a field has been set.

### GetErrorThresholdMinimumCount

`func (o *AddEntryCounterPluginCriteriaRequest) GetErrorThresholdMinimumCount() int64`

GetErrorThresholdMinimumCount returns the ErrorThresholdMinimumCount field if non-nil, zero value otherwise.

### GetErrorThresholdMinimumCountOk

`func (o *AddEntryCounterPluginCriteriaRequest) GetErrorThresholdMinimumCountOk() (*int64, bool)`

GetErrorThresholdMinimumCountOk returns a tuple with the ErrorThresholdMinimumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorThresholdMinimumCount

`func (o *AddEntryCounterPluginCriteriaRequest) SetErrorThresholdMinimumCount(v int64)`

SetErrorThresholdMinimumCount sets ErrorThresholdMinimumCount field to given value.

### HasErrorThresholdMinimumCount

`func (o *AddEntryCounterPluginCriteriaRequest) HasErrorThresholdMinimumCount() bool`

HasErrorThresholdMinimumCount returns a boolean if a field has been set.

### GetTrackMatchingEntrySize

`func (o *AddEntryCounterPluginCriteriaRequest) GetTrackMatchingEntrySize() bool`

GetTrackMatchingEntrySize returns the TrackMatchingEntrySize field if non-nil, zero value otherwise.

### GetTrackMatchingEntrySizeOk

`func (o *AddEntryCounterPluginCriteriaRequest) GetTrackMatchingEntrySizeOk() (*bool, bool)`

GetTrackMatchingEntrySizeOk returns a tuple with the TrackMatchingEntrySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackMatchingEntrySize

`func (o *AddEntryCounterPluginCriteriaRequest) SetTrackMatchingEntrySize(v bool)`

SetTrackMatchingEntrySize sets TrackMatchingEntrySize field to given value.

### HasTrackMatchingEntrySize

`func (o *AddEntryCounterPluginCriteriaRequest) HasTrackMatchingEntrySize() bool`

HasTrackMatchingEntrySize returns a boolean if a field has been set.

### GetCriteriaName

`func (o *AddEntryCounterPluginCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddEntryCounterPluginCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddEntryCounterPluginCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


