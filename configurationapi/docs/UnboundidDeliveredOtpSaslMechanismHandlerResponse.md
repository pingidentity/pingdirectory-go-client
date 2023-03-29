# UnboundidDeliveredOtpSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the SASL Mechanism Handler | 
**Schemas** | [**[]EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn**](EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn.md) |  | 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the user(s) targeted in the authentication and/or authorization identities contained in the bind request. This will only be used for \&quot;u:\&quot;-style identities. | 
**OtpValidityDuration** | **string** | The maximum length of time that a one-time password value should be considered valid. | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewUnboundidDeliveredOtpSaslMechanismHandlerResponse

`func NewUnboundidDeliveredOtpSaslMechanismHandlerResponse(id string, schemas []EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn, identityMapper string, otpValidityDuration string, enabled bool, ) *UnboundidDeliveredOtpSaslMechanismHandlerResponse`

NewUnboundidDeliveredOtpSaslMechanismHandlerResponse instantiates a new UnboundidDeliveredOtpSaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnboundidDeliveredOtpSaslMechanismHandlerResponseWithDefaults

`func NewUnboundidDeliveredOtpSaslMechanismHandlerResponseWithDefaults() *UnboundidDeliveredOtpSaslMechanismHandlerResponse`

NewUnboundidDeliveredOtpSaslMechanismHandlerResponseWithDefaults instantiates a new UnboundidDeliveredOtpSaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetSchemas() []EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) SetSchemas(v []EnumunboundidDeliveredOtpSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIdentityMapper

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetOtpValidityDuration

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetOtpValidityDuration() string`

GetOtpValidityDuration returns the OtpValidityDuration field if non-nil, zero value otherwise.

### GetOtpValidityDurationOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetOtpValidityDurationOk() (*string, bool)`

GetOtpValidityDurationOk returns a tuple with the OtpValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpValidityDuration

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) SetOtpValidityDuration(v string)`

SetOtpValidityDuration sets OtpValidityDuration field to given value.


### GetDescription

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidDeliveredOtpSaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


