# AddTopologyAdminUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | Name of the new Topology Admin User | 
**Schemas** | Pointer to [**[]EnumtopologyAdminUserSchemaUrn**](EnumtopologyAdminUserSchemaUrn.md) |  | [optional] 

## Methods

### NewAddTopologyAdminUserRequest

`func NewAddTopologyAdminUserRequest(userName string, ) *AddTopologyAdminUserRequest`

NewAddTopologyAdminUserRequest instantiates a new AddTopologyAdminUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTopologyAdminUserRequestWithDefaults

`func NewAddTopologyAdminUserRequestWithDefaults() *AddTopologyAdminUserRequest`

NewAddTopologyAdminUserRequestWithDefaults instantiates a new AddTopologyAdminUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *AddTopologyAdminUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AddTopologyAdminUserRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AddTopologyAdminUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetSchemas

`func (o *AddTopologyAdminUserRequest) GetSchemas() []EnumtopologyAdminUserSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTopologyAdminUserRequest) GetSchemasOk() (*[]EnumtopologyAdminUserSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTopologyAdminUserRequest) SetSchemas(v []EnumtopologyAdminUserSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddTopologyAdminUserRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


