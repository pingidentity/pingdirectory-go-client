# DeliverPasswordResetTokenExtendedOperationHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn**](EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn.md) |  | 
**PasswordGenerator** | **string** | The password generator that will be used to create the password reset token values to be delivered to the end user. | 
**DefaultTokenDeliveryMechanism** | **[]string** |  | 
**PasswordResetTokenValidityDuration** | **string** | The maximum length of time that a password reset token should be considered valid. | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewDeliverPasswordResetTokenExtendedOperationHandlerShared

`func NewDeliverPasswordResetTokenExtendedOperationHandlerShared(schemas []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn, passwordGenerator string, defaultTokenDeliveryMechanism []string, passwordResetTokenValidityDuration string, enabled bool, ) *DeliverPasswordResetTokenExtendedOperationHandlerShared`

NewDeliverPasswordResetTokenExtendedOperationHandlerShared instantiates a new DeliverPasswordResetTokenExtendedOperationHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverPasswordResetTokenExtendedOperationHandlerSharedWithDefaults

`func NewDeliverPasswordResetTokenExtendedOperationHandlerSharedWithDefaults() *DeliverPasswordResetTokenExtendedOperationHandlerShared`

NewDeliverPasswordResetTokenExtendedOperationHandlerSharedWithDefaults instantiates a new DeliverPasswordResetTokenExtendedOperationHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetSchemas() []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetSchemasOk() (*[]EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) SetSchemas(v []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordGenerator

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.


### GetDefaultTokenDeliveryMechanism

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetDefaultTokenDeliveryMechanism() []string`

GetDefaultTokenDeliveryMechanism returns the DefaultTokenDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultTokenDeliveryMechanismOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetDefaultTokenDeliveryMechanismOk() (*[]string, bool)`

GetDefaultTokenDeliveryMechanismOk returns a tuple with the DefaultTokenDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenDeliveryMechanism

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) SetDefaultTokenDeliveryMechanism(v []string)`

SetDefaultTokenDeliveryMechanism sets DefaultTokenDeliveryMechanism field to given value.


### GetPasswordResetTokenValidityDuration

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetPasswordResetTokenValidityDuration() string`

GetPasswordResetTokenValidityDuration returns the PasswordResetTokenValidityDuration field if non-nil, zero value otherwise.

### GetPasswordResetTokenValidityDurationOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetPasswordResetTokenValidityDurationOk() (*string, bool)`

GetPasswordResetTokenValidityDurationOk returns a tuple with the PasswordResetTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetTokenValidityDuration

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) SetPasswordResetTokenValidityDuration(v string)`

SetPasswordResetTokenValidityDuration sets PasswordResetTokenValidityDuration field to given value.


### GetDescription

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


