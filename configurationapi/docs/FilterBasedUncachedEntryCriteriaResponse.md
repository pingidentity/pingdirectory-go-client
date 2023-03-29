# FilterBasedUncachedEntryCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Uncached Entry Criteria | 
**Schemas** | [**[]EnumfilterBasedUncachedEntryCriteriaSchemaUrn**](EnumfilterBasedUncachedEntryCriteriaSchemaUrn.md) |  | 
**Filter** | **string** | Specifies the search filter that should be used to differentiate entries into cached and uncached sets. | 
**FilterIdentifiesUncachedEntries** | **bool** | Indicates whether the associated filter identifies those entries which should be stored in the uncached-id2entry database (if true) or entries which should be stored in the id2entry database (if false). | 
**Description** | Pointer to **string** | A description for this Uncached Entry Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Entry Criteria is enabled for use in the server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewFilterBasedUncachedEntryCriteriaResponse

`func NewFilterBasedUncachedEntryCriteriaResponse(id string, schemas []EnumfilterBasedUncachedEntryCriteriaSchemaUrn, filter string, filterIdentifiesUncachedEntries bool, enabled bool, ) *FilterBasedUncachedEntryCriteriaResponse`

NewFilterBasedUncachedEntryCriteriaResponse instantiates a new FilterBasedUncachedEntryCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterBasedUncachedEntryCriteriaResponseWithDefaults

`func NewFilterBasedUncachedEntryCriteriaResponseWithDefaults() *FilterBasedUncachedEntryCriteriaResponse`

NewFilterBasedUncachedEntryCriteriaResponseWithDefaults instantiates a new FilterBasedUncachedEntryCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterBasedUncachedEntryCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetSchemas() []EnumfilterBasedUncachedEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetSchemasOk() (*[]EnumfilterBasedUncachedEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FilterBasedUncachedEntryCriteriaResponse) SetSchemas(v []EnumfilterBasedUncachedEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFilter

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *FilterBasedUncachedEntryCriteriaResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetFilterIdentifiesUncachedEntries

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetFilterIdentifiesUncachedEntries() bool`

GetFilterIdentifiesUncachedEntries returns the FilterIdentifiesUncachedEntries field if non-nil, zero value otherwise.

### GetFilterIdentifiesUncachedEntriesOk

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetFilterIdentifiesUncachedEntriesOk() (*bool, bool)`

GetFilterIdentifiesUncachedEntriesOk returns a tuple with the FilterIdentifiesUncachedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIdentifiesUncachedEntries

`func (o *FilterBasedUncachedEntryCriteriaResponse) SetFilterIdentifiesUncachedEntries(v bool)`

SetFilterIdentifiesUncachedEntries sets FilterIdentifiesUncachedEntries field to given value.


### GetDescription

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FilterBasedUncachedEntryCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FilterBasedUncachedEntryCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FilterBasedUncachedEntryCriteriaResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FilterBasedUncachedEntryCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FilterBasedUncachedEntryCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *FilterBasedUncachedEntryCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *FilterBasedUncachedEntryCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *FilterBasedUncachedEntryCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


