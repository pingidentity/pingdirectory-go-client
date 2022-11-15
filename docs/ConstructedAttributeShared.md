# ConstructedAttributeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumconstructedAttributeSchemaUrn**](EnumconstructedAttributeSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Constructed Attribute | [optional] 
**AttributeType** | **string** | Specifies the attribute type for the attribute whose values are to be constructed. | 
**ValuePattern** | **[]string** |  | 

## Methods

### NewConstructedAttributeShared

`func NewConstructedAttributeShared(attributeType string, valuePattern []string, ) *ConstructedAttributeShared`

NewConstructedAttributeShared instantiates a new ConstructedAttributeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstructedAttributeSharedWithDefaults

`func NewConstructedAttributeSharedWithDefaults() *ConstructedAttributeShared`

NewConstructedAttributeSharedWithDefaults instantiates a new ConstructedAttributeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ConstructedAttributeShared) GetSchemas() []EnumconstructedAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConstructedAttributeShared) GetSchemasOk() (*[]EnumconstructedAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConstructedAttributeShared) SetSchemas(v []EnumconstructedAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ConstructedAttributeShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ConstructedAttributeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConstructedAttributeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConstructedAttributeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConstructedAttributeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeType

`func (o *ConstructedAttributeShared) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *ConstructedAttributeShared) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *ConstructedAttributeShared) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetValuePattern

`func (o *ConstructedAttributeShared) GetValuePattern() []string`

GetValuePattern returns the ValuePattern field if non-nil, zero value otherwise.

### GetValuePatternOk

`func (o *ConstructedAttributeShared) GetValuePatternOk() (*[]string, bool)`

GetValuePatternOk returns a tuple with the ValuePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePattern

`func (o *ConstructedAttributeShared) SetValuePattern(v []string)`

SetValuePattern sets ValuePattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


