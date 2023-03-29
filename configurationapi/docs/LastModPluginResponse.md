# LastModPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumlastModPluginSchemaUrn**](EnumlastModPluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin Root | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) | Specifies the set of plug-in types for the plug-in, which specifies the times at which the plug-in is invoked. | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the name or OID of an attribute type that must be updated in order for the modifiersName and modifyTimestamp attributes to be updated in the target entry. | [optional] 
**ExcludeAttribute** | Pointer to **[]string** | Specifies the name or OID of an attribute type which may be updated in a modify or modify DN operation without causing the modifiersName and modifyTimestamp values to be updated for that entry. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewLastModPluginResponse

`func NewLastModPluginResponse(schemas []EnumlastModPluginSchemaUrn, id string, pluginType []EnumpluginPluginTypeProp, enabled bool, ) *LastModPluginResponse`

NewLastModPluginResponse instantiates a new LastModPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastModPluginResponseWithDefaults

`func NewLastModPluginResponseWithDefaults() *LastModPluginResponse`

NewLastModPluginResponseWithDefaults instantiates a new LastModPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LastModPluginResponse) GetSchemas() []EnumlastModPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LastModPluginResponse) GetSchemasOk() (*[]EnumlastModPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LastModPluginResponse) SetSchemas(v []EnumlastModPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *LastModPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LastModPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LastModPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPluginType

`func (o *LastModPluginResponse) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *LastModPluginResponse) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *LastModPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetIncludeAttribute

`func (o *LastModPluginResponse) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *LastModPluginResponse) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *LastModPluginResponse) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *LastModPluginResponse) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetExcludeAttribute

`func (o *LastModPluginResponse) GetExcludeAttribute() []string`

GetExcludeAttribute returns the ExcludeAttribute field if non-nil, zero value otherwise.

### GetExcludeAttributeOk

`func (o *LastModPluginResponse) GetExcludeAttributeOk() (*[]string, bool)`

GetExcludeAttributeOk returns a tuple with the ExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttribute

`func (o *LastModPluginResponse) SetExcludeAttribute(v []string)`

SetExcludeAttribute sets ExcludeAttribute field to given value.

### HasExcludeAttribute

`func (o *LastModPluginResponse) HasExcludeAttribute() bool`

HasExcludeAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *LastModPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LastModPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LastModPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LastModPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LastModPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LastModPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LastModPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *LastModPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *LastModPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *LastModPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *LastModPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetMeta

`func (o *LastModPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LastModPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LastModPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LastModPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LastModPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LastModPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LastModPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LastModPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


