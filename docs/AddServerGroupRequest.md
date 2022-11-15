# AddServerGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** | Name of the new Server Group | 
**Schemas** | Pointer to [**[]EnumserverGroupSchemaUrn**](EnumserverGroupSchemaUrn.md) |  | [optional] 
**Member** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddServerGroupRequest

`func NewAddServerGroupRequest(groupName string, ) *AddServerGroupRequest`

NewAddServerGroupRequest instantiates a new AddServerGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServerGroupRequestWithDefaults

`func NewAddServerGroupRequestWithDefaults() *AddServerGroupRequest`

NewAddServerGroupRequestWithDefaults instantiates a new AddServerGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *AddServerGroupRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AddServerGroupRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AddServerGroupRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetSchemas

`func (o *AddServerGroupRequest) GetSchemas() []EnumserverGroupSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddServerGroupRequest) GetSchemasOk() (*[]EnumserverGroupSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddServerGroupRequest) SetSchemas(v []EnumserverGroupSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddServerGroupRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetMember

`func (o *AddServerGroupRequest) GetMember() []string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *AddServerGroupRequest) GetMemberOk() (*[]string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *AddServerGroupRequest) SetMember(v []string)`

SetMember sets Member field to given value.

### HasMember

`func (o *AddServerGroupRequest) HasMember() bool`

HasMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


