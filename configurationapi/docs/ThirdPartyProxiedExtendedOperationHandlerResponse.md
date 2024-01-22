# ThirdPartyProxiedExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumthirdPartyProxiedExtendedOperationHandlerSchemaUrn**](EnumthirdPartyProxiedExtendedOperationHandlerSchemaUrn.md) |  | 
**Id** | **string** | Name of the Extended Operation Handler | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Proxied Extended Operation Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Proxied Extended Operation Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**RouteToBackendSetBehavior** | Pointer to [**EnumextendedOperationHandlerRouteToBackendSetBehaviorProp**](EnumextendedOperationHandlerRouteToBackendSetBehaviorProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewThirdPartyProxiedExtendedOperationHandlerResponse

`func NewThirdPartyProxiedExtendedOperationHandlerResponse(schemas []EnumthirdPartyProxiedExtendedOperationHandlerSchemaUrn, id string, extensionClass string, enabled bool, ) *ThirdPartyProxiedExtendedOperationHandlerResponse`

NewThirdPartyProxiedExtendedOperationHandlerResponse instantiates a new ThirdPartyProxiedExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyProxiedExtendedOperationHandlerResponseWithDefaults

`func NewThirdPartyProxiedExtendedOperationHandlerResponseWithDefaults() *ThirdPartyProxiedExtendedOperationHandlerResponse`

NewThirdPartyProxiedExtendedOperationHandlerResponseWithDefaults instantiates a new ThirdPartyProxiedExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetSchemas() []EnumthirdPartyProxiedExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumthirdPartyProxiedExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) SetSchemas(v []EnumthirdPartyProxiedExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetExtensionClass

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetRouteToBackendSetBehavior

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetRouteToBackendSetBehavior() EnumextendedOperationHandlerRouteToBackendSetBehaviorProp`

GetRouteToBackendSetBehavior returns the RouteToBackendSetBehavior field if non-nil, zero value otherwise.

### GetRouteToBackendSetBehaviorOk

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetRouteToBackendSetBehaviorOk() (*EnumextendedOperationHandlerRouteToBackendSetBehaviorProp, bool)`

GetRouteToBackendSetBehaviorOk returns a tuple with the RouteToBackendSetBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteToBackendSetBehavior

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) SetRouteToBackendSetBehavior(v EnumextendedOperationHandlerRouteToBackendSetBehaviorProp)`

SetRouteToBackendSetBehavior sets RouteToBackendSetBehavior field to given value.

### HasRouteToBackendSetBehavior

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) HasRouteToBackendSetBehavior() bool`

HasRouteToBackendSetBehavior returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyProxiedExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


