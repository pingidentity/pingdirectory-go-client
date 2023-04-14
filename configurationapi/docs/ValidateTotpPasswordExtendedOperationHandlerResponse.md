# ValidateTotpPasswordExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Extended Operation Handler | 
**Schemas** | [**[]EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn**](EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn.md) |  | 
**SharedSecretAttributeType** | Pointer to **string** | The name or OID of the attribute that will be used to hold the shared secret key used during TOTP processing. | [optional] 
**TimeIntervalDuration** | Pointer to **string** | The duration of the time interval used for TOTP processing. | [optional] 
**AdjacentIntervalsToCheck** | Pointer to **int64** | The number of adjacent time intervals (both before and after the current time) that should be checked when performing authentication. | [optional] 
**PreventTOTPReuse** | Pointer to **bool** | Indicates whether to prevent clients from re-using TOTP passwords. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewValidateTotpPasswordExtendedOperationHandlerResponse

`func NewValidateTotpPasswordExtendedOperationHandlerResponse(id string, schemas []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn, enabled bool, ) *ValidateTotpPasswordExtendedOperationHandlerResponse`

NewValidateTotpPasswordExtendedOperationHandlerResponse instantiates a new ValidateTotpPasswordExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateTotpPasswordExtendedOperationHandlerResponseWithDefaults

`func NewValidateTotpPasswordExtendedOperationHandlerResponseWithDefaults() *ValidateTotpPasswordExtendedOperationHandlerResponse`

NewValidateTotpPasswordExtendedOperationHandlerResponseWithDefaults instantiates a new ValidateTotpPasswordExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetSchemas() []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) SetSchemas(v []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSharedSecretAttributeType

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetSharedSecretAttributeType() string`

GetSharedSecretAttributeType returns the SharedSecretAttributeType field if non-nil, zero value otherwise.

### GetSharedSecretAttributeTypeOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetSharedSecretAttributeTypeOk() (*string, bool)`

GetSharedSecretAttributeTypeOk returns a tuple with the SharedSecretAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretAttributeType

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) SetSharedSecretAttributeType(v string)`

SetSharedSecretAttributeType sets SharedSecretAttributeType field to given value.

### HasSharedSecretAttributeType

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) HasSharedSecretAttributeType() bool`

HasSharedSecretAttributeType returns a boolean if a field has been set.

### GetTimeIntervalDuration

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetTimeIntervalDuration() string`

GetTimeIntervalDuration returns the TimeIntervalDuration field if non-nil, zero value otherwise.

### GetTimeIntervalDurationOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetTimeIntervalDurationOk() (*string, bool)`

GetTimeIntervalDurationOk returns a tuple with the TimeIntervalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntervalDuration

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) SetTimeIntervalDuration(v string)`

SetTimeIntervalDuration sets TimeIntervalDuration field to given value.

### HasTimeIntervalDuration

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) HasTimeIntervalDuration() bool`

HasTimeIntervalDuration returns a boolean if a field has been set.

### GetAdjacentIntervalsToCheck

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetAdjacentIntervalsToCheck() int64`

GetAdjacentIntervalsToCheck returns the AdjacentIntervalsToCheck field if non-nil, zero value otherwise.

### GetAdjacentIntervalsToCheckOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetAdjacentIntervalsToCheckOk() (*int64, bool)`

GetAdjacentIntervalsToCheckOk returns a tuple with the AdjacentIntervalsToCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjacentIntervalsToCheck

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) SetAdjacentIntervalsToCheck(v int64)`

SetAdjacentIntervalsToCheck sets AdjacentIntervalsToCheck field to given value.

### HasAdjacentIntervalsToCheck

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) HasAdjacentIntervalsToCheck() bool`

HasAdjacentIntervalsToCheck returns a boolean if a field has been set.

### GetPreventTOTPReuse

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetPreventTOTPReuse() bool`

GetPreventTOTPReuse returns the PreventTOTPReuse field if non-nil, zero value otherwise.

### GetPreventTOTPReuseOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetPreventTOTPReuseOk() (*bool, bool)`

GetPreventTOTPReuseOk returns a tuple with the PreventTOTPReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventTOTPReuse

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) SetPreventTOTPReuse(v bool)`

SetPreventTOTPReuse sets PreventTOTPReuse field to given value.

### HasPreventTOTPReuse

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) HasPreventTOTPReuse() bool`

HasPreventTOTPReuse returns a boolean if a field has been set.

### GetDescription

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ValidateTotpPasswordExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


