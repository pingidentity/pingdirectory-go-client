# DeliverOtpExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Extended Operation Handler | 
**Schemas** | [**[]EnumdeliverOtpExtendedOperationHandlerSchemaUrn**](EnumdeliverOtpExtendedOperationHandlerSchemaUrn.md) |  | 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the user(s) targeted by the authentication identity contained in the extended request. This will only be used for \&quot;u:\&quot;-style authentication identities. | 
**PasswordGenerator** | **string** | The password generator that will be used to create the one-time password values to be delivered to the end user. | 
**DefaultOTPDeliveryMechanism** | **[]string** |  | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewDeliverOtpExtendedOperationHandlerResponse

`func NewDeliverOtpExtendedOperationHandlerResponse(id string, schemas []EnumdeliverOtpExtendedOperationHandlerSchemaUrn, identityMapper string, passwordGenerator string, defaultOTPDeliveryMechanism []string, enabled bool, ) *DeliverOtpExtendedOperationHandlerResponse`

NewDeliverOtpExtendedOperationHandlerResponse instantiates a new DeliverOtpExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverOtpExtendedOperationHandlerResponseWithDefaults

`func NewDeliverOtpExtendedOperationHandlerResponseWithDefaults() *DeliverOtpExtendedOperationHandlerResponse`

NewDeliverOtpExtendedOperationHandlerResponseWithDefaults instantiates a new DeliverOtpExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliverOtpExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetSchemas() []EnumdeliverOtpExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumdeliverOtpExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DeliverOtpExtendedOperationHandlerResponse) SetSchemas(v []EnumdeliverOtpExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIdentityMapper

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *DeliverOtpExtendedOperationHandlerResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetPasswordGenerator

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *DeliverOtpExtendedOperationHandlerResponse) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.


### GetDefaultOTPDeliveryMechanism

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetDefaultOTPDeliveryMechanism() []string`

GetDefaultOTPDeliveryMechanism returns the DefaultOTPDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultOTPDeliveryMechanismOk

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetDefaultOTPDeliveryMechanismOk() (*[]string, bool)`

GetDefaultOTPDeliveryMechanismOk returns a tuple with the DefaultOTPDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOTPDeliveryMechanism

`func (o *DeliverOtpExtendedOperationHandlerResponse) SetDefaultOTPDeliveryMechanism(v []string)`

SetDefaultOTPDeliveryMechanism sets DefaultOTPDeliveryMechanism field to given value.


### GetDescription

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeliverOtpExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeliverOtpExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeliverOtpExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeliverOtpExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


