# GenericDelegatedAdminAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Delegated Admin Attribute | 
**Schemas** | [**[]EnumgenericDelegatedAdminAttributeSchemaUrn**](EnumgenericDelegatedAdminAttributeSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Delegated Admin Attribute | [optional] 
**AttributeType** | **string** | Specifies the name or OID of the LDAP attribute type. | 
**DisplayName** | **string** | A human readable display name for this Delegated Admin Attribute. | 
**Mutability** | [**EnumdelegatedAdminAttributeMutabilityProp**](EnumdelegatedAdminAttributeMutabilityProp.md) |  | 
**MultiValued** | **bool** | Indicates whether this Delegated Admin Attribute may have multiple values. | 
**IncludeInSummary** | **bool** | Indicates whether this Delegated Admin Attribute is to be included in the summary display for a resource. | 
**AttributeCategory** | Pointer to **string** | Specifies which attribute category this attribute belongs to. | [optional] 
**DisplayOrderIndex** | **int64** | This property determines a display order for attributes within a given attribute category. Attributes are ordered within their category based on this index from least to greatest. | 
**ReferenceResourceType** | Pointer to **string** | For LDAP attributes with DN syntax, specifies what kind of resource is referenced. | [optional] 
**AttributePresentation** | Pointer to [**EnumdelegatedAdminAttributeAttributePresentationProp**](EnumdelegatedAdminAttributeAttributePresentationProp.md) |  | [optional] 
**DateTimeFormat** | Pointer to **string** | Specifies the format string that is used to present a date and/or time value to the user of the app. This property only applies to LDAP attribute types whose LDAP syntax is GeneralizedTime and is ignored if the attribute type has any other syntax. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewGenericDelegatedAdminAttributeResponse

`func NewGenericDelegatedAdminAttributeResponse(id string, schemas []EnumgenericDelegatedAdminAttributeSchemaUrn, attributeType string, displayName string, mutability EnumdelegatedAdminAttributeMutabilityProp, multiValued bool, includeInSummary bool, displayOrderIndex int64, ) *GenericDelegatedAdminAttributeResponse`

NewGenericDelegatedAdminAttributeResponse instantiates a new GenericDelegatedAdminAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericDelegatedAdminAttributeResponseWithDefaults

`func NewGenericDelegatedAdminAttributeResponseWithDefaults() *GenericDelegatedAdminAttributeResponse`

NewGenericDelegatedAdminAttributeResponseWithDefaults instantiates a new GenericDelegatedAdminAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenericDelegatedAdminAttributeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenericDelegatedAdminAttributeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenericDelegatedAdminAttributeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *GenericDelegatedAdminAttributeResponse) GetSchemas() []EnumgenericDelegatedAdminAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GenericDelegatedAdminAttributeResponse) GetSchemasOk() (*[]EnumgenericDelegatedAdminAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GenericDelegatedAdminAttributeResponse) SetSchemas(v []EnumgenericDelegatedAdminAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *GenericDelegatedAdminAttributeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericDelegatedAdminAttributeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericDelegatedAdminAttributeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenericDelegatedAdminAttributeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeType

`func (o *GenericDelegatedAdminAttributeResponse) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *GenericDelegatedAdminAttributeResponse) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *GenericDelegatedAdminAttributeResponse) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetDisplayName

`func (o *GenericDelegatedAdminAttributeResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GenericDelegatedAdminAttributeResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GenericDelegatedAdminAttributeResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMutability

`func (o *GenericDelegatedAdminAttributeResponse) GetMutability() EnumdelegatedAdminAttributeMutabilityProp`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *GenericDelegatedAdminAttributeResponse) GetMutabilityOk() (*EnumdelegatedAdminAttributeMutabilityProp, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *GenericDelegatedAdminAttributeResponse) SetMutability(v EnumdelegatedAdminAttributeMutabilityProp)`

SetMutability sets Mutability field to given value.


### GetMultiValued

`func (o *GenericDelegatedAdminAttributeResponse) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *GenericDelegatedAdminAttributeResponse) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *GenericDelegatedAdminAttributeResponse) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.


### GetIncludeInSummary

`func (o *GenericDelegatedAdminAttributeResponse) GetIncludeInSummary() bool`

