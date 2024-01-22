# AddAggregatePassThroughAuthenticationHandlerRequest

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
**HandlerName** | **string** | Name of the new Pass Through Authentication Handler | 

## Methods

### NewAddAggregatePassThroughAuthenticationHandlerRequest

`func NewAddAggregatePassThroughAuthenticationHandlerRequest(schemas []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn, subordinatePassThroughAuthenticationHandler []string, handlerName string, ) *AddAggregatePassThroughAuthenticationHandlerRequest`

NewAddAggregatePassThroughAuthenticationHandlerRequest instantiates a new AddAggregatePassThroughAuthenticationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAggregatePassThroughAuthenticationHandlerRequestWithDefaults

`func NewAddAggregatePassThroughAuthenticationHandlerRequestWithDefaults() *AddAggregatePassThroughAuthenticationHandlerRequest`

NewAddAggregatePassThroughAuthenticationHandlerRequestWithDefaults instantiates a new AddAggregatePassThroughAuthenticationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetSchemas() []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetSchemasOk() (*[]EnumaggregatePassThroughAuthenticationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetSchemas(v []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSubordinatePassThroughAuthenticationHandler

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetSubordinatePassThroughAuthenticationHandler() []string`

GetSubordinatePassThroughAuthenticationHandler returns the SubordinatePassThroughAuthenticationHandler field if non-nil, zero value otherwise.

### GetSubordinatePassThroughAuthenticationHandlerOk

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetSubordinatePassThroughAuthenticationHandlerOk() (*[]string, bool)`

GetSubordinatePassThroughAuthenticationHandlerOk returns a tuple with the SubordinatePassThroughAuthenticationHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubordinatePassThroughAuthenticationHandler

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetSubordinatePassThroughAuthenticationHandler(v []string)`

SetSubordinatePassThroughAuthenticationHandler sets SubordinatePassThroughAuthenticationHandler field to given value.


### GetContinueOnFailureType

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetContinueOnFailureType() []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp`

GetContinueOnFailureType returns the ContinueOnFailureType field if non-nil, zero value otherwise.

### GetContinueOnFailureTypeOk

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetContinueOnFailureTypeOk() (*[]EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp, bool)`

GetContinueOnFailureTypeOk returns a tuple with the ContinueOnFailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnFailureType

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetContinueOnFailureType(v []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp)`

SetContinueOnFailureType sets ContinueOnFailureType field to given value.

### HasContinueOnFailureType

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) HasContinueOnFailureType() bool`

HasContinueOnFailureType returns a boolean if a field has been set.

### GetDescription

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludedLocalEntryBaseDN

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetHandlerName

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


