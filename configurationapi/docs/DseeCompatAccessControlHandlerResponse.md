# DseeCompatAccessControlHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumdseeCompatAccessControlHandlerSchemaUrn**](EnumdseeCompatAccessControlHandlerSchemaUrn.md) |  | 
**GlobalACI** | Pointer to **[]string** | Defines global access control rules. | [optional] 
**AllowedBindControl** | Pointer to [**[]EnumaccessControlHandlerAllowedBindControlProp**](EnumaccessControlHandlerAllowedBindControlProp.md) |  | [optional] 
**AllowedBindControlOID** | Pointer to **[]string** | Specifies the OIDs of any additional controls (not covered by the allowed-bind-control property) that should be permitted in bind requests. | [optional] 
**Enabled** | **bool** | Indicates whether this Access Control Handler is enabled. If set to FALSE, then no access control is enforced, and any client (including unauthenticated or anonymous clients) could be allowed to perform any operation if not subject to other restrictions, such as those enforced by the privilege subsystem. | 

## Methods

### NewDseeCompatAccessControlHandlerResponse

`func NewDseeCompatAccessControlHandlerResponse(schemas []EnumdseeCompatAccessControlHandlerSchemaUrn, enabled bool, ) *DseeCompatAccessControlHandlerResponse`

NewDseeCompatAccessControlHandlerResponse instantiates a new DseeCompatAccessControlHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDseeCompatAccessControlHandlerResponseWithDefaults

`func NewDseeCompatAccessControlHandlerResponseWithDefaults() *DseeCompatAccessControlHandlerResponse`

NewDseeCompatAccessControlHandlerResponseWithDefaults instantiates a new DseeCompatAccessControlHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *DseeCompatAccessControlHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DseeCompatAccessControlHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DseeCompatAccessControlHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DseeCompatAccessControlHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DseeCompatAccessControlHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DseeCompatAccessControlHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DseeCompatAccessControlHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DseeCompatAccessControlHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *DseeCompatAccessControlHandlerResponse) GetSchemas() []EnumdseeCompatAccessControlHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DseeCompatAccessControlHandlerResponse) GetSchemasOk() (*[]EnumdseeCompatAccessControlHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DseeCompatAccessControlHandlerResponse) SetSchemas(v []EnumdseeCompatAccessControlHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetGlobalACI

`func (o *DseeCompatAccessControlHandlerResponse) GetGlobalACI() []string`

GetGlobalACI returns the GlobalACI field if non-nil, zero value otherwise.

### GetGlobalACIOk

`func (o *DseeCompatAccessControlHandlerResponse) GetGlobalACIOk() (*[]string, bool)`

GetGlobalACIOk returns a tuple with the GlobalACI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalACI

`func (o *DseeCompatAccessControlHandlerResponse) SetGlobalACI(v []string)`

SetGlobalACI sets GlobalACI field to given value.

### HasGlobalACI

`func (o *DseeCompatAccessControlHandlerResponse) HasGlobalACI() bool`

HasGlobalACI returns a boolean if a field has been set.

### GetAllowedBindControl

`func (o *DseeCompatAccessControlHandlerResponse) GetAllowedBindControl() []EnumaccessControlHandlerAllowedBindControlProp`

GetAllowedBindControl returns the AllowedBindControl field if non-nil, zero value otherwise.

### GetAllowedBindControlOk

`func (o *DseeCompatAccessControlHandlerResponse) GetAllowedBindControlOk() (*[]EnumaccessControlHandlerAllowedBindControlProp, bool)`

GetAllowedBindControlOk returns a tuple with the AllowedBindControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedBindControl

`func (o *DseeCompatAccessControlHandlerResponse) SetAllowedBindControl(v []EnumaccessControlHandlerAllowedBindControlProp)`

SetAllowedBindControl sets AllowedBindControl field to given value.

### HasAllowedBindControl

`func (o *DseeCompatAccessControlHandlerResponse) HasAllowedBindControl() bool`

HasAllowedBindControl returns a boolean if a field has been set.

### GetAllowedBindControlOID

`func (o *DseeCompatAccessControlHandlerResponse) GetAllowedBindControlOID() []string`

GetAllowedBindControlOID returns the AllowedBindControlOID field if non-nil, zero value otherwise.

### GetAllowedBindControlOIDOk

`func (o *DseeCompatAccessControlHandlerResponse) GetAllowedBindControlOIDOk() (*[]string, bool)`

GetAllowedBindControlOIDOk returns a tuple with the AllowedBindControlOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedBindControlOID

`func (o *DseeCompatAccessControlHandlerResponse) SetAllowedBindControlOID(v []string)`

SetAllowedBindControlOID sets AllowedBindControlOID field to given value.

### HasAllowedBindControlOID

`func (o *DseeCompatAccessControlHandlerResponse) HasAllowedBindControlOID() bool`

HasAllowedBindControlOID returns a boolean if a field has been set.

### GetEnabled

`func (o *DseeCompatAccessControlHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DseeCompatAccessControlHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DseeCompatAccessControlHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


