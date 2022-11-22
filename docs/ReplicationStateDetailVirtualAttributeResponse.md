# ReplicationStateDetailVirtualAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumreplicationStateDetailVirtualAttributeSchemaUrn**](EnumreplicationStateDetailVirtualAttributeSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Virtual Attribute is enabled for use. | 
**RequireExplicitRequestByName** | Pointer to **bool** | Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type. | [optional] 

## Methods

### NewReplicationStateDetailVirtualAttributeResponse

`func NewReplicationStateDetailVirtualAttributeResponse(schemas []EnumreplicationStateDetailVirtualAttributeSchemaUrn, enabled bool, ) *ReplicationStateDetailVirtualAttributeResponse`

NewReplicationStateDetailVirtualAttributeResponse instantiates a new ReplicationStateDetailVirtualAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationStateDetailVirtualAttributeResponseWithDefaults

`func NewReplicationStateDetailVirtualAttributeResponseWithDefaults() *ReplicationStateDetailVirtualAttributeResponse`

NewReplicationStateDetailVirtualAttributeResponseWithDefaults instantiates a new ReplicationStateDetailVirtualAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ReplicationStateDetailVirtualAttributeResponse) GetSchemas() []EnumreplicationStateDetailVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ReplicationStateDetailVirtualAttributeResponse) GetSchemasOk() (*[]EnumreplicationStateDetailVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ReplicationStateDetailVirtualAttributeResponse) SetSchemas(v []EnumreplicationStateDetailVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *ReplicationStateDetailVirtualAttributeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReplicationStateDetailVirtualAttributeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReplicationStateDetailVirtualAttributeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireExplicitRequestByName

`func (o *ReplicationStateDetailVirtualAttributeResponse) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *ReplicationStateDetailVirtualAttributeResponse) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *ReplicationStateDetailVirtualAttributeResponse) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *ReplicationStateDetailVirtualAttributeResponse) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


