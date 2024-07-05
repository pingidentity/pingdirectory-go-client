# ThirdPartyPolicyDecisionLogPublisherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumthirdPartyPolicyDecisionLogPublisherSchemaUrn**](EnumthirdPartyPolicyDecisionLogPublisherSchemaUrn.md) |  | 
**Id** | **string** | Name of the Log Publisher | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Policy Decision Log Publisher. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Policy Decision Log Publisher. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**IncludePDPResponse** | Pointer to **bool** | Indicates whether policy decision messages recorded by this log publisher will include the full response returned by the PDP. | [optional] 
**PolicyMessageType** | Pointer to [**[]EnumlogPublisherPolicyMessageTypeProp**](EnumlogPublisherPolicyMessageTypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewThirdPartyPolicyDecisionLogPublisherResponse

`func NewThirdPartyPolicyDecisionLogPublisherResponse(schemas []EnumthirdPartyPolicyDecisionLogPublisherSchemaUrn, id string, extensionClass string, enabled bool, ) *ThirdPartyPolicyDecisionLogPublisherResponse`

NewThirdPartyPolicyDecisionLogPublisherResponse instantiates a new ThirdPartyPolicyDecisionLogPublisherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyPolicyDecisionLogPublisherResponseWithDefaults

`func NewThirdPartyPolicyDecisionLogPublisherResponseWithDefaults() *ThirdPartyPolicyDecisionLogPublisherResponse`

NewThirdPartyPolicyDecisionLogPublisherResponseWithDefaults instantiates a new ThirdPartyPolicyDecisionLogPublisherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetSchemas() []EnumthirdPartyPolicyDecisionLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetSchemasOk() (*[]EnumthirdPartyPolicyDecisionLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetSchemas(v []EnumthirdPartyPolicyDecisionLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetId(v string)`

SetId sets Id field to given value.


### GetExtensionClass

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetIncludePDPResponse

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetIncludePDPResponse() bool`

GetIncludePDPResponse returns the IncludePDPResponse field if non-nil, zero value otherwise.

### GetIncludePDPResponseOk

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetIncludePDPResponseOk() (*bool, bool)`

GetIncludePDPResponseOk returns a tuple with the IncludePDPResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePDPResponse

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetIncludePDPResponse(v bool)`

SetIncludePDPResponse sets IncludePDPResponse field to given value.

### HasIncludePDPResponse

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) HasIncludePDPResponse() bool`

HasIncludePDPResponse returns a boolean if a field has been set.

### GetPolicyMessageType

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetPolicyMessageType() []EnumlogPublisherPolicyMessageTypeProp`

GetPolicyMessageType returns the PolicyMessageType field if non-nil, zero value otherwise.

### GetPolicyMessageTypeOk

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetPolicyMessageTypeOk() (*[]EnumlogPublisherPolicyMessageTypeProp, bool)`

GetPolicyMessageTypeOk returns a tuple with the PolicyMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMessageType

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetPolicyMessageType(v []EnumlogPublisherPolicyMessageTypeProp)`

SetPolicyMessageType sets PolicyMessageType field to given value.

### HasPolicyMessageType

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) HasPolicyMessageType() bool`

HasPolicyMessageType returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *ThirdPartyPolicyDecisionLogPublisherResponse) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


