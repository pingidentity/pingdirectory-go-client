# AddThirdPartyNotificationManagerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagerName** | **string** | Name of the new Notification Manager | 
**Schemas** | [**[]EnumthirdPartyNotificationManagerSchemaUrn**](EnumthirdPartyNotificationManagerSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Notification Manager. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Notification Manager. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Notification Manager | [optional] 
**Enabled** | **bool** | Indicates whether this Notification Manager is enabled within the server. | 
**SubscriptionBaseDN** | **string** | Specifies the DN of the entry below which subscription data is stored for this Notification Manager. This needs to be in the backend that has the data to be notified on, and must not be the same entry as the backend base DN. The subscription base DN entry does not need to exist as it will be created by the server. | 
**TransactionNotification** | Pointer to [**EnumnotificationManagerTransactionNotificationProp**](EnumnotificationManagerTransactionNotificationProp.md) |  | [optional] 
**MonitorEntriesEnabled** | Pointer to **bool** | Enables monitor entries for this Notification Manager. | [optional] 

## Methods

### NewAddThirdPartyNotificationManagerRequest

`func NewAddThirdPartyNotificationManagerRequest(managerName string, schemas []EnumthirdPartyNotificationManagerSchemaUrn, extensionClass string, enabled bool, subscriptionBaseDN string, ) *AddThirdPartyNotificationManagerRequest`

NewAddThirdPartyNotificationManagerRequest instantiates a new AddThirdPartyNotificationManagerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyNotificationManagerRequestWithDefaults

`func NewAddThirdPartyNotificationManagerRequestWithDefaults() *AddThirdPartyNotificationManagerRequest`

NewAddThirdPartyNotificationManagerRequestWithDefaults instantiates a new AddThirdPartyNotificationManagerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagerName

`func (o *AddThirdPartyNotificationManagerRequest) GetManagerName() string`

GetManagerName returns the ManagerName field if non-nil, zero value otherwise.

### GetManagerNameOk

`func (o *AddThirdPartyNotificationManagerRequest) GetManagerNameOk() (*string, bool)`

GetManagerNameOk returns a tuple with the ManagerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerName

`func (o *AddThirdPartyNotificationManagerRequest) SetManagerName(v string)`

SetManagerName sets ManagerName field to given value.


### GetSchemas

`func (o *AddThirdPartyNotificationManagerRequest) GetSchemas() []EnumthirdPartyNotificationManagerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyNotificationManagerRequest) GetSchemasOk() (*[]EnumthirdPartyNotificationManagerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyNotificationManagerRequest) SetSchemas(v []EnumthirdPartyNotificationManagerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyNotificationManagerRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyNotificationManagerRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyNotificationManagerRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyNotificationManagerRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyNotificationManagerRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyNotificationManagerRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyNotificationManagerRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyNotificationManagerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyNotificationManagerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyNotificationManagerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyNotificationManagerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyNotificationManagerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyNotificationManagerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyNotificationManagerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSubscriptionBaseDN

`func (o *AddThirdPartyNotificationManagerRequest) GetSubscriptionBaseDN() string`

GetSubscriptionBaseDN returns the SubscriptionBaseDN field if non-nil, zero value otherwise.

### GetSubscriptionBaseDNOk

`func (o *AddThirdPartyNotificationManagerRequest) GetSubscriptionBaseDNOk() (*string, bool)`

GetSubscriptionBaseDNOk returns a tuple with the SubscriptionBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionBaseDN

`func (o *AddThirdPartyNotificationManagerRequest) SetSubscriptionBaseDN(v string)`

SetSubscriptionBaseDN sets SubscriptionBaseDN field to given value.


### GetTransactionNotification

`func (o *AddThirdPartyNotificationManagerRequest) GetTransactionNotification() EnumnotificationManagerTransactionNotificationProp`

GetTransactionNotification returns the TransactionNotification field if non-nil, zero value otherwise.

### GetTransactionNotificationOk

`func (o *AddThirdPartyNotificationManagerRequest) GetTransactionNotificationOk() (*EnumnotificationManagerTransactionNotificationProp, bool)`

GetTransactionNotificationOk returns a tuple with the TransactionNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionNotification

`func (o *AddThirdPartyNotificationManagerRequest) SetTransactionNotification(v EnumnotificationManagerTransactionNotificationProp)`

SetTransactionNotification sets TransactionNotification field to given value.

### HasTransactionNotification

`func (o *AddThirdPartyNotificationManagerRequest) HasTransactionNotification() bool`

HasTransactionNotification returns a boolean if a field has been set.

### GetMonitorEntriesEnabled

`func (o *AddThirdPartyNotificationManagerRequest) GetMonitorEntriesEnabled() bool`

GetMonitorEntriesEnabled returns the MonitorEntriesEnabled field if non-nil, zero value otherwise.

### GetMonitorEntriesEnabledOk

`func (o *AddThirdPartyNotificationManagerRequest) GetMonitorEntriesEnabledOk() (*bool, bool)`

GetMonitorEntriesEnabledOk returns a tuple with the MonitorEntriesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorEntriesEnabled

`func (o *AddThirdPartyNotificationManagerRequest) SetMonitorEntriesEnabled(v bool)`

SetMonitorEntriesEnabled sets MonitorEntriesEnabled field to given value.

### HasMonitorEntriesEnabled

`func (o *AddThirdPartyNotificationManagerRequest) HasMonitorEntriesEnabled() bool`

HasMonitorEntriesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


