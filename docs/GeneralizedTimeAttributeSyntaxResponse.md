# GeneralizedTimeAttributeSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgeneralizedTimeAttributeSyntaxSchemaUrn**](EnumgeneralizedTimeAttributeSyntaxSchemaUrn.md) |  | 
**EnableCompaction** | Pointer to **bool** | Indicates whether values of attributes with this syntax should be compacted when stored in a local DB database. | [optional] 
**IncludeAttributeInCompaction** | Pointer to **[]string** |  | [optional] 
**ExcludeAttributeFromCompaction** | Pointer to **[]string** |  | [optional] 
**Enabled** | **bool** | Indicates whether the Attribute Syntax is enabled. | 
**RequireBinaryTransfer** | Pointer to **bool** | Indicates whether values of this attribute are required to have a \&quot;binary\&quot; transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \&quot;;binary\&quot; (e.g., \&quot;userCertificate;binary\&quot;). | [optional] 

## Methods

### NewGeneralizedTimeAttributeSyntaxResponse

`func NewGeneralizedTimeAttributeSyntaxResponse(schemas []EnumgeneralizedTimeAttributeSyntaxSchemaUrn, enabled bool, ) *GeneralizedTimeAttributeSyntaxResponse`

NewGeneralizedTimeAttributeSyntaxResponse instantiates a new GeneralizedTimeAttributeSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralizedTimeAttributeSyntaxResponseWithDefaults

`func NewGeneralizedTimeAttributeSyntaxResponseWithDefaults() *GeneralizedTimeAttributeSyntaxResponse`

NewGeneralizedTimeAttributeSyntaxResponseWithDefaults instantiates a new GeneralizedTimeAttributeSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetSchemas() []EnumgeneralizedTimeAttributeSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetSchemasOk() (*[]EnumgeneralizedTimeAttributeSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GeneralizedTimeAttributeSyntaxResponse) SetSchemas(v []EnumgeneralizedTimeAttributeSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnableCompaction

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetEnableCompaction() bool`

GetEnableCompaction returns the EnableCompaction field if non-nil, zero value otherwise.

### GetEnableCompactionOk

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetEnableCompactionOk() (*bool, bool)`

GetEnableCompactionOk returns a tuple with the EnableCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCompaction

`func (o *GeneralizedTimeAttributeSyntaxResponse) SetEnableCompaction(v bool)`

SetEnableCompaction sets EnableCompaction field to given value.

### HasEnableCompaction

`func (o *GeneralizedTimeAttributeSyntaxResponse) HasEnableCompaction() bool`

HasEnableCompaction returns a boolean if a field has been set.

### GetIncludeAttributeInCompaction

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetIncludeAttributeInCompaction() []string`

GetIncludeAttributeInCompaction returns the IncludeAttributeInCompaction field if non-nil, zero value otherwise.

### GetIncludeAttributeInCompactionOk

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetIncludeAttributeInCompactionOk() (*[]string, bool)`

GetIncludeAttributeInCompactionOk returns a tuple with the IncludeAttributeInCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttributeInCompaction

`func (o *GeneralizedTimeAttributeSyntaxResponse) SetIncludeAttributeInCompaction(v []string)`

SetIncludeAttributeInCompaction sets IncludeAttributeInCompaction field to given value.

### HasIncludeAttributeInCompaction

`func (o *GeneralizedTimeAttributeSyntaxResponse) HasIncludeAttributeInCompaction() bool`

HasIncludeAttributeInCompaction returns a boolean if a field has been set.

### GetExcludeAttributeFromCompaction

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetExcludeAttributeFromCompaction() []string`

GetExcludeAttributeFromCompaction returns the ExcludeAttributeFromCompaction field if non-nil, zero value otherwise.

### GetExcludeAttributeFromCompactionOk

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetExcludeAttributeFromCompactionOk() (*[]string, bool)`

GetExcludeAttributeFromCompactionOk returns a tuple with the ExcludeAttributeFromCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttributeFromCompaction

`func (o *GeneralizedTimeAttributeSyntaxResponse) SetExcludeAttributeFromCompaction(v []string)`

SetExcludeAttributeFromCompaction sets ExcludeAttributeFromCompaction field to given value.

### HasExcludeAttributeFromCompaction

`func (o *GeneralizedTimeAttributeSyntaxResponse) HasExcludeAttributeFromCompaction() bool`

HasExcludeAttributeFromCompaction returns a boolean if a field has been set.

### GetEnabled

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GeneralizedTimeAttributeSyntaxResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireBinaryTransfer

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetRequireBinaryTransfer() bool`

GetRequireBinaryTransfer returns the RequireBinaryTransfer field if non-nil, zero value otherwise.

### GetRequireBinaryTransferOk

`func (o *GeneralizedTimeAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool)`

GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBinaryTransfer

`func (o *GeneralizedTimeAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool)`

SetRequireBinaryTransfer sets RequireBinaryTransfer field to given value.

### HasRequireBinaryTransfer

`func (o *GeneralizedTimeAttributeSyntaxResponse) HasRequireBinaryTransfer() bool`

HasRequireBinaryTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


