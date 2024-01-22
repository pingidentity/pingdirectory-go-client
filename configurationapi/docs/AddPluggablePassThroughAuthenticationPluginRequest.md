# AddPluggablePassThroughAuthenticationPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumpluggablePassThroughAuthenticationPluginSchemaUrn**](EnumpluggablePassThroughAuthenticationPluginSchemaUrn.md) |  | 
**PassThroughAuthenticationHandler** | **string** | The component used to manage authentication with the external authentication service. | 
**IncludedLocalEntryBaseDN** | Pointer to **[]string** | The base DNs for the local users whose authentication attempts may be passed through to the external authentication service. | [optional] 
**ConnectionCriteria** | Pointer to **string** | A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**RequestCriteria** | Pointer to **string** | A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**TryLocalBind** | Pointer to **bool** | Indicates whether to attempt the bind in the local server first and only send the request to the external authentication service if the local bind attempt fails, or to only attempt the bind in the external service. | [optional] 
**OverrideLocalPassword** | Pointer to **bool** | Indicates whether to attempt the authentication in the external service if the local user entry includes a password. This property will be ignored if try-local-bind is false. | [optional] 
**UpdateLocalPassword** | Pointer to **bool** | Indicates whether to overwrite the user&#39;s local password if the local bind fails but the authentication attempt succeeds when attempted in the external service. This property may only be set to true if try-local-bind is also true. | [optional] 
**UpdateLocalPasswordDN** | Pointer to **string** | The DN of the authorization identity that will be used when updating the user&#39;s local password if update-local-password is true. This is primarily intended for use if the Data Sync Server will be used to synchronize passwords between the local server and the external service, and in that case, the DN used here should also be added to the ignore-changes-by-dn property in the appropriate Sync Source object in the Data Sync Server configuration. | [optional] 
**AllowLaxPassThroughAuthenticationPasswords** | Pointer to **bool** | Indicates whether to overwrite the user&#39;s local password even if the password used to authenticate to the external service would have failed validation if the user attempted to set it directly. | [optional] 
**IgnoredPasswordPolicyStateErrorCondition** | Pointer to [**[]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp**](EnumpluginIgnoredPasswordPolicyStateErrorConditionProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**PluginName** | **string** | Name of the new Plugin | 

## Methods

### NewAddPluggablePassThroughAuthenticationPluginRequest

`func NewAddPluggablePassThroughAuthenticationPluginRequest(schemas []EnumpluggablePassThroughAuthenticationPluginSchemaUrn, passThroughAuthenticationHandler string, enabled bool, pluginName string, ) *AddPluggablePassThroughAuthenticationPluginRequest`

NewAddPluggablePassThroughAuthenticationPluginRequest instantiates a new AddPluggablePassThroughAuthenticationPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPluggablePassThroughAuthenticationPluginRequestWithDefaults

`func NewAddPluggablePassThroughAuthenticationPluginRequestWithDefaults() *AddPluggablePassThroughAuthenticationPluginRequest`

NewAddPluggablePassThroughAuthenticationPluginRequestWithDefaults instantiates a new AddPluggablePassThroughAuthenticationPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetSchemas() []EnumpluggablePassThroughAuthenticationPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetSchemasOk() (*[]EnumpluggablePassThroughAuthenticationPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetSchemas(v []EnumpluggablePassThroughAuthenticationPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPassThroughAuthenticationHandler

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetPassThroughAuthenticationHandler() string`

GetPassThroughAuthenticationHandler returns the PassThroughAuthenticationHandler field if non-nil, zero value otherwise.

### GetPassThroughAuthenticationHandlerOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetPassThroughAuthenticationHandlerOk() (*string, bool)`

GetPassThroughAuthenticationHandlerOk returns a tuple with the PassThroughAuthenticationHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThroughAuthenticationHandler

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetPassThroughAuthenticationHandler(v string)`

SetPassThroughAuthenticationHandler sets PassThroughAuthenticationHandler field to given value.


### GetIncludedLocalEntryBaseDN

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetTryLocalBind

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetTryLocalBind() bool`

GetTryLocalBind returns the TryLocalBind field if non-nil, zero value otherwise.

### GetTryLocalBindOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetTryLocalBindOk() (*bool, bool)`

GetTryLocalBindOk returns a tuple with the TryLocalBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryLocalBind

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetTryLocalBind(v bool)`

SetTryLocalBind sets TryLocalBind field to given value.

### HasTryLocalBind

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasTryLocalBind() bool`

HasTryLocalBind returns a boolean if a field has been set.

### GetOverrideLocalPassword

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetOverrideLocalPassword() bool`

GetOverrideLocalPassword returns the OverrideLocalPassword field if non-nil, zero value otherwise.

### GetOverrideLocalPasswordOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetOverrideLocalPasswordOk() (*bool, bool)`

GetOverrideLocalPasswordOk returns a tuple with the OverrideLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideLocalPassword

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetOverrideLocalPassword(v bool)`

SetOverrideLocalPassword sets OverrideLocalPassword field to given value.

### HasOverrideLocalPassword

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasOverrideLocalPassword() bool`

HasOverrideLocalPassword returns a boolean if a field has been set.

### GetUpdateLocalPassword

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetUpdateLocalPassword() bool`

GetUpdateLocalPassword returns the UpdateLocalPassword field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetUpdateLocalPasswordOk() (*bool, bool)`

GetUpdateLocalPasswordOk returns a tuple with the UpdateLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPassword

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetUpdateLocalPassword(v bool)`

SetUpdateLocalPassword sets UpdateLocalPassword field to given value.

### HasUpdateLocalPassword

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasUpdateLocalPassword() bool`

HasUpdateLocalPassword returns a boolean if a field has been set.

### GetUpdateLocalPasswordDN

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetUpdateLocalPasswordDN() string`

GetUpdateLocalPasswordDN returns the UpdateLocalPasswordDN field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordDNOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetUpdateLocalPasswordDNOk() (*string, bool)`

GetUpdateLocalPasswordDNOk returns a tuple with the UpdateLocalPasswordDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPasswordDN

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetUpdateLocalPasswordDN(v string)`

SetUpdateLocalPasswordDN sets UpdateLocalPasswordDN field to given value.

### HasUpdateLocalPasswordDN

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasUpdateLocalPasswordDN() bool`

HasUpdateLocalPasswordDN returns a boolean if a field has been set.

### GetAllowLaxPassThroughAuthenticationPasswords

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetAllowLaxPassThroughAuthenticationPasswords() bool`

GetAllowLaxPassThroughAuthenticationPasswords returns the AllowLaxPassThroughAuthenticationPasswords field if non-nil, zero value otherwise.

### GetAllowLaxPassThroughAuthenticationPasswordsOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetAllowLaxPassThroughAuthenticationPasswordsOk() (*bool, bool)`

GetAllowLaxPassThroughAuthenticationPasswordsOk returns a tuple with the AllowLaxPassThroughAuthenticationPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLaxPassThroughAuthenticationPasswords

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetAllowLaxPassThroughAuthenticationPasswords(v bool)`

SetAllowLaxPassThroughAuthenticationPasswords sets AllowLaxPassThroughAuthenticationPasswords field to given value.

### HasAllowLaxPassThroughAuthenticationPasswords

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasAllowLaxPassThroughAuthenticationPasswords() bool`

HasAllowLaxPassThroughAuthenticationPasswords returns a boolean if a field has been set.

### GetIgnoredPasswordPolicyStateErrorCondition

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetIgnoredPasswordPolicyStateErrorCondition() []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp`

GetIgnoredPasswordPolicyStateErrorCondition returns the IgnoredPasswordPolicyStateErrorCondition field if non-nil, zero value otherwise.

### GetIgnoredPasswordPolicyStateErrorConditionOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetIgnoredPasswordPolicyStateErrorConditionOk() (*[]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp, bool)`

GetIgnoredPasswordPolicyStateErrorConditionOk returns a tuple with the IgnoredPasswordPolicyStateErrorCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredPasswordPolicyStateErrorCondition

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetIgnoredPasswordPolicyStateErrorCondition(v []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp)`

SetIgnoredPasswordPolicyStateErrorCondition sets IgnoredPasswordPolicyStateErrorCondition field to given value.

### HasIgnoredPasswordPolicyStateErrorCondition

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasIgnoredPasswordPolicyStateErrorCondition() bool`

HasIgnoredPasswordPolicyStateErrorCondition returns a boolean if a field has been set.

### GetDescription

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetPluginName

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddPluggablePassThroughAuthenticationPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


