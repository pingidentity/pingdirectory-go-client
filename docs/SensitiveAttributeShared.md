# SensitiveAttributeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumsensitiveAttributeSchemaUrn**](EnumsensitiveAttributeSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Sensitive Attribute | [optional] 
**AttributeType** | **[]string** |  | 
**IncludeDefaultSensitiveOperationalAttributes** | Pointer to **bool** | Indicates whether to automatically include any server-generated operational attributes that may contain sensitive data. | [optional] 
**AllowInReturnedEntries** | Pointer to [**EnumsensitiveAttributeAllowInReturnedEntriesProp**](EnumsensitiveAttributeAllowInReturnedEntriesProp.md) |  | [optional] 
**AllowInFilter** | Pointer to [**EnumsensitiveAttributeAllowInFilterProp**](EnumsensitiveAttributeAllowInFilterProp.md) |  | [optional] 
**AllowInAdd** | Pointer to [**EnumsensitiveAttributeAllowInAddProp**](EnumsensitiveAttributeAllowInAddProp.md) |  | [optional] 
**AllowInCompare** | Pointer to [**EnumsensitiveAttributeAllowInCompareProp**](EnumsensitiveAttributeAllowInCompareProp.md) |  | [optional] 
**AllowInModify** | Pointer to [**EnumsensitiveAttributeAllowInModifyProp**](EnumsensitiveAttributeAllowInModifyProp.md) |  | [optional] 

## Methods

### NewSensitiveAttributeShared

`func NewSensitiveAttributeShared(attributeType []string, ) *SensitiveAttributeShared`

NewSensitiveAttributeShared instantiates a new SensitiveAttributeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensitiveAttributeSharedWithDefaults

`func NewSensitiveAttributeSharedWithDefaults() *SensitiveAttributeShared`

