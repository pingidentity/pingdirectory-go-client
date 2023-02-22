# PasswordPolicyImportPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumpasswordPolicyImportPluginSchemaUrn**](EnumpasswordPolicyImportPluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**DefaultUserPasswordStorageScheme** | Pointer to **[]string** | Specifies the names of the password storage schemes to be used for encoding passwords contained in attributes with the user password syntax for entries that do not include the ds-pwp-password-policy-dn attribute specifying which password policy is to be used to govern them. | [optional] 
**DefaultAuthPasswordStorageScheme** | Pointer to **[]string** | Specifies the names of password storage schemes that to be used for encoding passwords contained in attributes with the auth password syntax for entries that do not include the ds-pwp-password-policy-dn attribute specifying which password policy should be used to govern them. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewPasswordPolicyImportPluginResponse

`func NewPasswordPolicyImportPluginResponse(schemas []EnumpasswordPolicyImportPluginSchemaUrn, id string, enabled bool, ) *PasswordPolicyImportPluginResponse`

NewPasswordPolicyImportPluginResponse instantiates a new PasswordPolicyImportPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyImportPluginResponseWithDefaults

`func NewPasswordPolicyImportPluginResponseWithDefaults() *PasswordPolicyImportPluginResponse`

NewPasswordPolicyImportPluginResponseWithDefaults instantiates a new PasswordPolicyImportPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PasswordPolicyImportPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PasswordPolicyImportPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PasswordPolicyImportPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PasswordPolicyImportPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PasswordPolicyImportPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PasswordPolicyImportPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PasswordPolicyImportPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PasswordPolicyImportPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *PasswordPolicyImportPluginResponse) GetSchemas() []EnumpasswordPolicyImportPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PasswordPolicyImportPluginResponse) GetSchemasOk() (*[]EnumpasswordPolicyImportPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PasswordPolicyImportPluginResponse) SetSchemas(v []EnumpasswordPolicyImportPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *PasswordPolicyImportPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PasswordPolicyImportPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PasswordPolicyImportPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInvokeForInternalOperations

`func (o *PasswordPolicyImportPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *PasswordPolicyImportPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *PasswordPolicyImportPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *PasswordPolicyImportPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetDefaultUserPasswordStorageScheme

`func (o *PasswordPolicyImportPluginResponse) GetDefaultUserPasswordStorageScheme() []string`

GetDefaultUserPasswordStorageScheme returns the DefaultUserPasswordStorageScheme field if non-nil, zero value otherwise.

### GetDefaultUserPasswordStorageSchemeOk

`func (o *PasswordPolicyImportPluginResponse) GetDefaultUserPasswordStorageSchemeOk() (*[]string, bool)`

GetDefaultUserPasswordStorageSchemeOk returns a tuple with the DefaultUserPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserPasswordStorageScheme

`func (o *PasswordPolicyImportPluginResponse) SetDefaultUserPasswordStorageScheme(v []string)`

SetDefaultUserPasswordStorageScheme sets DefaultUserPasswordStorageScheme field to given value.

### HasDefaultUserPasswordStorageScheme

`func (o *PasswordPolicyImportPluginResponse) HasDefaultUserPasswordStorageScheme() bool`

HasDefaultUserPasswordStorageScheme returns a boolean if a field has been set.

### GetDefaultAuthPasswordStorageScheme

`func (o *PasswordPolicyImportPluginResponse) GetDefaultAuthPasswordStorageScheme() []string`

GetDefaultAuthPasswordStorageScheme returns the DefaultAuthPasswordStorageScheme field if non-nil, zero value otherwise.

### GetDefaultAuthPasswordStorageSchemeOk

`func (o *PasswordPolicyImportPluginResponse) GetDefaultAuthPasswordStorageSchemeOk() (*[]string, bool)`

GetDefaultAuthPasswordStorageSchemeOk returns a tuple with the DefaultAuthPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAuthPasswordStorageScheme

`func (o *PasswordPolicyImportPluginResponse) SetDefaultAuthPasswordStorageScheme(v []string)`

SetDefaultAuthPasswordStorageScheme sets DefaultAuthPasswordStorageScheme field to given value.

### HasDefaultAuthPasswordStorageScheme

`func (o *PasswordPolicyImportPluginResponse) HasDefaultAuthPasswordStorageScheme() bool`

HasDefaultAuthPasswordStorageScheme returns a boolean if a field has been set.

### GetDescription

`func (o *PasswordPolicyImportPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PasswordPolicyImportPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PasswordPolicyImportPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PasswordPolicyImportPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PasswordPolicyImportPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PasswordPolicyImportPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PasswordPolicyImportPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


