# AddDelegatedAdminAttribute200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Delegated Admin Attribute | 
**Schemas** | [**[]EnumgenericDelegatedAdminAttributeSchemaUrn**](EnumgenericDelegatedAdminAttributeSchemaUrn.md) |  | 
**AllowedMIMEType** | Pointer to [**[]EnumdelegatedAdminAttributePhotoAllowedMIMETypeProp**](EnumdelegatedAdminAttributePhotoAllowedMIMETypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Attribute | [optional] 
**AttributeType** | **string** | Specifies the name or OID of the LDAP attribute type. | 
**DisplayName** | **string** | A human readable display name for this Delegated Admin Attribute. | 
**Mutability** | [**EnumdelegatedAdminAttributeMutabilityProp**](EnumdelegatedAdminAttributeMutabilityProp.md) |  | 
**MultiValued** | **bool** | Indicates whether this Delegated Admin Attribute may have multiple values. | 
**AttributeCategory** | Pointer to **string** | Specifies which attribute category this attribute belongs to. | [optional] 
**DisplayOrderIndex** | **int64** | This property determines a display order for attributes within a given attribute category. Attributes are ordered within their category based on this index from least to greatest. | 
**ReferenceResourceType** | Pointer to **string** | For LDAP attributes with DN syntax, specifies what kind of resource is referenced. | [optional] 
**AttributePresentation** | Pointer to [**EnumdelegatedAdminAttributeAttributePresentationProp**](EnumdelegatedAdminAttributeAttributePresentationProp.md) |  | [optional] 
**DateTimeFormat** | Pointer to **string** | Specifies the format string that is used to present a date and/or time value to the user of the app. This property only applies to LDAP attribute types whose LDAP syntax is GeneralizedTime and is ignored if the attribute type has any other syntax. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**IncludeInSummary** | **bool** | Indicates whether this Delegated Admin Attribute is to be included in the summary display for a resource. | 

## Methods

### NewAddDelegatedAdminAttribute200Response

`func NewAddDelegatedAdminAttribute200Response(id string, schemas []EnumgenericDelegatedAdminAttributeSchemaUrn, attributeType string, displayName string, mutability EnumdelegatedAdminAttributeMutabilityProp, multiValued bool, displayOrderIndex int64, includeInSummary bool, ) *AddDelegatedAdminAttribute200Response`

NewAddDelegatedAdminAttribute200Response instantiates a new AddDelegatedAdminAttribute200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDelegatedAdminAttribute200ResponseWithDefaults

`func NewAddDelegatedAdminAttribute200ResponseWithDefaults() *AddDelegatedAdminAttribute200Response`

NewAddDelegatedAdminAttribute200ResponseWithDefaults instantiates a new AddDelegatedAdminAttribute200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddDelegatedAdminAttribute200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddDelegatedAdminAttribute200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddDelegatedAdminAttribute200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddDelegatedAdminAttribute200Response) GetSchemas() []EnumgenericDelegatedAdminAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDelegatedAdminAttribute200Response) GetSchemasOk() (*[]EnumgenericDelegatedAdminAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDelegatedAdminAttribute200Response) SetSchemas(v []EnumgenericDelegatedAdminAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowedMIMEType

`func (o *AddDelegatedAdminAttribute200Response) GetAllowedMIMEType() []EnumdelegatedAdminAttributePhotoAllowedMIMETypeProp`

GetAllowedMIMEType returns the AllowedMIMEType field if non-nil, zero value otherwise.

### GetAllowedMIMETypeOk

`func (o *AddDelegatedAdminAttribute200Response) GetAllowedMIMETypeOk() (*[]EnumdelegatedAdminAttributePhotoAllowedMIMETypeProp, bool)`

GetAllowedMIMETypeOk returns a tuple with the AllowedMIMEType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMIMEType

`func (o *AddDelegatedAdminAttribute200Response) SetAllowedMIMEType(v []EnumdelegatedAdminAttributePhotoAllowedMIMETypeProp)`

SetAllowedMIMEType sets AllowedMIMEType field to given value.

### HasAllowedMIMEType

`func (o *AddDelegatedAdminAttribute200Response) HasAllowedMIMEType() bool`

HasAllowedMIMEType returns a boolean if a field has been set.

### GetDescription

`func (o *AddDelegatedAdminAttribute200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDelegatedAdminAttribute200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDelegatedAdminAttribute200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDelegatedAdminAttribute200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeType

`func (o *AddDelegatedAdminAttribute200Response) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddDelegatedAdminAttribute200Response) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddDelegatedAdminAttribute200Response) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetDisplayName

