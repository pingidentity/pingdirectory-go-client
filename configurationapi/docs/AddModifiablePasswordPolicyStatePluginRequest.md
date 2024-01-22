# AddModifiablePasswordPolicyStatePluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnummodifiablePasswordPolicyStatePluginSchemaUrn**](EnummodifiablePasswordPolicyStatePluginSchemaUrn.md) |  | 
**BaseDN** | Pointer to **[]string** | A base DN that may be used to identify entries that should support the ds-pwp-modifiable-state-json operational attribute. | [optional] 
**Filter** | Pointer to **[]string** | A filter that may be used to identify entries that should support the ds-pwp-modifiable-state-json operational attribute. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**PluginName** | **string** | Name of the new Plugin | 

## Methods

### NewAddModifiablePasswordPolicyStatePluginRequest

`func NewAddModifiablePasswordPolicyStatePluginRequest(schemas []EnummodifiablePasswordPolicyStatePluginSchemaUrn, enabled bool, pluginName string, ) *AddModifiablePasswordPolicyStatePluginRequest`

NewAddModifiablePasswordPolicyStatePluginRequest instantiates a new AddModifiablePasswordPolicyStatePluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddModifiablePasswordPolicyStatePluginRequestWithDefaults

`func NewAddModifiablePasswordPolicyStatePluginRequestWithDefaults() *AddModifiablePasswordPolicyStatePluginRequest`

NewAddModifiablePasswordPolicyStatePluginRequestWithDefaults instantiates a new AddModifiablePasswordPolicyStatePluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetSchemas() []EnummodifiablePasswordPolicyStatePluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetSchemasOk() (*[]EnummodifiablePasswordPolicyStatePluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddModifiablePasswordPolicyStatePluginRequest) SetSchemas(v []EnummodifiablePasswordPolicyStatePluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseDN

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddModifiablePasswordPolicyStatePluginRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddModifiablePasswordPolicyStatePluginRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetFilter

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddModifiablePasswordPolicyStatePluginRequest) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AddModifiablePasswordPolicyStatePluginRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDescription

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddModifiablePasswordPolicyStatePluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddModifiablePasswordPolicyStatePluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddModifiablePasswordPolicyStatePluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPluginName

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddModifiablePasswordPolicyStatePluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddModifiablePasswordPolicyStatePluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


