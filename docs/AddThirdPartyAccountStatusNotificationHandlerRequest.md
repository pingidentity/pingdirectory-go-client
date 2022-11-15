# AddThirdPartyAccountStatusNotificationHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Account Status Notification Handler | 
**Schemas** | [**[]EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn**](EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Account Status Notification Handler. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Account Status Notification Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification. | [optional] 
**AccountCreationNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which add requests should result in account creation notifications for this handler. | [optional] 
**AccountUpdateNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler. | [optional] 

## Methods

### NewAddThirdPartyAccountStatusNotificationHandlerRequest

`func NewAddThirdPartyAccountStatusNotificationHandlerRequest(handlerName string, schemas []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn, extensionClass string, enabled bool, ) *AddThirdPartyAccountStatusNotificationHandlerRequest`

NewAddThirdPartyAccountStatusNotificationHandlerRequest instantiates a new AddThirdPartyAccountStatusNotificationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyAccountStatusNotificationHandlerRequestWithDefaults

`func NewAddThirdPartyAccountStatusNotificationHandlerRequestWithDefaults() *AddThirdPartyAccountStatusNotificationHandlerRequest`

NewAddThirdPartyAccountStatusNotificationHandlerRequestWithDefaults instantiates a new AddThirdPartyAccountStatusNotificationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetSchemas() []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetSchemasOk() (*[]EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetSchemas(v []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *AddThirdPartyAccountStatusNotificationHandlerRequest) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


