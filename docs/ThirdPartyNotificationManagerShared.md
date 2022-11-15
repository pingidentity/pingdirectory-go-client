# ThirdPartyNotificationManagerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumthirdPartyNotificationManagerSchemaUrn**](EnumthirdPartyNotificationManagerSchemaUrn.md) |  | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Notification Manager. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Notification Manager | [optional] 
**Enabled** | **bool** | Indicates whether this Notification Manager is enabled within the server. | 
**SubscriptionBaseDN** | **string** | Specifies the DN of the entry below which subscription data is stored for this Notification Manager. This needs to be in the backend that has the data to be notified on, and must not be the same entry as the backend base DN. The subscription base DN entry does not need to exist as it will be created by the server. | 
**TransactionNotification** | [**EnumnotificationManagerTransactionNotificationProp**](EnumnotificationManagerTransactionNotificationProp.md) |  | 
**MonitorEntriesEnabled** | Pointer to **bool** | Enables monitor entries for this Notification Manager. | [optional] 

## Methods

### NewThirdPartyNotificationManagerShared

`func NewThirdPartyNotificationManagerShared(extensionClass string, enabled bool, subscriptionBaseDN string, transactionNotification EnumnotificationManagerTransactionNotificationProp, ) *ThirdPartyNotificationManagerShared`

NewThirdPartyNotificationManagerShared instantiates a new ThirdPartyNotificationManagerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyNotificationManagerSharedWithDefaults

`func NewThirdPartyNotificationManagerSharedWithDefaults() *ThirdPartyNotificationManagerShared`

NewThirdPartyNotificationManagerSharedWithDefaults instantiates a new ThirdPartyNotificationManagerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ThirdPartyNotificationManagerShared) GetSchemas() []EnumthirdPartyNotificationManagerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyNotificationManagerShared) GetSchemasOk() (*[]EnumthirdPartyNotificationManagerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyNotificationManagerShared) SetSchemas(v []EnumthirdPartyNotificationManagerSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ThirdPartyNotificationManagerShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetExtensionClass

`func (o *ThirdPartyNotificationManagerShared) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyNotificationManagerShared) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyNotificationManagerShared) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyNotificationManagerShared) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyNotificationManagerShared) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyNotificationManagerShared) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyNotificationManagerShared) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyNotificationManagerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyNotificationManagerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyNotificationManagerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyNotificationManagerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyNotificationManagerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyNotificationManagerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyNotificationManagerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSubscriptionBaseDN

`func (o *ThirdPartyNotificationManagerShared) GetSubscriptionBaseDN() string`

GetSubscriptionBaseDN returns the SubscriptionBaseDN field if non-nil, zero value otherwise.

### GetSubscriptionBaseDNOk

`func (o *ThirdPartyNotificationManagerShared) GetSubscriptionBaseDNOk() (*string, bool)`

GetSubscriptionBaseDNOk returns a tuple with the SubscriptionBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionBaseDN

`func (o *ThirdPartyNotificationManagerShared) SetSubscriptionBaseDN(v string)`

SetSubscriptionBaseDN sets SubscriptionBaseDN field to given value.


### GetTransactionNotification

`func (o *ThirdPartyNotificationManagerShared) GetTransactionNotification() EnumnotificationManagerTransactionNotificationProp`

GetTransactionNotification returns the TransactionNotification field if non-nil, zero value otherwise.

### GetTransactionNotificationOk

`func (o *ThirdPartyNotificationManagerShared) GetTransactionNotificationOk() (*EnumnotificationManagerTransactionNotificationProp, bool)`

GetTransactionNotificationOk returns a tuple with the TransactionNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionNotification

`func (o *ThirdPartyNotificationManagerShared) SetTransactionNotification(v EnumnotificationManagerTransactionNotificationProp)`

SetTransactionNotification sets TransactionNotification field to given value.


### GetMonitorEntriesEnabled

`func (o *ThirdPartyNotificationManagerShared) GetMonitorEntriesEnabled() bool`

GetMonitorEntriesEnabled returns the MonitorEntriesEnabled field if non-nil, zero value otherwise.

### GetMonitorEntriesEnabledOk

`func (o *ThirdPartyNotificationManagerShared) GetMonitorEntriesEnabledOk() (*bool, bool)`

GetMonitorEntriesEnabledOk returns a tuple with the MonitorEntriesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorEntriesEnabled

`func (o *ThirdPartyNotificationManagerShared) SetMonitorEntriesEnabled(v bool)`

SetMonitorEntriesEnabled sets MonitorEntriesEnabled field to given value.

### HasMonitorEntriesEnabled

`func (o *ThirdPartyNotificationManagerShared) HasMonitorEntriesEnabled() bool`

HasMonitorEntriesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


