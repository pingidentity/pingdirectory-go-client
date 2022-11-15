# ThirdPartyErrorLogPublisherShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyErrorLogPublisherSchemaUrn**](EnumthirdPartyErrorLogPublisherSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Error Log Publisher. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**DefaultSeverity** | Pointer to [**[]EnumlogPublisherDefaultSeverityProp**](EnumlogPublisherDefaultSeverityProp.md) |  | [optional] 
**OverrideSeverity** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewThirdPartyErrorLogPublisherShared

`func NewThirdPartyErrorLogPublisherShared(schemas []EnumthirdPartyErrorLogPublisherSchemaUrn, extensionClass string, enabled bool, ) *ThirdPartyErrorLogPublisherShared`

NewThirdPartyErrorLogPublisherShared instantiates a new ThirdPartyErrorLogPublisherShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyErrorLogPublisherSharedWithDefaults

`func NewThirdPartyErrorLogPublisherSharedWithDefaults() *ThirdPartyErrorLogPublisherShared`

NewThirdPartyErrorLogPublisherSharedWithDefaults instantiates a new ThirdPartyErrorLogPublisherShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ThirdPartyErrorLogPublisherShared) GetSchemas() []EnumthirdPartyErrorLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyErrorLogPublisherShared) GetSchemasOk() (*[]EnumthirdPartyErrorLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyErrorLogPublisherShared) SetSchemas(v []EnumthirdPartyErrorLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyErrorLogPublisherShared) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyErrorLogPublisherShared) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyErrorLogPublisherShared) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyErrorLogPublisherShared) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyErrorLogPublisherShared) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyErrorLogPublisherShared) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyErrorLogPublisherShared) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDefaultSeverity

`func (o *ThirdPartyErrorLogPublisherShared) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp`

GetDefaultSeverity returns the DefaultSeverity field if non-nil, zero value otherwise.

### GetDefaultSeverityOk

`func (o *ThirdPartyErrorLogPublisherShared) GetDefaultSeverityOk() (*[]EnumlogPublisherDefaultSeverityProp, bool)`

GetDefaultSeverityOk returns a tuple with the DefaultSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeverity

`func (o *ThirdPartyErrorLogPublisherShared) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp)`

SetDefaultSeverity sets DefaultSeverity field to given value.

### HasDefaultSeverity

`func (o *ThirdPartyErrorLogPublisherShared) HasDefaultSeverity() bool`

HasDefaultSeverity returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *ThirdPartyErrorLogPublisherShared) GetOverrideSeverity() []string`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *ThirdPartyErrorLogPublisherShared) GetOverrideSeverityOk() (*[]string, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *ThirdPartyErrorLogPublisherShared) SetOverrideSeverity(v []string)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *ThirdPartyErrorLogPublisherShared) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyErrorLogPublisherShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyErrorLogPublisherShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyErrorLogPublisherShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyErrorLogPublisherShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyErrorLogPublisherShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyErrorLogPublisherShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyErrorLogPublisherShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *ThirdPartyErrorLogPublisherShared) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *ThirdPartyErrorLogPublisherShared) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *ThirdPartyErrorLogPublisherShared) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *ThirdPartyErrorLogPublisherShared) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


