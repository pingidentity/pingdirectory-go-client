# InternalSearchRatePluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnuminternalSearchRatePluginSchemaUrn**](EnuminternalSearchRatePluginSchemaUrn.md) |  | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | 
**NumThreads** | **int64** | Specifies the number of concurrent threads that should be used to process the search operations. | 
**BaseDN** | **string** | Specifies the base DN to use for the searches to perform. | 
**LowerBound** | Pointer to **int64** | Specifies the lower bound for the numeric value which will be inserted into the search filter. | [optional] 
**UpperBound** | Pointer to **int64** | Specifies the upper bound for the numeric value which will be inserted into the search filter. | [optional] 
**FilterPrefix** | **string** | Specifies a prefix which will be used in front of the randomly-selected numeric value in all search filters used. If no upper bound is defined, then this should contain the entire filter string. | 
**FilterSuffix** | Pointer to **string** | Specifies a suffix which will be used after of the randomly-selected numeric value in all search filters used. If no upper bound is defined, then this should be omitted. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 

## Methods

### NewInternalSearchRatePluginResponse

`func NewInternalSearchRatePluginResponse(schemas []EnuminternalSearchRatePluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, numThreads int64, baseDN string, filterPrefix string, enabled bool, id string, ) *InternalSearchRatePluginResponse`

NewInternalSearchRatePluginResponse instantiates a new InternalSearchRatePluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalSearchRatePluginResponseWithDefaults

`func NewInternalSearchRatePluginResponseWithDefaults() *InternalSearchRatePluginResponse`

NewInternalSearchRatePluginResponseWithDefaults instantiates a new InternalSearchRatePluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *InternalSearchRatePluginResponse) GetSchemas() []EnuminternalSearchRatePluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *InternalSearchRatePluginResponse) GetSchemasOk() (*[]EnuminternalSearchRatePluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *InternalSearchRatePluginResponse) SetSchemas(v []EnuminternalSearchRatePluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *InternalSearchRatePluginResponse) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *InternalSearchRatePluginResponse) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *InternalSearchRatePluginResponse) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetNumThreads

`func (o *InternalSearchRatePluginResponse) GetNumThreads() int64`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *InternalSearchRatePluginResponse) GetNumThreadsOk() (*int64, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *InternalSearchRatePluginResponse) SetNumThreads(v int64)`

SetNumThreads sets NumThreads field to given value.


### GetBaseDN

`func (o *InternalSearchRatePluginResponse) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *InternalSearchRatePluginResponse) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *InternalSearchRatePluginResponse) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.


### GetLowerBound

`func (o *InternalSearchRatePluginResponse) GetLowerBound() int64`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *InternalSearchRatePluginResponse) GetLowerBoundOk() (*int64, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *InternalSearchRatePluginResponse) SetLowerBound(v int64)`

SetLowerBound sets LowerBound field to given value.

### HasLowerBound

`func (o *InternalSearchRatePluginResponse) HasLowerBound() bool`

HasLowerBound returns a boolean if a field has been set.

### GetUpperBound

`func (o *InternalSearchRatePluginResponse) GetUpperBound() int64`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *InternalSearchRatePluginResponse) GetUpperBoundOk() (*int64, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *InternalSearchRatePluginResponse) SetUpperBound(v int64)`

SetUpperBound sets UpperBound field to given value.

### HasUpperBound

`func (o *InternalSearchRatePluginResponse) HasUpperBound() bool`

HasUpperBound returns a boolean if a field has been set.

### GetFilterPrefix

`func (o *InternalSearchRatePluginResponse) GetFilterPrefix() string`

GetFilterPrefix returns the FilterPrefix field if non-nil, zero value otherwise.

### GetFilterPrefixOk

`func (o *InternalSearchRatePluginResponse) GetFilterPrefixOk() (*string, bool)`

GetFilterPrefixOk returns a tuple with the FilterPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPrefix

`func (o *InternalSearchRatePluginResponse) SetFilterPrefix(v string)`

SetFilterPrefix sets FilterPrefix field to given value.


### GetFilterSuffix

`func (o *InternalSearchRatePluginResponse) GetFilterSuffix() string`

GetFilterSuffix returns the FilterSuffix field if non-nil, zero value otherwise.

### GetFilterSuffixOk

`func (o *InternalSearchRatePluginResponse) GetFilterSuffixOk() (*string, bool)`

GetFilterSuffixOk returns a tuple with the FilterSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSuffix

`func (o *InternalSearchRatePluginResponse) SetFilterSuffix(v string)`

SetFilterSuffix sets FilterSuffix field to given value.

### HasFilterSuffix

`func (o *InternalSearchRatePluginResponse) HasFilterSuffix() bool`

HasFilterSuffix returns a boolean if a field has been set.

### GetDescription

`func (o *InternalSearchRatePluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalSearchRatePluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalSearchRatePluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalSearchRatePluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *InternalSearchRatePluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InternalSearchRatePluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InternalSearchRatePluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *InternalSearchRatePluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *InternalSearchRatePluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *InternalSearchRatePluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *InternalSearchRatePluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetMeta

`func (o *InternalSearchRatePluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InternalSearchRatePluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InternalSearchRatePluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InternalSearchRatePluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *InternalSearchRatePluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *InternalSearchRatePluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *InternalSearchRatePluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *InternalSearchRatePluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *InternalSearchRatePluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalSearchRatePluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalSearchRatePluginResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


