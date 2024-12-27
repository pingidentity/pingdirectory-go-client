# EntryCounterPluginCriteriaResponse

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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Entry Counter Plugin Criteria | 

## Methods

### NewEntryCounterPluginCriteriaResponse

`func NewEntryCounterPluginCriteriaResponse(filter string, id string, ) *EntryCounterPluginCriteriaResponse`

NewEntryCounterPluginCriteriaResponse instantiates a new EntryCounterPluginCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryCounterPluginCriteriaResponseWithDefaults

`func NewEntryCounterPluginCriteriaResponseWithDefaults() *EntryCounterPluginCriteriaResponse`

NewEntryCounterPluginCriteriaResponseWithDefaults instantiates a new EntryCounterPluginCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *EntryCounterPluginCriteriaResponse) GetSchemas() []EnumentryCounterPluginCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EntryCounterPluginCriteriaResponse) GetSchemasOk() (*[]EnumentryCounterPluginCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EntryCounterPluginCriteriaResponse) SetSchemas(v []EnumentryCounterPluginCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *EntryCounterPluginCriteriaResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetBaseDN

`func (o *EntryCounterPluginCriteriaResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *EntryCounterPluginCriteriaResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *EntryCounterPluginCriteriaResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *EntryCounterPluginCriteriaResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetFilter

`func (o *EntryCounterPluginCriteriaResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *EntryCounterPluginCriteriaResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *EntryCounterPluginCriteriaResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetNamedSubCategoryFilter

`func (o *EntryCounterPluginCriteriaResponse) GetNamedSubCategoryFilter() []string`

GetNamedSubCategoryFilter returns the NamedSubCategoryFilter field if non-nil, zero value otherwise.

### GetNamedSubCategoryFilterOk

`func (o *EntryCounterPluginCriteriaResponse) GetNamedSubCategoryFilterOk() (*[]string, bool)`

GetNamedSubCategoryFilterOk returns a tuple with the NamedSubCategoryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedSubCategoryFilter

`func (o *EntryCounterPluginCriteriaResponse) SetNamedSubCategoryFilter(v []string)`

SetNamedSubCategoryFilter sets NamedSubCategoryFilter field to given value.

### HasNamedSubCategoryFilter

`func (o *EntryCounterPluginCriteriaResponse) HasNamedSubCategoryFilter() bool`

HasNamedSubCategoryFilter returns a boolean if a field has been set.

### GetWarningThresholdMinimumCount

`func (o *EntryCounterPluginCriteriaResponse) GetWarningThresholdMinimumCount() int64`

GetWarningThresholdMinimumCount returns the WarningThresholdMinimumCount field if non-nil, zero value otherwise.

### GetWarningThresholdMinimumCountOk

`func (o *EntryCounterPluginCriteriaResponse) GetWarningThresholdMinimumCountOk() (*int64, bool)`

GetWarningThresholdMinimumCountOk returns a tuple with the WarningThresholdMinimumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdMinimumCount

`func (o *EntryCounterPluginCriteriaResponse) SetWarningThresholdMinimumCount(v int64)`

SetWarningThresholdMinimumCount sets WarningThresholdMinimumCount field to given value.

### HasWarningThresholdMinimumCount

`func (o *EntryCounterPluginCriteriaResponse) HasWarningThresholdMinimumCount() bool`

HasWarningThresholdMinimumCount returns a boolean if a field has been set.

### GetErrorThresholdMinimumCount

`func (o *EntryCounterPluginCriteriaResponse) GetErrorThresholdMinimumCount() int64`

GetErrorThresholdMinimumCount returns the ErrorThresholdMinimumCount field if non-nil, zero value otherwise.

### GetErrorThresholdMinimumCountOk

`func (o *EntryCounterPluginCriteriaResponse) GetErrorThresholdMinimumCountOk() (*int64, bool)`

GetErrorThresholdMinimumCountOk returns a tuple with the ErrorThresholdMinimumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorThresholdMinimumCount

`func (o *EntryCounterPluginCriteriaResponse) SetErrorThresholdMinimumCount(v int64)`

SetErrorThresholdMinimumCount sets ErrorThresholdMinimumCount field to given value.

### HasErrorThresholdMinimumCount

`func (o *EntryCounterPluginCriteriaResponse) HasErrorThresholdMinimumCount() bool`

HasErrorThresholdMinimumCount returns a boolean if a field has been set.

### GetTrackMatchingEntrySize

`func (o *EntryCounterPluginCriteriaResponse) GetTrackMatchingEntrySize() bool`

GetTrackMatchingEntrySize returns the TrackMatchingEntrySize field if non-nil, zero value otherwise.

### GetTrackMatchingEntrySizeOk

`func (o *EntryCounterPluginCriteriaResponse) GetTrackMatchingEntrySizeOk() (*bool, bool)`

GetTrackMatchingEntrySizeOk returns a tuple with the TrackMatchingEntrySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackMatchingEntrySize

`func (o *EntryCounterPluginCriteriaResponse) SetTrackMatchingEntrySize(v bool)`

SetTrackMatchingEntrySize sets TrackMatchingEntrySize field to given value.

### HasTrackMatchingEntrySize

`func (o *EntryCounterPluginCriteriaResponse) HasTrackMatchingEntrySize() bool`

HasTrackMatchingEntrySize returns a boolean if a field has been set.

### GetMeta

`func (o *EntryCounterPluginCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EntryCounterPluginCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EntryCounterPluginCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EntryCounterPluginCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *EntryCounterPluginCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *EntryCounterPluginCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *EntryCounterPluginCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *EntryCounterPluginCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *EntryCounterPluginCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryCounterPluginCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryCounterPluginCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


