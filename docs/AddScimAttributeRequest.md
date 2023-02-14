# AddScimAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | **string** | Name of the new SCIM Attribute | 
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

## Methods

### NewAddScimAttributeRequest

`func NewAddScimAttributeRequest(attributeName string, name string, type_ EnumscimAttributeTypeProp, required bool, caseExact bool, multiValued bool, mutability EnumscimAttributeMutabilityProp, returned EnumscimAttributeReturnedProp, ) *AddScimAttributeRequest`

NewAddScimAttributeRequest instantiates a new AddScimAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddScimAttributeRequestWithDefaults

`func NewAddScimAttributeRequestWithDefaults() *AddScimAttributeRequest`

NewAddScimAttributeRequestWithDefaults instantiates a new AddScimAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *AddScimAttributeRequest) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *AddScimAttributeRequest) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *AddScimAttributeRequest) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetSchemas

`func (o *AddScimAttributeRequest) GetSchemas() []EnumscimAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddScimAttributeRequest) GetSchemasOk() (*[]EnumscimAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddScimAttributeRequest) SetSchemas(v []EnumscimAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddScimAttributeRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddScimAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddScimAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddScimAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddScimAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AddScimAttributeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddScimAttributeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddScimAttributeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddScimAttributeRequest) GetType() EnumscimAttributeTypeProp`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddScimAttributeRequest) GetTypeOk() (*EnumscimAttributeTypeProp, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddScimAttributeRequest) SetType(v EnumscimAttributeTypeProp)`

SetType sets Type field to given value.


### GetRequired

`func (o *AddScimAttributeRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AddScimAttributeRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AddScimAttributeRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetCaseExact

`func (o *AddScimAttributeRequest) GetCaseExact() bool`

GetCaseExact returns the CaseExact field if non-nil, zero value otherwise.

### GetCaseExactOk

`func (o *AddScimAttributeRequest) GetCaseExactOk() (*bool, bool)`

GetCaseExactOk returns a tuple with the CaseExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExact

`func (o *AddScimAttributeRequest) SetCaseExact(v bool)`

SetCaseExact sets CaseExact field to given value.


### GetMultiValued

`func (o *AddScimAttributeRequest) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *AddScimAttributeRequest) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *AddScimAttributeRequest) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.


### GetCanonicalValue

`func (o *AddScimAttributeRequest) GetCanonicalValue() []string`

GetCanonicalValue returns the CanonicalValue field if non-nil, zero value otherwise.

### GetCanonicalValueOk

`func (o *AddScimAttributeRequest) GetCanonicalValueOk() (*[]string, bool)`

GetCanonicalValueOk returns a tuple with the CanonicalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalValue

`func (o *AddScimAttributeRequest) SetCanonicalValue(v []string)`

SetCanonicalValue sets CanonicalValue field to given value.

### HasCanonicalValue

`func (o *AddScimAttributeRequest) HasCanonicalValue() bool`

HasCanonicalValue returns a boolean if a field has been set.

### GetMutability

`func (o *AddScimAttributeRequest) GetMutability() EnumscimAttributeMutabilityProp`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *AddScimAttributeRequest) GetMutabilityOk() (*EnumscimAttributeMutabilityProp, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *AddScimAttributeRequest) SetMutability(v EnumscimAttributeMutabilityProp)`

SetMutability sets Mutability field to given value.


### GetReturned

`func (o *AddScimAttributeRequest) GetReturned() EnumscimAttributeReturnedProp`

GetReturned returns the Returned field if non-nil, zero value otherwise.

### GetReturnedOk

`func (o *AddScimAttributeRequest) GetReturnedOk() (*EnumscimAttributeReturnedProp, bool)`

GetReturnedOk returns a tuple with the Returned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturned

`func (o *AddScimAttributeRequest) SetReturned(v EnumscimAttributeReturnedProp)`

SetReturned sets Returned field to given value.


### GetReferenceType

`func (o *AddScimAttributeRequest) GetReferenceType() []string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *AddScimAttributeRequest) GetReferenceTypeOk() (*[]string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *AddScimAttributeRequest) SetReferenceType(v []string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *AddScimAttributeRequest) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


