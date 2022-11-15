# DirectoryStringAttributeSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdirectoryStringAttributeSyntaxSchemaUrn**](EnumdirectoryStringAttributeSyntaxSchemaUrn.md) |  | 
**AllowZeroLengthValues** | Pointer to **bool** | Indicates whether zero-length (that is, an empty string) values are allowed. | [optional] 
**Enabled** | **bool** | Indicates whether the Attribute Syntax is enabled. | 
**RequireBinaryTransfer** | Pointer to **bool** | Indicates whether values of this attribute are required to have a \&quot;binary\&quot; transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \&quot;;binary\&quot; (e.g., \&quot;userCertificate;binary\&quot;). | [optional] 

## Methods

### NewDirectoryStringAttributeSyntaxResponse

`func NewDirectoryStringAttributeSyntaxResponse(schemas []EnumdirectoryStringAttributeSyntaxSchemaUrn, enabled bool, ) *DirectoryStringAttributeSyntaxResponse`

NewDirectoryStringAttributeSyntaxResponse instantiates a new DirectoryStringAttributeSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryStringAttributeSyntaxResponseWithDefaults

`func NewDirectoryStringAttributeSyntaxResponseWithDefaults() *DirectoryStringAttributeSyntaxResponse`

NewDirectoryStringAttributeSyntaxResponseWithDefaults instantiates a new DirectoryStringAttributeSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DirectoryStringAttributeSyntaxResponse) GetSchemas() []EnumdirectoryStringAttributeSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DirectoryStringAttributeSyntaxResponse) GetSchemasOk() (*[]EnumdirectoryStringAttributeSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DirectoryStringAttributeSyntaxResponse) SetSchemas(v []EnumdirectoryStringAttributeSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowZeroLengthValues

`func (o *DirectoryStringAttributeSyntaxResponse) GetAllowZeroLengthValues() bool`

GetAllowZeroLengthValues returns the AllowZeroLengthValues field if non-nil, zero value otherwise.

### GetAllowZeroLengthValuesOk

`func (o *DirectoryStringAttributeSyntaxResponse) GetAllowZeroLengthValuesOk() (*bool, bool)`

GetAllowZeroLengthValuesOk returns a tuple with the AllowZeroLengthValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowZeroLengthValues

`func (o *DirectoryStringAttributeSyntaxResponse) SetAllowZeroLengthValues(v bool)`

SetAllowZeroLengthValues sets AllowZeroLengthValues field to given value.

### HasAllowZeroLengthValues

`func (o *DirectoryStringAttributeSyntaxResponse) HasAllowZeroLengthValues() bool`

HasAllowZeroLengthValues returns a boolean if a field has been set.

### GetEnabled

`func (o *DirectoryStringAttributeSyntaxResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DirectoryStringAttributeSyntaxResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DirectoryStringAttributeSyntaxResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireBinaryTransfer

`func (o *DirectoryStringAttributeSyntaxResponse) GetRequireBinaryTransfer() bool`

GetRequireBinaryTransfer returns the RequireBinaryTransfer field if non-nil, zero value otherwise.

### GetRequireBinaryTransferOk

`func (o *DirectoryStringAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool)`

GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBinaryTransfer

`func (o *DirectoryStringAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool)`

SetRequireBinaryTransfer sets RequireBinaryTransfer field to given value.

### HasRequireBinaryTransfer

`func (o *DirectoryStringAttributeSyntaxResponse) HasRequireBinaryTransfer() bool`

HasRequireBinaryTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


