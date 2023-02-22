# AddLoggingChangeSubscriptionHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Change Subscription Handler | 
**Schemas** | [**[]EnumloggingChangeSubscriptionHandlerSchemaUrn**](EnumloggingChangeSubscriptionHandlerSchemaUrn.md) |  | 
**LogFile** | **string** | Specifies the log file in which the change notification messages will be written. | 
**Description** | Pointer to **string** | A description for this Change Subscription Handler | [optional] 
**Enabled** | **bool** | Indicates whether this change subscription handler is enabled within the server. | 
**ChangeSubscription** | Pointer to **[]string** | The set of change subscriptions for which this change subscription handler should be notified. If no values are provided then it will be notified for all change subscriptions defined in the server. | [optional] 

## Methods

### NewAddLoggingChangeSubscriptionHandlerRequest

`func NewAddLoggingChangeSubscriptionHandlerRequest(handlerName string, schemas []EnumloggingChangeSubscriptionHandlerSchemaUrn, logFile string, enabled bool, ) *AddLoggingChangeSubscriptionHandlerRequest`

NewAddLoggingChangeSubscriptionHandlerRequest instantiates a new AddLoggingChangeSubscriptionHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLoggingChangeSubscriptionHandlerRequestWithDefaults

`func NewAddLoggingChangeSubscriptionHandlerRequestWithDefaults() *AddLoggingChangeSubscriptionHandlerRequest`

NewAddLoggingChangeSubscriptionHandlerRequestWithDefaults instantiates a new AddLoggingChangeSubscriptionHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddLoggingChangeSubscriptionHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetSchemas() []EnumloggingChangeSubscriptionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetSchemasOk() (*[]EnumloggingChangeSubscriptionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLoggingChangeSubscriptionHandlerRequest) SetSchemas(v []EnumloggingChangeSubscriptionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLogFile

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *AddLoggingChangeSubscriptionHandlerRequest) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.


### GetDescription

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLoggingChangeSubscriptionHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLoggingChangeSubscriptionHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddLoggingChangeSubscriptionHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetChangeSubscription

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetChangeSubscription() []string`

GetChangeSubscription returns the ChangeSubscription field if non-nil, zero value otherwise.

### GetChangeSubscriptionOk

`func (o *AddLoggingChangeSubscriptionHandlerRequest) GetChangeSubscriptionOk() (*[]string, bool)`

GetChangeSubscriptionOk returns a tuple with the ChangeSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSubscription

`func (o *AddLoggingChangeSubscriptionHandlerRequest) SetChangeSubscription(v []string)`

SetChangeSubscription sets ChangeSubscription field to given value.

### HasChangeSubscription

`func (o *AddLoggingChangeSubscriptionHandlerRequest) HasChangeSubscription() bool`

HasChangeSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


