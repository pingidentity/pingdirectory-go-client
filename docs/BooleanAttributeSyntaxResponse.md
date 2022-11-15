# BooleanAttributeSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumbooleanAttributeSyntaxSchemaUrn**](EnumbooleanAttributeSyntaxSchemaUrn.md) |  | 
**EnableCompaction** | Pointer to **bool** | Indicates whether values of attributes with this syntax should be compacted when stored in a local DB database. | [optional] 
**IncludeAttributeInCompaction** | Pointer to **[]string** |  | [optional] 
**ExcludeAttributeFromCompaction** | Pointer to **[]string** |  | [optional] 
**Enabled** | **bool** | Indicates whether the Attribute Syntax is enabled. | 
**RequireBinaryTransfer** | Pointer to **bool** | Indicates whether values of this attribute are required to have a \&quot;binary\&quot; transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \&quot;;binary\&quot; (e.g., \&quot;userCertificate;binary\&quot;). | [optional] 

## Methods

### NewBooleanAttributeSyntaxResponse

`func NewBooleanAttributeSyntaxResponse(schemas []EnumbooleanAttributeSyntaxSchemaUrn, enabled bool, ) *BooleanAttributeSyntaxResponse`

NewBooleanAttributeSyntaxResponse instantiates a new BooleanAttributeSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooleanAttributeSyntaxResponseWithDefaults

`func NewBooleanAttributeSyntaxResponseWithDefaults() *BooleanAttributeSyntaxResponse`

NewBooleanAttributeSyntaxResponseWithDefaults instantiates a new BooleanAttributeSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *BooleanAttributeSyntaxResponse) GetSchemas() []EnumbooleanAttributeSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BooleanAttributeSyntaxResponse) GetSchemasOk() (*[]EnumbooleanAttributeSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BooleanAttributeSyntaxResponse) SetSchemas(v []EnumbooleanAttributeSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnableCompaction

`func (o *BooleanAttributeSyntaxResponse) GetEnableCompaction() bool`

GetEnableCompaction returns the EnableCompaction field if non-nil, zero value otherwise.

### GetEnableCompactionOk

`func (o *BooleanAttributeSyntaxResponse) GetEnableCompactionOk() (*bool, bool)`

GetEnableCompactionOk returns a tuple with the EnableCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCompaction

`func (o *BooleanAttributeSyntaxResponse) SetEnableCompaction(v bool)`

SetEnableCompaction sets EnableCompaction field to given value.

### HasEnableCompaction

`func (o *BooleanAttributeSyntaxResponse) HasEnableCompaction() bool`

HasEnableCompaction returns a boolean if a field has been set.

### GetIncludeAttributeInCompaction

`func (o *BooleanAttributeSyntaxResponse) GetIncludeAttributeInCompaction() []string`

GetIncludeAttributeInCompaction returns the IncludeAttributeInCompaction field if non-nil, zero value otherwise.

### GetIncludeAttributeInCompactionOk

`func (o *BooleanAttributeSyntaxResponse) GetIncludeAttributeInCompactionOk() (*[]string, bool)`

GetIncludeAttributeInCompactionOk returns a tuple with the IncludeAttributeInCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttributeInCompaction

`func (o *BooleanAttributeSyntaxResponse) SetIncludeAttributeInCompaction(v []string)`

SetIncludeAttributeInCompaction sets IncludeAttributeInCompaction field to given value.

### HasIncludeAttributeInCompaction

`func (o *BooleanAttributeSyntaxResponse) HasIncludeAttributeInCompaction() bool`

HasIncludeAttributeInCompaction returns a boolean if a field has been set.

### GetExcludeAttributeFromCompaction

`func (o *BooleanAttributeSyntaxResponse) GetExcludeAttributeFromCompaction() []string`

GetExcludeAttributeFromCompaction returns the ExcludeAttributeFromCompaction field if non-nil, zero value otherwise.

### GetExcludeAttributeFromCompactionOk

`func (o *BooleanAttributeSyntaxResponse) GetExcludeAttributeFromCompactionOk() (*[]string, bool)`

GetExcludeAttributeFromCompactionOk returns a tuple with the ExcludeAttributeFromCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttributeFromCompaction

`func (o *BooleanAttributeSyntaxResponse) SetExcludeAttributeFromCompaction(v []string)`

SetExcludeAttributeFromCompaction sets ExcludeAttributeFromCompaction field to given value.

### HasExcludeAttributeFromCompaction

`func (o *BooleanAttributeSyntaxResponse) HasExcludeAttributeFromCompaction() bool`

HasExcludeAttributeFromCompaction returns a boolean if a field has been set.

### GetEnabled

`func (o *BooleanAttributeSyntaxResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BooleanAttributeSyntaxResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BooleanAttributeSyntaxResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireBinaryTransfer

`func (o *BooleanAttributeSyntaxResponse) GetRequireBinaryTransfer() bool`

GetRequireBinaryTransfer returns the RequireBinaryTransfer field if non-nil, zero value otherwise.

### GetRequireBinaryTransferOk

`func (o *BooleanAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool)`

GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBinaryTransfer

`func (o *BooleanAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool)`

SetRequireBinaryTransfer sets RequireBinaryTransfer field to given value.

### HasRequireBinaryTransfer

`func (o *BooleanAttributeSyntaxResponse) HasRequireBinaryTransfer() bool`

HasRequireBinaryTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


