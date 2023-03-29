# ScimSubattributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the SCIM Schema | 
**Schemas** | Pointer to [**[]EnumscimSubattributeSchemaUrn**](EnumscimSubattributeSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this SCIM Subattribute | [optional] 
**Type** | [**EnumscimSubattributeTypeProp**](EnumscimSubattributeTypeProp.md) |  | 
**Required** | **bool** | Specifies whether this sub-attribute is required. | 
**CaseExact** | **bool** | Specifies whether the sub-attribute values are case sensitive. | 
**MultiValued** | **bool** | Specifies whether this attribute may have multiple values. | 
**CanonicalValue** | Pointer to **[]string** | Specifies the suggested canonical type values for the sub-attribute. | [optional] 
**Mutability** | [**EnumscimSubattributeMutabilityProp**](EnumscimSubattributeMutabilityProp.md) |  | 
**Returned** | [**EnumscimSubattributeReturnedProp**](EnumscimSubattributeReturnedProp.md) |  | 
**ReferenceType** | Pointer to **[]string** | Specifies the SCIM resource types that may be referenced. This property is only applicable for sub-attributes that are of type &#39;reference&#39;. Valid values are: A SCIM resource type (e.g., &#39;User&#39; or &#39;Group&#39;), &#39;external&#39; - indicating the resource is an external resource (e.g., such as a photo), or &#39;uri&#39; - indicating that the reference is to a service endpoint or an identifier (such as a schema urn). | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewScimSubattributeResponse

`func NewScimSubattributeResponse(id string, type_ EnumscimSubattributeTypeProp, required bool, caseExact bool, multiValued bool, mutability EnumscimSubattributeMutabilityProp, returned EnumscimSubattributeReturnedProp, ) *ScimSubattributeResponse`

NewScimSubattributeResponse instantiates a new ScimSubattributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSubattributeResponseWithDefaults

`func NewScimSubattributeResponseWithDefaults() *ScimSubattributeResponse`

NewScimSubattributeResponseWithDefaults instantiates a new ScimSubattributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScimSubattributeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimSubattributeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimSubattributeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ScimSubattributeResponse) GetSchemas() []EnumscimSubattributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimSubattributeResponse) GetSchemasOk() (*[]EnumscimSubattributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimSubattributeResponse) SetSchemas(v []EnumscimSubattributeSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimSubattributeResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ScimSubattributeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimSubattributeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimSubattributeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScimSubattributeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ScimSubattributeResponse) GetType() EnumscimSubattributeTypeProp`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScimSubattributeResponse) GetTypeOk() (*EnumscimSubattributeTypeProp, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScimSubattributeResponse) SetType(v EnumscimSubattributeTypeProp)`

SetType sets Type field to given value.


### GetRequired

`func (o *ScimSubattributeResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ScimSubattributeResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ScimSubattributeResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetCaseExact

`func (o *ScimSubattributeResponse) GetCaseExact() bool`

GetCaseExact returns the CaseExact field if non-nil, zero value otherwise.

### GetCaseExactOk

`func (o *ScimSubattributeResponse) GetCaseExactOk() (*bool, bool)`

GetCaseExactOk returns a tuple with the CaseExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExact

`func (o *ScimSubattributeResponse) SetCaseExact(v bool)`

SetCaseExact sets CaseExact field to given value.


### GetMultiValued

`func (o *ScimSubattributeResponse) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *ScimSubattributeResponse) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *ScimSubattributeResponse) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.


### GetCanonicalValue

`func (o *ScimSubattributeResponse) GetCanonicalValue() []string`

GetCanonicalValue returns the CanonicalValue field if non-nil, zero value otherwise.

### GetCanonicalValueOk

`func (o *ScimSubattributeResponse) GetCanonicalValueOk() (*[]string, bool)`

GetCanonicalValueOk returns a tuple with the CanonicalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalValue

`func (o *ScimSubattributeResponse) SetCanonicalValue(v []string)`

SetCanonicalValue sets CanonicalValue field to given value.

### HasCanonicalValue

`func (o *ScimSubattributeResponse) HasCanonicalValue() bool`

HasCanonicalValue returns a boolean if a field has been set.

### GetMutability

`func (o *ScimSubattributeResponse) GetMutability() EnumscimSubattributeMutabilityProp`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *ScimSubattributeResponse) GetMutabilityOk() (*EnumscimSubattributeMutabilityProp, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *ScimSubattributeResponse) SetMutability(v EnumscimSubattributeMutabilityProp)`

SetMutability sets Mutability field to given value.


### GetReturned

`func (o *ScimSubattributeResponse) GetReturned() EnumscimSubattributeReturnedProp`

GetReturned returns the Returned field if non-nil, zero value otherwise.

### GetReturnedOk

`func (o *ScimSubattributeResponse) GetReturnedOk() (*EnumscimSubattributeReturnedProp, bool)`

GetReturnedOk returns a tuple with the Returned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturned

`func (o *ScimSubattributeResponse) SetReturned(v EnumscimSubattributeReturnedProp)`

SetReturned sets Returned field to given value.


### GetReferenceType

`func (o *ScimSubattributeResponse) GetReferenceType() []string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *ScimSubattributeResponse) GetReferenceTypeOk() (*[]string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *ScimSubattributeResponse) SetReferenceType(v []string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *ScimSubattributeResponse) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetMeta

`func (o *ScimSubattributeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimSubattributeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimSubattributeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScimSubattributeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ScimSubattributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ScimSubattributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ScimSubattributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ScimSubattributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


