# ServerGroupShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumserverGroupSchemaUrn**](EnumserverGroupSchemaUrn.md) |  | [optional] 
**Member** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServerGroupShared

`func NewServerGroupShared() *ServerGroupShared`

NewServerGroupShared instantiates a new ServerGroupShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerGroupSharedWithDefaults

`func NewServerGroupSharedWithDefaults() *ServerGroupShared`

NewServerGroupSharedWithDefaults instantiates a new ServerGroupShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ServerGroupShared) GetSchemas() []EnumserverGroupSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ServerGroupShared) GetSchemasOk() (*[]EnumserverGroupSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ServerGroupShared) SetSchemas(v []EnumserverGroupSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ServerGroupShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetMember

`func (o *ServerGroupShared) GetMember() []string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ServerGroupShared) GetMemberOk() (*[]string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ServerGroupShared) SetMember(v []string)`

SetMember sets Member field to given value.

### HasMember

`func (o *ServerGroupShared) HasMember() bool`

HasMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


