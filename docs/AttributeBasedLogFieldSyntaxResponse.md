# AttributeBasedLogFieldSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumattributeBasedLogFieldSyntaxSchemaUrn**](EnumattributeBasedLogFieldSyntaxSchemaUrn.md) |  | 
**IncludedSensitiveAttribute** | Pointer to **[]string** |  | [optional] 
**ExcludedSensitiveAttribute** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Syntax | [optional] 
**DefaultBehavior** | Pointer to [**EnumlogFieldSyntaxDefaultBehaviorProp**](EnumlogFieldSyntaxDefaultBehaviorProp.md) |  | [optional] 

## Methods

### NewAttributeBasedLogFieldSyntaxResponse

`func NewAttributeBasedLogFieldSyntaxResponse(schemas []EnumattributeBasedLogFieldSyntaxSchemaUrn, ) *AttributeBasedLogFieldSyntaxResponse`

NewAttributeBasedLogFieldSyntaxResponse instantiates a new AttributeBasedLogFieldSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeBasedLogFieldSyntaxResponseWithDefaults

`func NewAttributeBasedLogFieldSyntaxResponseWithDefaults() *AttributeBasedLogFieldSyntaxResponse`

NewAttributeBasedLogFieldSyntaxResponseWithDefaults instantiates a new AttributeBasedLogFieldSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AttributeBasedLogFieldSyntaxResponse) GetSchemas() []EnumattributeBasedLogFieldSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetSchemasOk() (*[]EnumattributeBasedLogFieldSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AttributeBasedLogFieldSyntaxResponse) SetSchemas(v []EnumattributeBasedLogFieldSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIncludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) GetIncludedSensitiveAttribute() []string`

GetIncludedSensitiveAttribute returns the IncludedSensitiveAttribute field if non-nil, zero value otherwise.

### GetIncludedSensitiveAttributeOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetIncludedSensitiveAttributeOk() (*[]string, bool)`

GetIncludedSensitiveAttributeOk returns a tuple with the IncludedSensitiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) SetIncludedSensitiveAttribute(v []string)`

SetIncludedSensitiveAttribute sets IncludedSensitiveAttribute field to given value.

### HasIncludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) HasIncludedSensitiveAttribute() bool`

HasIncludedSensitiveAttribute returns a boolean if a field has been set.

### GetExcludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) GetExcludedSensitiveAttribute() []string`

GetExcludedSensitiveAttribute returns the ExcludedSensitiveAttribute field if non-nil, zero value otherwise.

### GetExcludedSensitiveAttributeOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetExcludedSensitiveAttributeOk() (*[]string, bool)`

GetExcludedSensitiveAttributeOk returns a tuple with the ExcludedSensitiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) SetExcludedSensitiveAttribute(v []string)`

SetExcludedSensitiveAttribute sets ExcludedSensitiveAttribute field to given value.

### HasExcludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) HasExcludedSensitiveAttribute() bool`

HasExcludedSensitiveAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *AttributeBasedLogFieldSyntaxResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttributeBasedLogFieldSyntaxResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttributeBasedLogFieldSyntaxResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultBehavior

`func (o *AttributeBasedLogFieldSyntaxResponse) GetDefaultBehavior() EnumlogFieldSyntaxDefaultBehaviorProp`

GetDefaultBehavior returns the DefaultBehavior field if non-nil, zero value otherwise.

### GetDefaultBehaviorOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetDefaultBehaviorOk() (*EnumlogFieldSyntaxDefaultBehaviorProp, bool)`

GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehavior

`func (o *AttributeBasedLogFieldSyntaxResponse) SetDefaultBehavior(v EnumlogFieldSyntaxDefaultBehaviorProp)`

SetDefaultBehavior sets DefaultBehavior field to given value.

### HasDefaultBehavior

`func (o *AttributeBasedLogFieldSyntaxResponse) HasDefaultBehavior() bool`

HasDefaultBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


