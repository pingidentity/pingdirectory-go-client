# AddSensitiveAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | **string** | Name of the new Sensitive Attribute | 
**Schemas** | Pointer to [**[]EnumsensitiveAttributeSchemaUrn**](EnumsensitiveAttributeSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Sensitive Attribute | [optional] 
**AttributeType** | **[]string** | The name(s) or OID(s) of the attribute types for attributes whose values may be considered sensitive. | 
**IncludeDefaultSensitiveOperationalAttributes** | Pointer to **bool** | Indicates whether to automatically include any server-generated operational attributes that may contain sensitive data. | [optional] 
**AllowInReturnedEntries** | Pointer to [**EnumsensitiveAttributeAllowInReturnedEntriesProp**](EnumsensitiveAttributeAllowInReturnedEntriesProp.md) |  | [optional] 
**AllowInFilter** | Pointer to [**EnumsensitiveAttributeAllowInFilterProp**](EnumsensitiveAttributeAllowInFilterProp.md) |  | [optional] 
**AllowInAdd** | Pointer to [**EnumsensitiveAttributeAllowInAddProp**](EnumsensitiveAttributeAllowInAddProp.md) |  | [optional] 
**AllowInCompare** | Pointer to [**EnumsensitiveAttributeAllowInCompareProp**](EnumsensitiveAttributeAllowInCompareProp.md) |  | [optional] 
**AllowInModify** | Pointer to [**EnumsensitiveAttributeAllowInModifyProp**](EnumsensitiveAttributeAllowInModifyProp.md) |  | [optional] 

## Methods

### NewAddSensitiveAttributeRequest

`func NewAddSensitiveAttributeRequest(attributeName string, attributeType []string, ) *AddSensitiveAttributeRequest`

NewAddSensitiveAttributeRequest instantiates a new AddSensitiveAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSensitiveAttributeRequestWithDefaults

`func NewAddSensitiveAttributeRequestWithDefaults() *AddSensitiveAttributeRequest`

NewAddSensitiveAttributeRequestWithDefaults instantiates a new AddSensitiveAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *AddSensitiveAttributeRequest) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *AddSensitiveAttributeRequest) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *AddSensitiveAttributeRequest) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetSchemas

`func (o *AddSensitiveAttributeRequest) GetSchemas() []EnumsensitiveAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSensitiveAttributeRequest) GetSchemasOk() (*[]EnumsensitiveAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSensitiveAttributeRequest) SetSchemas(v []EnumsensitiveAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddSensitiveAttributeRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddSensitiveAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSensitiveAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSensitiveAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSensitiveAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeType

`func (o *AddSensitiveAttributeRequest) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddSensitiveAttributeRequest) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddSensitiveAttributeRequest) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetIncludeDefaultSensitiveOperationalAttributes

`func (o *AddSensitiveAttributeRequest) GetIncludeDefaultSensitiveOperationalAttributes() bool`

GetIncludeDefaultSensitiveOperationalAttributes returns the IncludeDefaultSensitiveOperationalAttributes field if non-nil, zero value otherwise.

### GetIncludeDefaultSensitiveOperationalAttributesOk

`func (o *AddSensitiveAttributeRequest) GetIncludeDefaultSensitiveOperationalAttributesOk() (*bool, bool)`

GetIncludeDefaultSensitiveOperationalAttributesOk returns a tuple with the IncludeDefaultSensitiveOperationalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDefaultSensitiveOperationalAttributes

`func (o *AddSensitiveAttributeRequest) SetIncludeDefaultSensitiveOperationalAttributes(v bool)`

SetIncludeDefaultSensitiveOperationalAttributes sets IncludeDefaultSensitiveOperationalAttributes field to given value.

### HasIncludeDefaultSensitiveOperationalAttributes

`func (o *AddSensitiveAttributeRequest) HasIncludeDefaultSensitiveOperationalAttributes() bool`

HasIncludeDefaultSensitiveOperationalAttributes returns a boolean if a field has been set.

### GetAllowInReturnedEntries

`func (o *AddSensitiveAttributeRequest) GetAllowInReturnedEntries() EnumsensitiveAttributeAllowInReturnedEntriesProp`

GetAllowInReturnedEntries returns the AllowInReturnedEntries field if non-nil, zero value otherwise.

