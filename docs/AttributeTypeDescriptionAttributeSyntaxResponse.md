# AttributeTypeDescriptionAttributeSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn**](EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn.md) |  | 
**StripSyntaxMinUpperBound** | Pointer to **bool** | Indicates whether the suggested minimum upper bound appended to an attribute&#39;s syntax OID in its schema definition Attribute Type Description should be stripped. | [optional] 
**Enabled** | **bool** | Indicates whether the Attribute Syntax is enabled. | 
**RequireBinaryTransfer** | Pointer to **bool** | Indicates whether values of this attribute are required to have a \&quot;binary\&quot; transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \&quot;;binary\&quot; (e.g., \&quot;userCertificate;binary\&quot;). | [optional] 

## Methods

### NewAttributeTypeDescriptionAttributeSyntaxResponse

`func NewAttributeTypeDescriptionAttributeSyntaxResponse(schemas []EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn, enabled bool, ) *AttributeTypeDescriptionAttributeSyntaxResponse`

NewAttributeTypeDescriptionAttributeSyntaxResponse instantiates a new AttributeTypeDescriptionAttributeSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeTypeDescriptionAttributeSyntaxResponseWithDefaults

`func NewAttributeTypeDescriptionAttributeSyntaxResponseWithDefaults() *AttributeTypeDescriptionAttributeSyntaxResponse`

NewAttributeTypeDescriptionAttributeSyntaxResponseWithDefaults instantiates a new AttributeTypeDescriptionAttributeSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetSchemas() []EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetSchemasOk() (*[]EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetSchemas(v []EnumattributeTypeDescriptionAttributeSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetStripSyntaxMinUpperBound

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetStripSyntaxMinUpperBound() bool`

GetStripSyntaxMinUpperBound returns the StripSyntaxMinUpperBound field if non-nil, zero value otherwise.

### GetStripSyntaxMinUpperBoundOk

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetStripSyntaxMinUpperBoundOk() (*bool, bool)`

GetStripSyntaxMinUpperBoundOk returns a tuple with the StripSyntaxMinUpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripSyntaxMinUpperBound

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetStripSyntaxMinUpperBound(v bool)`

SetStripSyntaxMinUpperBound sets StripSyntaxMinUpperBound field to given value.

### HasStripSyntaxMinUpperBound

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) HasStripSyntaxMinUpperBound() bool`

HasStripSyntaxMinUpperBound returns a boolean if a field has been set.

### GetEnabled

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireBinaryTransfer

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetRequireBinaryTransfer() bool`

GetRequireBinaryTransfer returns the RequireBinaryTransfer field if non-nil, zero value otherwise.

### GetRequireBinaryTransferOk

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool)`

GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBinaryTransfer

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool)`

SetRequireBinaryTransfer sets RequireBinaryTransfer field to given value.

### HasRequireBinaryTransfer

`func (o *AttributeTypeDescriptionAttributeSyntaxResponse) HasRequireBinaryTransfer() bool`

HasRequireBinaryTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


