# ThirdPartyAlertHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyAlertHandlerSchemaUrn**](EnumthirdPartyAlertHandlerSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Alert Handler. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Alert Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Alert Handler is enabled. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated. | [optional] 
**EnabledAlertSeverity** | Pointer to [**[]EnumalertHandlerEnabledAlertSeverityProp**](EnumalertHandlerEnabledAlertSeverityProp.md) |  | [optional] 
**EnabledAlertType** | Pointer to [**[]EnumalertHandlerEnabledAlertTypeProp**](EnumalertHandlerEnabledAlertTypeProp.md) |  | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumalertHandlerDisabledAlertTypeProp**](EnumalertHandlerDisabledAlertTypeProp.md) |  | [optional] 

## Methods

### NewThirdPartyAlertHandlerShared

`func NewThirdPartyAlertHandlerShared(schemas []EnumthirdPartyAlertHandlerSchemaUrn, extensionClass string, enabled bool, ) *ThirdPartyAlertHandlerShared`

NewThirdPartyAlertHandlerShared instantiates a new ThirdPartyAlertHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyAlertHandlerSharedWithDefaults

`func NewThirdPartyAlertHandlerSharedWithDefaults() *ThirdPartyAlertHandlerShared`

NewThirdPartyAlertHandlerSharedWithDefaults instantiates a new ThirdPartyAlertHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ThirdPartyAlertHandlerShared) GetSchemas() []EnumthirdPartyAlertHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyAlertHandlerShared) GetSchemasOk() (*[]EnumthirdPartyAlertHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyAlertHandlerShared) SetSchemas(v []EnumthirdPartyAlertHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyAlertHandlerShared) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyAlertHandlerShared) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyAlertHandlerShared) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyAlertHandlerShared) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyAlertHandlerShared) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyAlertHandlerShared) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyAlertHandlerShared) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyAlertHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyAlertHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyAlertHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyAlertHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyAlertHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyAlertHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyAlertHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *ThirdPartyAlertHandlerShared) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *ThirdPartyAlertHandlerShared) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *ThirdPartyAlertHandlerShared) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *ThirdPartyAlertHandlerShared) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetEnabledAlertSeverity

`func (o *ThirdPartyAlertHandlerShared) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp`

GetEnabledAlertSeverity returns the EnabledAlertSeverity field if non-nil, zero value otherwise.

### GetEnabledAlertSeverityOk

`func (o *ThirdPartyAlertHandlerShared) GetEnabledAlertSeverityOk() (*[]EnumalertHandlerEnabledAlertSeverityProp, bool)`

GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertSeverity

`func (o *ThirdPartyAlertHandlerShared) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp)`

SetEnabledAlertSeverity sets EnabledAlertSeverity field to given value.

### HasEnabledAlertSeverity

`func (o *ThirdPartyAlertHandlerShared) HasEnabledAlertSeverity() bool`

HasEnabledAlertSeverity returns a boolean if a field has been set.

### GetEnabledAlertType

`func (o *ThirdPartyAlertHandlerShared) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp`

GetEnabledAlertType returns the EnabledAlertType field if non-nil, zero value otherwise.

### GetEnabledAlertTypeOk

`func (o *ThirdPartyAlertHandlerShared) GetEnabledAlertTypeOk() (*[]EnumalertHandlerEnabledAlertTypeProp, bool)`

GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertType

`func (o *ThirdPartyAlertHandlerShared) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp)`

SetEnabledAlertType sets EnabledAlertType field to given value.

### HasEnabledAlertType

`func (o *ThirdPartyAlertHandlerShared) HasEnabledAlertType() bool`

HasEnabledAlertType returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *ThirdPartyAlertHandlerShared) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *ThirdPartyAlertHandlerShared) GetDisabledAlertTypeOk() (*[]EnumalertHandlerDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *ThirdPartyAlertHandlerShared) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *ThirdPartyAlertHandlerShared) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


