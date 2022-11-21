# TopologyAdminUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Topology Admin User | 
**Schemas** | Pointer to [**[]EnumtopologyAdminUserSchemaUrn**](EnumtopologyAdminUserSchemaUrn.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewTopologyAdminUserResponse

`func NewTopologyAdminUserResponse(id string, ) *TopologyAdminUserResponse`

NewTopologyAdminUserResponse instantiates a new TopologyAdminUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyAdminUserResponseWithDefaults

`func NewTopologyAdminUserResponseWithDefaults() *TopologyAdminUserResponse`

NewTopologyAdminUserResponseWithDefaults instantiates a new TopologyAdminUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TopologyAdminUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopologyAdminUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopologyAdminUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *TopologyAdminUserResponse) GetSchemas() []EnumtopologyAdminUserSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TopologyAdminUserResponse) GetSchemasOk() (*[]EnumtopologyAdminUserSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TopologyAdminUserResponse) SetSchemas(v []EnumtopologyAdminUserSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *TopologyAdminUserResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetMeta

`func (o *TopologyAdminUserResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TopologyAdminUserResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TopologyAdminUserResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TopologyAdminUserResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


