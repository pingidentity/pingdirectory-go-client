# ErrorLogFieldMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log Field Mapping | 
**Schemas** | [**[]EnumerrorLogFieldMappingSchemaUrn**](EnumerrorLogFieldMappingSchemaUrn.md) |  | 
**LogFieldTimestamp** | Pointer to **string** | The time that the log message was generated. | [optional] 
**LogFieldProductName** | Pointer to **string** | The name for this Directory Server product, which may be used to identify which product was used to log the message if multiple products log to the same database table. | [optional] 
**LogFieldInstanceName** | Pointer to **string** | A name that uniquely identifies this Directory Server instance, which may be used to identify which instance was used to log the message if multiple server instances log to the same database table. | [optional] 
**LogFieldStartupid** | Pointer to **string** | The startup ID for the Directory Server. A different value will be generated each time the server is started. | [optional] 
**LogFieldCategory** | Pointer to **string** | The category for the log message. | [optional] 
**LogFieldSeverity** | Pointer to **string** | The severity for the log message. | [optional] 
**LogFieldMessageID** | Pointer to **string** | The numeric value which uniquely identifies the type of message. | [optional] 
**LogFieldMessage** | Pointer to **string** | The text of the log message. | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Mapping | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewErrorLogFieldMappingResponse

`func NewErrorLogFieldMappingResponse(id string, schemas []EnumerrorLogFieldMappingSchemaUrn, ) *ErrorLogFieldMappingResponse`

NewErrorLogFieldMappingResponse instantiates a new ErrorLogFieldMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLogFieldMappingResponseWithDefaults

`func NewErrorLogFieldMappingResponseWithDefaults() *ErrorLogFieldMappingResponse`