### GetAllowInReturnedEntriesOk

`func (o *AddSensitiveAttributeRequest) GetAllowInReturnedEntriesOk() (*EnumsensitiveAttributeAllowInReturnedEntriesProp, bool)`

GetAllowInReturnedEntriesOk returns a tuple with the AllowInReturnedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInReturnedEntries

`func (o *AddSensitiveAttributeRequest) SetAllowInReturnedEntries(v EnumsensitiveAttributeAllowInReturnedEntriesProp)`

SetAllowInReturnedEntries sets AllowInReturnedEntries field to given value.

### HasAllowInReturnedEntries

`func (o *AddSensitiveAttributeRequest) HasAllowInReturnedEntries() bool`

HasAllowInReturnedEntries returns a boolean if a field has been set.

### GetAllowInFilter

`func (o *AddSensitiveAttributeRequest) GetAllowInFilter() EnumsensitiveAttributeAllowInFilterProp`

GetAllowInFilter returns the AllowInFilter field if non-nil, zero value otherwise.

### GetAllowInFilterOk

`func (o *AddSensitiveAttributeRequest) GetAllowInFilterOk() (*EnumsensitiveAttributeAllowInFilterProp, bool)`

GetAllowInFilterOk returns a tuple with the AllowInFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInFilter

`func (o *AddSensitiveAttributeRequest) SetAllowInFilter(v EnumsensitiveAttributeAllowInFilterProp)`

SetAllowInFilter sets AllowInFilter field to given value.

### HasAllowInFilter

`func (o *AddSensitiveAttributeRequest) HasAllowInFilter() bool`

HasAllowInFilter returns a boolean if a field has been set.

### GetAllowInAdd

`func (o *AddSensitiveAttributeRequest) GetAllowInAdd() EnumsensitiveAttributeAllowInAddProp`

GetAllowInAdd returns the AllowInAdd field if non-nil, zero value otherwise.

### GetAllowInAddOk

`func (o *AddSensitiveAttributeRequest) GetAllowInAddOk() (*EnumsensitiveAttributeAllowInAddProp, bool)`

GetAllowInAddOk returns a tuple with the AllowInAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInAdd

`func (o *AddSensitiveAttributeRequest) SetAllowInAdd(v EnumsensitiveAttributeAllowInAddProp)`

SetAllowInAdd sets AllowInAdd field to given value.

### HasAllowInAdd

`func (o *AddSensitiveAttributeRequest) HasAllowInAdd() bool`

HasAllowInAdd returns a boolean if a field has been set.

### GetAllowInCompare

`func (o *AddSensitiveAttributeRequest) GetAllowInCompare() EnumsensitiveAttributeAllowInCompareProp`

GetAllowInCompare returns the AllowInCompare field if non-nil, zero value otherwise.

### GetAllowInCompareOk

`func (o *AddSensitiveAttributeRequest) GetAllowInCompareOk() (*EnumsensitiveAttributeAllowInCompareProp, bool)`

GetAllowInCompareOk returns a tuple with the AllowInCompare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInCompare

`func (o *AddSensitiveAttributeRequest) SetAllowInCompare(v EnumsensitiveAttributeAllowInCompareProp)`

SetAllowInCompare sets AllowInCompare field to given value.

### HasAllowInCompare

`func (o *AddSensitiveAttributeRequest) HasAllowInCompare() bool`

HasAllowInCompare returns a boolean if a field has been set.

### GetAllowInModify

`func (o *AddSensitiveAttributeRequest) GetAllowInModify() EnumsensitiveAttributeAllowInModifyProp`

GetAllowInModify returns the AllowInModify field if non-nil, zero value otherwise.

### GetAllowInModifyOk

`func (o *AddSensitiveAttributeRequest) GetAllowInModifyOk() (*EnumsensitiveAttributeAllowInModifyProp, bool)`

GetAllowInModifyOk returns a tuple with the AllowInModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInModify

`func (o *AddSensitiveAttributeRequest) SetAllowInModify(v EnumsensitiveAttributeAllowInModifyProp)`

SetAllowInModify sets AllowInModify field to given value.

### HasAllowInModify

`func (o *AddSensitiveAttributeRequest) HasAllowInModify() bool`

HasAllowInModify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


