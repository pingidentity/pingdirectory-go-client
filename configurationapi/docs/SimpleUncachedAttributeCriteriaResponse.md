# SimpleUncachedAttributeCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Uncached Attribute Criteria | 
**Schemas** | [**[]EnumsimpleUncachedAttributeCriteriaSchemaUrn**](EnumsimpleUncachedAttributeCriteriaSchemaUrn.md) |  | 
**AttributeType** | **[]string** | Specifies the attribute types for attributes that may be written to the uncached-id2entry database. | 
**MinValueCount** | Pointer to **int32** | Specifies the minimum number of values that an attribute must have before it will be written into the uncached-id2entry database. | [optional] 
**MinTotalValueSize** | Pointer to **string** | Specifies the minimum total value size (i.e., the sum of the sizes of all values) that an attribute must have before it will be written into the uncached-id2entry database. | [optional] 
**Description** | Pointer to **string** | A description for this Uncached Attribute Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Attribute Criteria is enabled for use in the server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewSimpleUncachedAttributeCriteriaResponse

`func NewSimpleUncachedAttributeCriteriaResponse(id string, schemas []EnumsimpleUncachedAttributeCriteriaSchemaUrn, attributeType []string, enabled bool, ) *SimpleUncachedAttributeCriteriaResponse`

NewSimpleUncachedAttributeCriteriaResponse instantiates a new SimpleUncachedAttributeCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleUncachedAttributeCriteriaResponseWithDefaults

`func NewSimpleUncachedAttributeCriteriaResponseWithDefaults() *SimpleUncachedAttributeCriteriaResponse`

NewSimpleUncachedAttributeCriteriaResponseWithDefaults instantiates a new SimpleUncachedAttributeCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleUncachedAttributeCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleUncachedAttributeCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleUncachedAttributeCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SimpleUncachedAttributeCriteriaResponse) GetSchemas() []EnumsimpleUncachedAttributeCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SimpleUncachedAttributeCriteriaResponse) GetSchemasOk() (*[]EnumsimpleUncachedAttributeCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SimpleUncachedAttributeCriteriaResponse) SetSchemas(v []EnumsimpleUncachedAttributeCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAttributeType

`func (o *SimpleUncachedAttributeCriteriaResponse) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *SimpleUncachedAttributeCriteriaResponse) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *SimpleUncachedAttributeCriteriaResponse) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetMinValueCount

`func (o *SimpleUncachedAttributeCriteriaResponse) GetMinValueCount() int32`

GetMinValueCount returns the MinValueCount field if non-nil, zero value otherwise.

### GetMinValueCountOk

`func (o *SimpleUncachedAttributeCriteriaResponse) GetMinValueCountOk() (*int32, bool)`

GetMinValueCountOk returns a tuple with the MinValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValueCount

`func (o *SimpleUncachedAttributeCriteriaResponse) SetMinValueCount(v int32)`

SetMinValueCount sets MinValueCount field to given value.

### HasMinValueCount

`func (o *SimpleUncachedAttributeCriteriaResponse) HasMinValueCount() bool`

HasMinValueCount returns a boolean if a field has been set.

### GetMinTotalValueSize

`func (o *SimpleUncachedAttributeCriteriaResponse) GetMinTotalValueSize() string`

GetMinTotalValueSize returns the MinTotalValueSize field if non-nil, zero value otherwise.

### GetMinTotalValueSizeOk

`func (o *SimpleUncachedAttributeCriteriaResponse) GetMinTotalValueSizeOk() (*string, bool)`

GetMinTotalValueSizeOk returns a tuple with the MinTotalValueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTotalValueSize

`func (o *SimpleUncachedAttributeCriteriaResponse) SetMinTotalValueSize(v string)`

SetMinTotalValueSize sets MinTotalValueSize field to given value.

### HasMinTotalValueSize

`func (o *SimpleUncachedAttributeCriteriaResponse) HasMinTotalValueSize() bool`

HasMinTotalValueSize returns a boolean if a field has been set.

### GetDescription

`func (o *SimpleUncachedAttributeCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleUncachedAttributeCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleUncachedAttributeCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimpleUncachedAttributeCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SimpleUncachedAttributeCriteriaResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SimpleUncachedAttributeCriteriaResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SimpleUncachedAttributeCriteriaResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *SimpleUncachedAttributeCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SimpleUncachedAttributeCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SimpleUncachedAttributeCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SimpleUncachedAttributeCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleUncachedAttributeCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SimpleUncachedAttributeCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleUncachedAttributeCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleUncachedAttributeCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