NewErrorLogFieldMappingResponseWithDefaults instantiates a new ErrorLogFieldMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ErrorLogFieldMappingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorLogFieldMappingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorLogFieldMappingResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ErrorLogFieldMappingResponse) GetSchemas() []EnumerrorLogFieldMappingSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ErrorLogFieldMappingResponse) GetSchemasOk() (*[]EnumerrorLogFieldMappingSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ErrorLogFieldMappingResponse) SetSchemas(v []EnumerrorLogFieldMappingSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLogFieldTimestamp

`func (o *ErrorLogFieldMappingResponse) GetLogFieldTimestamp() string`

GetLogFieldTimestamp returns the LogFieldTimestamp field if non-nil, zero value otherwise.

### GetLogFieldTimestampOk

`func (o *ErrorLogFieldMappingResponse) GetLogFieldTimestampOk() (*string, bool)`

GetLogFieldTimestampOk returns a tuple with the LogFieldTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTimestamp

`func (o *ErrorLogFieldMappingResponse) SetLogFieldTimestamp(v string)`

SetLogFieldTimestamp sets LogFieldTimestamp field to given value.

### HasLogFieldTimestamp

`func (o *ErrorLogFieldMappingResponse) HasLogFieldTimestamp() bool`

HasLogFieldTimestamp returns a boolean if a field has been set.

### GetLogFieldProductName

`func (o *ErrorLogFieldMappingResponse) GetLogFieldProductName() string`

GetLogFieldProductName returns the LogFieldProductName field if non-nil, zero value otherwise.

### GetLogFieldProductNameOk

`func (o *ErrorLogFieldMappingResponse) GetLogFieldProductNameOk() (*string, bool)`

GetLogFieldProductNameOk returns a tuple with the LogFieldProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProductName

`func (o *ErrorLogFieldMappingResponse) SetLogFieldProductName(v string)`

SetLogFieldProductName sets LogFieldProductName field to given value.

### HasLogFieldProductName

`func (o *ErrorLogFieldMappingResponse) HasLogFieldProductName() bool`

HasLogFieldProductName returns a boolean if a field has been set.

### GetLogFieldInstanceName

`func (o *ErrorLogFieldMappingResponse) GetLogFieldInstanceName() string`

GetLogFieldInstanceName returns the LogFieldInstanceName field if non-nil, zero value otherwise.

### GetLogFieldInstanceNameOk

`func (o *ErrorLogFieldMappingResponse) GetLogFieldInstanceNameOk() (*string, bool)`

GetLogFieldInstanceNameOk returns a tuple with the LogFieldInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldInstanceName

`func (o *ErrorLogFieldMappingResponse) SetLogFieldInstanceName(v string)`

SetLogFieldInstanceName sets LogFieldInstanceName field to given value.

### HasLogFieldInstanceName

`func (o *ErrorLogFieldMappingResponse) HasLogFieldInstanceName() bool`

HasLogFieldInstanceName returns a boolean if a field has been set.

### GetLogFieldStartupid

`func (o *ErrorLogFieldMappingResponse) GetLogFieldStartupid() string`

GetLogFieldStartupid returns the LogFieldStartupid field if non-nil, zero value otherwise.

### GetLogFieldStartupidOk

`func (o *ErrorLogFieldMappingResponse) GetLogFieldStartupidOk() (*string, bool)`

GetLogFieldStartupidOk returns a tuple with the LogFieldStartupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldStartupid

`func (o *ErrorLogFieldMappingResponse) SetLogFieldStartupid(v string)`

SetLogFieldStartupid sets LogFieldStartupid field to given value.

### HasLogFieldStartupid

`func (o *ErrorLogFieldMappingResponse) HasLogFieldStartupid() bool`

HasLogFieldStartupid returns a boolean if a field has been set.

### GetLogFieldCategory

`func (o *ErrorLogFieldMappingResponse) GetLogFieldCategory() string`

GetLogFieldCategory returns the LogFieldCategory field if non-nil, zero value otherwise.

### GetLogFieldCategoryOk

`func (o *ErrorLogFieldMappingResponse) GetLogFieldCategoryOk() (*string, bool)`

GetLogFieldCategoryOk returns a tuple with the LogFieldCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldCategory

`func (o *ErrorLogFieldMappingResponse) SetLogFieldCategory(v string)`

SetLogFieldCategory sets LogFieldCategory field to given value.

### HasLogFieldCategory

`func (o *ErrorLogFieldMappingResponse) HasLogFieldCategory() bool`

HasLogFieldCategory returns a boolean if a field has been set.

### GetLogFieldSeverity

`func (o *ErrorLogFieldMappingResponse) GetLogFieldSeverity() string`

GetLogFieldSeverity returns the LogFieldSeverity field if non-nil, zero value otherwise.

### GetLogFieldSeverityOk

`func (o *ErrorLogFieldMappingResponse) GetLogFieldSeverityOk() (*string, bool)`

GetLogFieldSeverityOk returns a tuple with the LogFieldSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldSeverity

`func (o *ErrorLogFieldMappingResponse) SetLogFieldSeverity(v string)`

SetLogFieldSeverity sets LogFieldSeverity field to given value.

### HasLogFieldSeverity

`func (o *ErrorLogFieldMappingResponse) HasLogFieldSeverity() bool`

HasLogFieldSeverity returns a boolean if a field has been set.

### GetLogFieldMessageID

`func (o *ErrorLogFieldMappingResponse) GetLogFieldMessageID() string`

GetLogFieldMessageID returns the LogFieldMessageID field if non-nil, zero value otherwise.

### GetLogFieldMessageIDOk

`func (o *ErrorLogFieldMappingResponse) GetLogFieldMessageIDOk() (*string, bool)`

GetLogFieldMessageIDOk returns a tuple with the LogFieldMessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessageID

`func (o *ErrorLogFieldMappingResponse) SetLogFieldMessageID(v string)`

SetLogFieldMessageID sets LogFieldMessageID field to given value.

### HasLogFieldMessageID

`func (o *ErrorLogFieldMappingResponse) HasLogFieldMessageID() bool`

HasLogFieldMessageID returns a boolean if a field has been set.

### GetLogFieldMessage

`func (o *ErrorLogFieldMappingResponse) GetLogFieldMessage() string`

GetLogFieldMessage returns the LogFieldMessage field if non-nil, zero value otherwise.

### GetLogFieldMessageOk

`func (o *ErrorLogFieldMappingResponse) GetLogFieldMessageOk() (*string, bool)`

GetLogFieldMessageOk returns a tuple with the LogFieldMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessage

`func (o *ErrorLogFieldMappingResponse) SetLogFieldMessage(v string)`

SetLogFieldMessage sets LogFieldMessage field to given value.

### HasLogFieldMessage

`func (o *ErrorLogFieldMappingResponse) HasLogFieldMessage() bool`

HasLogFieldMessage returns a boolean if a field has been set.

### GetDescription

`func (o *ErrorLogFieldMappingResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorLogFieldMappingResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorLogFieldMappingResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorLogFieldMappingResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *ErrorLogFieldMappingResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ErrorLogFieldMappingResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ErrorLogFieldMappingResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ErrorLogFieldMappingResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


