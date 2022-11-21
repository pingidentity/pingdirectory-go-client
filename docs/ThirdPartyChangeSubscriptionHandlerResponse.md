# ThirdPartyChangeSubscriptionHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Change Subscription Handler | 
**Schemas** | [**[]EnumthirdPartyChangeSubscriptionHandlerSchemaUrn**](EnumthirdPartyChangeSubscriptionHandlerSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Change Subscription Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Change Subscription Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Change Subscription Handler | [optional] 
**Enabled** | **bool** | Indicates whether this change subscription handler is enabled within the server. | 
**ChangeSubscription** | Pointer to **[]string** | The set of change subscriptions for which this change subscription handler should be notified. If no values are provided then it will be notified for all change subscriptions defined in the server. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewThirdPartyChangeSubscriptionHandlerResponse

`func NewThirdPartyChangeSubscriptionHandlerResponse(id string, schemas []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, extensionClass string, enabled bool, ) *ThirdPartyChangeSubscriptionHandlerResponse`

NewThirdPartyChangeSubscriptionHandlerResponse instantiates a new ThirdPartyChangeSubscriptionHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyChangeSubscriptionHandlerResponseWithDefaults

`func NewThirdPartyChangeSubscriptionHandlerResponseWithDefaults() *ThirdPartyChangeSubscriptionHandlerResponse`

NewThirdPartyChangeSubscriptionHandlerResponseWithDefaults instantiates a new ThirdPartyChangeSubscriptionHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetSchemas() []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetSchemasOk() (*[]EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) SetSchemas(v []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetChangeSubscription

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetChangeSubscription() []string`

GetChangeSubscription returns the ChangeSubscription field if non-nil, zero value otherwise.

### GetChangeSubscriptionOk

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetChangeSubscriptionOk() (*[]string, bool)`

GetChangeSubscriptionOk returns a tuple with the ChangeSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSubscription

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) SetChangeSubscription(v []string)`

SetChangeSubscription sets ChangeSubscription field to given value.

### HasChangeSubscription

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) HasChangeSubscription() bool`

HasChangeSubscription returns a boolean if a field has been set.

### GetMeta

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ThirdPartyChangeSubscriptionHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


