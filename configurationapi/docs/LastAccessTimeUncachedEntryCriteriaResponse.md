# LastAccessTimeUncachedEntryCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn**](EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn.md) |  | 
**AccessTimeThreshold** | **string** | Specifies the maximum length of time that has passed since an entry was last accessed that it should still be included in the id2entry database. Entries that have not been accessed in more than this length of time may be written into the uncached-id2entry database. | 
**Description** | Pointer to **string** | A description for this Uncached Entry Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Entry Criteria is enabled for use in the server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Uncached Entry Criteria | 

## Methods

### NewLastAccessTimeUncachedEntryCriteriaResponse

`func NewLastAccessTimeUncachedEntryCriteriaResponse(schemas []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn, accessTimeThreshold string, enabled bool, id string, ) *LastAccessTimeUncachedEntryCriteriaResponse`

NewLastAccessTimeUncachedEntryCriteriaResponse instantiates a new LastAccessTimeUncachedEntryCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastAccessTimeUncachedEntryCriteriaResponseWithDefaults

`func NewLastAccessTimeUncachedEntryCriteriaResponseWithDefaults() *LastAccessTimeUncachedEntryCriteriaResponse`

NewLastAccessTimeUncachedEntryCriteriaResponseWithDefaults instantiates a new LastAccessTimeUncachedEntryCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetSchemas() []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetSchemasOk() (*[]EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetSchemas(v []EnumlastAccessTimeUncachedEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccessTimeThreshold

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetAccessTimeThreshold() string`

GetAccessTimeThreshold returns the AccessTimeThreshold field if non-nil, zero value otherwise.

### GetAccessTimeThresholdOk

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetAccessTimeThresholdOk() (*string, bool)`

GetAccessTimeThresholdOk returns a tuple with the AccessTimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTimeThreshold

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetAccessTimeThreshold(v string)`

SetAccessTimeThreshold sets AccessTimeThreshold field to given value.


### GetDescription

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LastAccessTimeUncachedEntryCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


