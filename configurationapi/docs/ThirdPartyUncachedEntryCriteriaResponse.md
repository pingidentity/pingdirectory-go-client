# ThirdPartyUncachedEntryCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyUncachedEntryCriteriaSchemaUrn**](EnumthirdPartyUncachedEntryCriteriaSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Uncached Entry Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Uncached Entry Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Uncached Entry Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Entry Criteria is enabled for use in the server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Uncached Entry Criteria | 

## Methods

### NewThirdPartyUncachedEntryCriteriaResponse

`func NewThirdPartyUncachedEntryCriteriaResponse(schemas []EnumthirdPartyUncachedEntryCriteriaSchemaUrn, extensionClass string, enabled bool, id string, ) *ThirdPartyUncachedEntryCriteriaResponse`

NewThirdPartyUncachedEntryCriteriaResponse instantiates a new ThirdPartyUncachedEntryCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyUncachedEntryCriteriaResponseWithDefaults

`func NewThirdPartyUncachedEntryCriteriaResponseWithDefaults() *ThirdPartyUncachedEntryCriteriaResponse`

NewThirdPartyUncachedEntryCriteriaResponseWithDefaults instantiates a new ThirdPartyUncachedEntryCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetSchemas() []EnumthirdPartyUncachedEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetSchemasOk() (*[]EnumthirdPartyUncachedEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyUncachedEntryCriteriaResponse) SetSchemas(v []EnumthirdPartyUncachedEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyUncachedEntryCriteriaResponse) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyUncachedEntryCriteriaResponse) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyUncachedEntryCriteriaResponse) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyUncachedEntryCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyUncachedEntryCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyUncachedEntryCriteriaResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ThirdPartyUncachedEntryCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ThirdPartyUncachedEntryCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyUncachedEntryCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyUncachedEntryCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyUncachedEntryCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyUncachedEntryCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


