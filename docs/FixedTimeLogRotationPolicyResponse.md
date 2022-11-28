# FixedTimeLogRotationPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Log Rotation Policy | 
**Schemas** | [**[]EnumfixedTimeLogRotationPolicySchemaUrn**](EnumfixedTimeLogRotationPolicySchemaUrn.md) |  | 
**TimeOfDay** | **[]string** | Specifies the time of day at which log rotation should occur. | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 

## Methods

### NewFixedTimeLogRotationPolicyResponse

`func NewFixedTimeLogRotationPolicyResponse(id string, schemas []EnumfixedTimeLogRotationPolicySchemaUrn, timeOfDay []string, ) *FixedTimeLogRotationPolicyResponse`

NewFixedTimeLogRotationPolicyResponse instantiates a new FixedTimeLogRotationPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedTimeLogRotationPolicyResponseWithDefaults

`func NewFixedTimeLogRotationPolicyResponseWithDefaults() *FixedTimeLogRotationPolicyResponse`

NewFixedTimeLogRotationPolicyResponseWithDefaults instantiates a new FixedTimeLogRotationPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *FixedTimeLogRotationPolicyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FixedTimeLogRotationPolicyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FixedTimeLogRotationPolicyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FixedTimeLogRotationPolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *FixedTimeLogRotationPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *FixedTimeLogRotationPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *FixedTimeLogRotationPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *FixedTimeLogRotationPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *FixedTimeLogRotationPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FixedTimeLogRotationPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FixedTimeLogRotationPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *FixedTimeLogRotationPolicyResponse) GetSchemas() []EnumfixedTimeLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FixedTimeLogRotationPolicyResponse) GetSchemasOk() (*[]EnumfixedTimeLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FixedTimeLogRotationPolicyResponse) SetSchemas(v []EnumfixedTimeLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTimeOfDay

`func (o *FixedTimeLogRotationPolicyResponse) GetTimeOfDay() []string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *FixedTimeLogRotationPolicyResponse) GetTimeOfDayOk() (*[]string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *FixedTimeLogRotationPolicyResponse) SetTimeOfDay(v []string)`

SetTimeOfDay sets TimeOfDay field to given value.


### GetDescription

`func (o *FixedTimeLogRotationPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FixedTimeLogRotationPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FixedTimeLogRotationPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FixedTimeLogRotationPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


