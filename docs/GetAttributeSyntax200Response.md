# GetAttributeSyntax200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumnameAndOptionalUidAttributeSyntaxSchemaUrn**](EnumnameAndOptionalUidAttributeSyntaxSchemaUrn.md) |  | 
**Id** | **string** | Name of the Attribute Syntax | 
**StripSyntaxMinUpperBound** | Pointer to **bool** | Indicates whether the suggested minimum upper bound appended to an attribute&#39;s syntax OID in its schema definition Attribute Type Description should be stripped. | [optional] 
**Enabled** | **bool** | Indicates whether the Attribute Syntax is enabled. | 
**RequireBinaryTransfer** | Pointer to **bool** | Indicates whether values of this attribute are required to have a \&quot;binary\&quot; transfer option as described in RFC 4522. Attributes with this syntax will generally be referenced with names including \&quot;;binary\&quot; (e.g., \&quot;userCertificate;binary\&quot;). | [optional] 
**AllowZeroLengthValues** | Pointer to **bool** | Indicates whether zero-length (that is, an empty string) values are allowed. | [optional] 
**StrictFormat** | Pointer to **bool** | Indicates whether values for attributes with this syntax will be required to be in the valid LDAP URL format. If this is set to false, then arbitrary strings will be allowed. | [optional] 
**EnableCompaction** | Pointer to **bool** | Indicates whether values of attributes with this syntax should be compacted when stored in a local DB database. | [optional] 
**IncludeAttributeInCompaction** | Pointer to **[]string** | Specifies the specific attributes (which should be associated with this syntax) whose values should be compacted. If one or more include attributes are specified, then only those attributes will have their values compacted. If not set then all attributes will have their values compacted. The exclude-attribute-from-compaction property takes precedence over this property. | [optional] 
**ExcludeAttributeFromCompaction** | Pointer to **[]string** | Specifies the specific attributes (which should be associated with this syntax) whose values should not be compacted. If one or more exclude attributes are specified, then values of those attributes will not have their values compacted. This property takes precedence over the include-attribute-in-compaction property. | [optional] 

## Methods

### NewGetAttributeSyntax200Response

`func NewGetAttributeSyntax200Response(schemas []EnumnameAndOptionalUidAttributeSyntaxSchemaUrn, id string, enabled bool, ) *GetAttributeSyntax200Response`

NewGetAttributeSyntax200Response instantiates a new GetAttributeSyntax200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAttributeSyntax200ResponseWithDefaults

`func NewGetAttributeSyntax200ResponseWithDefaults() *GetAttributeSyntax200Response`

