# SingleUseTokensExtendedOperationHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsingleUseTokensExtendedOperationHandlerSchemaUrn**](EnumsingleUseTokensExtendedOperationHandlerSchemaUrn.md) |  | 
**PasswordGenerator** | **string** | The password generator that will be used to create the single-use token values to be delivered to the end user. | 
**DefaultOTPDeliveryMechanism** | **[]string** |  | 
**DefaultSingleUseTokenValidityDuration** | Pointer to **string** | The default length of time that a single-use token will be considered valid by the server if the client doesn&#39;t specify a duration in the deliver single-use token request. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewSingleUseTokensExtendedOperationHandlerShared

`func NewSingleUseTokensExtendedOperationHandlerShared(schemas []EnumsingleUseTokensExtendedOperationHandlerSchemaUrn, passwordGenerator string, defaultOTPDeliveryMechanism []string, enabled bool, ) *SingleUseTokensExtendedOperationHandlerShared`

NewSingleUseTokensExtendedOperationHandlerShared instantiates a new SingleUseTokensExtendedOperationHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleUseTokensExtendedOperationHandlerSharedWithDefaults

`func NewSingleUseTokensExtendedOperationHandlerSharedWithDefaults() *SingleUseTokensExtendedOperationHandlerShared`

NewSingleUseTokensExtendedOperationHandlerSharedWithDefaults instantiates a new SingleUseTokensExtendedOperationHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetSchemas() []EnumsingleUseTokensExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetSchemasOk() (*[]EnumsingleUseTokensExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SingleUseTokensExtendedOperationHandlerShared) SetSchemas(v []EnumsingleUseTokensExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordGenerator

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *SingleUseTokensExtendedOperationHandlerShared) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.


### GetDefaultOTPDeliveryMechanism

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetDefaultOTPDeliveryMechanism() []string`

GetDefaultOTPDeliveryMechanism returns the DefaultOTPDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultOTPDeliveryMechanismOk

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetDefaultOTPDeliveryMechanismOk() (*[]string, bool)`

GetDefaultOTPDeliveryMechanismOk returns a tuple with the DefaultOTPDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOTPDeliveryMechanism

`func (o *SingleUseTokensExtendedOperationHandlerShared) SetDefaultOTPDeliveryMechanism(v []string)`

SetDefaultOTPDeliveryMechanism sets DefaultOTPDeliveryMechanism field to given value.


### GetDefaultSingleUseTokenValidityDuration

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetDefaultSingleUseTokenValidityDuration() string`

GetDefaultSingleUseTokenValidityDuration returns the DefaultSingleUseTokenValidityDuration field if non-nil, zero value otherwise.

### GetDefaultSingleUseTokenValidityDurationOk

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetDefaultSingleUseTokenValidityDurationOk() (*string, bool)`

GetDefaultSingleUseTokenValidityDurationOk returns a tuple with the DefaultSingleUseTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSingleUseTokenValidityDuration

`func (o *SingleUseTokensExtendedOperationHandlerShared) SetDefaultSingleUseTokenValidityDuration(v string)`

SetDefaultSingleUseTokenValidityDuration sets DefaultSingleUseTokenValidityDuration field to given value.

### HasDefaultSingleUseTokenValidityDuration

`func (o *SingleUseTokensExtendedOperationHandlerShared) HasDefaultSingleUseTokenValidityDuration() bool`

HasDefaultSingleUseTokenValidityDuration returns a boolean if a field has been set.

### GetDescription

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SingleUseTokensExtendedOperationHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SingleUseTokensExtendedOperationHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SingleUseTokensExtendedOperationHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SingleUseTokensExtendedOperationHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


