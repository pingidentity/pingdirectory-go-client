# AddSingleUseTokensExtendedOperationHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Extended Operation Handler | 
**Schemas** | [**[]EnumsingleUseTokensExtendedOperationHandlerSchemaUrn**](EnumsingleUseTokensExtendedOperationHandlerSchemaUrn.md) |  | 
**PasswordGenerator** | **string** | The password generator that will be used to create the single-use token values to be delivered to the end user. | 
**DefaultOTPDeliveryMechanism** | **[]string** | The set of delivery mechanisms that may be used to deliver single-use tokens to users in requests that do not specify one or more preferred delivery mechanisms. | 
**DefaultSingleUseTokenValidityDuration** | Pointer to **string** | The default length of time that a single-use token will be considered valid by the server if the client doesn&#39;t specify a duration in the deliver single-use token request. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewAddSingleUseTokensExtendedOperationHandlerRequest

`func NewAddSingleUseTokensExtendedOperationHandlerRequest(handlerName string, schemas []EnumsingleUseTokensExtendedOperationHandlerSchemaUrn, passwordGenerator string, defaultOTPDeliveryMechanism []string, enabled bool, ) *AddSingleUseTokensExtendedOperationHandlerRequest`

NewAddSingleUseTokensExtendedOperationHandlerRequest instantiates a new AddSingleUseTokensExtendedOperationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSingleUseTokensExtendedOperationHandlerRequestWithDefaults

`func NewAddSingleUseTokensExtendedOperationHandlerRequestWithDefaults() *AddSingleUseTokensExtendedOperationHandlerRequest`

NewAddSingleUseTokensExtendedOperationHandlerRequestWithDefaults instantiates a new AddSingleUseTokensExtendedOperationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetSchemas() []EnumsingleUseTokensExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetSchemasOk() (*[]EnumsingleUseTokensExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) SetSchemas(v []EnumsingleUseTokensExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordGenerator

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.


### GetDefaultOTPDeliveryMechanism

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetDefaultOTPDeliveryMechanism() []string`

GetDefaultOTPDeliveryMechanism returns the DefaultOTPDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultOTPDeliveryMechanismOk

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetDefaultOTPDeliveryMechanismOk() (*[]string, bool)`

GetDefaultOTPDeliveryMechanismOk returns a tuple with the DefaultOTPDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOTPDeliveryMechanism

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) SetDefaultOTPDeliveryMechanism(v []string)`

SetDefaultOTPDeliveryMechanism sets DefaultOTPDeliveryMechanism field to given value.


### GetDefaultSingleUseTokenValidityDuration

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetDefaultSingleUseTokenValidityDuration() string`

GetDefaultSingleUseTokenValidityDuration returns the DefaultSingleUseTokenValidityDuration field if non-nil, zero value otherwise.

### GetDefaultSingleUseTokenValidityDurationOk

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetDefaultSingleUseTokenValidityDurationOk() (*string, bool)`

GetDefaultSingleUseTokenValidityDurationOk returns a tuple with the DefaultSingleUseTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSingleUseTokenValidityDuration

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) SetDefaultSingleUseTokenValidityDuration(v string)`

SetDefaultSingleUseTokenValidityDuration sets DefaultSingleUseTokenValidityDuration field to given value.

### HasDefaultSingleUseTokenValidityDuration

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) HasDefaultSingleUseTokenValidityDuration() bool`

HasDefaultSingleUseTokenValidityDuration returns a boolean if a field has been set.

### GetDescription

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSingleUseTokensExtendedOperationHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


