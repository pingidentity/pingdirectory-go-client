# ThirdPartyOtpDeliveryMechanismResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the OTP Delivery Mechanism | 
**Schemas** | [**[]EnumthirdPartyOtpDeliveryMechanismSchemaUrn**](EnumthirdPartyOtpDeliveryMechanismSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party OTP Delivery Mechanism. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this OTP Delivery Mechanism | [optional] 
**Enabled** | **bool** | Indicates whether this OTP Delivery Mechanism is enabled for use in the server. | 

## Methods

### NewThirdPartyOtpDeliveryMechanismResponse

`func NewThirdPartyOtpDeliveryMechanismResponse(id string, schemas []EnumthirdPartyOtpDeliveryMechanismSchemaUrn, extensionClass string, enabled bool, ) *ThirdPartyOtpDeliveryMechanismResponse`

NewThirdPartyOtpDeliveryMechanismResponse instantiates a new ThirdPartyOtpDeliveryMechanismResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyOtpDeliveryMechanismResponseWithDefaults

`func NewThirdPartyOtpDeliveryMechanismResponseWithDefaults() *ThirdPartyOtpDeliveryMechanismResponse`

NewThirdPartyOtpDeliveryMechanismResponseWithDefaults instantiates a new ThirdPartyOtpDeliveryMechanismResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyOtpDeliveryMechanismResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetSchemas() []EnumthirdPartyOtpDeliveryMechanismSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetSchemasOk() (*[]EnumthirdPartyOtpDeliveryMechanismSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyOtpDeliveryMechanismResponse) SetSchemas(v []EnumthirdPartyOtpDeliveryMechanismSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyOtpDeliveryMechanismResponse) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyOtpDeliveryMechanismResponse) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyOtpDeliveryMechanismResponse) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyOtpDeliveryMechanismResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyOtpDeliveryMechanismResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyOtpDeliveryMechanismResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyOtpDeliveryMechanismResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


