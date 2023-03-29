# AddCertificateDelegatedAdminAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeType** | **string** | Specifies the name or OID of the LDAP attribute type. | 
**Schemas** | [**[]EnumcertificateDelegatedAdminAttributeSchemaUrn**](EnumcertificateDelegatedAdminAttributeSchemaUrn.md) |  | 
**AllowedMIMEType** | Pointer to [**[]EnumdelegatedAdminAttributeCertificateAllowedMIMETypeProp**](EnumdelegatedAdminAttributeCertificateAllowedMIMETypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Attribute | [optional] 
**DisplayName** | **string** | A human readable display name for this Delegated Admin Attribute. | 
**Mutability** | Pointer to [**EnumdelegatedAdminAttributeMutabilityProp**](EnumdelegatedAdminAttributeMutabilityProp.md) |  | [optional] 
**MultiValued** | Pointer to **bool** | Indicates whether this Delegated Admin Attribute may have multiple values. | [optional] 
**AttributeCategory** | Pointer to **string** | Specifies which attribute category this attribute belongs to. | [optional] 
**DisplayOrderIndex** | Pointer to **int32** | This property determines a display order for attributes within a given attribute category. Attributes are ordered within their category based on this index from least to greatest. | [optional] 
**ReferenceResourceType** | Pointer to **string** | For LDAP attributes with DN syntax, specifies what kind of resource is referenced. | [optional] 
**AttributePresentation** | Pointer to [**EnumdelegatedAdminAttributeAttributePresentationProp**](EnumdelegatedAdminAttributeAttributePresentationProp.md) |  | [optional] 
**DateTimeFormat** | Pointer to **string** | Specifies the format string that is used to present a date and/or time value to the user of the app. This property only applies to LDAP attribute types whose LDAP syntax is GeneralizedTime and is ignored if the attribute type has any other syntax. | [optional] 

## Methods

### NewAddCertificateDelegatedAdminAttributeRequest

`func NewAddCertificateDelegatedAdminAttributeRequest(attributeType string, schemas []EnumcertificateDelegatedAdminAttributeSchemaUrn, displayName string, ) *AddCertificateDelegatedAdminAttributeRequest`

NewAddCertificateDelegatedAdminAttributeRequest instantiates a new AddCertificateDelegatedAdminAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCertificateDelegatedAdminAttributeRequestWithDefaults

`func NewAddCertificateDelegatedAdminAttributeRequestWithDefaults() *AddCertificateDelegatedAdminAttributeRequest`

NewAddCertificateDelegatedAdminAttributeRequestWithDefaults instantiates a new AddCertificateDelegatedAdminAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeType

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetSchemas

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetSchemas() []EnumcertificateDelegatedAdminAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetSchemasOk() (*[]EnumcertificateDelegatedAdminAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetSchemas(v []EnumcertificateDelegatedAdminAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowedMIMEType

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetAllowedMIMEType() []EnumdelegatedAdminAttributeCertificateAllowedMIMETypeProp`

GetAllowedMIMEType returns the AllowedMIMEType field if non-nil, zero value otherwise.

### GetAllowedMIMETypeOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetAllowedMIMETypeOk() (*[]EnumdelegatedAdminAttributeCertificateAllowedMIMETypeProp, bool)`

GetAllowedMIMETypeOk returns a tuple with the AllowedMIMEType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMIMEType

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetAllowedMIMEType(v []EnumdelegatedAdminAttributeCertificateAllowedMIMETypeProp)`

SetAllowedMIMEType sets AllowedMIMEType field to given value.

### HasAllowedMIMEType

`func (o *AddCertificateDelegatedAdminAttributeRequest) HasAllowedMIMEType() bool`

HasAllowedMIMEType returns a boolean if a field has been set.

### GetDescription

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCertificateDelegatedAdminAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMutability

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetMutability() EnumdelegatedAdminAttributeMutabilityProp`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetMutabilityOk() (*EnumdelegatedAdminAttributeMutabilityProp, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetMutability(v EnumdelegatedAdminAttributeMutabilityProp)`

SetMutability sets Mutability field to given value.

### HasMutability

`func (o *AddCertificateDelegatedAdminAttributeRequest) HasMutability() bool`

HasMutability returns a boolean if a field has been set.

### GetMultiValued

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *AddCertificateDelegatedAdminAttributeRequest) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetAttributeCategory

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetAttributeCategory() string`

GetAttributeCategory returns the AttributeCategory field if non-nil, zero value otherwise.

### GetAttributeCategoryOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetAttributeCategoryOk() (*string, bool)`

GetAttributeCategoryOk returns a tuple with the AttributeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeCategory

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetAttributeCategory(v string)`

SetAttributeCategory sets AttributeCategory field to given value.

### HasAttributeCategory

`func (o *AddCertificateDelegatedAdminAttributeRequest) HasAttributeCategory() bool`

HasAttributeCategory returns a boolean if a field has been set.

### GetDisplayOrderIndex

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetDisplayOrderIndex() int32`

GetDisplayOrderIndex returns the DisplayOrderIndex field if non-nil, zero value otherwise.

### GetDisplayOrderIndexOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetDisplayOrderIndexOk() (*int32, bool)`

GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrderIndex

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetDisplayOrderIndex(v int32)`

SetDisplayOrderIndex sets DisplayOrderIndex field to given value.

### HasDisplayOrderIndex

`func (o *AddCertificateDelegatedAdminAttributeRequest) HasDisplayOrderIndex() bool`

HasDisplayOrderIndex returns a boolean if a field has been set.

### GetReferenceResourceType

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetReferenceResourceType() string`

GetReferenceResourceType returns the ReferenceResourceType field if non-nil, zero value otherwise.

### GetReferenceResourceTypeOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetReferenceResourceTypeOk() (*string, bool)`

GetReferenceResourceTypeOk returns a tuple with the ReferenceResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceResourceType

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetReferenceResourceType(v string)`

SetReferenceResourceType sets ReferenceResourceType field to given value.

### HasReferenceResourceType

`func (o *AddCertificateDelegatedAdminAttributeRequest) HasReferenceResourceType() bool`

HasReferenceResourceType returns a boolean if a field has been set.

### GetAttributePresentation

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetAttributePresentation() EnumdelegatedAdminAttributeAttributePresentationProp`

GetAttributePresentation returns the AttributePresentation field if non-nil, zero value otherwise.

### GetAttributePresentationOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetAttributePresentationOk() (*EnumdelegatedAdminAttributeAttributePresentationProp, bool)`

GetAttributePresentationOk returns a tuple with the AttributePresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributePresentation

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetAttributePresentation(v EnumdelegatedAdminAttributeAttributePresentationProp)`

SetAttributePresentation sets AttributePresentation field to given value.

### HasAttributePresentation

`func (o *AddCertificateDelegatedAdminAttributeRequest) HasAttributePresentation() bool`

HasAttributePresentation returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *AddCertificateDelegatedAdminAttributeRequest) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *AddCertificateDelegatedAdminAttributeRequest) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *AddCertificateDelegatedAdminAttributeRequest) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


