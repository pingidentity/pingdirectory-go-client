# EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn**](EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn.md) |  | 
**CheckFrequency** | **string** | The frequency with which this monitor provider should confirm the ability to access the server&#39;s encryption settings database. | 
**ProlongedOutageDuration** | Pointer to **string** | The minimum length of time that an outage should persist before it is considered a prolonged outage. If an outage lasts at least as long as this duration, then the server will take the action indicated by the prolonged-outage-behavior property. | [optional] 
**ProlongedOutageBehavior** | Pointer to [**EnummonitorProviderProlongedOutageBehaviorProp**](EnummonitorProviderProlongedOutageBehaviorProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Monitor Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Monitor Provider is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Monitor Provider | 

## Methods

### NewEncryptionSettingsDatabaseAccessibilityMonitorProviderResponse

`func NewEncryptionSettingsDatabaseAccessibilityMonitorProviderResponse(schemas []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn, checkFrequency string, enabled bool, id string, ) *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse`

NewEncryptionSettingsDatabaseAccessibilityMonitorProviderResponse instantiates a new EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionSettingsDatabaseAccessibilityMonitorProviderResponseWithDefaults

`func NewEncryptionSettingsDatabaseAccessibilityMonitorProviderResponseWithDefaults() *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse`

NewEncryptionSettingsDatabaseAccessibilityMonitorProviderResponseWithDefaults instantiates a new EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetSchemas() []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetSchemasOk() (*[]EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) SetSchemas(v []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCheckFrequency

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetCheckFrequency() string`

GetCheckFrequency returns the CheckFrequency field if non-nil, zero value otherwise.

### GetCheckFrequencyOk

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetCheckFrequencyOk() (*string, bool)`

GetCheckFrequencyOk returns a tuple with the CheckFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckFrequency

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) SetCheckFrequency(v string)`

SetCheckFrequency sets CheckFrequency field to given value.


### GetProlongedOutageDuration

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetProlongedOutageDuration() string`

GetProlongedOutageDuration returns the ProlongedOutageDuration field if non-nil, zero value otherwise.

### GetProlongedOutageDurationOk

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetProlongedOutageDurationOk() (*string, bool)`

GetProlongedOutageDurationOk returns a tuple with the ProlongedOutageDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongedOutageDuration

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) SetProlongedOutageDuration(v string)`

SetProlongedOutageDuration sets ProlongedOutageDuration field to given value.

### HasProlongedOutageDuration

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) HasProlongedOutageDuration() bool`

HasProlongedOutageDuration returns a boolean if a field has been set.

### GetProlongedOutageBehavior

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetProlongedOutageBehavior() EnummonitorProviderProlongedOutageBehaviorProp`

GetProlongedOutageBehavior returns the ProlongedOutageBehavior field if non-nil, zero value otherwise.

### GetProlongedOutageBehaviorOk

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetProlongedOutageBehaviorOk() (*EnummonitorProviderProlongedOutageBehaviorProp, bool)`

GetProlongedOutageBehaviorOk returns a tuple with the ProlongedOutageBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongedOutageBehavior

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) SetProlongedOutageBehavior(v EnummonitorProviderProlongedOutageBehaviorProp)`

SetProlongedOutageBehavior sets ProlongedOutageBehavior field to given value.

### HasProlongedOutageBehavior

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) HasProlongedOutageBehavior() bool`

HasProlongedOutageBehavior returns a boolean if a field has been set.

### GetDescription

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EncryptionSettingsDatabaseAccessibilityMonitorProviderResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


