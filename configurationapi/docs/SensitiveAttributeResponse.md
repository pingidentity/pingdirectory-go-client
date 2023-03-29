# SensitiveAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Sensitive Attribute | 
**Schemas** | Pointer to [**[]EnumsensitiveAttributeSchemaUrn**](EnumsensitiveAttributeSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Sensitive Attribute | [optional] 
**AttributeType** | **[]string** | The name(s) or OID(s) of the attribute types for attributes whose values may be considered sensitive. | 
**IncludeDefaultSensitiveOperationalAttributes** | Pointer to **bool** | Indicates whether to automatically include any server-generated operational attributes that may contain sensitive data. | [optional] 
**AllowInReturnedEntries** | Pointer to [**EnumsensitiveAttributeAllowInReturnedEntriesProp**](EnumsensitiveAttributeAllowInReturnedEntriesProp.md) |  | [optional] 
**AllowInFilter** | Pointer to [**EnumsensitiveAttributeAllowInFilterProp**](EnumsensitiveAttributeAllowInFilterProp.md) |  | [optional] 
**AllowInAdd** | Pointer to [**EnumsensitiveAttributeAllowInAddProp**](EnumsensitiveAttributeAllowInAddProp.md) |  | [optional] 
**AllowInCompare** | Pointer to [**EnumsensitiveAttributeAllowInCompareProp**](EnumsensitiveAttributeAllowInCompareProp.md) |  | [optional] 
**AllowInModify** | Pointer to [**EnumsensitiveAttributeAllowInModifyProp**](EnumsensitiveAttributeAllowInModifyProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewSensitiveAttributeResponse

`func NewSensitiveAttributeResponse(id string, attributeType []string, ) *SensitiveAttributeResponse`

NewSensitiveAttributeResponse instantiates a new SensitiveAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensitiveAttributeResponseWithDefaults

`func NewSensitiveAttributeResponseWithDefaults() *SensitiveAttributeResponse`

NewSensitiveAttributeResponseWithDefaults instantiates a new SensitiveAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SensitiveAttributeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SensitiveAttributeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SensitiveAttributeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SensitiveAttributeResponse) GetSchemas() []EnumsensitiveAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SensitiveAttributeResponse) GetSchemasOk() (*[]EnumsensitiveAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SensitiveAttributeResponse) SetSchemas(v []EnumsensitiveAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *SensitiveAttributeResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *SensitiveAttributeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SensitiveAttributeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SensitiveAttributeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SensitiveAttributeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeType

`func (o *SensitiveAttributeResponse) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *SensitiveAttributeResponse) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *SensitiveAttributeResponse) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetIncludeDefaultSensitiveOperationalAttributes

`func (o *SensitiveAttributeResponse) GetIncludeDefaultSensitiveOperationalAttributes() bool`

GetIncludeDefaultSensitiveOperationalAttributes returns the IncludeDefaultSensitiveOperationalAttributes field if non-nil, zero value otherwise.

### GetIncludeDefaultSensitiveOperationalAttributesOk

`func (o *SensitiveAttributeResponse) GetIncludeDefaultSensitiveOperationalAttributesOk() (*bool, bool)`

GetIncludeDefaultSensitiveOperationalAttributesOk returns a tuple with the IncludeDefaultSensitiveOperationalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDefaultSensitiveOperationalAttributes

`func (o *SensitiveAttributeResponse) SetIncludeDefaultSensitiveOperationalAttributes(v bool)`

SetIncludeDefaultSensitiveOperationalAttributes sets IncludeDefaultSensitiveOperationalAttributes field to given value.

### HasIncludeDefaultSensitiveOperationalAttributes

`func (o *SensitiveAttributeResponse) HasIncludeDefaultSensitiveOperationalAttributes() bool`

HasIncludeDefaultSensitiveOperationalAttributes returns a boolean if a field has been set.

### GetAllowInReturnedEntries

`func (o *SensitiveAttributeResponse) GetAllowInReturnedEntries() EnumsensitiveAttributeAllowInReturnedEntriesProp`

GetAllowInReturnedEntries returns the AllowInReturnedEntries field if non-nil, zero value otherwise.

### GetAllowInReturnedEntriesOk

`func (o *SensitiveAttributeResponse) GetAllowInReturnedEntriesOk() (*EnumsensitiveAttributeAllowInReturnedEntriesProp, bool)`

GetAllowInReturnedEntriesOk returns a tuple with the AllowInReturnedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInReturnedEntries

`func (o *SensitiveAttributeResponse) SetAllowInReturnedEntries(v EnumsensitiveAttributeAllowInReturnedEntriesProp)`

SetAllowInReturnedEntries sets AllowInReturnedEntries field to given value.

### HasAllowInReturnedEntries

`func (o *SensitiveAttributeResponse) HasAllowInReturnedEntries() bool`

HasAllowInReturnedEntries returns a boolean if a field has been set.

### GetAllowInFilter

`func (o *SensitiveAttributeResponse) GetAllowInFilter() EnumsensitiveAttributeAllowInFilterProp`

GetAllowInFilter returns the AllowInFilter field if non-nil, zero value otherwise.

### GetAllowInFilterOk

`func (o *SensitiveAttributeResponse) GetAllowInFilterOk() (*EnumsensitiveAttributeAllowInFilterProp, bool)`

GetAllowInFilterOk returns a tuple with the AllowInFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInFilter

`func (o *SensitiveAttributeResponse) SetAllowInFilter(v EnumsensitiveAttributeAllowInFilterProp)`

SetAllowInFilter sets AllowInFilter field to given value.

### HasAllowInFilter

`func (o *SensitiveAttributeResponse) HasAllowInFilter() bool`

HasAllowInFilter returns a boolean if a field has been set.

### GetAllowInAdd

`func (o *SensitiveAttributeResponse) GetAllowInAdd() EnumsensitiveAttributeAllowInAddProp`

GetAllowInAdd returns the AllowInAdd field if non-nil, zero value otherwise.

### GetAllowInAddOk

`func (o *SensitiveAttributeResponse) GetAllowInAddOk() (*EnumsensitiveAttributeAllowInAddProp, bool)`

GetAllowInAddOk returns a tuple with the AllowInAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInAdd

`func (o *SensitiveAttributeResponse) SetAllowInAdd(v EnumsensitiveAttributeAllowInAddProp)`

SetAllowInAdd sets AllowInAdd field to given value.

### HasAllowInAdd

`func (o *SensitiveAttributeResponse) HasAllowInAdd() bool`

HasAllowInAdd returns a boolean if a field has been set.

### GetAllowInCompare

`func (o *SensitiveAttributeResponse) GetAllowInCompare() EnumsensitiveAttributeAllowInCompareProp`

GetAllowInCompare returns the AllowInCompare field if non-nil, zero value otherwise.

### GetAllowInCompareOk

`func (o *SensitiveAttributeResponse) GetAllowInCompareOk() (*EnumsensitiveAttributeAllowInCompareProp, bool)`

GetAllowInCompareOk returns a tuple with the AllowInCompare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInCompare

`func (o *SensitiveAttributeResponse) SetAllowInCompare(v EnumsensitiveAttributeAllowInCompareProp)`

SetAllowInCompare sets AllowInCompare field to given value.

### HasAllowInCompare

`func (o *SensitiveAttributeResponse) HasAllowInCompare() bool`

HasAllowInCompare returns a boolean if a field has been set.

### GetAllowInModify

`func (o *SensitiveAttributeResponse) GetAllowInModify() EnumsensitiveAttributeAllowInModifyProp`

GetAllowInModify returns the AllowInModify field if non-nil, zero value otherwise.

### GetAllowInModifyOk

`func (o *SensitiveAttributeResponse) GetAllowInModifyOk() (*EnumsensitiveAttributeAllowInModifyProp, bool)`

GetAllowInModifyOk returns a tuple with the AllowInModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInModify

`func (o *SensitiveAttributeResponse) SetAllowInModify(v EnumsensitiveAttributeAllowInModifyProp)`

SetAllowInModify sets AllowInModify field to given value.

### HasAllowInModify

`func (o *SensitiveAttributeResponse) HasAllowInModify() bool`

HasAllowInModify returns a boolean if a field has been set.

### GetMeta

`func (o *SensitiveAttributeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SensitiveAttributeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SensitiveAttributeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SensitiveAttributeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SensitiveAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SensitiveAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SensitiveAttributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SensitiveAttributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


