# IntegerAttributeSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumintegerAttributeSyntaxSchemaUrn**](EnumintegerAttributeSyntaxSchemaUrn.md) |  | 
**EnableCompaction** | Pointer to **bool** | Indicates whether values of attributes with this syntax should be compacted when stored in a local DB database. | [optional] 
**IncludeAttributeInCompaction** | Pointer to **[]string** |  | [optional] 
**ExcludeAttributeFromCompaction** | Pointer to **[]string** |  | [optional] 
**Enabled** | **bool** | Indicates whether the Attribute Syntax is enabled. | 
**RequireBinaryTransfer** | Pointer to **bool** | Indicates whether values of this attribute are required to have a \&quot;binary\&quot; transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \&quot;;binary\&quot; (e.g., \&quot;userCertificate;binary\&quot;). | [optional] 

## Methods

### NewIntegerAttributeSyntaxResponse

`func NewIntegerAttributeSyntaxResponse(schemas []EnumintegerAttributeSyntaxSchemaUrn, enabled bool, ) *IntegerAttributeSyntaxResponse`

NewIntegerAttributeSyntaxResponse instantiates a new IntegerAttributeSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegerAttributeSyntaxResponseWithDefaults

`func NewIntegerAttributeSyntaxResponseWithDefaults() *IntegerAttributeSyntaxResponse`

NewIntegerAttributeSyntaxResponseWithDefaults instantiates a new IntegerAttributeSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *IntegerAttributeSyntaxResponse) GetSchemas() []EnumintegerAttributeSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *IntegerAttributeSyntaxResponse) GetSchemasOk() (*[]EnumintegerAttributeSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *IntegerAttributeSyntaxResponse) SetSchemas(v []EnumintegerAttributeSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnableCompaction

`func (o *IntegerAttributeSyntaxResponse) GetEnableCompaction() bool`

GetEnableCompaction returns the EnableCompaction field if non-nil, zero value otherwise.

### GetEnableCompactionOk

`func (o *IntegerAttributeSyntaxResponse) GetEnableCompactionOk() (*bool, bool)`

GetEnableCompactionOk returns a tuple with the EnableCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCompaction

`func (o *IntegerAttributeSyntaxResponse) SetEnableCompaction(v bool)`

SetEnableCompaction sets EnableCompaction field to given value.

### HasEnableCompaction

`func (o *IntegerAttributeSyntaxResponse) HasEnableCompaction() bool`

HasEnableCompaction returns a boolean if a field has been set.

### GetIncludeAttributeInCompaction

`func (o *IntegerAttributeSyntaxResponse) GetIncludeAttributeInCompaction() []string`

GetIncludeAttributeInCompaction returns the IncludeAttributeInCompaction field if non-nil, zero value otherwise.

### GetIncludeAttributeInCompactionOk

`func (o *IntegerAttributeSyntaxResponse) GetIncludeAttributeInCompactionOk() (*[]string, bool)`

GetIncludeAttributeInCompactionOk returns a tuple with the IncludeAttributeInCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttributeInCompaction

`func (o *IntegerAttributeSyntaxResponse) SetIncludeAttributeInCompaction(v []string)`

SetIncludeAttributeInCompaction sets IncludeAttributeInCompaction field to given value.

### HasIncludeAttributeInCompaction

`func (o *IntegerAttributeSyntaxResponse) HasIncludeAttributeInCompaction() bool`

HasIncludeAttributeInCompaction returns a boolean if a field has been set.

### GetExcludeAttributeFromCompaction

`func (o *IntegerAttributeSyntaxResponse) GetExcludeAttributeFromCompaction() []string`

GetExcludeAttributeFromCompaction returns the ExcludeAttributeFromCompaction field if non-nil, zero value otherwise.

### GetExcludeAttributeFromCompactionOk

`func (o *IntegerAttributeSyntaxResponse) GetExcludeAttributeFromCompactionOk() (*[]string, bool)`

GetExcludeAttributeFromCompactionOk returns a tuple with the ExcludeAttributeFromCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttributeFromCompaction

`func (o *IntegerAttributeSyntaxResponse) SetExcludeAttributeFromCompaction(v []string)`

SetExcludeAttributeFromCompaction sets ExcludeAttributeFromCompaction field to given value.

### HasExcludeAttributeFromCompaction

`func (o *IntegerAttributeSyntaxResponse) HasExcludeAttributeFromCompaction() bool`

HasExcludeAttributeFromCompaction returns a boolean if a field has been set.

### GetEnabled

`func (o *IntegerAttributeSyntaxResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegerAttributeSyntaxResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegerAttributeSyntaxResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireBinaryTransfer

`func (o *IntegerAttributeSyntaxResponse) GetRequireBinaryTransfer() bool`

GetRequireBinaryTransfer returns the RequireBinaryTransfer field if non-nil, zero value otherwise.

### GetRequireBinaryTransferOk

`func (o *IntegerAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool)`

GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBinaryTransfer

`func (o *IntegerAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool)`

SetRequireBinaryTransfer sets RequireBinaryTransfer field to given value.

### HasRequireBinaryTransfer

`func (o *IntegerAttributeSyntaxResponse) HasRequireBinaryTransfer() bool`

HasRequireBinaryTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


