# ThirdPartyNotificationManagerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Notification Manager | 
**Schemas** | Pointer to [**[]EnumthirdPartyNotificationManagerSchemaUrn**](EnumthirdPartyNotificationManagerSchemaUrn.md) |  | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Notification Manager. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Notification Manager | [optional] 
**Enabled** | **bool** | Indicates whether this Notification Manager is enabled within the server. | 
**SubscriptionBaseDN** | **string** | Specifies the DN of the entry below which subscription data is stored for this Notification Manager. This needs to be in the backend that has the data to be notified on, and must not be the same entry as the backend base DN. The subscription base DN entry does not need to exist as it will be created by the server. | 
**TransactionNotification** | [**EnumnotificationManagerTransactionNotificationProp**](EnumnotificationManagerTransactionNotificationProp.md) |  | 
**MonitorEntriesEnabled** | Pointer to **bool** | Enables monitor entries for this Notification Manager. | [optional] 

## Methods

### NewThirdPartyNotificationManagerResponse

`func NewThirdPartyNotificationManagerResponse(id string, extensionClass string, enabled bool, subscriptionBaseDN string, transactionNotification EnumnotificationManagerTransactionNotificationProp, ) *ThirdPartyNotificationManagerResponse`

NewThirdPartyNotificationManagerResponse instantiates a new ThirdPartyNotificationManagerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyNotificationManagerResponseWithDefaults

`func NewThirdPartyNotificationManagerResponseWithDefaults() *ThirdPartyNotificationManagerResponse`

NewThirdPartyNotificationManagerResponseWithDefaults instantiates a new ThirdPartyNotificationManagerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyNotificationManagerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyNotificationManagerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyNotificationManagerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ThirdPartyNotificationManagerResponse) GetSchemas() []EnumthirdPartyNotificationManagerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyNotificationManagerResponse) GetSchemasOk() (*[]EnumthirdPartyNotificationManagerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyNotificationManagerResponse) SetSchemas(v []EnumthirdPartyNotificationManagerSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ThirdPartyNotificationManagerResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetExtensionClass

`func (o *ThirdPartyNotificationManagerResponse) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyNotificationManagerResponse) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyNotificationManagerResponse) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyNotificationManagerResponse) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyNotificationManagerResponse) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyNotificationManagerResponse) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyNotificationManagerResponse) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyNotificationManagerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyNotificationManagerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyNotificationManagerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyNotificationManagerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyNotificationManagerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyNotificationManagerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyNotificationManagerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSubscriptionBaseDN

`func (o *ThirdPartyNotificationManagerResponse) GetSubscriptionBaseDN() string`

GetSubscriptionBaseDN returns the SubscriptionBaseDN field if non-nil, zero value otherwise.

### GetSubscriptionBaseDNOk

`func (o *ThirdPartyNotificationManagerResponse) GetSubscriptionBaseDNOk() (*string, bool)`

GetSubscriptionBaseDNOk returns a tuple with the SubscriptionBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionBaseDN

`func (o *ThirdPartyNotificationManagerResponse) SetSubscriptionBaseDN(v string)`

SetSubscriptionBaseDN sets SubscriptionBaseDN field to given value.


### GetTransactionNotification

`func (o *ThirdPartyNotificationManagerResponse) GetTransactionNotification() EnumnotificationManagerTransactionNotificationProp`

GetTransactionNotification returns the TransactionNotification field if non-nil, zero value otherwise.

### GetTransactionNotificationOk

`func (o *ThirdPartyNotificationManagerResponse) GetTransactionNotificationOk() (*EnumnotificationManagerTransactionNotificationProp, bool)`

GetTransactionNotificationOk returns a tuple with the TransactionNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionNotification

`func (o *ThirdPartyNotificationManagerResponse) SetTransactionNotification(v EnumnotificationManagerTransactionNotificationProp)`

SetTransactionNotification sets TransactionNotification field to given value.


### GetMonitorEntriesEnabled

`func (o *ThirdPartyNotificationManagerResponse) GetMonitorEntriesEnabled() bool`

GetMonitorEntriesEnabled returns the MonitorEntriesEnabled field if non-nil, zero value otherwise.

### GetMonitorEntriesEnabledOk

`func (o *ThirdPartyNotificationManagerResponse) GetMonitorEntriesEnabledOk() (*bool, bool)`

GetMonitorEntriesEnabledOk returns a tuple with the MonitorEntriesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorEntriesEnabled

`func (o *ThirdPartyNotificationManagerResponse) SetMonitorEntriesEnabled(v bool)`

SetMonitorEntriesEnabled sets MonitorEntriesEnabled field to given value.

### HasMonitorEntriesEnabled

`func (o *ThirdPartyNotificationManagerResponse) HasMonitorEntriesEnabled() bool`

HasMonitorEntriesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


