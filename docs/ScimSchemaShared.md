# ScimSchemaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumscimSchemaSchemaUrn**](EnumscimSchemaSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this SCIM Schema | [optional] 
**SchemaURN** | **string** | The URN which identifies this SCIM Schema. | 
**DisplayName** | Pointer to **string** | The human readable name for this SCIM Schema. | [optional] 

## Methods

### NewScimSchemaShared

`func NewScimSchemaShared(schemaURN string, ) *ScimSchemaShared`

NewScimSchemaShared instantiates a new ScimSchemaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSchemaSharedWithDefaults

`func NewScimSchemaSharedWithDefaults() *ScimSchemaShared`

NewScimSchemaSharedWithDefaults instantiates a new ScimSchemaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimSchemaShared) GetSchemas() []EnumscimSchemaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimSchemaShared) GetSchemasOk() (*[]EnumscimSchemaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimSchemaShared) SetSchemas(v []EnumscimSchemaSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimSchemaShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ScimSchemaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimSchemaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimSchemaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScimSchemaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchemaURN

`func (o *ScimSchemaShared) GetSchemaURN() string`

GetSchemaURN returns the SchemaURN field if non-nil, zero value otherwise.

### GetSchemaURNOk

`func (o *ScimSchemaShared) GetSchemaURNOk() (*string, bool)`

GetSchemaURNOk returns a tuple with the SchemaURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaURN

`func (o *ScimSchemaShared) SetSchemaURN(v string)`

SetSchemaURN sets SchemaURN field to given value.


### GetDisplayName

`func (o *ScimSchemaShared) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimSchemaShared) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimSchemaShared) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ScimSchemaShared) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


