# CoalesceModificationsPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumcoalesceModificationsPluginSchemaUrn**](EnumcoalesceModificationsPluginSchemaUrn.md) |  | 
**RequestCriteria** | **string** | A reference to request criteria that indicates which modify requests should be coalesced. | 
**AllowedRequestControl** | Pointer to **[]string** | Specifies the OIDs of the controls that are allowed to be present in operations to coalesce. These controls are passed through when the request is validated, but they will not be included when the background thread applies the coalesced modify requests. | [optional] 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 

## Methods

### NewCoalesceModificationsPluginResponse

`func NewCoalesceModificationsPluginResponse(schemas []EnumcoalesceModificationsPluginSchemaUrn, requestCriteria string, enabled bool, id string, ) *CoalesceModificationsPluginResponse`

NewCoalesceModificationsPluginResponse instantiates a new CoalesceModificationsPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoalesceModificationsPluginResponseWithDefaults

`func NewCoalesceModificationsPluginResponseWithDefaults() *CoalesceModificationsPluginResponse`

NewCoalesceModificationsPluginResponseWithDefaults instantiates a new CoalesceModificationsPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *CoalesceModificationsPluginResponse) GetSchemas() []EnumcoalesceModificationsPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CoalesceModificationsPluginResponse) GetSchemasOk() (*[]EnumcoalesceModificationsPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CoalesceModificationsPluginResponse) SetSchemas(v []EnumcoalesceModificationsPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *CoalesceModificationsPluginResponse) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *CoalesceModificationsPluginResponse) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *CoalesceModificationsPluginResponse) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.


### GetAllowedRequestControl

`func (o *CoalesceModificationsPluginResponse) GetAllowedRequestControl() []string`

GetAllowedRequestControl returns the AllowedRequestControl field if non-nil, zero value otherwise.

### GetAllowedRequestControlOk

`func (o *CoalesceModificationsPluginResponse) GetAllowedRequestControlOk() (*[]string, bool)`

GetAllowedRequestControlOk returns a tuple with the AllowedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequestControl

`func (o *CoalesceModificationsPluginResponse) SetAllowedRequestControl(v []string)`

SetAllowedRequestControl sets AllowedRequestControl field to given value.

### HasAllowedRequestControl

`func (o *CoalesceModificationsPluginResponse) HasAllowedRequestControl() bool`

HasAllowedRequestControl returns a boolean if a field has been set.

### GetInvokeForInternalOperations

`func (o *CoalesceModificationsPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *CoalesceModificationsPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *CoalesceModificationsPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *CoalesceModificationsPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetDescription

`func (o *CoalesceModificationsPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoalesceModificationsPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoalesceModificationsPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoalesceModificationsPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CoalesceModificationsPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CoalesceModificationsPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CoalesceModificationsPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *CoalesceModificationsPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CoalesceModificationsPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CoalesceModificationsPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CoalesceModificationsPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CoalesceModificationsPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CoalesceModificationsPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CoalesceModificationsPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CoalesceModificationsPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *CoalesceModificationsPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoalesceModificationsPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoalesceModificationsPluginResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


