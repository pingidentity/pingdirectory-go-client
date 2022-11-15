# UnboundidDeliveredOtpSaslMechanismHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn**](EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn.md) |  | 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the user(s) targeted in the authentication and/or authorization identities contained in the bind request. This will only be used for \&quot;u:\&quot;-style identities. | 
**OtpValidityDuration** | **string** | The maximum length of time that a one-time password value should be considered valid. | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewUnboundidDeliveredOtpSaslMechanismHandlerShared

`func NewUnboundidDeliveredOtpSaslMechanismHandlerShared(schemas []EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn, identityMapper string, otpValidityDuration string, enabled bool, ) *UnboundidDeliveredOtpSaslMechanismHandlerShared`

NewUnboundidDeliveredOtpSaslMechanismHandlerShared instantiates a new UnboundidDeliveredOtpSaslMechanismHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnboundidDeliveredOtpSaslMechanismHandlerSharedWithDefaults

`func NewUnboundidDeliveredOtpSaslMechanismHandlerSharedWithDefaults() *UnboundidDeliveredOtpSaslMechanismHandlerShared`

NewUnboundidDeliveredOtpSaslMechanismHandlerSharedWithDefaults instantiates a new UnboundidDeliveredOtpSaslMechanismHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) GetSchemas() []EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) GetSchemasOk() (*[]EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) SetSchemas(v []EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIdentityMapper

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetOtpValidityDuration

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) GetOtpValidityDuration() string`

GetOtpValidityDuration returns the OtpValidityDuration field if non-nil, zero value otherwise.

### GetOtpValidityDurationOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) GetOtpValidityDurationOk() (*string, bool)`

GetOtpValidityDurationOk returns a tuple with the OtpValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpValidityDuration

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) SetOtpValidityDuration(v string)`

SetOtpValidityDuration sets OtpValidityDuration field to given value.


### GetDescription

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


