# AddDelegatedAdminAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeType** | **string** | Specifies the name or OID of the LDAP attribute type. | 
**Schemas** | [**[]EnumgenericDelegatedAdminAttributeSchemaUrn**](EnumgenericDelegatedAdminAttributeSchemaUrn.md) |  | 
**AllowedMIMEType** | Pointer to [**[]EnumdelegatedAdminAttributePhotoAllowedMIMETypeProp**](EnumdelegatedAdminAttributePhotoAllowedMIMETypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Attribute | [optional] 
**DisplayName** | **string** | A human readable display name for this Delegated Admin Attribute. | 
**Mutability** | Pointer to [**EnumdelegatedAdminAttributeMutabilityProp**](EnumdelegatedAdminAttributeMutabilityProp.md) |  | [optional] 
**MultiValued** | Pointer to **bool** | Indicates whether this Delegated Admin Attribute may have multiple values. | [optional] 
**AttributeCategory** | Pointer to **string** | Specifies which attribute category this attribute belongs to. | [optional] 
**DisplayOrderIndex** | Pointer to **int32** | This property determines a display order for attributes within a given attribute category. Attributes are ordered within their category based on this index from least to greatest. | [optional] 
**ReferenceResourceType** | Pointer to **string** | For LDAP attributes with DN syntax, specifies what kind of resource is referenced. | [optional] 
**AttributePresentation** | Pointer to [**EnumdelegatedAdminAttributeAttributePresentationProp**](EnumdelegatedAdminAttributeAttributePresentationProp.md) |  | [optional] 
**DateTimeFormat** | Pointer to **string** | Specifies the format string that is used to present a date and/or time value to the user of the app. This property only applies to LDAP attribute types whose LDAP syntax is GeneralizedTime and is ignored if the attribute type has any other syntax. | [optional] 
**IncludeInSummary** | Pointer to **bool** | Indicates whether this Delegated Admin Attribute is to be included in the summary display for a resource. | [optional] 

## Methods

### NewAddDelegatedAdminAttributeRequest

`func NewAddDelegatedAdminAttributeRequest(attributeType string, schemas []EnumgenericDelegatedAdminAttributeSchemaUrn, displayName string, ) *AddDelegatedAdminAttributeRequest`

NewAddDelegatedAdminAttributeRequest instantiates a new AddDelegatedAdminAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDelegatedAdminAttributeRequestWithDefaults

`func NewAddDelegatedAdminAttributeRequestWithDefaults() *AddDelegatedAdminAttributeRequest`

NewAddDelegatedAdminAttributeRequestWithDefaults instantiates a new AddDelegatedAdminAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeType

`func (o *AddDelegatedAdminAttributeRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddDelegatedAdminAttributeRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddDelegatedAdminAttributeRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetSchemas

`func (o *AddDelegatedAdminAttributeRequest) GetSchemas() []EnumgenericDelegatedAdminAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDelegatedAdminAttributeRequest) GetSchemasOk() (*[]EnumgenericDelegatedAdminAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDelegatedAdminAttributeRequest) SetSchemas(v []EnumgenericDelegatedAdminAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowedMIMEType

`func (o *AddDelegatedAdminAttributeRequest) GetAllowedMIMEType() []EnumdelegatedAdminAttributePhotoAllowedMIMETypeProp`

GetAllowedMIMEType returns the AllowedMIMEType field if non-nil, zero value otherwise.

### GetAllowedMIMETypeOk

`func (o *AddDelegatedAdminAttributeRequest) GetAllowedMIMETypeOk() (*[]EnumdelegatedAdminAttributePhotoAllowedMIMETypeProp, bool)`

GetAllowedMIMETypeOk returns a tuple with the AllowedMIMEType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMIMEType

`func (o *AddDelegatedAdminAttributeRequest) SetAllowedMIMEType(v []EnumdelegatedAdminAttributePhotoAllowedMIMETypeProp)`

SetAllowedMIMEType sets AllowedMIMEType field to given value.

### HasAllowedMIMEType

`func (o *AddDelegatedAdminAttributeRequest) HasAllowedMIMEType() bool`

HasAllowedMIMEType returns a boolean if a field has been set.

### GetDescription

`func (o *AddDelegatedAdminAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDelegatedAdminAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDelegatedAdminAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDelegatedAdminAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddDelegatedAdminAttributeRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddDelegatedAdminAttributeRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddDelegatedAdminAttributeRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMutability

`func (o *AddDelegatedAdminAttributeRequest) GetMutability() EnumdelegatedAdminAttributeMutabilityProp`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *AddDelegatedAdminAttributeRequest) GetMutabilityOk() (*EnumdelegatedAdminAttributeMutabilityProp, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *AddDelegatedAdminAttributeRequest) SetMutability(v EnumdelegatedAdminAttributeMutabilityProp)`

SetMutability sets Mutability field to given value.

### HasMutability

`func (o *AddDelegatedAdminAttributeRequest) HasMutability() bool`

HasMutability returns a boolean if a field has been set.

### GetMultiValued

`func (o *AddDelegatedAdminAttributeRequest) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *AddDelegatedAdminAttributeRequest) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *AddDelegatedAdminAttributeRequest) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *AddDelegatedAdminAttributeRequest) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetAttributeCategory

`func (o *AddDelegatedAdminAttributeRequest) GetAttributeCategory() string`

GetAttributeCategory returns the AttributeCategory field if non-nil, zero value otherwise.

### GetAttributeCategoryOk

`func (o *AddDelegatedAdminAttributeRequest) GetAttributeCategoryOk() (*string, bool)`

GetAttributeCategoryOk returns a tuple with the AttributeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeCategory

`func (o *AddDelegatedAdminAttributeRequest) SetAttributeCategory(v string)`

SetAttributeCategory sets AttributeCategory field to given value.

### HasAttributeCategory

`func (o *AddDelegatedAdminAttributeRequest) HasAttributeCategory() bool`

HasAttributeCategory returns a boolean if a field has been set.

### GetDisplayOrderIndex

`func (o *AddDelegatedAdminAttributeRequest) GetDisplayOrderIndex() int32`

GetDisplayOrderIndex returns the DisplayOrderIndex field if non-nil, zero value otherwise.

### GetDisplayOrderIndexOk

`func (o *AddDelegatedAdminAttributeRequest) GetDisplayOrderIndexOk() (*int32, bool)`

GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrderIndex

`func (o *AddDelegatedAdminAttributeRequest) SetDisplayOrderIndex(v int32)`

SetDisplayOrderIndex sets DisplayOrderIndex field to given value.

### HasDisplayOrderIndex

`func (o *AddDelegatedAdminAttributeRequest) HasDisplayOrderIndex() bool`

HasDisplayOrderIndex returns a boolean if a field has been set.

### GetReferenceResourceType

`func (o *AddDelegatedAdminAttributeRequest) GetReferenceResourceType() string`

GetReferenceResourceType returns the ReferenceResourceType field if non-nil, zero value otherwise.

### GetReferenceResourceTypeOk

`func (o *AddDelegatedAdminAttributeRequest) GetReferenceResourceTypeOk() (*string, bool)`

GetReferenceResourceTypeOk returns a tuple with the ReferenceResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceResourceType

`func (o *AddDelegatedAdminAttributeRequest) SetReferenceResourceType(v string)`

SetReferenceResourceType sets ReferenceResourceType field to given value.

### HasReferenceResourceType

`func (o *AddDelegatedAdminAttributeRequest) HasReferenceResourceType() bool`

HasReferenceResourceType returns a boolean if a field has been set.

### GetAttributePresentation

`func (o *AddDelegatedAdminAttributeRequest) GetAttributePresentation() EnumdelegatedAdminAttributeAttributePresentationProp`

GetAttributePresentation returns the AttributePresentation field if non-nil, zero value otherwise.

### GetAttributePresentationOk

`func (o *AddDelegatedAdminAttributeRequest) GetAttributePresentationOk() (*EnumdelegatedAdminAttributeAttributePresentationProp, bool)`

GetAttributePresentationOk returns a tuple with the AttributePresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributePresentation

`func (o *AddDelegatedAdminAttributeRequest) SetAttributePresentation(v EnumdelegatedAdminAttributeAttributePresentationProp)`

SetAttributePresentation sets AttributePresentation field to given value.

### HasAttributePresentation

`func (o *AddDelegatedAdminAttributeRequest) HasAttributePresentation() bool`

HasAttributePresentation returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *AddDelegatedAdminAttributeRequest) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *AddDelegatedAdminAttributeRequest) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *AddDelegatedAdminAttributeRequest) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *AddDelegatedAdminAttributeRequest) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### GetIncludeInSummary

`func (o *AddDelegatedAdminAttributeRequest) GetIncludeInSummary() bool`

GetIncludeInSummary returns the IncludeInSummary field if non-nil, zero value otherwise.

### GetIncludeInSummaryOk

`func (o *AddDelegatedAdminAttributeRequest) GetIncludeInSummaryOk() (*bool, bool)`

GetIncludeInSummaryOk returns a tuple with the IncludeInSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSummary

`func (o *AddDelegatedAdminAttributeRequest) SetIncludeInSummary(v bool)`

SetIncludeInSummary sets IncludeInSummary field to given value.

### HasIncludeInSummary

`func (o *AddDelegatedAdminAttributeRequest) HasIncludeInSummary() bool`

HasIncludeInSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


