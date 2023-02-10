# PhotoDelegatedAdminAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Delegated Admin Attribute | 
**Schemas** | [**[]EnumphotoDelegatedAdminAttributeSchemaUrn**](EnumphotoDelegatedAdminAttributeSchemaUrn.md) |  | 
**AllowedMIMEType** | Pointer to [**[]EnumdelegatedAdminAttributeAllowedMIMETypeProp**](EnumdelegatedAdminAttributeAllowedMIMETypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Attribute | [optional] 
**AttributeType** | **string** | Specifies the name or OID of the LDAP attribute type. | 
**DisplayName** | **string** | A human readable display name for this Delegated Admin Attribute. | 
**Mutability** | [**EnumdelegatedAdminAttributeMutabilityProp**](EnumdelegatedAdminAttributeMutabilityProp.md) |  | 
**MultiValued** | **bool** | Indicates whether this Delegated Admin Attribute may have multiple values. | 
**AttributeCategory** | Pointer to **string** | Specifies which attribute category this attribute belongs to. | [optional] 
**DisplayOrderIndex** | **int32** | This property determines a display order for attributes within a given attribute category. Attributes are ordered within their category based on this index from least to greatest. | 
**ReferenceResourceType** | Pointer to **string** | For LDAP attributes with DN syntax, specifies what kind of resource is referenced. | [optional] 
**AttributePresentation** | Pointer to [**EnumdelegatedAdminAttributeAttributePresentationProp**](EnumdelegatedAdminAttributeAttributePresentationProp.md) |  | [optional] 
**DateTimeFormat** | Pointer to **string** | Specifies the format string that is used to present a date and/or time value to the user of the app. This property only applies to LDAP attribute types whose LDAP syntax is GeneralizedTime and is ignored if the attribute type has any other syntax. | [optional] 

## Methods

### NewPhotoDelegatedAdminAttributeResponse

`func NewPhotoDelegatedAdminAttributeResponse(id string, schemas []EnumphotoDelegatedAdminAttributeSchemaUrn, attributeType string, displayName string, mutability EnumdelegatedAdminAttributeMutabilityProp, multiValued bool, displayOrderIndex int32, ) *PhotoDelegatedAdminAttributeResponse`

NewPhotoDelegatedAdminAttributeResponse instantiates a new PhotoDelegatedAdminAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhotoDelegatedAdminAttributeResponseWithDefaults

`func NewPhotoDelegatedAdminAttributeResponseWithDefaults() *PhotoDelegatedAdminAttributeResponse`

NewPhotoDelegatedAdminAttributeResponseWithDefaults instantiates a new PhotoDelegatedAdminAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PhotoDelegatedAdminAttributeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PhotoDelegatedAdminAttributeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PhotoDelegatedAdminAttributeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PhotoDelegatedAdminAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PhotoDelegatedAdminAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PhotoDelegatedAdminAttributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PhotoDelegatedAdminAttributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *PhotoDelegatedAdminAttributeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhotoDelegatedAdminAttributeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PhotoDelegatedAdminAttributeResponse) GetSchemas() []EnumphotoDelegatedAdminAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetSchemasOk() (*[]EnumphotoDelegatedAdminAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PhotoDelegatedAdminAttributeResponse) SetSchemas(v []EnumphotoDelegatedAdminAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowedMIMEType

`func (o *PhotoDelegatedAdminAttributeResponse) GetAllowedMIMEType() []EnumdelegatedAdminAttributeAllowedMIMETypeProp`

GetAllowedMIMEType returns the AllowedMIMEType field if non-nil, zero value otherwise.

### GetAllowedMIMETypeOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetAllowedMIMETypeOk() (*[]EnumdelegatedAdminAttributeAllowedMIMETypeProp, bool)`

GetAllowedMIMETypeOk returns a tuple with the AllowedMIMEType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMIMEType

`func (o *PhotoDelegatedAdminAttributeResponse) SetAllowedMIMEType(v []EnumdelegatedAdminAttributeAllowedMIMETypeProp)`

SetAllowedMIMEType sets AllowedMIMEType field to given value.

### HasAllowedMIMEType

`func (o *PhotoDelegatedAdminAttributeResponse) HasAllowedMIMEType() bool`

HasAllowedMIMEType returns a boolean if a field has been set.

### GetDescription

`func (o *PhotoDelegatedAdminAttributeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PhotoDelegatedAdminAttributeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PhotoDelegatedAdminAttributeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeType

`func (o *PhotoDelegatedAdminAttributeResponse) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *PhotoDelegatedAdminAttributeResponse) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetDisplayName

