# ValidateTotpPasswordExtendedOperationHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn**](EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn.md) |  | 
**SharedSecretAttributeType** | Pointer to **string** | The name or OID of the attribute that will be used to hold the shared secret key used during TOTP processing. | [optional] 
**TimeIntervalDuration** | Pointer to **string** | The duration of the time interval used for TOTP processing. | [optional] 
**AdjacentIntervalsToCheck** | Pointer to **int32** | The number of adjacent time intervals (both before and after the current time) that should be checked when performing authentication. | [optional] 
**PreventTOTPReuse** | Pointer to **bool** | Indicates whether to prevent clients from re-using TOTP passwords. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewValidateTotpPasswordExtendedOperationHandlerShared

`func NewValidateTotpPasswordExtendedOperationHandlerShared(schemas []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn, enabled bool, ) *ValidateTotpPasswordExtendedOperationHandlerShared`

NewValidateTotpPasswordExtendedOperationHandlerShared instantiates a new ValidateTotpPasswordExtendedOperationHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateTotpPasswordExtendedOperationHandlerSharedWithDefaults

`func NewValidateTotpPasswordExtendedOperationHandlerSharedWithDefaults() *ValidateTotpPasswordExtendedOperationHandlerShared`

NewValidateTotpPasswordExtendedOperationHandlerSharedWithDefaults instantiates a new ValidateTotpPasswordExtendedOperationHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetSchemas() []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetSchemasOk() (*[]EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) SetSchemas(v []EnumvalidateTotpPasswordExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSharedSecretAttributeType

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetSharedSecretAttributeType() string`

GetSharedSecretAttributeType returns the SharedSecretAttributeType field if non-nil, zero value otherwise.

### GetSharedSecretAttributeTypeOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetSharedSecretAttributeTypeOk() (*string, bool)`

GetSharedSecretAttributeTypeOk returns a tuple with the SharedSecretAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretAttributeType

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) SetSharedSecretAttributeType(v string)`

SetSharedSecretAttributeType sets SharedSecretAttributeType field to given value.

### HasSharedSecretAttributeType

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) HasSharedSecretAttributeType() bool`

HasSharedSecretAttributeType returns a boolean if a field has been set.

### GetTimeIntervalDuration

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetTimeIntervalDuration() string`

GetTimeIntervalDuration returns the TimeIntervalDuration field if non-nil, zero value otherwise.

### GetTimeIntervalDurationOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetTimeIntervalDurationOk() (*string, bool)`

GetTimeIntervalDurationOk returns a tuple with the TimeIntervalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntervalDuration

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) SetTimeIntervalDuration(v string)`

SetTimeIntervalDuration sets TimeIntervalDuration field to given value.

### HasTimeIntervalDuration

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) HasTimeIntervalDuration() bool`

HasTimeIntervalDuration returns a boolean if a field has been set.

### GetAdjacentIntervalsToCheck

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetAdjacentIntervalsToCheck() int32`

GetAdjacentIntervalsToCheck returns the AdjacentIntervalsToCheck field if non-nil, zero value otherwise.

### GetAdjacentIntervalsToCheckOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetAdjacentIntervalsToCheckOk() (*int32, bool)`

GetAdjacentIntervalsToCheckOk returns a tuple with the AdjacentIntervalsToCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjacentIntervalsToCheck

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) SetAdjacentIntervalsToCheck(v int32)`

SetAdjacentIntervalsToCheck sets AdjacentIntervalsToCheck field to given value.

### HasAdjacentIntervalsToCheck

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) HasAdjacentIntervalsToCheck() bool`

HasAdjacentIntervalsToCheck returns a boolean if a field has been set.

### GetPreventTOTPReuse

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetPreventTOTPReuse() bool`

GetPreventTOTPReuse returns the PreventTOTPReuse field if non-nil, zero value otherwise.

### GetPreventTOTPReuseOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetPreventTOTPReuseOk() (*bool, bool)`

GetPreventTOTPReuseOk returns a tuple with the PreventTOTPReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventTOTPReuse

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) SetPreventTOTPReuse(v bool)`

SetPreventTOTPReuse sets PreventTOTPReuse field to given value.

### HasPreventTOTPReuse

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) HasPreventTOTPReuse() bool`

HasPreventTOTPReuse returns a boolean if a field has been set.

### GetDescription

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ValidateTotpPasswordExtendedOperationHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


