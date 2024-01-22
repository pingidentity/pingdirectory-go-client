# AddConstructedAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumconstructedAttributeSchemaUrn**](EnumconstructedAttributeSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Constructed Attribute | [optional] 
**AttributeType** | **string** | Specifies the attribute type for the attribute whose values are to be constructed. | 
**ValuePattern** | **[]string** | Specifies a pattern for constructing the attribute value using fixed text and attribute values from the entry. | 
**AttributeName** | **string** | Name of the new Constructed Attribute | 

## Methods

### NewAddConstructedAttributeRequest

`func NewAddConstructedAttributeRequest(attributeType string, valuePattern []string, attributeName string, ) *AddConstructedAttributeRequest`

NewAddConstructedAttributeRequest instantiates a new AddConstructedAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddConstructedAttributeRequestWithDefaults

`func NewAddConstructedAttributeRequestWithDefaults() *AddConstructedAttributeRequest`

NewAddConstructedAttributeRequestWithDefaults instantiates a new AddConstructedAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddConstructedAttributeRequest) GetSchemas() []EnumconstructedAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddConstructedAttributeRequest) GetSchemasOk() (*[]EnumconstructedAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddConstructedAttributeRequest) SetSchemas(v []EnumconstructedAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddConstructedAttributeRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddConstructedAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddConstructedAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddConstructedAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddConstructedAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeType

`func (o *AddConstructedAttributeRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddConstructedAttributeRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddConstructedAttributeRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetValuePattern

`func (o *AddConstructedAttributeRequest) GetValuePattern() []string`

GetValuePattern returns the ValuePattern field if non-nil, zero value otherwise.

### GetValuePatternOk

`func (o *AddConstructedAttributeRequest) GetValuePatternOk() (*[]string, bool)`

GetValuePatternOk returns a tuple with the ValuePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePattern

`func (o *AddConstructedAttributeRequest) SetValuePattern(v []string)`

SetValuePattern sets ValuePattern field to given value.


### GetAttributeName

`func (o *AddConstructedAttributeRequest) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *AddConstructedAttributeRequest) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *AddConstructedAttributeRequest) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


