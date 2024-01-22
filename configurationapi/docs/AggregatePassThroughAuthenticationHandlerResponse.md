# AggregatePassThroughAuthenticationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaggregatePassThroughAuthenticationHandlerSchemaUrn**](EnumaggregatePassThroughAuthenticationHandlerSchemaUrn.md) |  | 
**SubordinatePassThroughAuthenticationHandler** | **[]string** | The set of subordinate pass-through authentication handlers that may be used to perform the authentication processing. Handlers will be invoked in order until one is found for which the bind operation matches the associated criteria and either succeeds or fails in a manner that should not be ignored. | 
**ContinueOnFailureType** | Pointer to [**[]EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp**](EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Pass Through Authentication Handler | [optional] 
**IncludedLocalEntryBaseDN** | Pointer to **[]string** | The base DNs for the local users whose authentication attempts may be passed through to the external authentication service. | [optional] 
**ConnectionCriteria** | Pointer to **string** | A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**RequestCriteria** | Pointer to **string** | A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Pass Through Authentication Handler | 

## Methods

### NewAggregatePassThroughAuthenticationHandlerResponse

`func NewAggregatePassThroughAuthenticationHandlerResponse(schemas []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn, subordinatePassThroughAuthenticationHandler []string, id string, ) *AggregatePassThroughAuthenticationHandlerResponse`

NewAggregatePassThroughAuthenticationHandlerResponse instantiates a new AggregatePassThroughAuthenticationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatePassThroughAuthenticationHandlerResponseWithDefaults

`func NewAggregatePassThroughAuthenticationHandlerResponseWithDefaults() *AggregatePassThroughAuthenticationHandlerResponse`

NewAggregatePassThroughAuthenticationHandlerResponseWithDefaults instantiates a new AggregatePassThroughAuthenticationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetSchemas() []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetSchemasOk() (*[]EnumaggregatePassThroughAuthenticationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregatePassThroughAuthenticationHandlerResponse) SetSchemas(v []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSubordinatePassThroughAuthenticationHandler

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetSubordinatePassThroughAuthenticationHandler() []string`

GetSubordinatePassThroughAuthenticationHandler returns the SubordinatePassThroughAuthenticationHandler field if non-nil, zero value otherwise.

### GetSubordinatePassThroughAuthenticationHandlerOk

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetSubordinatePassThroughAuthenticationHandlerOk() (*[]string, bool)`

GetSubordinatePassThroughAuthenticationHandlerOk returns a tuple with the SubordinatePassThroughAuthenticationHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubordinatePassThroughAuthenticationHandler

`func (o *AggregatePassThroughAuthenticationHandlerResponse) SetSubordinatePassThroughAuthenticationHandler(v []string)`

SetSubordinatePassThroughAuthenticationHandler sets SubordinatePassThroughAuthenticationHandler field to given value.


### GetContinueOnFailureType

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetContinueOnFailureType() []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp`

GetContinueOnFailureType returns the ContinueOnFailureType field if non-nil, zero value otherwise.

### GetContinueOnFailureTypeOk

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetContinueOnFailureTypeOk() (*[]EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp, bool)`

GetContinueOnFailureTypeOk returns a tuple with the ContinueOnFailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnFailureType

`func (o *AggregatePassThroughAuthenticationHandlerResponse) SetContinueOnFailureType(v []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp)`

SetContinueOnFailureType sets ContinueOnFailureType field to given value.

### HasContinueOnFailureType

`func (o *AggregatePassThroughAuthenticationHandlerResponse) HasContinueOnFailureType() bool`

HasContinueOnFailureType returns a boolean if a field has been set.

### GetDescription

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregatePassThroughAuthenticationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregatePassThroughAuthenticationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludedLocalEntryBaseDN

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *AggregatePassThroughAuthenticationHandlerResponse) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *AggregatePassThroughAuthenticationHandlerResponse) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AggregatePassThroughAuthenticationHandlerResponse) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AggregatePassThroughAuthenticationHandlerResponse) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AggregatePassThroughAuthenticationHandlerResponse) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AggregatePassThroughAuthenticationHandlerResponse) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetMeta

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AggregatePassThroughAuthenticationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AggregatePassThroughAuthenticationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AggregatePassThroughAuthenticationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AggregatePassThroughAuthenticationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregatePassThroughAuthenticationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregatePassThroughAuthenticationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