GetIncludeInSummary returns the IncludeInSummary field if non-nil, zero value otherwise.

### GetIncludeInSummaryOk

`func (o *GenericDelegatedAdminAttributeResponse) GetIncludeInSummaryOk() (*bool, bool)`

GetIncludeInSummaryOk returns a tuple with the IncludeInSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSummary

`func (o *GenericDelegatedAdminAttributeResponse) SetIncludeInSummary(v bool)`

SetIncludeInSummary sets IncludeInSummary field to given value.


### GetAttributeCategory

`func (o *GenericDelegatedAdminAttributeResponse) GetAttributeCategory() string`

GetAttributeCategory returns the AttributeCategory field if non-nil, zero value otherwise.

### GetAttributeCategoryOk

`func (o *GenericDelegatedAdminAttributeResponse) GetAttributeCategoryOk() (*string, bool)`

GetAttributeCategoryOk returns a tuple with the AttributeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeCategory

`func (o *GenericDelegatedAdminAttributeResponse) SetAttributeCategory(v string)`

SetAttributeCategory sets AttributeCategory field to given value.

### HasAttributeCategory

`func (o *GenericDelegatedAdminAttributeResponse) HasAttributeCategory() bool`

HasAttributeCategory returns a boolean if a field has been set.

### GetDisplayOrderIndex

`func (o *GenericDelegatedAdminAttributeResponse) GetDisplayOrderIndex() int64`

GetDisplayOrderIndex returns the DisplayOrderIndex field if non-nil, zero value otherwise.

### GetDisplayOrderIndexOk

`func (o *GenericDelegatedAdminAttributeResponse) GetDisplayOrderIndexOk() (*int64, bool)`

GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrderIndex

`func (o *GenericDelegatedAdminAttributeResponse) SetDisplayOrderIndex(v int64)`

SetDisplayOrderIndex sets DisplayOrderIndex field to given value.


### GetReferenceResourceType

`func (o *GenericDelegatedAdminAttributeResponse) GetReferenceResourceType() string`

GetReferenceResourceType returns the ReferenceResourceType field if non-nil, zero value otherwise.

### GetReferenceResourceTypeOk

`func (o *GenericDelegatedAdminAttributeResponse) GetReferenceResourceTypeOk() (*string, bool)`

GetReferenceResourceTypeOk returns a tuple with the ReferenceResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceResourceType

`func (o *GenericDelegatedAdminAttributeResponse) SetReferenceResourceType(v string)`

SetReferenceResourceType sets ReferenceResourceType field to given value.

### HasReferenceResourceType

`func (o *GenericDelegatedAdminAttributeResponse) HasReferenceResourceType() bool`

HasReferenceResourceType returns a boolean if a field has been set.

### GetAttributePresentation

`func (o *GenericDelegatedAdminAttributeResponse) GetAttributePresentation() EnumdelegatedAdminAttributeAttributePresentationProp`

GetAttributePresentation returns the AttributePresentation field if non-nil, zero value otherwise.

### GetAttributePresentationOk

`func (o *GenericDelegatedAdminAttributeResponse) GetAttributePresentationOk() (*EnumdelegatedAdminAttributeAttributePresentationProp, bool)`

GetAttributePresentationOk returns a tuple with the AttributePresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributePresentation

`func (o *GenericDelegatedAdminAttributeResponse) SetAttributePresentation(v EnumdelegatedAdminAttributeAttributePresentationProp)`

SetAttributePresentation sets AttributePresentation field to given value.

### HasAttributePresentation

`func (o *GenericDelegatedAdminAttributeResponse) HasAttributePresentation() bool`

HasAttributePresentation returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *GenericDelegatedAdminAttributeResponse) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *GenericDelegatedAdminAttributeResponse) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *GenericDelegatedAdminAttributeResponse) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *GenericDelegatedAdminAttributeResponse) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### GetMeta

`func (o *GenericDelegatedAdminAttributeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GenericDelegatedAdminAttributeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GenericDelegatedAdminAttributeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GenericDelegatedAdminAttributeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GenericDelegatedAdminAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GenericDelegatedAdminAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GenericDelegatedAdminAttributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GenericDelegatedAdminAttributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


