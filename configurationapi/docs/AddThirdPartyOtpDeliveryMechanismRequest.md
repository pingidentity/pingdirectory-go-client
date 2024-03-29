# AddThirdPartyOtpDeliveryMechanismRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyOtpDeliveryMechanismSchemaUrn**](EnumthirdPartyOtpDeliveryMechanismSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party OTP Delivery Mechanism. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party OTP Delivery Mechanism. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this OTP Delivery Mechanism | [optional] 
**Enabled** | **bool** | Indicates whether this OTP Delivery Mechanism is enabled for use in the server. | 
**MechanismName** | **string** | Name of the new OTP Delivery Mechanism | 

## Methods

### NewAddThirdPartyOtpDeliveryMechanismRequest

`func NewAddThirdPartyOtpDeliveryMechanismRequest(schemas []EnumthirdPartyOtpDeliveryMechanismSchemaUrn, extensionClass string, enabled bool, mechanismName string, ) *AddThirdPartyOtpDeliveryMechanismRequest`

NewAddThirdPartyOtpDeliveryMechanismRequest instantiates a new AddThirdPartyOtpDeliveryMechanismRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyOtpDeliveryMechanismRequestWithDefaults

`func NewAddThirdPartyOtpDeliveryMechanismRequestWithDefaults() *AddThirdPartyOtpDeliveryMechanismRequest`

NewAddThirdPartyOtpDeliveryMechanismRequestWithDefaults instantiates a new AddThirdPartyOtpDeliveryMechanismRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetSchemas() []EnumthirdPartyOtpDeliveryMechanismSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetSchemasOk() (*[]EnumthirdPartyOtpDeliveryMechanismSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetSchemas(v []EnumthirdPartyOtpDeliveryMechanismSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMechanismName

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetMechanismName() string`

GetMechanismName returns the MechanismName field if non-nil, zero value otherwise.

### GetMechanismNameOk

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) GetMechanismNameOk() (*string, bool)`

GetMechanismNameOk returns a tuple with the MechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanismName

`func (o *AddThirdPartyOtpDeliveryMechanismRequest) SetMechanismName(v string)`

SetMechanismName sets MechanismName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


