# ScimSchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumscimSchemaSchemaUrn**](EnumscimSchemaSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this SCIM Schema | [optional] 
**SchemaURN** | **string** | The URN which identifies this SCIM Schema. | 
**DisplayName** | Pointer to **string** | The human readable name for this SCIM Schema. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the SCIM Schema | 

## Methods

### NewScimSchemaResponse

`func NewScimSchemaResponse(schemaURN string, id string, ) *ScimSchemaResponse`

NewScimSchemaResponse instantiates a new ScimSchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSchemaResponseWithDefaults

`func NewScimSchemaResponseWithDefaults() *ScimSchemaResponse`

NewScimSchemaResponseWithDefaults instantiates a new ScimSchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimSchemaResponse) GetSchemas() []EnumscimSchemaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimSchemaResponse) GetSchemasOk() (*[]EnumscimSchemaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimSchemaResponse) SetSchemas(v []EnumscimSchemaSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimSchemaResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ScimSchemaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimSchemaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimSchemaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScimSchemaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchemaURN

`func (o *ScimSchemaResponse) GetSchemaURN() string`

GetSchemaURN returns the SchemaURN field if non-nil, zero value otherwise.

### GetSchemaURNOk

`func (o *ScimSchemaResponse) GetSchemaURNOk() (*string, bool)`

GetSchemaURNOk returns a tuple with the SchemaURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaURN

`func (o *ScimSchemaResponse) SetSchemaURN(v string)`

SetSchemaURN sets SchemaURN field to given value.


### GetDisplayName

`func (o *ScimSchemaResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimSchemaResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimSchemaResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ScimSchemaResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMeta

`func (o *ScimSchemaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimSchemaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimSchemaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScimSchemaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ScimSchemaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ScimSchemaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ScimSchemaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ScimSchemaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ScimSchemaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimSchemaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimSchemaResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


