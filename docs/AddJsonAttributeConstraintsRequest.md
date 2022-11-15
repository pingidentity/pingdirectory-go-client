# AddJsonAttributeConstraintsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeType** | **string** | The name or OID of the LDAP attribute type whose values will be subject to the associated field constraints. This attribute type must be defined in the server schema, and it must have a \&quot;JSON object\&quot; syntax. | 
**Schemas** | Pointer to [**[]EnumjsonAttributeConstraintsSchemaUrn**](EnumjsonAttributeConstraintsSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this JSON Attribute Constraints | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this JSON Attribute Constraints is enabled. | [optional] 
**AllowUnnamedFields** | Pointer to **bool** | Indicates whether JSON objects stored as values of attributes with the associated attribute-type will be permitted to include fields for which there is no subordinate json-field-constraints definition. If unnamed fields are allowed, then no constraints will be imposed on the values of those fields. However, if unnamed fields are not allowed, then the server will reject any attempt to store a JSON object with a field for which there is no corresponding json-fields-constraints definition. | [optional] 

## Methods

### NewAddJsonAttributeConstraintsRequest

`func NewAddJsonAttributeConstraintsRequest(attributeType string, ) *AddJsonAttributeConstraintsRequest`

NewAddJsonAttributeConstraintsRequest instantiates a new AddJsonAttributeConstraintsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJsonAttributeConstraintsRequestWithDefaults

`func NewAddJsonAttributeConstraintsRequestWithDefaults() *AddJsonAttributeConstraintsRequest`

NewAddJsonAttributeConstraintsRequestWithDefaults instantiates a new AddJsonAttributeConstraintsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeType

`func (o *AddJsonAttributeConstraintsRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddJsonAttributeConstraintsRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddJsonAttributeConstraintsRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetSchemas

`func (o *AddJsonAttributeConstraintsRequest) GetSchemas() []EnumjsonAttributeConstraintsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddJsonAttributeConstraintsRequest) GetSchemasOk() (*[]EnumjsonAttributeConstraintsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddJsonAttributeConstraintsRequest) SetSchemas(v []EnumjsonAttributeConstraintsSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddJsonAttributeConstraintsRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddJsonAttributeConstraintsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddJsonAttributeConstraintsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddJsonAttributeConstraintsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddJsonAttributeConstraintsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddJsonAttributeConstraintsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddJsonAttributeConstraintsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddJsonAttributeConstraintsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddJsonAttributeConstraintsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAllowUnnamedFields

`func (o *AddJsonAttributeConstraintsRequest) GetAllowUnnamedFields() bool`

GetAllowUnnamedFields returns the AllowUnnamedFields field if non-nil, zero value otherwise.

### GetAllowUnnamedFieldsOk

`func (o *AddJsonAttributeConstraintsRequest) GetAllowUnnamedFieldsOk() (*bool, bool)`

GetAllowUnnamedFieldsOk returns a tuple with the AllowUnnamedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnnamedFields

`func (o *AddJsonAttributeConstraintsRequest) SetAllowUnnamedFields(v bool)`

SetAllowUnnamedFields sets AllowUnnamedFields field to given value.

### HasAllowUnnamedFields

`func (o *AddJsonAttributeConstraintsRequest) HasAllowUnnamedFields() bool`

HasAllowUnnamedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


