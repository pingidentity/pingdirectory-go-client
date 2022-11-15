# LdapUrlAttributeSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumldapUrlAttributeSyntaxSchemaUrn**](EnumldapUrlAttributeSyntaxSchemaUrn.md) |  | 
**StrictFormat** | Pointer to **bool** | Indicates whether values for attributes with this syntax will be required to be in the valid LDAP URL format. If this is set to false, then arbitrary strings will be allowed. | [optional] 
**Enabled** | **bool** | Indicates whether the Attribute Syntax is enabled. | 
**RequireBinaryTransfer** | Pointer to **bool** | Indicates whether values of this attribute are required to have a \&quot;binary\&quot; transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \&quot;;binary\&quot; (e.g., \&quot;userCertificate;binary\&quot;). | [optional] 

## Methods

### NewLdapUrlAttributeSyntaxResponse

`func NewLdapUrlAttributeSyntaxResponse(schemas []EnumldapUrlAttributeSyntaxSchemaUrn, enabled bool, ) *LdapUrlAttributeSyntaxResponse`

NewLdapUrlAttributeSyntaxResponse instantiates a new LdapUrlAttributeSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapUrlAttributeSyntaxResponseWithDefaults

`func NewLdapUrlAttributeSyntaxResponseWithDefaults() *LdapUrlAttributeSyntaxResponse`

NewLdapUrlAttributeSyntaxResponseWithDefaults instantiates a new LdapUrlAttributeSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LdapUrlAttributeSyntaxResponse) GetSchemas() []EnumldapUrlAttributeSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdapUrlAttributeSyntaxResponse) GetSchemasOk() (*[]EnumldapUrlAttributeSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdapUrlAttributeSyntaxResponse) SetSchemas(v []EnumldapUrlAttributeSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetStrictFormat

`func (o *LdapUrlAttributeSyntaxResponse) GetStrictFormat() bool`

GetStrictFormat returns the StrictFormat field if non-nil, zero value otherwise.

### GetStrictFormatOk

`func (o *LdapUrlAttributeSyntaxResponse) GetStrictFormatOk() (*bool, bool)`

GetStrictFormatOk returns a tuple with the StrictFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictFormat

`func (o *LdapUrlAttributeSyntaxResponse) SetStrictFormat(v bool)`

SetStrictFormat sets StrictFormat field to given value.

### HasStrictFormat

`func (o *LdapUrlAttributeSyntaxResponse) HasStrictFormat() bool`

HasStrictFormat returns a boolean if a field has been set.

### GetEnabled

`func (o *LdapUrlAttributeSyntaxResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LdapUrlAttributeSyntaxResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LdapUrlAttributeSyntaxResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireBinaryTransfer

`func (o *LdapUrlAttributeSyntaxResponse) GetRequireBinaryTransfer() bool`

GetRequireBinaryTransfer returns the RequireBinaryTransfer field if non-nil, zero value otherwise.

### GetRequireBinaryTransferOk

`func (o *LdapUrlAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool)`

GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBinaryTransfer

`func (o *LdapUrlAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool)`

SetRequireBinaryTransfer sets RequireBinaryTransfer field to given value.

### HasRequireBinaryTransfer

`func (o *LdapUrlAttributeSyntaxResponse) HasRequireBinaryTransfer() bool`

HasRequireBinaryTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


