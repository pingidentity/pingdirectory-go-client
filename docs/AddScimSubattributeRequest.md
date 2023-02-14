# AddScimSubattributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubattributeName** | **string** | Name of the new SCIM Subattribute | 
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

## Methods

### NewAddScimSubattributeRequest

`func NewAddScimSubattributeRequest(subattributeName string, type_ EnumscimSubattributeTypeProp, required bool, caseExact bool, multiValued bool, mutability EnumscimSubattributeMutabilityProp, returned EnumscimSubattributeReturnedProp, ) *AddScimSubattributeRequest`

NewAddScimSubattributeRequest instantiates a new AddScimSubattributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddScimSubattributeRequestWithDefaults

`func NewAddScimSubattributeRequestWithDefaults() *AddScimSubattributeRequest`

NewAddScimSubattributeRequestWithDefaults instantiates a new AddScimSubattributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubattributeName

`func (o *AddScimSubattributeRequest) GetSubattributeName() string`

GetSubattributeName returns the SubattributeName field if non-nil, zero value otherwise.

### GetSubattributeNameOk

`func (o *AddScimSubattributeRequest) GetSubattributeNameOk() (*string, bool)`

GetSubattributeNameOk returns a tuple with the SubattributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubattributeName

`func (o *AddScimSubattributeRequest) SetSubattributeName(v string)`

SetSubattributeName sets SubattributeName field to given value.


### GetSchemas

`func (o *AddScimSubattributeRequest) GetSchemas() []EnumscimSubattributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddScimSubattributeRequest) GetSchemasOk() (*[]EnumscimSubattributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddScimSubattributeRequest) SetSchemas(v []EnumscimSubattributeSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddScimSubattributeRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddScimSubattributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddScimSubattributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddScimSubattributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddScimSubattributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *AddScimSubattributeRequest) GetType() EnumscimSubattributeTypeProp`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddScimSubattributeRequest) GetTypeOk() (*EnumscimSubattributeTypeProp, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddScimSubattributeRequest) SetType(v EnumscimSubattributeTypeProp)`

SetType sets Type field to given value.


### GetRequired

`func (o *AddScimSubattributeRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AddScimSubattributeRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AddScimSubattributeRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetCaseExact

`func (o *AddScimSubattributeRequest) GetCaseExact() bool`

GetCaseExact returns the CaseExact field if non-nil, zero value otherwise.

### GetCaseExactOk

`func (o *AddScimSubattributeRequest) GetCaseExactOk() (*bool, bool)`

GetCaseExactOk returns a tuple with the CaseExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExact

`func (o *AddScimSubattributeRequest) SetCaseExact(v bool)`

SetCaseExact sets CaseExact field to given value.


### GetMultiValued

`func (o *AddScimSubattributeRequest) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *AddScimSubattributeRequest) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *AddScimSubattributeRequest) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.


### GetCanonicalValue

`func (o *AddScimSubattributeRequest) GetCanonicalValue() []string`

GetCanonicalValue returns the CanonicalValue field if non-nil, zero value otherwise.

### GetCanonicalValueOk

`func (o *AddScimSubattributeRequest) GetCanonicalValueOk() (*[]string, bool)`

GetCanonicalValueOk returns a tuple with the CanonicalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalValue

`func (o *AddScimSubattributeRequest) SetCanonicalValue(v []string)`

SetCanonicalValue sets CanonicalValue field to given value.

### HasCanonicalValue

`func (o *AddScimSubattributeRequest) HasCanonicalValue() bool`

HasCanonicalValue returns a boolean if a field has been set.

### GetMutability

`func (o *AddScimSubattributeRequest) GetMutability() EnumscimSubattributeMutabilityProp`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *AddScimSubattributeRequest) GetMutabilityOk() (*EnumscimSubattributeMutabilityProp, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *AddScimSubattributeRequest) SetMutability(v EnumscimSubattributeMutabilityProp)`

SetMutability sets Mutability field to given value.


### GetReturned

`func (o *AddScimSubattributeRequest) GetReturned() EnumscimSubattributeReturnedProp`

GetReturned returns the Returned field if non-nil, zero value otherwise.

### GetReturnedOk

`func (o *AddScimSubattributeRequest) GetReturnedOk() (*EnumscimSubattributeReturnedProp, bool)`

GetReturnedOk returns a tuple with the Returned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturned

`func (o *AddScimSubattributeRequest) SetReturned(v EnumscimSubattributeReturnedProp)`

SetReturned sets Returned field to given value.


### GetReferenceType

`func (o *AddScimSubattributeRequest) GetReferenceType() []string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *AddScimSubattributeRequest) GetReferenceTypeOk() (*[]string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *AddScimSubattributeRequest) SetReferenceType(v []string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *AddScimSubattributeRequest) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


