# AddSimpleUncachedAttributeCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Uncached Attribute Criteria | 
**Schemas** | [**[]EnumsimpleUncachedAttributeCriteriaSchemaUrn**](EnumsimpleUncachedAttributeCriteriaSchemaUrn.md) |  | 
**AttributeType** | **[]string** | Specifies the attribute types for attributes that may be written to the uncached-id2entry database. | 
**MinValueCount** | Pointer to **int32** | Specifies the minimum number of values that an attribute must have before it will be written into the uncached-id2entry database. | [optional] 
**MinTotalValueSize** | Pointer to **string** | Specifies the minimum total value size (i.e., the sum of the sizes of all values) that an attribute must have before it will be written into the uncached-id2entry database. | [optional] 
**Description** | Pointer to **string** | A description for this Uncached Attribute Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Attribute Criteria is enabled for use in the server. | 

## Methods

### NewAddSimpleUncachedAttributeCriteriaRequest

`func NewAddSimpleUncachedAttributeCriteriaRequest(criteriaName string, schemas []EnumsimpleUncachedAttributeCriteriaSchemaUrn, attributeType []string, enabled bool, ) *AddSimpleUncachedAttributeCriteriaRequest`

NewAddSimpleUncachedAttributeCriteriaRequest instantiates a new AddSimpleUncachedAttributeCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSimpleUncachedAttributeCriteriaRequestWithDefaults

`func NewAddSimpleUncachedAttributeCriteriaRequestWithDefaults() *AddSimpleUncachedAttributeCriteriaRequest`

NewAddSimpleUncachedAttributeCriteriaRequestWithDefaults instantiates a new AddSimpleUncachedAttributeCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddSimpleUncachedAttributeCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetSchemas() []EnumsimpleUncachedAttributeCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetSchemasOk() (*[]EnumsimpleUncachedAttributeCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSimpleUncachedAttributeCriteriaRequest) SetSchemas(v []EnumsimpleUncachedAttributeCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAttributeType

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddSimpleUncachedAttributeCriteriaRequest) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetMinValueCount

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetMinValueCount() int32`

GetMinValueCount returns the MinValueCount field if non-nil, zero value otherwise.

### GetMinValueCountOk

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetMinValueCountOk() (*int32, bool)`

GetMinValueCountOk returns a tuple with the MinValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValueCount

`func (o *AddSimpleUncachedAttributeCriteriaRequest) SetMinValueCount(v int32)`

SetMinValueCount sets MinValueCount field to given value.

### HasMinValueCount

`func (o *AddSimpleUncachedAttributeCriteriaRequest) HasMinValueCount() bool`

HasMinValueCount returns a boolean if a field has been set.

### GetMinTotalValueSize

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetMinTotalValueSize() string`

GetMinTotalValueSize returns the MinTotalValueSize field if non-nil, zero value otherwise.

### GetMinTotalValueSizeOk

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetMinTotalValueSizeOk() (*string, bool)`

GetMinTotalValueSizeOk returns a tuple with the MinTotalValueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTotalValueSize

`func (o *AddSimpleUncachedAttributeCriteriaRequest) SetMinTotalValueSize(v string)`

SetMinTotalValueSize sets MinTotalValueSize field to given value.

### HasMinTotalValueSize

`func (o *AddSimpleUncachedAttributeCriteriaRequest) HasMinTotalValueSize() bool`

HasMinTotalValueSize returns a boolean if a field has been set.

### GetDescription

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSimpleUncachedAttributeCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSimpleUncachedAttributeCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSimpleUncachedAttributeCriteriaRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSimpleUncachedAttributeCriteriaRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


