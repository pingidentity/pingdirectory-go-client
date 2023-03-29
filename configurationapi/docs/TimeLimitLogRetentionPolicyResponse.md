# TimeLimitLogRetentionPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log Retention Policy | 
**Schemas** | [**[]EnumtimeLimitLogRetentionPolicySchemaUrn**](EnumtimeLimitLogRetentionPolicySchemaUrn.md) |  | 
**RetainDuration** | **string** | Specifies the desired minimum length of time that each log file should be retained. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewTimeLimitLogRetentionPolicyResponse

`func NewTimeLimitLogRetentionPolicyResponse(id string, schemas []EnumtimeLimitLogRetentionPolicySchemaUrn, retainDuration string, ) *TimeLimitLogRetentionPolicyResponse`

NewTimeLimitLogRetentionPolicyResponse instantiates a new TimeLimitLogRetentionPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeLimitLogRetentionPolicyResponseWithDefaults

`func NewTimeLimitLogRetentionPolicyResponseWithDefaults() *TimeLimitLogRetentionPolicyResponse`

NewTimeLimitLogRetentionPolicyResponseWithDefaults instantiates a new TimeLimitLogRetentionPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeLimitLogRetentionPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeLimitLogRetentionPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeLimitLogRetentionPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *TimeLimitLogRetentionPolicyResponse) GetSchemas() []EnumtimeLimitLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TimeLimitLogRetentionPolicyResponse) GetSchemasOk() (*[]EnumtimeLimitLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TimeLimitLogRetentionPolicyResponse) SetSchemas(v []EnumtimeLimitLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRetainDuration

`func (o *TimeLimitLogRetentionPolicyResponse) GetRetainDuration() string`

GetRetainDuration returns the RetainDuration field if non-nil, zero value otherwise.

### GetRetainDurationOk

`func (o *TimeLimitLogRetentionPolicyResponse) GetRetainDurationOk() (*string, bool)`

GetRetainDurationOk returns a tuple with the RetainDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainDuration

`func (o *TimeLimitLogRetentionPolicyResponse) SetRetainDuration(v string)`

SetRetainDuration sets RetainDuration field to given value.


### GetDescription

`func (o *TimeLimitLogRetentionPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TimeLimitLogRetentionPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TimeLimitLogRetentionPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TimeLimitLogRetentionPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *TimeLimitLogRetentionPolicyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TimeLimitLogRetentionPolicyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TimeLimitLogRetentionPolicyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TimeLimitLogRetentionPolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *TimeLimitLogRetentionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *TimeLimitLogRetentionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *TimeLimitLogRetentionPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *TimeLimitLogRetentionPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


