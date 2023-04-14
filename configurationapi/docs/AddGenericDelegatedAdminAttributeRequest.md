# AddGenericDelegatedAdminAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeType** | **string** | Specifies the name or OID of the LDAP attribute type. | 
**Schemas** | [**[]EnumgenericDelegatedAdminAttributeSchemaUrn**](EnumgenericDelegatedAdminAttributeSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Delegated Admin Attribute | [optional] 
**DisplayName** | **string** | A human readable display name for this Delegated Admin Attribute. | 
**Mutability** | Pointer to [**EnumdelegatedAdminAttributeMutabilityProp**](EnumdelegatedAdminAttributeMutabilityProp.md) |  | [optional] 
**MultiValued** | Pointer to **bool** | Indicates whether this Delegated Admin Attribute may have multiple values. | [optional] 
**IncludeInSummary** | Pointer to **bool** | Indicates whether this Delegated Admin Attribute is to be included in the summary display for a resource. | [optional] 
**AttributeCategory** | Pointer to **string** | Specifies which attribute category this attribute belongs to. | [optional] 
**DisplayOrderIndex** | Pointer to **int64** | This property determines a display order for attributes within a given attribute category. Attributes are ordered within their category based on this index from least to greatest. | [optional] 
**ReferenceResourceType** | Pointer to **string** | For LDAP attributes with DN syntax, specifies what kind of resource is referenced. | [optional] 
**AttributePresentation** | Pointer to [**EnumdelegatedAdminAttributeAttributePresentationProp**](EnumdelegatedAdminAttributeAttributePresentationProp.md) |  | [optional] 
**DateTimeFormat** | Pointer to **string** | Specifies the format string that is used to present a date and/or time value to the user of the app. This property only applies to LDAP attribute types whose LDAP syntax is GeneralizedTime and is ignored if the attribute type has any other syntax. | [optional] 

## Methods

### NewAddGenericDelegatedAdminAttributeRequest

`func NewAddGenericDelegatedAdminAttributeRequest(attributeType string, schemas []EnumgenericDelegatedAdminAttributeSchemaUrn, displayName string, ) *AddGenericDelegatedAdminAttributeRequest`

NewAddGenericDelegatedAdminAttributeRequest instantiates a new AddGenericDelegatedAdminAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGenericDelegatedAdminAttributeRequestWithDefaults

`func NewAddGenericDelegatedAdminAttributeRequestWithDefaults() *AddGenericDelegatedAdminAttributeRequest`

NewAddGenericDelegatedAdminAttributeRequestWithDefaults instantiates a new AddGenericDelegatedAdminAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeType

`func (o *AddGenericDelegatedAdminAttributeRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddGenericDelegatedAdminAttributeRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetSchemas

`func (o *AddGenericDelegatedAdminAttributeRequest) GetSchemas() []EnumgenericDelegatedAdminAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetSchemasOk() (*[]EnumgenericDelegatedAdminAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGenericDelegatedAdminAttributeRequest) SetSchemas(v []EnumgenericDelegatedAdminAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddGenericDelegatedAdminAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGenericDelegatedAdminAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGenericDelegatedAdminAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddGenericDelegatedAdminAttributeRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddGenericDelegatedAdminAttributeRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMutability

`func (o *AddGenericDelegatedAdminAttributeRequest) GetMutability() EnumdelegatedAdminAttributeMutabilityProp`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetMutabilityOk() (*EnumdelegatedAdminAttributeMutabilityProp, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *AddGenericDelegatedAdminAttributeRequest) SetMutability(v EnumdelegatedAdminAttributeMutabilityProp)`

SetMutability sets Mutability field to given value.

### HasMutability

`func (o *AddGenericDelegatedAdminAttributeRequest) HasMutability() bool`

HasMutability returns a boolean if a field has been set.

### GetMultiValued

`func (o *AddGenericDelegatedAdminAttributeRequest) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *AddGenericDelegatedAdminAttributeRequest) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *AddGenericDelegatedAdminAttributeRequest) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetIncludeInSummary

