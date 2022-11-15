# RootDnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumrootDnSchemaUrn**](EnumrootDnSchemaUrn.md) |  | [optional] 
**DefaultRootPrivilegeName** | Pointer to [**[]EnumrootDnDefaultRootPrivilegeNameProp**](EnumrootDnDefaultRootPrivilegeNameProp.md) |  | [optional] 

## Methods

### NewRootDnResponse

`func NewRootDnResponse() *RootDnResponse`

NewRootDnResponse instantiates a new RootDnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootDnResponseWithDefaults

`func NewRootDnResponseWithDefaults() *RootDnResponse`

NewRootDnResponseWithDefaults instantiates a new RootDnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *RootDnResponse) GetSchemas() []EnumrootDnSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RootDnResponse) GetSchemasOk() (*[]EnumrootDnSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RootDnResponse) SetSchemas(v []EnumrootDnSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *RootDnResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDefaultRootPrivilegeName

`func (o *RootDnResponse) GetDefaultRootPrivilegeName() []EnumrootDnDefaultRootPrivilegeNameProp`

GetDefaultRootPrivilegeName returns the DefaultRootPrivilegeName field if non-nil, zero value otherwise.

### GetDefaultRootPrivilegeNameOk

`func (o *RootDnResponse) GetDefaultRootPrivilegeNameOk() (*[]EnumrootDnDefaultRootPrivilegeNameProp, bool)`

GetDefaultRootPrivilegeNameOk returns a tuple with the DefaultRootPrivilegeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRootPrivilegeName

`func (o *RootDnResponse) SetDefaultRootPrivilegeName(v []EnumrootDnDefaultRootPrivilegeNameProp)`

SetDefaultRootPrivilegeName sets DefaultRootPrivilegeName field to given value.

### HasDefaultRootPrivilegeName

`func (o *RootDnResponse) HasDefaultRootPrivilegeName() bool`

HasDefaultRootPrivilegeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


