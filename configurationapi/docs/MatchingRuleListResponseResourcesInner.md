# MatchingRuleListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumapproximateMatchingRuleSchemaUrn**](EnumapproximateMatchingRuleSchemaUrn.md) |  | 
**Id** | **string** | Name of the Matching Rule | 
**Enabled** | **bool** | Indicates whether the Matching Rule is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewMatchingRuleListResponseResourcesInner

`func NewMatchingRuleListResponseResourcesInner(schemas []EnumapproximateMatchingRuleSchemaUrn, id string, enabled bool, ) *MatchingRuleListResponseResourcesInner`

NewMatchingRuleListResponseResourcesInner instantiates a new MatchingRuleListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchingRuleListResponseResourcesInnerWithDefaults

`func NewMatchingRuleListResponseResourcesInnerWithDefaults() *MatchingRuleListResponseResourcesInner`

NewMatchingRuleListResponseResourcesInnerWithDefaults instantiates a new MatchingRuleListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *MatchingRuleListResponseResourcesInner) GetSchemas() []EnumapproximateMatchingRuleSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MatchingRuleListResponseResourcesInner) GetSchemasOk() (*[]EnumapproximateMatchingRuleSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MatchingRuleListResponseResourcesInner) SetSchemas(v []EnumapproximateMatchingRuleSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *MatchingRuleListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MatchingRuleListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MatchingRuleListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetEnabled

`func (o *MatchingRuleListResponseResourcesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MatchingRuleListResponseResourcesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MatchingRuleListResponseResourcesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *MatchingRuleListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MatchingRuleListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MatchingRuleListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MatchingRuleListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *MatchingRuleListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *MatchingRuleListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *MatchingRuleListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *MatchingRuleListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


