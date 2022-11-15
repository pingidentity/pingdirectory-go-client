# ThirdPartyChangeSubscriptionHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyChangeSubscriptionHandlerSchemaUrn**](EnumthirdPartyChangeSubscriptionHandlerSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Change Subscription Handler. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Change Subscription Handler | [optional] 
**Enabled** | **bool** | Indicates whether this change subscription handler is enabled within the server. | 
**ChangeSubscription** | Pointer to **[]string** |  | [optional] 

## Methods

### NewThirdPartyChangeSubscriptionHandlerShared

`func NewThirdPartyChangeSubscriptionHandlerShared(schemas []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, extensionClass string, enabled bool, ) *ThirdPartyChangeSubscriptionHandlerShared`

NewThirdPartyChangeSubscriptionHandlerShared instantiates a new ThirdPartyChangeSubscriptionHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyChangeSubscriptionHandlerSharedWithDefaults

`func NewThirdPartyChangeSubscriptionHandlerSharedWithDefaults() *ThirdPartyChangeSubscriptionHandlerShared`

NewThirdPartyChangeSubscriptionHandlerSharedWithDefaults instantiates a new ThirdPartyChangeSubscriptionHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetSchemas() []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetSchemasOk() (*[]EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyChangeSubscriptionHandlerShared) SetSchemas(v []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyChangeSubscriptionHandlerShared) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyChangeSubscriptionHandlerShared) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyChangeSubscriptionHandlerShared) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyChangeSubscriptionHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyChangeSubscriptionHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyChangeSubscriptionHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetChangeSubscription

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetChangeSubscription() []string`

GetChangeSubscription returns the ChangeSubscription field if non-nil, zero value otherwise.

### GetChangeSubscriptionOk

`func (o *ThirdPartyChangeSubscriptionHandlerShared) GetChangeSubscriptionOk() (*[]string, bool)`

GetChangeSubscriptionOk returns a tuple with the ChangeSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSubscription

`func (o *ThirdPartyChangeSubscriptionHandlerShared) SetChangeSubscription(v []string)`

SetChangeSubscription sets ChangeSubscription field to given value.

### HasChangeSubscription

`func (o *ThirdPartyChangeSubscriptionHandlerShared) HasChangeSubscription() bool`

HasChangeSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


