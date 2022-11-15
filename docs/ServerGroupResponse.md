# ServerGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Server Group | 
**Schemas** | Pointer to [**[]EnumserverGroupSchemaUrn**](EnumserverGroupSchemaUrn.md) |  | [optional] 
**Member** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServerGroupResponse

`func NewServerGroupResponse(id string, ) *ServerGroupResponse`

NewServerGroupResponse instantiates a new ServerGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerGroupResponseWithDefaults

`func NewServerGroupResponseWithDefaults() *ServerGroupResponse`

NewServerGroupResponseWithDefaults instantiates a new ServerGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ServerGroupResponse) GetSchemas() []EnumserverGroupSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ServerGroupResponse) GetSchemasOk() (*[]EnumserverGroupSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ServerGroupResponse) SetSchemas(v []EnumserverGroupSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ServerGroupResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetMember

`func (o *ServerGroupResponse) GetMember() []string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ServerGroupResponse) GetMemberOk() (*[]string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ServerGroupResponse) SetMember(v []string)`

SetMember sets Member field to given value.

### HasMember

`func (o *ServerGroupResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


