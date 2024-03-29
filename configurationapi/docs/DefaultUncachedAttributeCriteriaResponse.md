# DefaultUncachedAttributeCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdefaultUncachedAttributeCriteriaSchemaUrn**](EnumdefaultUncachedAttributeCriteriaSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Uncached Attribute Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Attribute Criteria is enabled for use in the server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Uncached Attribute Criteria | 

## Methods

### NewDefaultUncachedAttributeCriteriaResponse

`func NewDefaultUncachedAttributeCriteriaResponse(schemas []EnumdefaultUncachedAttributeCriteriaSchemaUrn, enabled bool, id string, ) *DefaultUncachedAttributeCriteriaResponse`

NewDefaultUncachedAttributeCriteriaResponse instantiates a new DefaultUncachedAttributeCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultUncachedAttributeCriteriaResponseWithDefaults

`func NewDefaultUncachedAttributeCriteriaResponseWithDefaults() *DefaultUncachedAttributeCriteriaResponse`

NewDefaultUncachedAttributeCriteriaResponseWithDefaults instantiates a new DefaultUncachedAttributeCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DefaultUncachedAttributeCriteriaResponse) GetSchemas() []EnumdefaultUncachedAttributeCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DefaultUncachedAttributeCriteriaResponse) GetSchemasOk() (*[]EnumdefaultUncachedAttributeCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DefaultUncachedAttributeCriteriaResponse) SetSchemas(v []EnumdefaultUncachedAttributeCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *DefaultUncachedAttributeCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DefaultUncachedAttributeCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DefaultUncachedAttributeCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DefaultUncachedAttributeCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DefaultUncachedAttributeCriteriaResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DefaultUncachedAttributeCriteriaResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DefaultUncachedAttributeCriteriaResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *DefaultUncachedAttributeCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DefaultUncachedAttributeCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DefaultUncachedAttributeCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DefaultUncachedAttributeCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DefaultUncachedAttributeCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DefaultUncachedAttributeCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DefaultUncachedAttributeCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DefaultUncachedAttributeCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *DefaultUncachedAttributeCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DefaultUncachedAttributeCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DefaultUncachedAttributeCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


