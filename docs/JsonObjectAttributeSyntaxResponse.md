# JsonObjectAttributeSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumjsonObjectAttributeSyntaxSchemaUrn**](EnumjsonObjectAttributeSyntaxSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Attribute Syntax is enabled. | 
**RequireBinaryTransfer** | Pointer to **bool** | Indicates whether values of this attribute are required to have a \&quot;binary\&quot; transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \&quot;;binary\&quot; (e.g., \&quot;userCertificate;binary\&quot;). | [optional] 

## Methods

### NewJsonObjectAttributeSyntaxResponse

`func NewJsonObjectAttributeSyntaxResponse(schemas []EnumjsonObjectAttributeSyntaxSchemaUrn, enabled bool, ) *JsonObjectAttributeSyntaxResponse`

NewJsonObjectAttributeSyntaxResponse instantiates a new JsonObjectAttributeSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonObjectAttributeSyntaxResponseWithDefaults

`func NewJsonObjectAttributeSyntaxResponseWithDefaults() *JsonObjectAttributeSyntaxResponse`

NewJsonObjectAttributeSyntaxResponseWithDefaults instantiates a new JsonObjectAttributeSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *JsonObjectAttributeSyntaxResponse) GetSchemas() []EnumjsonObjectAttributeSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JsonObjectAttributeSyntaxResponse) GetSchemasOk() (*[]EnumjsonObjectAttributeSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JsonObjectAttributeSyntaxResponse) SetSchemas(v []EnumjsonObjectAttributeSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *JsonObjectAttributeSyntaxResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JsonObjectAttributeSyntaxResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JsonObjectAttributeSyntaxResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireBinaryTransfer

`func (o *JsonObjectAttributeSyntaxResponse) GetRequireBinaryTransfer() bool`

GetRequireBinaryTransfer returns the RequireBinaryTransfer field if non-nil, zero value otherwise.

### GetRequireBinaryTransferOk

`func (o *JsonObjectAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool)`

GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBinaryTransfer

`func (o *JsonObjectAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool)`

SetRequireBinaryTransfer sets RequireBinaryTransfer field to given value.

### HasRequireBinaryTransfer

`func (o *JsonObjectAttributeSyntaxResponse) HasRequireBinaryTransfer() bool`

HasRequireBinaryTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


