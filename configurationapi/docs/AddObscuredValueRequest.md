# AddObscuredValueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumobscuredValueSchemaUrn**](EnumobscuredValueSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Obscured Value | [optional] 
**ObscuredValue** | **string** | The value to be stored in an obscured form. | 
**ValueName** | **string** | Name of the new Obscured Value | 

## Methods

### NewAddObscuredValueRequest

`func NewAddObscuredValueRequest(obscuredValue string, valueName string, ) *AddObscuredValueRequest`

NewAddObscuredValueRequest instantiates a new AddObscuredValueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddObscuredValueRequestWithDefaults

`func NewAddObscuredValueRequestWithDefaults() *AddObscuredValueRequest`

NewAddObscuredValueRequestWithDefaults instantiates a new AddObscuredValueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddObscuredValueRequest) GetSchemas() []EnumobscuredValueSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddObscuredValueRequest) GetSchemasOk() (*[]EnumobscuredValueSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddObscuredValueRequest) SetSchemas(v []EnumobscuredValueSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddObscuredValueRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddObscuredValueRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddObscuredValueRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddObscuredValueRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddObscuredValueRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObscuredValue

`func (o *AddObscuredValueRequest) GetObscuredValue() string`

GetObscuredValue returns the ObscuredValue field if non-nil, zero value otherwise.

### GetObscuredValueOk

`func (o *AddObscuredValueRequest) GetObscuredValueOk() (*string, bool)`

GetObscuredValueOk returns a tuple with the ObscuredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObscuredValue

`func (o *AddObscuredValueRequest) SetObscuredValue(v string)`

SetObscuredValue sets ObscuredValue field to given value.


### GetValueName

`func (o *AddObscuredValueRequest) GetValueName() string`

GetValueName returns the ValueName field if non-nil, zero value otherwise.

### GetValueNameOk

`func (o *AddObscuredValueRequest) GetValueNameOk() (*string, bool)`

GetValueNameOk returns a tuple with the ValueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueName

`func (o *AddObscuredValueRequest) SetValueName(v string)`

SetValueName sets ValueName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


