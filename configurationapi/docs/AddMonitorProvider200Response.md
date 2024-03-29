# AddMonitorProvider200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Monitor Provider | 
**Schemas** | [**[]EnumthirdPartyMonitorProviderSchemaUrn**](EnumthirdPartyMonitorProviderSchemaUrn.md) |  | 
**CheckFrequency** | **string** | The frequency with which this monitor provider should confirm the ability to access the server&#39;s encryption settings database. | 
**ProlongedOutageDuration** | Pointer to **string** | The minimum length of time that an outage should persist before it is considered a prolonged outage. If an outage lasts at least as long as this duration, then the server will take the action indicated by the prolonged-outage-behavior property. | [optional] 
**ProlongedOutageBehavior** | Pointer to [**EnummonitorProviderProlongedOutageBehaviorProp**](EnummonitorProviderProlongedOutageBehaviorProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Monitor Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Monitor Provider is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Monitor Provider. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Monitor Provider. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddMonitorProvider200Response

`func NewAddMonitorProvider200Response(id string, schemas []EnumthirdPartyMonitorProviderSchemaUrn, checkFrequency string, enabled bool, extensionClass string, ) *AddMonitorProvider200Response`

NewAddMonitorProvider200Response instantiates a new AddMonitorProvider200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMonitorProvider200ResponseWithDefaults

`func NewAddMonitorProvider200ResponseWithDefaults() *AddMonitorProvider200Response`

NewAddMonitorProvider200ResponseWithDefaults instantiates a new AddMonitorProvider200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddMonitorProvider200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddMonitorProvider200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddMonitorProvider200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddMonitorProvider200Response) GetSchemas() []EnumthirdPartyMonitorProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddMonitorProvider200Response) GetSchemasOk() (*[]EnumthirdPartyMonitorProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddMonitorProvider200Response) SetSchemas(v []EnumthirdPartyMonitorProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCheckFrequency

`func (o *AddMonitorProvider200Response) GetCheckFrequency() string`

GetCheckFrequency returns the CheckFrequency field if non-nil, zero value otherwise.

### GetCheckFrequencyOk

`func (o *AddMonitorProvider200Response) GetCheckFrequencyOk() (*string, bool)`

GetCheckFrequencyOk returns a tuple with the CheckFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckFrequency

`func (o *AddMonitorProvider200Response) SetCheckFrequency(v string)`

SetCheckFrequency sets CheckFrequency field to given value.


### GetProlongedOutageDuration

`func (o *AddMonitorProvider200Response) GetProlongedOutageDuration() string`

GetProlongedOutageDuration returns the ProlongedOutageDuration field if non-nil, zero value otherwise.

### GetProlongedOutageDurationOk

`func (o *AddMonitorProvider200Response) GetProlongedOutageDurationOk() (*string, bool)`

GetProlongedOutageDurationOk returns a tuple with the ProlongedOutageDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongedOutageDuration

`func (o *AddMonitorProvider200Response) SetProlongedOutageDuration(v string)`

SetProlongedOutageDuration sets ProlongedOutageDuration field to given value.

### HasProlongedOutageDuration

`func (o *AddMonitorProvider200Response) HasProlongedOutageDuration() bool`

HasProlongedOutageDuration returns a boolean if a field has been set.

### GetProlongedOutageBehavior

`func (o *AddMonitorProvider200Response) GetProlongedOutageBehavior() EnummonitorProviderProlongedOutageBehaviorProp`

GetProlongedOutageBehavior returns the ProlongedOutageBehavior field if non-nil, zero value otherwise.

### GetProlongedOutageBehaviorOk

`func (o *AddMonitorProvider200Response) GetProlongedOutageBehaviorOk() (*EnummonitorProviderProlongedOutageBehaviorProp, bool)`

GetProlongedOutageBehaviorOk returns a tuple with the ProlongedOutageBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongedOutageBehavior

`func (o *AddMonitorProvider200Response) SetProlongedOutageBehavior(v EnummonitorProviderProlongedOutageBehaviorProp)`

SetProlongedOutageBehavior sets ProlongedOutageBehavior field to given value.

### HasProlongedOutageBehavior

`func (o *AddMonitorProvider200Response) HasProlongedOutageBehavior() bool`

HasProlongedOutageBehavior returns a boolean if a field has been set.

### GetDescription

`func (o *AddMonitorProvider200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddMonitorProvider200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddMonitorProvider200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddMonitorProvider200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddMonitorProvider200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddMonitorProvider200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddMonitorProvider200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *AddMonitorProvider200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddMonitorProvider200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddMonitorProvider200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddMonitorProvider200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddMonitorProvider200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddMonitorProvider200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddMonitorProvider200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddMonitorProvider200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddMonitorProvider200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddMonitorProvider200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddMonitorProvider200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddMonitorProvider200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddMonitorProvider200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddMonitorProvider200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddMonitorProvider200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


