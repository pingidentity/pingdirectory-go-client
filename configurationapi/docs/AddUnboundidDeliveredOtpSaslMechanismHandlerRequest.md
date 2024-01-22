# AddUnboundidDeliveredOtpSaslMechanismHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn**](EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn.md) |  | 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the user(s) targeted in the authentication and/or authorization identities contained in the bind request. This will only be used for \&quot;u:\&quot;-style identities. | 
**OtpValidityDuration** | Pointer to **string** | The maximum length of time that a one-time password value should be considered valid. | [optional] 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 
**HandlerName** | **string** | Name of the new SASL Mechanism Handler | 

## Methods

### NewAddUnboundidDeliveredOtpSaslMechanismHandlerRequest

`func NewAddUnboundidDeliveredOtpSaslMechanismHandlerRequest(schemas []EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool, handlerName string, ) *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest`

NewAddUnboundidDeliveredOtpSaslMechanismHandlerRequest instantiates a new AddUnboundidDeliveredOtpSaslMechanismHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUnboundidDeliveredOtpSaslMechanismHandlerRequestWithDefaults

`func NewAddUnboundidDeliveredOtpSaslMechanismHandlerRequestWithDefaults() *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest`

NewAddUnboundidDeliveredOtpSaslMechanismHandlerRequestWithDefaults instantiates a new AddUnboundidDeliveredOtpSaslMechanismHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetSchemas() []EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetSchemasOk() (*[]EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) SetSchemas(v []EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIdentityMapper

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetOtpValidityDuration

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetOtpValidityDuration() string`

GetOtpValidityDuration returns the OtpValidityDuration field if non-nil, zero value otherwise.

### GetOtpValidityDurationOk

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetOtpValidityDurationOk() (*string, bool)`

GetOtpValidityDurationOk returns a tuple with the OtpValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpValidityDuration

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) SetOtpValidityDuration(v string)`

SetOtpValidityDuration sets OtpValidityDuration field to given value.

### HasOtpValidityDuration

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) HasOtpValidityDuration() bool`

HasOtpValidityDuration returns a boolean if a field has been set.

### GetDescription

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHandlerName

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddUnboundidDeliveredOtpSaslMechanismHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