`func (o *PhotoDelegatedAdminAttributeResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PhotoDelegatedAdminAttributeResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMutability

`func (o *PhotoDelegatedAdminAttributeResponse) GetMutability() EnumdelegatedAdminAttributeMutabilityProp`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetMutabilityOk() (*EnumdelegatedAdminAttributeMutabilityProp, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *PhotoDelegatedAdminAttributeResponse) SetMutability(v EnumdelegatedAdminAttributeMutabilityProp)`

SetMutability sets Mutability field to given value.


### GetMultiValued

`func (o *PhotoDelegatedAdminAttributeResponse) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *PhotoDelegatedAdminAttributeResponse) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.


### GetAttributeCategory

`func (o *PhotoDelegatedAdminAttributeResponse) GetAttributeCategory() string`

GetAttributeCategory returns the AttributeCategory field if non-nil, zero value otherwise.

### GetAttributeCategoryOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetAttributeCategoryOk() (*string, bool)`

GetAttributeCategoryOk returns a tuple with the AttributeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeCategory

`func (o *PhotoDelegatedAdminAttributeResponse) SetAttributeCategory(v string)`

SetAttributeCategory sets AttributeCategory field to given value.

### HasAttributeCategory

`func (o *PhotoDelegatedAdminAttributeResponse) HasAttributeCategory() bool`

HasAttributeCategory returns a boolean if a field has been set.

### GetDisplayOrderIndex

`func (o *PhotoDelegatedAdminAttributeResponse) GetDisplayOrderIndex() int32`

GetDisplayOrderIndex returns the DisplayOrderIndex field if non-nil, zero value otherwise.

### GetDisplayOrderIndexOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetDisplayOrderIndexOk() (*int32, bool)`

GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrderIndex

`func (o *PhotoDelegatedAdminAttributeResponse) SetDisplayOrderIndex(v int32)`

SetDisplayOrderIndex sets DisplayOrderIndex field to given value.


### GetReferenceResourceType

`func (o *PhotoDelegatedAdminAttributeResponse) GetReferenceResourceType() string`

GetReferenceResourceType returns the ReferenceResourceType field if non-nil, zero value otherwise.

### GetReferenceResourceTypeOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetReferenceResourceTypeOk() (*string, bool)`

GetReferenceResourceTypeOk returns a tuple with the ReferenceResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceResourceType

`func (o *PhotoDelegatedAdminAttributeResponse) SetReferenceResourceType(v string)`

SetReferenceResourceType sets ReferenceResourceType field to given value.

### HasReferenceResourceType

`func (o *PhotoDelegatedAdminAttributeResponse) HasReferenceResourceType() bool`

HasReferenceResourceType returns a boolean if a field has been set.

### GetAttributePresentation

`func (o *PhotoDelegatedAdminAttributeResponse) GetAttributePresentation() EnumdelegatedAdminAttributeAttributePresentationProp`

GetAttributePresentation returns the AttributePresentation field if non-nil, zero value otherwise.

### GetAttributePresentationOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetAttributePresentationOk() (*EnumdelegatedAdminAttributeAttributePresentationProp, bool)`

GetAttributePresentationOk returns a tuple with the AttributePresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributePresentation

`func (o *PhotoDelegatedAdminAttributeResponse) SetAttributePresentation(v EnumdelegatedAdminAttributeAttributePresentationProp)`

SetAttributePresentation sets AttributePresentation field to given value.

### HasAttributePresentation

`func (o *PhotoDelegatedAdminAttributeResponse) HasAttributePresentation() bool`

HasAttributePresentation returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *PhotoDelegatedAdminAttributeResponse) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *PhotoDelegatedAdminAttributeResponse) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *PhotoDelegatedAdminAttributeResponse) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *PhotoDelegatedAdminAttributeResponse) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