`func (o *AddDelegatedAdminAttribute200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddDelegatedAdminAttribute200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddDelegatedAdminAttribute200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMutability

`func (o *AddDelegatedAdminAttribute200Response) GetMutability() EnumdelegatedAdminAttributeMutabilityProp`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *AddDelegatedAdminAttribute200Response) GetMutabilityOk() (*EnumdelegatedAdminAttributeMutabilityProp, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *AddDelegatedAdminAttribute200Response) SetMutability(v EnumdelegatedAdminAttributeMutabilityProp)`

SetMutability sets Mutability field to given value.


### GetMultiValued

`func (o *AddDelegatedAdminAttribute200Response) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *AddDelegatedAdminAttribute200Response) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *AddDelegatedAdminAttribute200Response) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.


### GetAttributeCategory

`func (o *AddDelegatedAdminAttribute200Response) GetAttributeCategory() string`

GetAttributeCategory returns the AttributeCategory field if non-nil, zero value otherwise.

### GetAttributeCategoryOk

`func (o *AddDelegatedAdminAttribute200Response) GetAttributeCategoryOk() (*string, bool)`

GetAttributeCategoryOk returns a tuple with the AttributeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeCategory

`func (o *AddDelegatedAdminAttribute200Response) SetAttributeCategory(v string)`

SetAttributeCategory sets AttributeCategory field to given value.

### HasAttributeCategory

`func (o *AddDelegatedAdminAttribute200Response) HasAttributeCategory() bool`

HasAttributeCategory returns a boolean if a field has been set.

### GetDisplayOrderIndex

`func (o *AddDelegatedAdminAttribute200Response) GetDisplayOrderIndex() int64`

GetDisplayOrderIndex returns the DisplayOrderIndex field if non-nil, zero value otherwise.

### GetDisplayOrderIndexOk

`func (o *AddDelegatedAdminAttribute200Response) GetDisplayOrderIndexOk() (*int64, bool)`

GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrderIndex

`func (o *AddDelegatedAdminAttribute200Response) SetDisplayOrderIndex(v int64)`

SetDisplayOrderIndex sets DisplayOrderIndex field to given value.


### GetReferenceResourceType

`func (o *AddDelegatedAdminAttribute200Response) GetReferenceResourceType() string`

GetReferenceResourceType returns the ReferenceResourceType field if non-nil, zero value otherwise.

### GetReferenceResourceTypeOk

`func (o *AddDelegatedAdminAttribute200Response) GetReferenceResourceTypeOk() (*string, bool)`

GetReferenceResourceTypeOk returns a tuple with the ReferenceResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceResourceType

`func (o *AddDelegatedAdminAttribute200Response) SetReferenceResourceType(v string)`

SetReferenceResourceType sets ReferenceResourceType field to given value.

### HasReferenceResourceType

`func (o *AddDelegatedAdminAttribute200Response) HasReferenceResourceType() bool`

HasReferenceResourceType returns a boolean if a field has been set.

### GetAttributePresentation

`func (o *AddDelegatedAdminAttribute200Response) GetAttributePresentation() EnumdelegatedAdminAttributeAttributePresentationProp`

GetAttributePresentation returns the AttributePresentation field if non-nil, zero value otherwise.

### GetAttributePresentationOk

`func (o *AddDelegatedAdminAttribute200Response) GetAttributePresentationOk() (*EnumdelegatedAdminAttributeAttributePresentationProp, bool)`

GetAttributePresentationOk returns a tuple with the AttributePresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributePresentation

`func (o *AddDelegatedAdminAttribute200Response) SetAttributePresentation(v EnumdelegatedAdminAttributeAttributePresentationProp)`

SetAttributePresentation sets AttributePresentation field to given value.

### HasAttributePresentation

`func (o *AddDelegatedAdminAttribute200Response) HasAttributePresentation() bool`

HasAttributePresentation returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *AddDelegatedAdminAttribute200Response) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *AddDelegatedAdminAttribute200Response) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *AddDelegatedAdminAttribute200Response) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *AddDelegatedAdminAttribute200Response) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### GetMeta

`func (o *AddDelegatedAdminAttribute200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddDelegatedAdminAttribute200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddDelegatedAdminAttribute200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddDelegatedAdminAttribute200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddDelegatedAdminAttribute200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddDelegatedAdminAttribute200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddDelegatedAdminAttribute200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddDelegatedAdminAttribute200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetIncludeInSummary

`func (o *AddDelegatedAdminAttribute200Response) GetIncludeInSummary() bool`

GetIncludeInSummary returns the IncludeInSummary field if non-nil, zero value otherwise.

### GetIncludeInSummaryOk

`func (o *AddDelegatedAdminAttribute200Response) GetIncludeInSummaryOk() (*bool, bool)`

GetIncludeInSummaryOk returns a tuple with the IncludeInSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSummary

`func (o *AddDelegatedAdminAttribute200Response) SetIncludeInSummary(v bool)`

SetIncludeInSummary sets IncludeInSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


