# GroupImplementationListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnuminvertedStaticGroupImplementationSchemaUrn**](EnuminvertedStaticGroupImplementationSchemaUrn.md) |  | 
**Id** | **string** | Name of the Group Implementation | 
**Description** | Pointer to **string** | A description for this Group Implementation | [optional] 
**Enabled** | **bool** | Indicates whether the Group Implementation is enabled. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewGroupImplementationListResponseResourcesInner

`func NewGroupImplementationListResponseResourcesInner(schemas []EnuminvertedStaticGroupImplementationSchemaUrn, id string, enabled bool, ) *GroupImplementationListResponseResourcesInner`

NewGroupImplementationListResponseResourcesInner instantiates a new GroupImplementationListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupImplementationListResponseResourcesInnerWithDefaults

`func NewGroupImplementationListResponseResourcesInnerWithDefaults() *GroupImplementationListResponseResourcesInner`

NewGroupImplementationListResponseResourcesInnerWithDefaults instantiates a new GroupImplementationListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GroupImplementationListResponseResourcesInner) GetSchemas() []EnuminvertedStaticGroupImplementationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroupImplementationListResponseResourcesInner) GetSchemasOk() (*[]EnuminvertedStaticGroupImplementationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroupImplementationListResponseResourcesInner) SetSchemas(v []EnuminvertedStaticGroupImplementationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GroupImplementationListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupImplementationListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupImplementationListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *GroupImplementationListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupImplementationListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupImplementationListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupImplementationListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GroupImplementationListResponseResourcesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroupImplementationListResponseResourcesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroupImplementationListResponseResourcesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *GroupImplementationListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GroupImplementationListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GroupImplementationListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GroupImplementationListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GroupImplementationListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GroupImplementationListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GroupImplementationListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GroupImplementationListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


