# AddDeliverPasswordResetTokenExtendedOperationHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Extended Operation Handler | 
**Schemas** | [**[]EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn**](EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn.md) |  | 
**PasswordGenerator** | **string** | The password generator that will be used to create the password reset token values to be delivered to the end user. | 
**DefaultTokenDeliveryMechanism** | **[]string** | The set of delivery mechanisms that may be used to deliver password reset tokens to users for requests that do not specify one or more preferred delivery mechanisms. | 
**PasswordResetTokenValidityDuration** | **string** | The maximum length of time that a password reset token should be considered valid. | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewAddDeliverPasswordResetTokenExtendedOperationHandlerRequest

`func NewAddDeliverPasswordResetTokenExtendedOperationHandlerRequest(handlerName string, schemas []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn, passwordGenerator string, defaultTokenDeliveryMechanism []string, passwordResetTokenValidityDuration string, enabled bool, ) *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest`

NewAddDeliverPasswordResetTokenExtendedOperationHandlerRequest instantiates a new AddDeliverPasswordResetTokenExtendedOperationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDeliverPasswordResetTokenExtendedOperationHandlerRequestWithDefaults

`func NewAddDeliverPasswordResetTokenExtendedOperationHandlerRequestWithDefaults() *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest`

NewAddDeliverPasswordResetTokenExtendedOperationHandlerRequestWithDefaults instantiates a new AddDeliverPasswordResetTokenExtendedOperationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetSchemas() []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetSchemasOk() (*[]EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetSchemas(v []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordGenerator

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.


### GetDefaultTokenDeliveryMechanism

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetDefaultTokenDeliveryMechanism() []string`

GetDefaultTokenDeliveryMechanism returns the DefaultTokenDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultTokenDeliveryMechanismOk

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetDefaultTokenDeliveryMechanismOk() (*[]string, bool)`

GetDefaultTokenDeliveryMechanismOk returns a tuple with the DefaultTokenDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenDeliveryMechanism

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetDefaultTokenDeliveryMechanism(v []string)`

SetDefaultTokenDeliveryMechanism sets DefaultTokenDeliveryMechanism field to given value.


### GetPasswordResetTokenValidityDuration

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetPasswordResetTokenValidityDuration() string`

GetPasswordResetTokenValidityDuration returns the PasswordResetTokenValidityDuration field if non-nil, zero value otherwise.

### GetPasswordResetTokenValidityDurationOk

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetPasswordResetTokenValidityDurationOk() (*string, bool)`

GetPasswordResetTokenValidityDurationOk returns a tuple with the PasswordResetTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetTokenValidityDuration

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetPasswordResetTokenValidityDuration(v string)`

SetPasswordResetTokenValidityDuration sets PasswordResetTokenValidityDuration field to given value.


### GetDescription

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddDeliverPasswordResetTokenExtendedOperationHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


