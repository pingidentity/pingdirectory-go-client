# UnboundidTotpSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumunboundidTotpSaslMechanismHandlerSchemaUrn**](EnumunboundidTotpSaslMechanismHandlerSchemaUrn.md) |  | 
**Id** | Pointer to **string** | Name of the SASL Mechanism Handler | [optional] 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the user(s) targeted in the authentication and/or authorization identities contained in the bind request. This will only be used for \&quot;u:\&quot;-style identities. | 
**SharedSecretAttributeType** | Pointer to **string** | The name or OID of the attribute that will be used to hold the shared secret key used during TOTP processing. | [optional] 
**TimeIntervalDuration** | Pointer to **string** | The duration of the time interval used for TOTP processing. | [optional] 
**AdjacentIntervalsToCheck** | Pointer to **int32** | The number of adjacent time intervals (both before and after the current time) that should be checked when performing authentication. | [optional] 
**RequireStaticPassword** | Pointer to **bool** | Indicates whether to require a static password (as might be held in the userPassword attribute, or whatever password attribute is defined in the password policy governing the user) in addition to the one-time password. | [optional] 
**PreventTOTPReuse** | Pointer to **bool** | Indicates whether to prevent clients from re-using TOTP passwords. | [optional] 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewUnboundidTotpSaslMechanismHandlerResponse

`func NewUnboundidTotpSaslMechanismHandlerResponse(schemas []EnumunboundidTotpSaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool, ) *UnboundidTotpSaslMechanismHandlerResponse`

NewUnboundidTotpSaslMechanismHandlerResponse instantiates a new UnboundidTotpSaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnboundidTotpSaslMechanismHandlerResponseWithDefaults

`func NewUnboundidTotpSaslMechanismHandlerResponseWithDefaults() *UnboundidTotpSaslMechanismHandlerResponse`

NewUnboundidTotpSaslMechanismHandlerResponseWithDefaults instantiates a new UnboundidTotpSaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UnboundidTotpSaslMechanismHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidTotpSaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetSchemas() []EnumunboundidTotpSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumunboundidTotpSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetSchemas(v []EnumunboundidTotpSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UnboundidTotpSaslMechanismHandlerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetSharedSecretAttributeType

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetSharedSecretAttributeType() string`

GetSharedSecretAttributeType returns the SharedSecretAttributeType field if non-nil, zero value otherwise.

### GetSharedSecretAttributeTypeOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetSharedSecretAttributeTypeOk() (*string, bool)`

GetSharedSecretAttributeTypeOk returns a tuple with the SharedSecretAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretAttributeType

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetSharedSecretAttributeType(v string)`

SetSharedSecretAttributeType sets SharedSecretAttributeType field to given value.

### HasSharedSecretAttributeType

`func (o *UnboundidTotpSaslMechanismHandlerResponse) HasSharedSecretAttributeType() bool`

HasSharedSecretAttributeType returns a boolean if a field has been set.

### GetTimeIntervalDuration

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetTimeIntervalDuration() string`

GetTimeIntervalDuration returns the TimeIntervalDuration field if non-nil, zero value otherwise.

### GetTimeIntervalDurationOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetTimeIntervalDurationOk() (*string, bool)`

GetTimeIntervalDurationOk returns a tuple with the TimeIntervalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntervalDuration

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetTimeIntervalDuration(v string)`

SetTimeIntervalDuration sets TimeIntervalDuration field to given value.

### HasTimeIntervalDuration

`func (o *UnboundidTotpSaslMechanismHandlerResponse) HasTimeIntervalDuration() bool`

HasTimeIntervalDuration returns a boolean if a field has been set.

### GetAdjacentIntervalsToCheck

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetAdjacentIntervalsToCheck() int32`

GetAdjacentIntervalsToCheck returns the AdjacentIntervalsToCheck field if non-nil, zero value otherwise.

### GetAdjacentIntervalsToCheckOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetAdjacentIntervalsToCheckOk() (*int32, bool)`

GetAdjacentIntervalsToCheckOk returns a tuple with the AdjacentIntervalsToCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjacentIntervalsToCheck

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetAdjacentIntervalsToCheck(v int32)`

SetAdjacentIntervalsToCheck sets AdjacentIntervalsToCheck field to given value.

### HasAdjacentIntervalsToCheck

`func (o *UnboundidTotpSaslMechanismHandlerResponse) HasAdjacentIntervalsToCheck() bool`

HasAdjacentIntervalsToCheck returns a boolean if a field has been set.

### GetRequireStaticPassword

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetRequireStaticPassword() bool`

GetRequireStaticPassword returns the RequireStaticPassword field if non-nil, zero value otherwise.

### GetRequireStaticPasswordOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetRequireStaticPasswordOk() (*bool, bool)`

GetRequireStaticPasswordOk returns a tuple with the RequireStaticPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireStaticPassword

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetRequireStaticPassword(v bool)`

SetRequireStaticPassword sets RequireStaticPassword field to given value.

### HasRequireStaticPassword

`func (o *UnboundidTotpSaslMechanismHandlerResponse) HasRequireStaticPassword() bool`

HasRequireStaticPassword returns a boolean if a field has been set.

### GetPreventTOTPReuse

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetPreventTOTPReuse() bool`

GetPreventTOTPReuse returns the PreventTOTPReuse field if non-nil, zero value otherwise.

### GetPreventTOTPReuseOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetPreventTOTPReuseOk() (*bool, bool)`

GetPreventTOTPReuseOk returns a tuple with the PreventTOTPReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventTOTPReuse

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetPreventTOTPReuse(v bool)`

SetPreventTOTPReuse sets PreventTOTPReuse field to given value.

### HasPreventTOTPReuse

`func (o *UnboundidTotpSaslMechanismHandlerResponse) HasPreventTOTPReuse() bool`

HasPreventTOTPReuse returns a boolean if a field has been set.

### GetDescription

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnboundidTotpSaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UnboundidTotpSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UnboundidTotpSaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


