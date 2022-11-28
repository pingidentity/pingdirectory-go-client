# TelephoneNumberAttributeSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumtelephoneNumberAttributeSyntaxSchemaUrn**](EnumtelephoneNumberAttributeSyntaxSchemaUrn.md) |  | 
**Id** | Pointer to **string** | Name of the Attribute Syntax | [optional] 
**StrictFormat** | Pointer to **bool** | Indicates whether to require telephone number values to strictly comply with the standard definition for this syntax. | [optional] 
**Enabled** | **bool** | Indicates whether the Attribute Syntax is enabled. | 
**RequireBinaryTransfer** | Pointer to **bool** | Indicates whether values of this attribute are required to have a \&quot;binary\&quot; transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \&quot;;binary\&quot; (e.g., \&quot;userCertificate;binary\&quot;). | [optional] 

## Methods

### NewTelephoneNumberAttributeSyntaxResponse

`func NewTelephoneNumberAttributeSyntaxResponse(schemas []EnumtelephoneNumberAttributeSyntaxSchemaUrn, enabled bool, ) *TelephoneNumberAttributeSyntaxResponse`

NewTelephoneNumberAttributeSyntaxResponse instantiates a new TelephoneNumberAttributeSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephoneNumberAttributeSyntaxResponseWithDefaults

`func NewTelephoneNumberAttributeSyntaxResponseWithDefaults() *TelephoneNumberAttributeSyntaxResponse`

NewTelephoneNumberAttributeSyntaxResponseWithDefaults instantiates a new TelephoneNumberAttributeSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *TelephoneNumberAttributeSyntaxResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TelephoneNumberAttributeSyntaxResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TelephoneNumberAttributeSyntaxResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TelephoneNumberAttributeSyntaxResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *TelephoneNumberAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *TelephoneNumberAttributeSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *TelephoneNumberAttributeSyntaxResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *TelephoneNumberAttributeSyntaxResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *TelephoneNumberAttributeSyntaxResponse) GetSchemas() []EnumtelephoneNumberAttributeSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TelephoneNumberAttributeSyntaxResponse) GetSchemasOk() (*[]EnumtelephoneNumberAttributeSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TelephoneNumberAttributeSyntaxResponse) SetSchemas(v []EnumtelephoneNumberAttributeSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *TelephoneNumberAttributeSyntaxResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TelephoneNumberAttributeSyntaxResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TelephoneNumberAttributeSyntaxResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TelephoneNumberAttributeSyntaxResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStrictFormat

`func (o *TelephoneNumberAttributeSyntaxResponse) GetStrictFormat() bool`

GetStrictFormat returns the StrictFormat field if non-nil, zero value otherwise.

### GetStrictFormatOk

`func (o *TelephoneNumberAttributeSyntaxResponse) GetStrictFormatOk() (*bool, bool)`

GetStrictFormatOk returns a tuple with the StrictFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictFormat

`func (o *TelephoneNumberAttributeSyntaxResponse) SetStrictFormat(v bool)`

SetStrictFormat sets StrictFormat field to given value.

### HasStrictFormat

`func (o *TelephoneNumberAttributeSyntaxResponse) HasStrictFormat() bool`

HasStrictFormat returns a boolean if a field has been set.

### GetEnabled

`func (o *TelephoneNumberAttributeSyntaxResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TelephoneNumberAttributeSyntaxResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TelephoneNumberAttributeSyntaxResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireBinaryTransfer

`func (o *TelephoneNumberAttributeSyntaxResponse) GetRequireBinaryTransfer() bool`

GetRequireBinaryTransfer returns the RequireBinaryTransfer field if non-nil, zero value otherwise.

### GetRequireBinaryTransferOk

`func (o *TelephoneNumberAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool)`

GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBinaryTransfer

`func (o *TelephoneNumberAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool)`

SetRequireBinaryTransfer sets RequireBinaryTransfer field to given value.

### HasRequireBinaryTransfer

`func (o *TelephoneNumberAttributeSyntaxResponse) HasRequireBinaryTransfer() bool`

HasRequireBinaryTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


