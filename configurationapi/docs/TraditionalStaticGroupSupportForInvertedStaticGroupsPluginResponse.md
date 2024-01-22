# TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn**](EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn.md) |  | 
**TraditionalStaticGroupObjectClass** | Pointer to [**EnumpluginTraditionalStaticGroupObjectClassProp**](EnumpluginTraditionalStaticGroupObjectClassProp.md) |  | [optional] 
**MaximumMembershipUpdatesPerModify** | Pointer to **int64** | An integer property that specifies the maximum number of membership changes that will be supported in a single modify operation. A value of zero indicates that modify operations targeting the group entry should not be permitted to alter the set of members for the group. | [optional] 
**ReadOperationSupport** | Pointer to [**EnumpluginReadOperationSupportProp**](EnumpluginReadOperationSupportProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 

## Methods

### NewTraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse

`func NewTraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse(schemas []EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn, enabled bool, id string, ) *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse`

NewTraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse instantiates a new TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponseWithDefaults

`func NewTraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponseWithDefaults() *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse`

NewTraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponseWithDefaults instantiates a new TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetSchemas() []EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetSchemasOk() (*[]EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) SetSchemas(v []EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTraditionalStaticGroupObjectClass

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetTraditionalStaticGroupObjectClass() EnumpluginTraditionalStaticGroupObjectClassProp`

GetTraditionalStaticGroupObjectClass returns the TraditionalStaticGroupObjectClass field if non-nil, zero value otherwise.

### GetTraditionalStaticGroupObjectClassOk

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetTraditionalStaticGroupObjectClassOk() (*EnumpluginTraditionalStaticGroupObjectClassProp, bool)`

GetTraditionalStaticGroupObjectClassOk returns a tuple with the TraditionalStaticGroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraditionalStaticGroupObjectClass

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) SetTraditionalStaticGroupObjectClass(v EnumpluginTraditionalStaticGroupObjectClassProp)`

SetTraditionalStaticGroupObjectClass sets TraditionalStaticGroupObjectClass field to given value.

### HasTraditionalStaticGroupObjectClass

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) HasTraditionalStaticGroupObjectClass() bool`

HasTraditionalStaticGroupObjectClass returns a boolean if a field has been set.

### GetMaximumMembershipUpdatesPerModify

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetMaximumMembershipUpdatesPerModify() int64`

GetMaximumMembershipUpdatesPerModify returns the MaximumMembershipUpdatesPerModify field if non-nil, zero value otherwise.

### GetMaximumMembershipUpdatesPerModifyOk

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetMaximumMembershipUpdatesPerModifyOk() (*int64, bool)`

GetMaximumMembershipUpdatesPerModifyOk returns a tuple with the MaximumMembershipUpdatesPerModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMembershipUpdatesPerModify

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) SetMaximumMembershipUpdatesPerModify(v int64)`

SetMaximumMembershipUpdatesPerModify sets MaximumMembershipUpdatesPerModify field to given value.

### HasMaximumMembershipUpdatesPerModify

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) HasMaximumMembershipUpdatesPerModify() bool`

HasMaximumMembershipUpdatesPerModify returns a boolean if a field has been set.

### GetReadOperationSupport

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetReadOperationSupport() EnumpluginReadOperationSupportProp`

GetReadOperationSupport returns the ReadOperationSupport field if non-nil, zero value otherwise.

### GetReadOperationSupportOk

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetReadOperationSupportOk() (*EnumpluginReadOperationSupportProp, bool)`

GetReadOperationSupportOk returns a tuple with the ReadOperationSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOperationSupport

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) SetReadOperationSupport(v EnumpluginReadOperationSupportProp)`

SetReadOperationSupport sets ReadOperationSupport field to given value.

### HasReadOperationSupport

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) HasReadOperationSupport() bool`

HasReadOperationSupport returns a boolean if a field has been set.

### GetDescription

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetMeta

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TraditionalStaticGroupSupportForInvertedStaticGroupsPluginResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


