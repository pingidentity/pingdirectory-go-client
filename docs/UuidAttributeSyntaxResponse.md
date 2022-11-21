# UuidAttributeSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumuuidAttributeSyntaxSchemaUrn**](EnumuuidAttributeSyntaxSchemaUrn.md) |  | 
**EnableCompaction** | Pointer to **bool** | Indicates whether values of attributes with this syntax should be compacted when stored in a local DB database. | [optional] 
**IncludeAttributeInCompaction** | Pointer to **[]string** | Specifies the specific attributes (which should be associated with this syntax) whose values should be compacted. If one or more include attributes are specified, then only those attributes will have their values compacted. If not set then all attributes will have their values compacted. The exclude-attribute-from-compaction property takes precedence over this property. | [optional] 
**ExcludeAttributeFromCompaction** | Pointer to **[]string** | Specifies the specific attributes (which should be associated with this syntax) whose values should not be compacted. If one or more exclude attributes are specified, then values of those attributes will not have their values compacted. This property takes precedence over the include-attribute-in-compaction property. | [optional] 
**Enabled** | **bool** | Indicates whether the Attribute Syntax is enabled. | 
**RequireBinaryTransfer** | Pointer to **bool** | Indicates whether values of this attribute are required to have a \&quot;binary\&quot; transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \&quot;;binary\&quot; (e.g., \&quot;userCertificate;binary\&quot;). | [optional] 

## Methods

### NewUuidAttributeSyntaxResponse

`func NewUuidAttributeSyntaxResponse(schemas []EnumuuidAttributeSyntaxSchemaUrn, enabled bool, ) *UuidAttributeSyntaxResponse`

NewUuidAttributeSyntaxResponse instantiates a new UuidAttributeSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidAttributeSyntaxResponseWithDefaults

`func NewUuidAttributeSyntaxResponseWithDefaults() *UuidAttributeSyntaxResponse`

NewUuidAttributeSyntaxResponseWithDefaults instantiates a new UuidAttributeSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *UuidAttributeSyntaxResponse) GetSchemas() []EnumuuidAttributeSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UuidAttributeSyntaxResponse) GetSchemasOk() (*[]EnumuuidAttributeSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UuidAttributeSyntaxResponse) SetSchemas(v []EnumuuidAttributeSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnableCompaction

`func (o *UuidAttributeSyntaxResponse) GetEnableCompaction() bool`

GetEnableCompaction returns the EnableCompaction field if non-nil, zero value otherwise.

### GetEnableCompactionOk

`func (o *UuidAttributeSyntaxResponse) GetEnableCompactionOk() (*bool, bool)`

GetEnableCompactionOk returns a tuple with the EnableCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCompaction

`func (o *UuidAttributeSyntaxResponse) SetEnableCompaction(v bool)`

SetEnableCompaction sets EnableCompaction field to given value.

### HasEnableCompaction

`func (o *UuidAttributeSyntaxResponse) HasEnableCompaction() bool`

HasEnableCompaction returns a boolean if a field has been set.

### GetIncludeAttributeInCompaction

`func (o *UuidAttributeSyntaxResponse) GetIncludeAttributeInCompaction() []string`

GetIncludeAttributeInCompaction returns the IncludeAttributeInCompaction field if non-nil, zero value otherwise.

### GetIncludeAttributeInCompactionOk

`func (o *UuidAttributeSyntaxResponse) GetIncludeAttributeInCompactionOk() (*[]string, bool)`

GetIncludeAttributeInCompactionOk returns a tuple with the IncludeAttributeInCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttributeInCompaction

`func (o *UuidAttributeSyntaxResponse) SetIncludeAttributeInCompaction(v []string)`

SetIncludeAttributeInCompaction sets IncludeAttributeInCompaction field to given value.

### HasIncludeAttributeInCompaction

`func (o *UuidAttributeSyntaxResponse) HasIncludeAttributeInCompaction() bool`

HasIncludeAttributeInCompaction returns a boolean if a field has been set.

### GetExcludeAttributeFromCompaction

`func (o *UuidAttributeSyntaxResponse) GetExcludeAttributeFromCompaction() []string`

GetExcludeAttributeFromCompaction returns the ExcludeAttributeFromCompaction field if non-nil, zero value otherwise.

### GetExcludeAttributeFromCompactionOk

`func (o *UuidAttributeSyntaxResponse) GetExcludeAttributeFromCompactionOk() (*[]string, bool)`

GetExcludeAttributeFromCompactionOk returns a tuple with the ExcludeAttributeFromCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttributeFromCompaction

`func (o *UuidAttributeSyntaxResponse) SetExcludeAttributeFromCompaction(v []string)`

SetExcludeAttributeFromCompaction sets ExcludeAttributeFromCompaction field to given value.

### HasExcludeAttributeFromCompaction

`func (o *UuidAttributeSyntaxResponse) HasExcludeAttributeFromCompaction() bool`

HasExcludeAttributeFromCompaction returns a boolean if a field has been set.

### GetEnabled

`func (o *UuidAttributeSyntaxResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UuidAttributeSyntaxResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UuidAttributeSyntaxResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireBinaryTransfer

`func (o *UuidAttributeSyntaxResponse) GetRequireBinaryTransfer() bool`

GetRequireBinaryTransfer returns the RequireBinaryTransfer field if non-nil, zero value otherwise.

### GetRequireBinaryTransferOk

`func (o *UuidAttributeSyntaxResponse) GetRequireBinaryTransferOk() (*bool, bool)`

GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBinaryTransfer

`func (o *UuidAttributeSyntaxResponse) SetRequireBinaryTransfer(v bool)`

SetRequireBinaryTransfer sets RequireBinaryTransfer field to given value.

### HasRequireBinaryTransfer

`func (o *UuidAttributeSyntaxResponse) HasRequireBinaryTransfer() bool`

HasRequireBinaryTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


