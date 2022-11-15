# ErrorLogAccountStatusNotificationHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumerrorLogAccountStatusNotificationHandlerSchemaUrn**](EnumerrorLogAccountStatusNotificationHandlerSchemaUrn.md) |  | 
**AccountStatusNotificationType** | [**[]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp**](EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp.md) |  | 
**Description** | Pointer to **string** | A description for this Account Status Notification Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification. | [optional] 
**AccountCreationNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which add requests should result in account creation notifications for this handler. | [optional] 
**AccountUpdateNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler. | [optional] 

## Methods

### NewErrorLogAccountStatusNotificationHandlerShared

`func NewErrorLogAccountStatusNotificationHandlerShared(schemas []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn, accountStatusNotificationType []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, enabled bool, ) *ErrorLogAccountStatusNotificationHandlerShared`

NewErrorLogAccountStatusNotificationHandlerShared instantiates a new ErrorLogAccountStatusNotificationHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLogAccountStatusNotificationHandlerSharedWithDefaults

`func NewErrorLogAccountStatusNotificationHandlerSharedWithDefaults() *ErrorLogAccountStatusNotificationHandlerShared`

NewErrorLogAccountStatusNotificationHandlerSharedWithDefaults instantiates a new ErrorLogAccountStatusNotificationHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetSchemas() []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetSchemasOk() (*[]EnumerrorLogAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ErrorLogAccountStatusNotificationHandlerShared) SetSchemas(v []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccountStatusNotificationType

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetAccountStatusNotificationType() []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp`

GetAccountStatusNotificationType returns the AccountStatusNotificationType field if non-nil, zero value otherwise.

### GetAccountStatusNotificationTypeOk

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetAccountStatusNotificationTypeOk() (*[]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, bool)`

GetAccountStatusNotificationTypeOk returns a tuple with the AccountStatusNotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatusNotificationType

`func (o *ErrorLogAccountStatusNotificationHandlerShared) SetAccountStatusNotificationType(v []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp)`

SetAccountStatusNotificationType sets AccountStatusNotificationType field to given value.


### GetDescription

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorLogAccountStatusNotificationHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorLogAccountStatusNotificationHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ErrorLogAccountStatusNotificationHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *ErrorLogAccountStatusNotificationHandlerShared) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *ErrorLogAccountStatusNotificationHandlerShared) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *ErrorLogAccountStatusNotificationHandlerShared) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *ErrorLogAccountStatusNotificationHandlerShared) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *ErrorLogAccountStatusNotificationHandlerShared) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *ErrorLogAccountStatusNotificationHandlerShared) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *ErrorLogAccountStatusNotificationHandlerShared) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


