# AddScimSchemaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaName** | **string** | Name of the new SCIM Schema | 
**Schemas** | Pointer to [**[]EnumscimSchemaSchemaUrn**](EnumscimSchemaSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this SCIM Schema | [optional] 
**SchemaURN** | **string** | The URN which identifies this SCIM Schema. | 
**DisplayName** | Pointer to **string** | The human readable name for this SCIM Schema. | [optional] 

## Methods

### NewAddScimSchemaRequest

`func NewAddScimSchemaRequest(schemaName string, schemaURN string, ) *AddScimSchemaRequest`

NewAddScimSchemaRequest instantiates a new AddScimSchemaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddScimSchemaRequestWithDefaults

`func NewAddScimSchemaRequestWithDefaults() *AddScimSchemaRequest`

NewAddScimSchemaRequestWithDefaults instantiates a new AddScimSchemaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaName

`func (o *AddScimSchemaRequest) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *AddScimSchemaRequest) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *AddScimSchemaRequest) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.


### GetSchemas

`func (o *AddScimSchemaRequest) GetSchemas() []EnumscimSchemaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddScimSchemaRequest) GetSchemasOk() (*[]EnumscimSchemaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddScimSchemaRequest) SetSchemas(v []EnumscimSchemaSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddScimSchemaRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddScimSchemaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddScimSchemaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddScimSchemaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddScimSchemaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchemaURN

`func (o *AddScimSchemaRequest) GetSchemaURN() string`

GetSchemaURN returns the SchemaURN field if non-nil, zero value otherwise.

### GetSchemaURNOk

`func (o *AddScimSchemaRequest) GetSchemaURNOk() (*string, bool)`

GetSchemaURNOk returns a tuple with the SchemaURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaURN

`func (o *AddScimSchemaRequest) SetSchemaURN(v string)`

SetSchemaURN sets SchemaURN field to given value.


### GetDisplayName

`func (o *AddScimSchemaRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddScimSchemaRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddScimSchemaRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddScimSchemaRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


