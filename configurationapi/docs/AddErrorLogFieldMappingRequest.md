# AddErrorLogFieldMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**MappingName** | **string** | Name of the new Log Field Mapping | 

## Methods

### NewAddErrorLogFieldMappingRequest

`func NewAddErrorLogFieldMappingRequest(schemas []EnumerrorLogFieldMappingSchemaUrn, mappingName string, ) *AddErrorLogFieldMappingRequest`

NewAddErrorLogFieldMappingRequest instantiates a new AddErrorLogFieldMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddErrorLogFieldMappingRequestWithDefaults

`func NewAddErrorLogFieldMappingRequestWithDefaults() *AddErrorLogFieldMappingRequest`

NewAddErrorLogFieldMappingRequestWithDefaults instantiates a new AddErrorLogFieldMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddErrorLogFieldMappingRequest) GetSchemas() []EnumerrorLogFieldMappingSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddErrorLogFieldMappingRequest) GetSchemasOk() (*[]EnumerrorLogFieldMappingSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddErrorLogFieldMappingRequest) SetSchemas(v []EnumerrorLogFieldMappingSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLogFieldTimestamp

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldTimestamp() string`

GetLogFieldTimestamp returns the LogFieldTimestamp field if non-nil, zero value otherwise.

### GetLogFieldTimestampOk

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldTimestampOk() (*string, bool)`

GetLogFieldTimestampOk returns a tuple with the LogFieldTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTimestamp

`func (o *AddErrorLogFieldMappingRequest) SetLogFieldTimestamp(v string)`

SetLogFieldTimestamp sets LogFieldTimestamp field to given value.

### HasLogFieldTimestamp

`func (o *AddErrorLogFieldMappingRequest) HasLogFieldTimestamp() bool`

HasLogFieldTimestamp returns a boolean if a field has been set.

### GetLogFieldProductName

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldProductName() string`

GetLogFieldProductName returns the LogFieldProductName field if non-nil, zero value otherwise.

### GetLogFieldProductNameOk

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldProductNameOk() (*string, bool)`

GetLogFieldProductNameOk returns a tuple with the LogFieldProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProductName

`func (o *AddErrorLogFieldMappingRequest) SetLogFieldProductName(v string)`

SetLogFieldProductName sets LogFieldProductName field to given value.

### HasLogFieldProductName

`func (o *AddErrorLogFieldMappingRequest) HasLogFieldProductName() bool`

HasLogFieldProductName returns a boolean if a field has been set.

### GetLogFieldInstanceName

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldInstanceName() string`

GetLogFieldInstanceName returns the LogFieldInstanceName field if non-nil, zero value otherwise.

### GetLogFieldInstanceNameOk

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldInstanceNameOk() (*string, bool)`

GetLogFieldInstanceNameOk returns a tuple with the LogFieldInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldInstanceName

`func (o *AddErrorLogFieldMappingRequest) SetLogFieldInstanceName(v string)`

SetLogFieldInstanceName sets LogFieldInstanceName field to given value.

### HasLogFieldInstanceName

`func (o *AddErrorLogFieldMappingRequest) HasLogFieldInstanceName() bool`

HasLogFieldInstanceName returns a boolean if a field has been set.

### GetLogFieldStartupid

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldStartupid() string`

GetLogFieldStartupid returns the LogFieldStartupid field if non-nil, zero value otherwise.

### GetLogFieldStartupidOk

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldStartupidOk() (*string, bool)`

GetLogFieldStartupidOk returns a tuple with the LogFieldStartupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldStartupid

`func (o *AddErrorLogFieldMappingRequest) SetLogFieldStartupid(v string)`

SetLogFieldStartupid sets LogFieldStartupid field to given value.

### HasLogFieldStartupid

`func (o *AddErrorLogFieldMappingRequest) HasLogFieldStartupid() bool`

HasLogFieldStartupid returns a boolean if a field has been set.

### GetLogFieldCategory

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldCategory() string`

GetLogFieldCategory returns the LogFieldCategory field if non-nil, zero value otherwise.

### GetLogFieldCategoryOk

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldCategoryOk() (*string, bool)`

GetLogFieldCategoryOk returns a tuple with the LogFieldCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldCategory

`func (o *AddErrorLogFieldMappingRequest) SetLogFieldCategory(v string)`

SetLogFieldCategory sets LogFieldCategory field to given value.

### HasLogFieldCategory

`func (o *AddErrorLogFieldMappingRequest) HasLogFieldCategory() bool`

HasLogFieldCategory returns a boolean if a field has been set.

### GetLogFieldSeverity

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldSeverity() string`

GetLogFieldSeverity returns the LogFieldSeverity field if non-nil, zero value otherwise.

### GetLogFieldSeverityOk

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldSeverityOk() (*string, bool)`

GetLogFieldSeverityOk returns a tuple with the LogFieldSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldSeverity

`func (o *AddErrorLogFieldMappingRequest) SetLogFieldSeverity(v string)`

SetLogFieldSeverity sets LogFieldSeverity field to given value.

### HasLogFieldSeverity

`func (o *AddErrorLogFieldMappingRequest) HasLogFieldSeverity() bool`

HasLogFieldSeverity returns a boolean if a field has been set.

### GetLogFieldMessageID

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldMessageID() string`

GetLogFieldMessageID returns the LogFieldMessageID field if non-nil, zero value otherwise.

### GetLogFieldMessageIDOk

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldMessageIDOk() (*string, bool)`

GetLogFieldMessageIDOk returns a tuple with the LogFieldMessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessageID

`func (o *AddErrorLogFieldMappingRequest) SetLogFieldMessageID(v string)`

SetLogFieldMessageID sets LogFieldMessageID field to given value.

### HasLogFieldMessageID

`func (o *AddErrorLogFieldMappingRequest) HasLogFieldMessageID() bool`

HasLogFieldMessageID returns a boolean if a field has been set.

### GetLogFieldMessage

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldMessage() string`

GetLogFieldMessage returns the LogFieldMessage field if non-nil, zero value otherwise.

### GetLogFieldMessageOk

`func (o *AddErrorLogFieldMappingRequest) GetLogFieldMessageOk() (*string, bool)`

GetLogFieldMessageOk returns a tuple with the LogFieldMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessage

`func (o *AddErrorLogFieldMappingRequest) SetLogFieldMessage(v string)`

SetLogFieldMessage sets LogFieldMessage field to given value.

### HasLogFieldMessage

`func (o *AddErrorLogFieldMappingRequest) HasLogFieldMessage() bool`

HasLogFieldMessage returns a boolean if a field has been set.

### GetDescription

`func (o *AddErrorLogFieldMappingRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddErrorLogFieldMappingRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddErrorLogFieldMappingRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddErrorLogFieldMappingRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMappingName

`func (o *AddErrorLogFieldMappingRequest) GetMappingName() string`

GetMappingName returns the MappingName field if non-nil, zero value otherwise.

### GetMappingNameOk

`func (o *AddErrorLogFieldMappingRequest) GetMappingNameOk() (*string, bool)`

GetMappingNameOk returns a tuple with the MappingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingName

`func (o *AddErrorLogFieldMappingRequest) SetMappingName(v string)`

SetMappingName sets MappingName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


