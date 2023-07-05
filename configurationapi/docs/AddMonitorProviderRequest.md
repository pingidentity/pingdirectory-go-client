# AddMonitorProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Monitor Provider | 
**Schemas** | [**[]EnumthirdPartyMonitorProviderSchemaUrn**](EnumthirdPartyMonitorProviderSchemaUrn.md) |  | 
**CheckFrequency** | Pointer to **string** | The frequency with which this monitor provider should confirm the ability to access the server&#39;s encryption settings database. | [optional] 
**ProlongedOutageDuration** | Pointer to **string** | The minimum length of time that an outage should persist before it is considered a prolonged outage. If an outage lasts at least as long as this duration, then the server will take the action indicated by the prolonged-outage-behavior property. | [optional] 
**ProlongedOutageBehavior** | Pointer to [**EnummonitorProviderProlongedOutageBehaviorProp**](EnummonitorProviderProlongedOutageBehaviorProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Monitor Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Monitor Provider is enabled for use. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Monitor Provider. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Monitor Provider. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddMonitorProviderRequest

`func NewAddMonitorProviderRequest(providerName string, schemas []EnumthirdPartyMonitorProviderSchemaUrn, enabled bool, extensionClass string, ) *AddMonitorProviderRequest`

NewAddMonitorProviderRequest instantiates a new AddMonitorProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMonitorProviderRequestWithDefaults

`func NewAddMonitorProviderRequestWithDefaults() *AddMonitorProviderRequest`

NewAddMonitorProviderRequestWithDefaults instantiates a new AddMonitorProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddMonitorProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddMonitorProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddMonitorProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddMonitorProviderRequest) GetSchemas() []EnumthirdPartyMonitorProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddMonitorProviderRequest) GetSchemasOk() (*[]EnumthirdPartyMonitorProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddMonitorProviderRequest) SetSchemas(v []EnumthirdPartyMonitorProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCheckFrequency

`func (o *AddMonitorProviderRequest) GetCheckFrequency() string`

GetCheckFrequency returns the CheckFrequency field if non-nil, zero value otherwise.

### GetCheckFrequencyOk

`func (o *AddMonitorProviderRequest) GetCheckFrequencyOk() (*string, bool)`

GetCheckFrequencyOk returns a tuple with the CheckFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckFrequency

`func (o *AddMonitorProviderRequest) SetCheckFrequency(v string)`

SetCheckFrequency sets CheckFrequency field to given value.

### HasCheckFrequency

`func (o *AddMonitorProviderRequest) HasCheckFrequency() bool`

HasCheckFrequency returns a boolean if a field has been set.

### GetProlongedOutageDuration

`func (o *AddMonitorProviderRequest) GetProlongedOutageDuration() string`

GetProlongedOutageDuration returns the ProlongedOutageDuration field if non-nil, zero value otherwise.

### GetProlongedOutageDurationOk

`func (o *AddMonitorProviderRequest) GetProlongedOutageDurationOk() (*string, bool)`

GetProlongedOutageDurationOk returns a tuple with the ProlongedOutageDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongedOutageDuration

`func (o *AddMonitorProviderRequest) SetProlongedOutageDuration(v string)`

SetProlongedOutageDuration sets ProlongedOutageDuration field to given value.

### HasProlongedOutageDuration

`func (o *AddMonitorProviderRequest) HasProlongedOutageDuration() bool`

HasProlongedOutageDuration returns a boolean if a field has been set.

### GetProlongedOutageBehavior

`func (o *AddMonitorProviderRequest) GetProlongedOutageBehavior() EnummonitorProviderProlongedOutageBehaviorProp`

GetProlongedOutageBehavior returns the ProlongedOutageBehavior field if non-nil, zero value otherwise.

### GetProlongedOutageBehaviorOk

`func (o *AddMonitorProviderRequest) GetProlongedOutageBehaviorOk() (*EnummonitorProviderProlongedOutageBehaviorProp, bool)`

GetProlongedOutageBehaviorOk returns a tuple with the ProlongedOutageBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongedOutageBehavior

`func (o *AddMonitorProviderRequest) SetProlongedOutageBehavior(v EnummonitorProviderProlongedOutageBehaviorProp)`

SetProlongedOutageBehavior sets ProlongedOutageBehavior field to given value.

### HasProlongedOutageBehavior

`func (o *AddMonitorProviderRequest) HasProlongedOutageBehavior() bool`

HasProlongedOutageBehavior returns a boolean if a field has been set.

### GetDescription

`func (o *AddMonitorProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddMonitorProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddMonitorProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddMonitorProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddMonitorProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddMonitorProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddMonitorProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExtensionClass

`func (o *AddMonitorProviderRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddMonitorProviderRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddMonitorProviderRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddMonitorProviderRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddMonitorProviderRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddMonitorProviderRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddMonitorProviderRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


