# DeliverPasswordResetTokenExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Extended Operation Handler | 
**Schemas** | [**[]EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn**](EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn.md) |  | 
**PasswordGenerator** | **string** | The password generator that will be used to create the password reset token values to be delivered to the end user. | 
**DefaultTokenDeliveryMechanism** | **[]string** |  | 
**PasswordResetTokenValidityDuration** | **string** | The maximum length of time that a password reset token should be considered valid. | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewDeliverPasswordResetTokenExtendedOperationHandlerResponse

`func NewDeliverPasswordResetTokenExtendedOperationHandlerResponse(id string, schemas []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn, passwordGenerator string, defaultTokenDeliveryMechanism []string, passwordResetTokenValidityDuration string, enabled bool, ) *DeliverPasswordResetTokenExtendedOperationHandlerResponse`

NewDeliverPasswordResetTokenExtendedOperationHandlerResponse instantiates a new DeliverPasswordResetTokenExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverPasswordResetTokenExtendedOperationHandlerResponseWithDefaults

`func NewDeliverPasswordResetTokenExtendedOperationHandlerResponseWithDefaults() *DeliverPasswordResetTokenExtendedOperationHandlerResponse`

NewDeliverPasswordResetTokenExtendedOperationHandlerResponseWithDefaults instantiates a new DeliverPasswordResetTokenExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetSchemas() []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) SetSchemas(v []EnumdeliverPasswordResetTokenExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordGenerator

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.


### GetDefaultTokenDeliveryMechanism

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetDefaultTokenDeliveryMechanism() []string`

GetDefaultTokenDeliveryMechanism returns the DefaultTokenDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultTokenDeliveryMechanismOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetDefaultTokenDeliveryMechanismOk() (*[]string, bool)`

GetDefaultTokenDeliveryMechanismOk returns a tuple with the DefaultTokenDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenDeliveryMechanism

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) SetDefaultTokenDeliveryMechanism(v []string)`

SetDefaultTokenDeliveryMechanism sets DefaultTokenDeliveryMechanism field to given value.


### GetPasswordResetTokenValidityDuration

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetPasswordResetTokenValidityDuration() string`

GetPasswordResetTokenValidityDuration returns the PasswordResetTokenValidityDuration field if non-nil, zero value otherwise.

### GetPasswordResetTokenValidityDurationOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetPasswordResetTokenValidityDurationOk() (*string, bool)`

GetPasswordResetTokenValidityDurationOk returns a tuple with the PasswordResetTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetTokenValidityDuration

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) SetPasswordResetTokenValidityDuration(v string)`

SetPasswordResetTokenValidityDuration sets PasswordResetTokenValidityDuration field to given value.


### GetDescription

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeliverPasswordResetTokenExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


