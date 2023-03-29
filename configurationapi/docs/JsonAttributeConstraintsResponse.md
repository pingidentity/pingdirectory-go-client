# JsonAttributeConstraintsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the JSON Attribute Constraints | 
**Schemas** | Pointer to [**[]EnumjsonAttributeConstraintsSchemaUrn**](EnumjsonAttributeConstraintsSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this JSON Attribute Constraints | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this JSON Attribute Constraints is enabled. | [optional] 
**AttributeType** | **string** | The name or OID of the LDAP attribute type whose values will be subject to the associated field constraints. This attribute type must be defined in the server schema, and it must have a \&quot;JSON object\&quot; syntax. | 
**AllowUnnamedFields** | Pointer to **bool** | Indicates whether JSON objects stored as values of attributes with the associated attribute-type will be permitted to include fields for which there is no subordinate json-field-constraints definition. If unnamed fields are allowed, then no constraints will be imposed on the values of those fields. However, if unnamed fields are not allowed, then the server will reject any attempt to store a JSON object with a field for which there is no corresponding json-fields-constraints definition. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewJsonAttributeConstraintsResponse

`func NewJsonAttributeConstraintsResponse(id string, attributeType string, ) *JsonAttributeConstraintsResponse`

NewJsonAttributeConstraintsResponse instantiates a new JsonAttributeConstraintsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonAttributeConstraintsResponseWithDefaults

`func NewJsonAttributeConstraintsResponseWithDefaults() *JsonAttributeConstraintsResponse`

NewJsonAttributeConstraintsResponseWithDefaults instantiates a new JsonAttributeConstraintsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JsonAttributeConstraintsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonAttributeConstraintsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonAttributeConstraintsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *JsonAttributeConstraintsResponse) GetSchemas() []EnumjsonAttributeConstraintsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JsonAttributeConstraintsResponse) GetSchemasOk() (*[]EnumjsonAttributeConstraintsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JsonAttributeConstraintsResponse) SetSchemas(v []EnumjsonAttributeConstraintsSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *JsonAttributeConstraintsResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *JsonAttributeConstraintsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JsonAttributeConstraintsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JsonAttributeConstraintsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JsonAttributeConstraintsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JsonAttributeConstraintsResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JsonAttributeConstraintsResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JsonAttributeConstraintsResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JsonAttributeConstraintsResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAttributeType

`func (o *JsonAttributeConstraintsResponse) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *JsonAttributeConstraintsResponse) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *JsonAttributeConstraintsResponse) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetAllowUnnamedFields

`func (o *JsonAttributeConstraintsResponse) GetAllowUnnamedFields() bool`

GetAllowUnnamedFields returns the AllowUnnamedFields field if non-nil, zero value otherwise.

### GetAllowUnnamedFieldsOk

`func (o *JsonAttributeConstraintsResponse) GetAllowUnnamedFieldsOk() (*bool, bool)`

GetAllowUnnamedFieldsOk returns a tuple with the AllowUnnamedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnnamedFields

`func (o *JsonAttributeConstraintsResponse) SetAllowUnnamedFields(v bool)`

SetAllowUnnamedFields sets AllowUnnamedFields field to given value.

### HasAllowUnnamedFields

`func (o *JsonAttributeConstraintsResponse) HasAllowUnnamedFields() bool`

HasAllowUnnamedFields returns a boolean if a field has been set.

### GetMeta

`func (o *JsonAttributeConstraintsResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JsonAttributeConstraintsResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JsonAttributeConstraintsResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JsonAttributeConstraintsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *JsonAttributeConstraintsResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *JsonAttributeConstraintsResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *JsonAttributeConstraintsResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *JsonAttributeConstraintsResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