`func (o *AddGenericDelegatedAdminAttributeRequest) GetIncludeInSummary() bool`

GetIncludeInSummary returns the IncludeInSummary field if non-nil, zero value otherwise.

### GetIncludeInSummaryOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetIncludeInSummaryOk() (*bool, bool)`

GetIncludeInSummaryOk returns a tuple with the IncludeInSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSummary

`func (o *AddGenericDelegatedAdminAttributeRequest) SetIncludeInSummary(v bool)`

SetIncludeInSummary sets IncludeInSummary field to given value.

### HasIncludeInSummary

`func (o *AddGenericDelegatedAdminAttributeRequest) HasIncludeInSummary() bool`

HasIncludeInSummary returns a boolean if a field has been set.

### GetAttributeCategory

`func (o *AddGenericDelegatedAdminAttributeRequest) GetAttributeCategory() string`

GetAttributeCategory returns the AttributeCategory field if non-nil, zero value otherwise.

### GetAttributeCategoryOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetAttributeCategoryOk() (*string, bool)`

GetAttributeCategoryOk returns a tuple with the AttributeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeCategory

`func (o *AddGenericDelegatedAdminAttributeRequest) SetAttributeCategory(v string)`

SetAttributeCategory sets AttributeCategory field to given value.

### HasAttributeCategory

`func (o *AddGenericDelegatedAdminAttributeRequest) HasAttributeCategory() bool`

HasAttributeCategory returns a boolean if a field has been set.

### GetDisplayOrderIndex

`func (o *AddGenericDelegatedAdminAttributeRequest) GetDisplayOrderIndex() int64`

GetDisplayOrderIndex returns the DisplayOrderIndex field if non-nil, zero value otherwise.

### GetDisplayOrderIndexOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetDisplayOrderIndexOk() (*int64, bool)`

GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrderIndex

`func (o *AddGenericDelegatedAdminAttributeRequest) SetDisplayOrderIndex(v int64)`

SetDisplayOrderIndex sets DisplayOrderIndex field to given value.

### HasDisplayOrderIndex

`func (o *AddGenericDelegatedAdminAttributeRequest) HasDisplayOrderIndex() bool`

HasDisplayOrderIndex returns a boolean if a field has been set.

### GetReferenceResourceType

`func (o *AddGenericDelegatedAdminAttributeRequest) GetReferenceResourceType() string`

GetReferenceResourceType returns the ReferenceResourceType field if non-nil, zero value otherwise.

### GetReferenceResourceTypeOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetReferenceResourceTypeOk() (*string, bool)`

GetReferenceResourceTypeOk returns a tuple with the ReferenceResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceResourceType

`func (o *AddGenericDelegatedAdminAttributeRequest) SetReferenceResourceType(v string)`

SetReferenceResourceType sets ReferenceResourceType field to given value.

### HasReferenceResourceType

`func (o *AddGenericDelegatedAdminAttributeRequest) HasReferenceResourceType() bool`

HasReferenceResourceType returns a boolean if a field has been set.

### GetAttributePresentation

`func (o *AddGenericDelegatedAdminAttributeRequest) GetAttributePresentation() EnumdelegatedAdminAttributeAttributePresentationProp`

GetAttributePresentation returns the AttributePresentation field if non-nil, zero value otherwise.

### GetAttributePresentationOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetAttributePresentationOk() (*EnumdelegatedAdminAttributeAttributePresentationProp, bool)`

GetAttributePresentationOk returns a tuple with the AttributePresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributePresentation

`func (o *AddGenericDelegatedAdminAttributeRequest) SetAttributePresentation(v EnumdelegatedAdminAttributeAttributePresentationProp)`

SetAttributePresentation sets AttributePresentation field to given value.

### HasAttributePresentation

`func (o *AddGenericDelegatedAdminAttributeRequest) HasAttributePresentation() bool`

HasAttributePresentation returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *AddGenericDelegatedAdminAttributeRequest) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *AddGenericDelegatedAdminAttributeRequest) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *AddGenericDelegatedAdminAttributeRequest) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *AddGenericDelegatedAdminAttributeRequest) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