NewSensitiveAttributeSharedWithDefaults instantiates a new SensitiveAttributeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SensitiveAttributeShared) GetSchemas() []EnumsensitiveAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SensitiveAttributeShared) GetSchemasOk() (*[]EnumsensitiveAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SensitiveAttributeShared) SetSchemas(v []EnumsensitiveAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *SensitiveAttributeShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *SensitiveAttributeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SensitiveAttributeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SensitiveAttributeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SensitiveAttributeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeType

`func (o *SensitiveAttributeShared) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *SensitiveAttributeShared) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *SensitiveAttributeShared) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetIncludeDefaultSensitiveOperationalAttributes

`func (o *SensitiveAttributeShared) GetIncludeDefaultSensitiveOperationalAttributes() bool`

GetIncludeDefaultSensitiveOperationalAttributes returns the IncludeDefaultSensitiveOperationalAttributes field if non-nil, zero value otherwise.

### GetIncludeDefaultSensitiveOperationalAttributesOk

`func (o *SensitiveAttributeShared) GetIncludeDefaultSensitiveOperationalAttributesOk() (*bool, bool)`

GetIncludeDefaultSensitiveOperationalAttributesOk returns a tuple with the IncludeDefaultSensitiveOperationalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDefaultSensitiveOperationalAttributes

`func (o *SensitiveAttributeShared) SetIncludeDefaultSensitiveOperationalAttributes(v bool)`

SetIncludeDefaultSensitiveOperationalAttributes sets IncludeDefaultSensitiveOperationalAttributes field to given value.

### HasIncludeDefaultSensitiveOperationalAttributes

`func (o *SensitiveAttributeShared) HasIncludeDefaultSensitiveOperationalAttributes() bool`

HasIncludeDefaultSensitiveOperationalAttributes returns a boolean if a field has been set.

### GetAllowInReturnedEntries

`func (o *SensitiveAttributeShared) GetAllowInReturnedEntries() EnumsensitiveAttributeAllowInReturnedEntriesProp`

GetAllowInReturnedEntries returns the AllowInReturnedEntries field if non-nil, zero value otherwise.

### GetAllowInReturnedEntriesOk

`func (o *SensitiveAttributeShared) GetAllowInReturnedEntriesOk() (*EnumsensitiveAttributeAllowInReturnedEntriesProp, bool)`

GetAllowInReturnedEntriesOk returns a tuple with the AllowInReturnedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInReturnedEntries

`func (o *SensitiveAttributeShared) SetAllowInReturnedEntries(v EnumsensitiveAttributeAllowInReturnedEntriesProp)`

SetAllowInReturnedEntries sets AllowInReturnedEntries field to given value.

### HasAllowInReturnedEntries

`func (o *SensitiveAttributeShared) HasAllowInReturnedEntries() bool`

HasAllowInReturnedEntries returns a boolean if a field has been set.

### GetAllowInFilter

`func (o *SensitiveAttributeShared) GetAllowInFilter() EnumsensitiveAttributeAllowInFilterProp`

GetAllowInFilter returns the AllowInFilter field if non-nil, zero value otherwise.

### GetAllowInFilterOk

`func (o *SensitiveAttributeShared) GetAllowInFilterOk() (*EnumsensitiveAttributeAllowInFilterProp, bool)`

GetAllowInFilterOk returns a tuple with the AllowInFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInFilter

`func (o *SensitiveAttributeShared) SetAllowInFilter(v EnumsensitiveAttributeAllowInFilterProp)`

SetAllowInFilter sets AllowInFilter field to given value.

### HasAllowInFilter

`func (o *SensitiveAttributeShared) HasAllowInFilter() bool`

HasAllowInFilter returns a boolean if a field has been set.

### GetAllowInAdd

`func (o *SensitiveAttributeShared) GetAllowInAdd() EnumsensitiveAttributeAllowInAddProp`

GetAllowInAdd returns the AllowInAdd field if non-nil, zero value otherwise.

### GetAllowInAddOk

`func (o *SensitiveAttributeShared) GetAllowInAddOk() (*EnumsensitiveAttributeAllowInAddProp, bool)`

GetAllowInAddOk returns a tuple with the AllowInAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInAdd

`func (o *SensitiveAttributeShared) SetAllowInAdd(v EnumsensitiveAttributeAllowInAddProp)`

SetAllowInAdd sets AllowInAdd field to given value.

### HasAllowInAdd

`func (o *SensitiveAttributeShared) HasAllowInAdd() bool`

HasAllowInAdd returns a boolean if a field has been set.

### GetAllowInCompare

`func (o *SensitiveAttributeShared) GetAllowInCompare() EnumsensitiveAttributeAllowInCompareProp`

GetAllowInCompare returns the AllowInCompare field if non-nil, zero value otherwise.

### GetAllowInCompareOk

`func (o *SensitiveAttributeShared) GetAllowInCompareOk() (*EnumsensitiveAttributeAllowInCompareProp, bool)`

GetAllowInCompareOk returns a tuple with the AllowInCompare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInCompare

`func (o *SensitiveAttributeShared) SetAllowInCompare(v EnumsensitiveAttributeAllowInCompareProp)`

SetAllowInCompare sets AllowInCompare field to given value.

### HasAllowInCompare

`func (o *SensitiveAttributeShared) HasAllowInCompare() bool`

HasAllowInCompare returns a boolean if a field has been set.

### GetAllowInModify

`func (o *SensitiveAttributeShared) GetAllowInModify() EnumsensitiveAttributeAllowInModifyProp`

GetAllowInModify returns the AllowInModify field if non-nil, zero value otherwise.

### GetAllowInModifyOk

`func (o *SensitiveAttributeShared) GetAllowInModifyOk() (*EnumsensitiveAttributeAllowInModifyProp, bool)`

GetAllowInModifyOk returns a tuple with the AllowInModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInModify

`func (o *SensitiveAttributeShared) SetAllowInModify(v EnumsensitiveAttributeAllowInModifyProp)`

SetAllowInModify sets AllowInModify field to given value.

### HasAllowInModify

`func (o *SensitiveAttributeShared) HasAllowInModify() bool`

HasAllowInModify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


