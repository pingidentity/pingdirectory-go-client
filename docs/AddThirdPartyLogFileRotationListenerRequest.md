# AddThirdPartyLogFileRotationListenerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListenerName** | **string** | Name of the new Log File Rotation Listener | 
**Schemas** | [**[]EnumthirdPartyLogFileRotationListenerSchemaUrn**](EnumthirdPartyLogFileRotationListenerSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Log File Rotation Listener. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Log File Rotation Listener. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Log File Rotation Listener | [optional] 
**Enabled** | **bool** | Indicates whether the Log File Rotation Listener is enabled for use. | 

## Methods

### NewAddThirdPartyLogFileRotationListenerRequest

`func NewAddThirdPartyLogFileRotationListenerRequest(listenerName string, schemas []EnumthirdPartyLogFileRotationListenerSchemaUrn, extensionClass string, enabled bool, ) *AddThirdPartyLogFileRotationListenerRequest`

NewAddThirdPartyLogFileRotationListenerRequest instantiates a new AddThirdPartyLogFileRotationListenerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyLogFileRotationListenerRequestWithDefaults

`func NewAddThirdPartyLogFileRotationListenerRequestWithDefaults() *AddThirdPartyLogFileRotationListenerRequest`

NewAddThirdPartyLogFileRotationListenerRequestWithDefaults instantiates a new AddThirdPartyLogFileRotationListenerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListenerName

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetListenerName() string`

GetListenerName returns the ListenerName field if non-nil, zero value otherwise.

### GetListenerNameOk

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetListenerNameOk() (*string, bool)`

GetListenerNameOk returns a tuple with the ListenerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerName

`func (o *AddThirdPartyLogFileRotationListenerRequest) SetListenerName(v string)`

SetListenerName sets ListenerName field to given value.


### GetSchemas

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetSchemas() []EnumthirdPartyLogFileRotationListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetSchemasOk() (*[]EnumthirdPartyLogFileRotationListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyLogFileRotationListenerRequest) SetSchemas(v []EnumthirdPartyLogFileRotationListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyLogFileRotationListenerRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyLogFileRotationListenerRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyLogFileRotationListenerRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyLogFileRotationListenerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyLogFileRotationListenerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyLogFileRotationListenerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyLogFileRotationListenerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


