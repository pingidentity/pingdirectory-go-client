# NotificationSubscriptionExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumnotificationSubscriptionExtendedOperationHandlerSchemaUrn**](EnumnotificationSubscriptionExtendedOperationHandlerSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewNotificationSubscriptionExtendedOperationHandlerResponse

`func NewNotificationSubscriptionExtendedOperationHandlerResponse(schemas []EnumnotificationSubscriptionExtendedOperationHandlerSchemaUrn, enabled bool, ) *NotificationSubscriptionExtendedOperationHandlerResponse`

NewNotificationSubscriptionExtendedOperationHandlerResponse instantiates a new NotificationSubscriptionExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSubscriptionExtendedOperationHandlerResponseWithDefaults

`func NewNotificationSubscriptionExtendedOperationHandlerResponseWithDefaults() *NotificationSubscriptionExtendedOperationHandlerResponse`

NewNotificationSubscriptionExtendedOperationHandlerResponseWithDefaults instantiates a new NotificationSubscriptionExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *NotificationSubscriptionExtendedOperationHandlerResponse) GetSchemas() []EnumnotificationSubscriptionExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *NotificationSubscriptionExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumnotificationSubscriptionExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *NotificationSubscriptionExtendedOperationHandlerResponse) SetSchemas(v []EnumnotificationSubscriptionExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *NotificationSubscriptionExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationSubscriptionExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationSubscriptionExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationSubscriptionExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *NotificationSubscriptionExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationSubscriptionExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationSubscriptionExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


