# JsonAttributeConstraintsShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumjsonAttributeConstraintsSchemaUrn**](EnumjsonAttributeConstraintsSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this JSON Attribute Constraints | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this JSON Attribute Constraints is enabled. | [optional] 
**AttributeType** | **string** | The name or OID of the LDAP attribute type whose values will be subject to the associated field constraints. This attribute type must be defined in the server schema, and it must have a \&quot;JSON object\&quot; syntax. | 
**AllowUnnamedFields** | Pointer to **bool** | Indicates whether JSON objects stored as values of attributes with the associated attribute-type will be permitted to include fields for which there is no subordinate json-field-constraints definition. If unnamed fields are allowed, then no constraints will be imposed on the values of those fields. However, if unnamed fields are not allowed, then the server will reject any attempt to store a JSON object with a field for which there is no corresponding json-fields-constraints definition. | [optional] 

## Methods

### NewJsonAttributeConstraintsShared

`func NewJsonAttributeConstraintsShared(attributeType string, ) *JsonAttributeConstraintsShared`

NewJsonAttributeConstraintsShared instantiates a new JsonAttributeConstraintsShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonAttributeConstraintsSharedWithDefaults

`func NewJsonAttributeConstraintsSharedWithDefaults() *JsonAttributeConstraintsShared`

NewJsonAttributeConstraintsSharedWithDefaults instantiates a new JsonAttributeConstraintsShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *JsonAttributeConstraintsShared) GetSchemas() []EnumjsonAttributeConstraintsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JsonAttributeConstraintsShared) GetSchemasOk() (*[]EnumjsonAttributeConstraintsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JsonAttributeConstraintsShared) SetSchemas(v []EnumjsonAttributeConstraintsSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *JsonAttributeConstraintsShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *JsonAttributeConstraintsShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JsonAttributeConstraintsShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JsonAttributeConstraintsShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JsonAttributeConstraintsShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JsonAttributeConstraintsShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JsonAttributeConstraintsShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JsonAttributeConstraintsShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JsonAttributeConstraintsShared) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAttributeType

`func (o *JsonAttributeConstraintsShared) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *JsonAttributeConstraintsShared) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *JsonAttributeConstraintsShared) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetAllowUnnamedFields

`func (o *JsonAttributeConstraintsShared) GetAllowUnnamedFields() bool`

GetAllowUnnamedFields returns the AllowUnnamedFields field if non-nil, zero value otherwise.

### GetAllowUnnamedFieldsOk

`func (o *JsonAttributeConstraintsShared) GetAllowUnnamedFieldsOk() (*bool, bool)`

GetAllowUnnamedFieldsOk returns a tuple with the AllowUnnamedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnnamedFields

`func (o *JsonAttributeConstraintsShared) SetAllowUnnamedFields(v bool)`

SetAllowUnnamedFields sets AllowUnnamedFields field to given value.

### HasAllowUnnamedFields

`func (o *JsonAttributeConstraintsShared) HasAllowUnnamedFields() bool`

HasAllowUnnamedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


