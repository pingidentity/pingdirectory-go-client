# SimpleUncachedAttributeCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsimpleUncachedAttributeCriteriaSchemaUrn**](EnumsimpleUncachedAttributeCriteriaSchemaUrn.md) |  | 
**AttributeType** | **[]string** |  | 
**MinValueCount** | Pointer to **int32** | Specifies the minimum number of values that an attribute must have before it will be written into the uncached-id2entry database. | [optional] 
**MinTotalValueSize** | Pointer to **string** | Specifies the minimum total value size (i.e., the sum of the sizes of all values) that an attribute must have before it will be written into the uncached-id2entry database. | [optional] 
**Description** | Pointer to **string** | A description for this Uncached Attribute Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Attribute Criteria is enabled for use in the server. | 

## Methods

### NewSimpleUncachedAttributeCriteriaShared

`func NewSimpleUncachedAttributeCriteriaShared(schemas []EnumsimpleUncachedAttributeCriteriaSchemaUrn, attributeType []string, enabled bool, ) *SimpleUncachedAttributeCriteriaShared`

NewSimpleUncachedAttributeCriteriaShared instantiates a new SimpleUncachedAttributeCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleUncachedAttributeCriteriaSharedWithDefaults

`func NewSimpleUncachedAttributeCriteriaSharedWithDefaults() *SimpleUncachedAttributeCriteriaShared`

NewSimpleUncachedAttributeCriteriaSharedWithDefaults instantiates a new SimpleUncachedAttributeCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SimpleUncachedAttributeCriteriaShared) GetSchemas() []EnumsimpleUncachedAttributeCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SimpleUncachedAttributeCriteriaShared) GetSchemasOk() (*[]EnumsimpleUncachedAttributeCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SimpleUncachedAttributeCriteriaShared) SetSchemas(v []EnumsimpleUncachedAttributeCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAttributeType

`func (o *SimpleUncachedAttributeCriteriaShared) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *SimpleUncachedAttributeCriteriaShared) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *SimpleUncachedAttributeCriteriaShared) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetMinValueCount

`func (o *SimpleUncachedAttributeCriteriaShared) GetMinValueCount() int32`

GetMinValueCount returns the MinValueCount field if non-nil, zero value otherwise.

### GetMinValueCountOk

`func (o *SimpleUncachedAttributeCriteriaShared) GetMinValueCountOk() (*int32, bool)`

GetMinValueCountOk returns a tuple with the MinValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValueCount

`func (o *SimpleUncachedAttributeCriteriaShared) SetMinValueCount(v int32)`

SetMinValueCount sets MinValueCount field to given value.

### HasMinValueCount

`func (o *SimpleUncachedAttributeCriteriaShared) HasMinValueCount() bool`

HasMinValueCount returns a boolean if a field has been set.

### GetMinTotalValueSize

`func (o *SimpleUncachedAttributeCriteriaShared) GetMinTotalValueSize() string`

GetMinTotalValueSize returns the MinTotalValueSize field if non-nil, zero value otherwise.

### GetMinTotalValueSizeOk

`func (o *SimpleUncachedAttributeCriteriaShared) GetMinTotalValueSizeOk() (*string, bool)`

GetMinTotalValueSizeOk returns a tuple with the MinTotalValueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTotalValueSize

`func (o *SimpleUncachedAttributeCriteriaShared) SetMinTotalValueSize(v string)`

SetMinTotalValueSize sets MinTotalValueSize field to given value.

### HasMinTotalValueSize

`func (o *SimpleUncachedAttributeCriteriaShared) HasMinTotalValueSize() bool`

HasMinTotalValueSize returns a boolean if a field has been set.

### GetDescription

`func (o *SimpleUncachedAttributeCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleUncachedAttributeCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleUncachedAttributeCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimpleUncachedAttributeCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SimpleUncachedAttributeCriteriaShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SimpleUncachedAttributeCriteriaShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SimpleUncachedAttributeCriteriaShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