NewGetAttributeSyntax200ResponseWithDefaults instantiates a new GetAttributeSyntax200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetAttributeSyntax200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAttributeSyntax200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAttributeSyntax200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetAttributeSyntax200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetAttributeSyntax200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetAttributeSyntax200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetAttributeSyntax200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetAttributeSyntax200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *GetAttributeSyntax200Response) GetSchemas() []EnumnameAndOptionalUidAttributeSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetAttributeSyntax200Response) GetSchemasOk() (*[]EnumnameAndOptionalUidAttributeSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetAttributeSyntax200Response) SetSchemas(v []EnumnameAndOptionalUidAttributeSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetAttributeSyntax200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAttributeSyntax200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAttributeSyntax200Response) SetId(v string)`

SetId sets Id field to given value.


### GetStripSyntaxMinUpperBound

`func (o *GetAttributeSyntax200Response) GetStripSyntaxMinUpperBound() bool`

GetStripSyntaxMinUpperBound returns the StripSyntaxMinUpperBound field if non-nil, zero value otherwise.

### GetStripSyntaxMinUpperBoundOk

`func (o *GetAttributeSyntax200Response) GetStripSyntaxMinUpperBoundOk() (*bool, bool)`

GetStripSyntaxMinUpperBoundOk returns a tuple with the StripSyntaxMinUpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripSyntaxMinUpperBound

`func (o *GetAttributeSyntax200Response) SetStripSyntaxMinUpperBound(v bool)`

SetStripSyntaxMinUpperBound sets StripSyntaxMinUpperBound field to given value.

### HasStripSyntaxMinUpperBound

`func (o *GetAttributeSyntax200Response) HasStripSyntaxMinUpperBound() bool`

HasStripSyntaxMinUpperBound returns a boolean if a field has been set.

### GetEnabled

`func (o *GetAttributeSyntax200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetAttributeSyntax200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetAttributeSyntax200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireBinaryTransfer

`func (o *GetAttributeSyntax200Response) GetRequireBinaryTransfer() bool`

GetRequireBinaryTransfer returns the RequireBinaryTransfer field if non-nil, zero value otherwise.

### GetRequireBinaryTransferOk

`func (o *GetAttributeSyntax200Response) GetRequireBinaryTransferOk() (*bool, bool)`

GetRequireBinaryTransferOk returns a tuple with the RequireBinaryTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBinaryTransfer

`func (o *GetAttributeSyntax200Response) SetRequireBinaryTransfer(v bool)`

SetRequireBinaryTransfer sets RequireBinaryTransfer field to given value.

### HasRequireBinaryTransfer

`func (o *GetAttributeSyntax200Response) HasRequireBinaryTransfer() bool`

HasRequireBinaryTransfer returns a boolean if a field has been set.

### GetAllowZeroLengthValues

`func (o *GetAttributeSyntax200Response) GetAllowZeroLengthValues() bool`

GetAllowZeroLengthValues returns the AllowZeroLengthValues field if non-nil, zero value otherwise.

### GetAllowZeroLengthValuesOk

`func (o *GetAttributeSyntax200Response) GetAllowZeroLengthValuesOk() (*bool, bool)`

GetAllowZeroLengthValuesOk returns a tuple with the AllowZeroLengthValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowZeroLengthValues

`func (o *GetAttributeSyntax200Response) SetAllowZeroLengthValues(v bool)`

SetAllowZeroLengthValues sets AllowZeroLengthValues field to given value.

### HasAllowZeroLengthValues

`func (o *GetAttributeSyntax200Response) HasAllowZeroLengthValues() bool`

HasAllowZeroLengthValues returns a boolean if a field has been set.

### GetStrictFormat

`func (o *GetAttributeSyntax200Response) GetStrictFormat() bool`

GetStrictFormat returns the StrictFormat field if non-nil, zero value otherwise.

### GetStrictFormatOk

`func (o *GetAttributeSyntax200Response) GetStrictFormatOk() (*bool, bool)`

GetStrictFormatOk returns a tuple with the StrictFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictFormat

`func (o *GetAttributeSyntax200Response) SetStrictFormat(v bool)`

SetStrictFormat sets StrictFormat field to given value.

### HasStrictFormat

`func (o *GetAttributeSyntax200Response) HasStrictFormat() bool`

HasStrictFormat returns a boolean if a field has been set.

### GetEnableCompaction

`func (o *GetAttributeSyntax200Response) GetEnableCompaction() bool`

GetEnableCompaction returns the EnableCompaction field if non-nil, zero value otherwise.

### GetEnableCompactionOk

`func (o *GetAttributeSyntax200Response) GetEnableCompactionOk() (*bool, bool)`

GetEnableCompactionOk returns a tuple with the EnableCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCompaction

`func (o *GetAttributeSyntax200Response) SetEnableCompaction(v bool)`

SetEnableCompaction sets EnableCompaction field to given value.

### HasEnableCompaction

`func (o *GetAttributeSyntax200Response) HasEnableCompaction() bool`

HasEnableCompaction returns a boolean if a field has been set.

### GetIncludeAttributeInCompaction

`func (o *GetAttributeSyntax200Response) GetIncludeAttributeInCompaction() []string`

GetIncludeAttributeInCompaction returns the IncludeAttributeInCompaction field if non-nil, zero value otherwise.

### GetIncludeAttributeInCompactionOk

`func (o *GetAttributeSyntax200Response) GetIncludeAttributeInCompactionOk() (*[]string, bool)`

GetIncludeAttributeInCompactionOk returns a tuple with the IncludeAttributeInCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttributeInCompaction

`func (o *GetAttributeSyntax200Response) SetIncludeAttributeInCompaction(v []string)`

SetIncludeAttributeInCompaction sets IncludeAttributeInCompaction field to given value.

### HasIncludeAttributeInCompaction

`func (o *GetAttributeSyntax200Response) HasIncludeAttributeInCompaction() bool`

HasIncludeAttributeInCompaction returns a boolean if a field has been set.

### GetExcludeAttributeFromCompaction

`func (o *GetAttributeSyntax200Response) GetExcludeAttributeFromCompaction() []string`

GetExcludeAttributeFromCompaction returns the ExcludeAttributeFromCompaction field if non-nil, zero value otherwise.

### GetExcludeAttributeFromCompactionOk

`func (o *GetAttributeSyntax200Response) GetExcludeAttributeFromCompactionOk() (*[]string, bool)`

GetExcludeAttributeFromCompactionOk returns a tuple with the ExcludeAttributeFromCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttributeFromCompaction

`func (o *GetAttributeSyntax200Response) SetExcludeAttributeFromCompaction(v []string)`

SetExcludeAttributeFromCompaction sets ExcludeAttributeFromCompaction field to given value.

### HasExcludeAttributeFromCompaction

`func (o *GetAttributeSyntax200Response) HasExcludeAttributeFromCompaction() bool`

HasExcludeAttributeFromCompaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


