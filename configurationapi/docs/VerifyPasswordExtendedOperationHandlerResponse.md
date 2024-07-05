# VerifyPasswordExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumverifyPasswordExtendedOperationHandlerSchemaUrn**](EnumverifyPasswordExtendedOperationHandlerSchemaUrn.md) |  | 
**RejectInsecureRequests** | Pointer to **bool** | Indicates whether the server should reject attempts to use this extended operation over an insecure connection. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Extended Operation Handler | 

## Methods

### NewVerifyPasswordExtendedOperationHandlerResponse

`func NewVerifyPasswordExtendedOperationHandlerResponse(schemas []EnumverifyPasswordExtendedOperationHandlerSchemaUrn, enabled bool, id string, ) *VerifyPasswordExtendedOperationHandlerResponse`

NewVerifyPasswordExtendedOperationHandlerResponse instantiates a new VerifyPasswordExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyPasswordExtendedOperationHandlerResponseWithDefaults

`func NewVerifyPasswordExtendedOperationHandlerResponseWithDefaults() *VerifyPasswordExtendedOperationHandlerResponse`

NewVerifyPasswordExtendedOperationHandlerResponseWithDefaults instantiates a new VerifyPasswordExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetSchemas() []EnumverifyPasswordExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumverifyPasswordExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VerifyPasswordExtendedOperationHandlerResponse) SetSchemas(v []EnumverifyPasswordExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRejectInsecureRequests

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetRejectInsecureRequests() bool`

GetRejectInsecureRequests returns the RejectInsecureRequests field if non-nil, zero value otherwise.

### GetRejectInsecureRequestsOk

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetRejectInsecureRequestsOk() (*bool, bool)`

GetRejectInsecureRequestsOk returns a tuple with the RejectInsecureRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectInsecureRequests

`func (o *VerifyPasswordExtendedOperationHandlerResponse) SetRejectInsecureRequests(v bool)`

SetRejectInsecureRequests sets RejectInsecureRequests field to given value.

### HasRejectInsecureRequests

`func (o *VerifyPasswordExtendedOperationHandlerResponse) HasRejectInsecureRequests() bool`

HasRejectInsecureRequests returns a boolean if a field has been set.

### GetDescription

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VerifyPasswordExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VerifyPasswordExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VerifyPasswordExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VerifyPasswordExtendedOperationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VerifyPasswordExtendedOperationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *VerifyPasswordExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *VerifyPasswordExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerifyPasswordExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerifyPasswordExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


