# ScimAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the SCIM Attribute | 
**Schemas** | Pointer to [**[]EnumscimAttributeSchemaUrn**](EnumscimAttributeSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this SCIM Attribute | [optional] 
**Name** | **string** | The name of the attribute. | 
**Type** | [**EnumscimAttributeTypeProp**](EnumscimAttributeTypeProp.md) |  | 
**Required** | **bool** | Specifies whether this attribute is required. | 
**CaseExact** | **bool** | Specifies whether the attribute values are case sensitive. | 
**MultiValued** | **bool** | Specifies whether this attribute may have multiple values. | 
**CanonicalValue** | Pointer to **[]string** | Specifies the suggested canonical type values for the attribute. | [optional] 
**Mutability** | [**EnumscimAttributeMutabilityProp**](EnumscimAttributeMutabilityProp.md) |  | 
**Returned** | [**EnumscimAttributeReturnedProp**](EnumscimAttributeReturnedProp.md) |  | 
**ReferenceType** | Pointer to **[]string** | Specifies the SCIM resource types that may be referenced. This property is only applicable for attributes that are of type &#39;reference&#39;. Valid values are: A SCIM resource type (e.g., &#39;User&#39; or &#39;Group&#39;), &#39;external&#39; - indicating the resource is an external resource (e.g., such as a photo), or &#39;uri&#39; - indicating that the reference is to a service endpoint or an identifier (such as a schema urn). | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewScimAttributeResponse

`func NewScimAttributeResponse(id string, name string, type_ EnumscimAttributeTypeProp, required bool, caseExact bool, multiValued bool, mutability EnumscimAttributeMutabilityProp, returned EnumscimAttributeReturnedProp, ) *ScimAttributeResponse`

NewScimAttributeResponse instantiates a new ScimAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimAttributeResponseWithDefaults

`func NewScimAttributeResponseWithDefaults() *ScimAttributeResponse`

NewScimAttributeResponseWithDefaults instantiates a new ScimAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScimAttributeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimAttributeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimAttributeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ScimAttributeResponse) GetSchemas() []EnumscimAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimAttributeResponse) GetSchemasOk() (*[]EnumscimAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimAttributeResponse) SetSchemas(v []EnumscimAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimAttributeResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ScimAttributeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimAttributeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimAttributeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScimAttributeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ScimAttributeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimAttributeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimAttributeResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ScimAttributeResponse) GetType() EnumscimAttributeTypeProp`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScimAttributeResponse) GetTypeOk() (*EnumscimAttributeTypeProp, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScimAttributeResponse) SetType(v EnumscimAttributeTypeProp)`

SetType sets Type field to given value.


### GetRequired

`func (o *ScimAttributeResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ScimAttributeResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ScimAttributeResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetCaseExact

`func (o *ScimAttributeResponse) GetCaseExact() bool`

GetCaseExact returns the CaseExact field if non-nil, zero value otherwise.

### GetCaseExactOk

`func (o *ScimAttributeResponse) GetCaseExactOk() (*bool, bool)`

GetCaseExactOk returns a tuple with the CaseExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExact

`func (o *ScimAttributeResponse) SetCaseExact(v bool)`

SetCaseExact sets CaseExact field to given value.


### GetMultiValued

`func (o *ScimAttributeResponse) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *ScimAttributeResponse) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *ScimAttributeResponse) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.


### GetCanonicalValue

`func (o *ScimAttributeResponse) GetCanonicalValue() []string`

GetCanonicalValue returns the CanonicalValue field if non-nil, zero value otherwise.

### GetCanonicalValueOk

`func (o *ScimAttributeResponse) GetCanonicalValueOk() (*[]string, bool)`

GetCanonicalValueOk returns a tuple with the CanonicalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalValue

`func (o *ScimAttributeResponse) SetCanonicalValue(v []string)`

SetCanonicalValue sets CanonicalValue field to given value.

### HasCanonicalValue

`func (o *ScimAttributeResponse) HasCanonicalValue() bool`

HasCanonicalValue returns a boolean if a field has been set.

### GetMutability

`func (o *ScimAttributeResponse) GetMutability() EnumscimAttributeMutabilityProp`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *ScimAttributeResponse) GetMutabilityOk() (*EnumscimAttributeMutabilityProp, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *ScimAttributeResponse) SetMutability(v EnumscimAttributeMutabilityProp)`

SetMutability sets Mutability field to given value.


### GetReturned

`func (o *ScimAttributeResponse) GetReturned() EnumscimAttributeReturnedProp`

GetReturned returns the Returned field if non-nil, zero value otherwise.

### GetReturnedOk

`func (o *ScimAttributeResponse) GetReturnedOk() (*EnumscimAttributeReturnedProp, bool)`

GetReturnedOk returns a tuple with the Returned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturned

`func (o *ScimAttributeResponse) SetReturned(v EnumscimAttributeReturnedProp)`

SetReturned sets Returned field to given value.


### GetReferenceType

`func (o *ScimAttributeResponse) GetReferenceType() []string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *ScimAttributeResponse) GetReferenceTypeOk() (*[]string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *ScimAttributeResponse) SetReferenceType(v []string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *ScimAttributeResponse) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetMeta

`func (o *ScimAttributeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimAttributeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimAttributeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScimAttributeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ScimAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ScimAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ScimAttributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ScimAttributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


