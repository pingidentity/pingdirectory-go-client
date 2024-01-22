# ModifiablePasswordPolicyStatePluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnummodifiablePasswordPolicyStatePluginSchemaUrn**](EnummodifiablePasswordPolicyStatePluginSchemaUrn.md) |  | 
**BaseDN** | Pointer to **[]string** | A base DN that may be used to identify entries that should support the ds-pwp-modifiable-state-json operational attribute. | [optional] 
**Filter** | Pointer to **[]string** | A filter that may be used to identify entries that should support the ds-pwp-modifiable-state-json operational attribute. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 

## Methods

### NewModifiablePasswordPolicyStatePluginResponse

`func NewModifiablePasswordPolicyStatePluginResponse(schemas []EnummodifiablePasswordPolicyStatePluginSchemaUrn, enabled bool, id string, ) *ModifiablePasswordPolicyStatePluginResponse`

NewModifiablePasswordPolicyStatePluginResponse instantiates a new ModifiablePasswordPolicyStatePluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifiablePasswordPolicyStatePluginResponseWithDefaults

`func NewModifiablePasswordPolicyStatePluginResponseWithDefaults() *ModifiablePasswordPolicyStatePluginResponse`

NewModifiablePasswordPolicyStatePluginResponseWithDefaults instantiates a new ModifiablePasswordPolicyStatePluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetSchemas() []EnummodifiablePasswordPolicyStatePluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetSchemasOk() (*[]EnummodifiablePasswordPolicyStatePluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ModifiablePasswordPolicyStatePluginResponse) SetSchemas(v []EnummodifiablePasswordPolicyStatePluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseDN

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *ModifiablePasswordPolicyStatePluginResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *ModifiablePasswordPolicyStatePluginResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetFilter

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ModifiablePasswordPolicyStatePluginResponse) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ModifiablePasswordPolicyStatePluginResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDescription

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModifiablePasswordPolicyStatePluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModifiablePasswordPolicyStatePluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModifiablePasswordPolicyStatePluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ModifiablePasswordPolicyStatePluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ModifiablePasswordPolicyStatePluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ModifiablePasswordPolicyStatePluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ModifiablePasswordPolicyStatePluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifiablePasswordPolicyStatePluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifiablePasswordPolicyStatePluginResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


