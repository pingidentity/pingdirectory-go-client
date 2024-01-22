# AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn**](EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn.md) |  | 
**CheckFrequency** | Pointer to **string** | The frequency with which this monitor provider should confirm the ability to access the server&#39;s encryption settings database. | [optional] 
**ProlongedOutageDuration** | Pointer to **string** | The minimum length of time that an outage should persist before it is considered a prolonged outage. If an outage lasts at least as long as this duration, then the server will take the action indicated by the prolonged-outage-behavior property. | [optional] 
**ProlongedOutageBehavior** | Pointer to [**EnummonitorProviderProlongedOutageBehaviorProp**](EnummonitorProviderProlongedOutageBehaviorProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Monitor Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Monitor Provider is enabled for use. | 
**ProviderName** | **string** | Name of the new Monitor Provider | 

## Methods

### NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest

`func NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest(schemas []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn, enabled bool, providerName string, ) *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest`

NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest instantiates a new AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequestWithDefaults

`func NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequestWithDefaults() *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest`

NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequestWithDefaults instantiates a new AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetSchemas() []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetSchemasOk() (*[]EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetSchemas(v []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCheckFrequency

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetCheckFrequency() string`

GetCheckFrequency returns the CheckFrequency field if non-nil, zero value otherwise.

### GetCheckFrequencyOk

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetCheckFrequencyOk() (*string, bool)`

GetCheckFrequencyOk returns a tuple with the CheckFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckFrequency

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetCheckFrequency(v string)`

SetCheckFrequency sets CheckFrequency field to given value.

### HasCheckFrequency

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) HasCheckFrequency() bool`

HasCheckFrequency returns a boolean if a field has been set.

### GetProlongedOutageDuration

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProlongedOutageDuration() string`

GetProlongedOutageDuration returns the ProlongedOutageDuration field if non-nil, zero value otherwise.

### GetProlongedOutageDurationOk

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProlongedOutageDurationOk() (*string, bool)`

GetProlongedOutageDurationOk returns a tuple with the ProlongedOutageDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongedOutageDuration

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetProlongedOutageDuration(v string)`

SetProlongedOutageDuration sets ProlongedOutageDuration field to given value.

### HasProlongedOutageDuration

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) HasProlongedOutageDuration() bool`

HasProlongedOutageDuration returns a boolean if a field has been set.

### GetProlongedOutageBehavior

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProlongedOutageBehavior() EnummonitorProviderProlongedOutageBehaviorProp`

GetProlongedOutageBehavior returns the ProlongedOutageBehavior field if non-nil, zero value otherwise.

### GetProlongedOutageBehaviorOk

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProlongedOutageBehaviorOk() (*EnummonitorProviderProlongedOutageBehaviorProp, bool)`

GetProlongedOutageBehaviorOk returns a tuple with the ProlongedOutageBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongedOutageBehavior

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetProlongedOutageBehavior(v EnummonitorProviderProlongedOutageBehaviorProp)`

SetProlongedOutageBehavior sets ProlongedOutageBehavior field to given value.

### HasProlongedOutageBehavior

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) HasProlongedOutageBehavior() bool`

HasProlongedOutageBehavior returns a boolean if a field has been set.

### GetDescription

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProviderName

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


