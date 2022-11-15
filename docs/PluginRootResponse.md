# PluginRootResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumpluginRootSchemaUrn**](EnumpluginRootSchemaUrn.md) |  | [optional] 
**PluginOrderStartup** | Pointer to **string** | Specifies the order in which startup plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderShutdown** | Pointer to **string** | Specifies the order in which shutdown plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostConnect** | Pointer to **string** | Specifies the order in which post-connect plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostDisconnect** | Pointer to **string** | Specifies the order in which post-disconnect plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderLDIFImport** | Pointer to **string** | Specifies the order in which LDIF import plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderLDIFExport** | Pointer to **string** | Specifies the order in which LDIF export plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreParseAbandon** | Pointer to **string** | Specifies the order in which pre-parse abandon plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreParseAdd** | Pointer to **string** | Specifies the order in which pre-parse add plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreParseBind** | Pointer to **string** | Specifies the order in which pre-parse bind plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreParseCompare** | Pointer to **string** | Specifies the order in which pre-parse compare plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreParseDelete** | Pointer to **string** | Specifies the order in which pre-parse delete plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreParseExtended** | Pointer to **string** | Specifies the order in which pre-parse extended operation plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreParseModify** | Pointer to **string** | Specifies the order in which pre-parse modify plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreParseModifyDN** | Pointer to **string** | Specifies the order in which pre-parse modify DN plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreParseSearch** | Pointer to **string** | Specifies the order in which pre-parse search plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreParseUnbind** | Pointer to **string** | Specifies the order in which pre-parse unbind plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreOperationAdd** | Pointer to **string** | Specifies the order in which pre-operation add plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreOperationBind** | Pointer to **string** | Specifies the order in which pre-operation bind plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreOperationCompare** | Pointer to **string** | Specifies the order in which pre-operation compare plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreOperationDelete** | Pointer to **string** | Specifies the order in which pre-operation delete plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreOperationExtended** | Pointer to **string** | Specifies the order in which pre-operation extended operation plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreOperationModify** | Pointer to **string** | Specifies the order in which pre-operation modify plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreOperationModifyDN** | Pointer to **string** | Specifies the order in which pre-operation modify DN plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPreOperationSearch** | Pointer to **string** | Specifies the order in which pre-operation search plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostOperationAbandon** | Pointer to **string** | Specifies the order in which post-operation abandon plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostOperationAdd** | Pointer to **string** | Specifies the order in which post-operation add plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostOperationBind** | Pointer to **string** | Specifies the order in which post-operation bind plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostOperationCompare** | Pointer to **string** | Specifies the order in which post-operation compare plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostOperationDelete** | Pointer to **string** | Specifies the order in which post-operation delete plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostOperationExtended** | Pointer to **string** | Specifies the order in which post-operation extended operation plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostOperationModify** | Pointer to **string** | Specifies the order in which post-operation modify plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostOperationModifyDN** | Pointer to **string** | Specifies the order in which post-operation modify DN plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostOperationSearch** | Pointer to **string** | Specifies the order in which post-operation search plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostOperationUnbind** | Pointer to **string** | Specifies the order in which post-operation unbind plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostResponseAdd** | Pointer to **string** | Specifies the order in which post-response add plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostResponseBind** | Pointer to **string** | Specifies the order in which post-response bind plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostResponseCompare** | Pointer to **string** | Specifies the order in which post-response compare plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostResponseDelete** | Pointer to **string** | Specifies the order in which post-response delete plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostResponseExtended** | Pointer to **string** | Specifies the order in which post-response extended operation plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostResponseModify** | Pointer to **string** | Specifies the order in which post-response modify plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostResponseModifyDN** | Pointer to **string** | Specifies the order in which post-response modify DN plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostSynchronizationAdd** | Pointer to **string** | Specifies the order in which post-synchronization add plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostSynchronizationDelete** | Pointer to **string** | Specifies the order in which post-synchronization delete plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostSynchronizationModify** | Pointer to **string** | Specifies the order in which post-synchronization modify plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostSynchronizationModifyDN** | Pointer to **string** | Specifies the order in which post-synchronization modify DN plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderPostResponseSearch** | Pointer to **string** | Specifies the order in which post-response search plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderSearchResultEntry** | Pointer to **string** | Specifies the order in which search result entry plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderSearchResultReference** | Pointer to **string** | Specifies the order in which search result reference plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderSubordinateModifyDN** | Pointer to **string** | Specifies the order in which subordinate modify DN plug-ins are to be loaded and invoked. | [optional] 
**PluginOrderIntermediateResponse** | Pointer to **string** | Specifies the order in which intermediate response plug-ins are to be loaded and invoked. | [optional] 

## Methods

### NewPluginRootResponse

`func NewPluginRootResponse() *PluginRootResponse`

NewPluginRootResponse instantiates a new PluginRootResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginRootResponseWithDefaults

`func NewPluginRootResponseWithDefaults() *PluginRootResponse`

NewPluginRootResponseWithDefaults instantiates a new PluginRootResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PluginRootResponse) GetSchemas() []EnumpluginRootSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PluginRootResponse) GetSchemasOk() (*[]EnumpluginRootSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PluginRootResponse) SetSchemas(v []EnumpluginRootSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *PluginRootResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetPluginOrderStartup

`func (o *PluginRootResponse) GetPluginOrderStartup() string`

GetPluginOrderStartup returns the PluginOrderStartup field if non-nil, zero value otherwise.

### GetPluginOrderStartupOk

`func (o *PluginRootResponse) GetPluginOrderStartupOk() (*string, bool)`

GetPluginOrderStartupOk returns a tuple with the PluginOrderStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderStartup

`func (o *PluginRootResponse) SetPluginOrderStartup(v string)`

SetPluginOrderStartup sets PluginOrderStartup field to given value.

### HasPluginOrderStartup

`func (o *PluginRootResponse) HasPluginOrderStartup() bool`

HasPluginOrderStartup returns a boolean if a field has been set.

### GetPluginOrderShutdown

`func (o *PluginRootResponse) GetPluginOrderShutdown() string`

GetPluginOrderShutdown returns the PluginOrderShutdown field if non-nil, zero value otherwise.

### GetPluginOrderShutdownOk

`func (o *PluginRootResponse) GetPluginOrderShutdownOk() (*string, bool)`

GetPluginOrderShutdownOk returns a tuple with the PluginOrderShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderShutdown

`func (o *PluginRootResponse) SetPluginOrderShutdown(v string)`

SetPluginOrderShutdown sets PluginOrderShutdown field to given value.

### HasPluginOrderShutdown

`func (o *PluginRootResponse) HasPluginOrderShutdown() bool`

HasPluginOrderShutdown returns a boolean if a field has been set.

### GetPluginOrderPostConnect

`func (o *PluginRootResponse) GetPluginOrderPostConnect() string`

GetPluginOrderPostConnect returns the PluginOrderPostConnect field if non-nil, zero value otherwise.

### GetPluginOrderPostConnectOk

`func (o *PluginRootResponse) GetPluginOrderPostConnectOk() (*string, bool)`

GetPluginOrderPostConnectOk returns a tuple with the PluginOrderPostConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostConnect

`func (o *PluginRootResponse) SetPluginOrderPostConnect(v string)`

SetPluginOrderPostConnect sets PluginOrderPostConnect field to given value.

### HasPluginOrderPostConnect

`func (o *PluginRootResponse) HasPluginOrderPostConnect() bool`

HasPluginOrderPostConnect returns a boolean if a field has been set.

### GetPluginOrderPostDisconnect

`func (o *PluginRootResponse) GetPluginOrderPostDisconnect() string`

GetPluginOrderPostDisconnect returns the PluginOrderPostDisconnect field if non-nil, zero value otherwise.

### GetPluginOrderPostDisconnectOk

`func (o *PluginRootResponse) GetPluginOrderPostDisconnectOk() (*string, bool)`

GetPluginOrderPostDisconnectOk returns a tuple with the PluginOrderPostDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostDisconnect

`func (o *PluginRootResponse) SetPluginOrderPostDisconnect(v string)`

SetPluginOrderPostDisconnect sets PluginOrderPostDisconnect field to given value.

### HasPluginOrderPostDisconnect

`func (o *PluginRootResponse) HasPluginOrderPostDisconnect() bool`

HasPluginOrderPostDisconnect returns a boolean if a field has been set.

### GetPluginOrderLDIFImport

`func (o *PluginRootResponse) GetPluginOrderLDIFImport() string`

GetPluginOrderLDIFImport returns the PluginOrderLDIFImport field if non-nil, zero value otherwise.

### GetPluginOrderLDIFImportOk

`func (o *PluginRootResponse) GetPluginOrderLDIFImportOk() (*string, bool)`

GetPluginOrderLDIFImportOk returns a tuple with the PluginOrderLDIFImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderLDIFImport

`func (o *PluginRootResponse) SetPluginOrderLDIFImport(v string)`

SetPluginOrderLDIFImport sets PluginOrderLDIFImport field to given value.

### HasPluginOrderLDIFImport

`func (o *PluginRootResponse) HasPluginOrderLDIFImport() bool`

HasPluginOrderLDIFImport returns a boolean if a field has been set.

### GetPluginOrderLDIFExport

`func (o *PluginRootResponse) GetPluginOrderLDIFExport() string`

GetPluginOrderLDIFExport returns the PluginOrderLDIFExport field if non-nil, zero value otherwise.

### GetPluginOrderLDIFExportOk

`func (o *PluginRootResponse) GetPluginOrderLDIFExportOk() (*string, bool)`

GetPluginOrderLDIFExportOk returns a tuple with the PluginOrderLDIFExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderLDIFExport

`func (o *PluginRootResponse) SetPluginOrderLDIFExport(v string)`

SetPluginOrderLDIFExport sets PluginOrderLDIFExport field to given value.

### HasPluginOrderLDIFExport

`func (o *PluginRootResponse) HasPluginOrderLDIFExport() bool`

HasPluginOrderLDIFExport returns a boolean if a field has been set.

### GetPluginOrderPreParseAbandon

`func (o *PluginRootResponse) GetPluginOrderPreParseAbandon() string`

GetPluginOrderPreParseAbandon returns the PluginOrderPreParseAbandon field if non-nil, zero value otherwise.

### GetPluginOrderPreParseAbandonOk

`func (o *PluginRootResponse) GetPluginOrderPreParseAbandonOk() (*string, bool)`

GetPluginOrderPreParseAbandonOk returns a tuple with the PluginOrderPreParseAbandon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreParseAbandon

`func (o *PluginRootResponse) SetPluginOrderPreParseAbandon(v string)`

SetPluginOrderPreParseAbandon sets PluginOrderPreParseAbandon field to given value.

### HasPluginOrderPreParseAbandon

`func (o *PluginRootResponse) HasPluginOrderPreParseAbandon() bool`

HasPluginOrderPreParseAbandon returns a boolean if a field has been set.

### GetPluginOrderPreParseAdd

`func (o *PluginRootResponse) GetPluginOrderPreParseAdd() string`

GetPluginOrderPreParseAdd returns the PluginOrderPreParseAdd field if non-nil, zero value otherwise.

### GetPluginOrderPreParseAddOk

`func (o *PluginRootResponse) GetPluginOrderPreParseAddOk() (*string, bool)`

GetPluginOrderPreParseAddOk returns a tuple with the PluginOrderPreParseAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreParseAdd

`func (o *PluginRootResponse) SetPluginOrderPreParseAdd(v string)`

SetPluginOrderPreParseAdd sets PluginOrderPreParseAdd field to given value.

### HasPluginOrderPreParseAdd

`func (o *PluginRootResponse) HasPluginOrderPreParseAdd() bool`

HasPluginOrderPreParseAdd returns a boolean if a field has been set.

### GetPluginOrderPreParseBind

`func (o *PluginRootResponse) GetPluginOrderPreParseBind() string`

GetPluginOrderPreParseBind returns the PluginOrderPreParseBind field if non-nil, zero value otherwise.

### GetPluginOrderPreParseBindOk

`func (o *PluginRootResponse) GetPluginOrderPreParseBindOk() (*string, bool)`

GetPluginOrderPreParseBindOk returns a tuple with the PluginOrderPreParseBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreParseBind

`func (o *PluginRootResponse) SetPluginOrderPreParseBind(v string)`

SetPluginOrderPreParseBind sets PluginOrderPreParseBind field to given value.

### HasPluginOrderPreParseBind

`func (o *PluginRootResponse) HasPluginOrderPreParseBind() bool`

HasPluginOrderPreParseBind returns a boolean if a field has been set.

### GetPluginOrderPreParseCompare

`func (o *PluginRootResponse) GetPluginOrderPreParseCompare() string`

GetPluginOrderPreParseCompare returns the PluginOrderPreParseCompare field if non-nil, zero value otherwise.

### GetPluginOrderPreParseCompareOk

`func (o *PluginRootResponse) GetPluginOrderPreParseCompareOk() (*string, bool)`

GetPluginOrderPreParseCompareOk returns a tuple with the PluginOrderPreParseCompare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreParseCompare

`func (o *PluginRootResponse) SetPluginOrderPreParseCompare(v string)`

SetPluginOrderPreParseCompare sets PluginOrderPreParseCompare field to given value.

### HasPluginOrderPreParseCompare

`func (o *PluginRootResponse) HasPluginOrderPreParseCompare() bool`

HasPluginOrderPreParseCompare returns a boolean if a field has been set.

### GetPluginOrderPreParseDelete

`func (o *PluginRootResponse) GetPluginOrderPreParseDelete() string`

GetPluginOrderPreParseDelete returns the PluginOrderPreParseDelete field if non-nil, zero value otherwise.

### GetPluginOrderPreParseDeleteOk

`func (o *PluginRootResponse) GetPluginOrderPreParseDeleteOk() (*string, bool)`

GetPluginOrderPreParseDeleteOk returns a tuple with the PluginOrderPreParseDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreParseDelete

`func (o *PluginRootResponse) SetPluginOrderPreParseDelete(v string)`

SetPluginOrderPreParseDelete sets PluginOrderPreParseDelete field to given value.

### HasPluginOrderPreParseDelete

`func (o *PluginRootResponse) HasPluginOrderPreParseDelete() bool`

HasPluginOrderPreParseDelete returns a boolean if a field has been set.

### GetPluginOrderPreParseExtended

`func (o *PluginRootResponse) GetPluginOrderPreParseExtended() string`

GetPluginOrderPreParseExtended returns the PluginOrderPreParseExtended field if non-nil, zero value otherwise.

### GetPluginOrderPreParseExtendedOk

`func (o *PluginRootResponse) GetPluginOrderPreParseExtendedOk() (*string, bool)`

GetPluginOrderPreParseExtendedOk returns a tuple with the PluginOrderPreParseExtended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreParseExtended

`func (o *PluginRootResponse) SetPluginOrderPreParseExtended(v string)`

SetPluginOrderPreParseExtended sets PluginOrderPreParseExtended field to given value.

### HasPluginOrderPreParseExtended

`func (o *PluginRootResponse) HasPluginOrderPreParseExtended() bool`

HasPluginOrderPreParseExtended returns a boolean if a field has been set.

### GetPluginOrderPreParseModify

`func (o *PluginRootResponse) GetPluginOrderPreParseModify() string`

GetPluginOrderPreParseModify returns the PluginOrderPreParseModify field if non-nil, zero value otherwise.

### GetPluginOrderPreParseModifyOk

`func (o *PluginRootResponse) GetPluginOrderPreParseModifyOk() (*string, bool)`

GetPluginOrderPreParseModifyOk returns a tuple with the PluginOrderPreParseModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreParseModify

`func (o *PluginRootResponse) SetPluginOrderPreParseModify(v string)`

SetPluginOrderPreParseModify sets PluginOrderPreParseModify field to given value.

### HasPluginOrderPreParseModify

`func (o *PluginRootResponse) HasPluginOrderPreParseModify() bool`

HasPluginOrderPreParseModify returns a boolean if a field has been set.

### GetPluginOrderPreParseModifyDN

`func (o *PluginRootResponse) GetPluginOrderPreParseModifyDN() string`

GetPluginOrderPreParseModifyDN returns the PluginOrderPreParseModifyDN field if non-nil, zero value otherwise.

### GetPluginOrderPreParseModifyDNOk

`func (o *PluginRootResponse) GetPluginOrderPreParseModifyDNOk() (*string, bool)`

GetPluginOrderPreParseModifyDNOk returns a tuple with the PluginOrderPreParseModifyDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreParseModifyDN

`func (o *PluginRootResponse) SetPluginOrderPreParseModifyDN(v string)`

SetPluginOrderPreParseModifyDN sets PluginOrderPreParseModifyDN field to given value.

### HasPluginOrderPreParseModifyDN

`func (o *PluginRootResponse) HasPluginOrderPreParseModifyDN() bool`

HasPluginOrderPreParseModifyDN returns a boolean if a field has been set.

### GetPluginOrderPreParseSearch

`func (o *PluginRootResponse) GetPluginOrderPreParseSearch() string`

GetPluginOrderPreParseSearch returns the PluginOrderPreParseSearch field if non-nil, zero value otherwise.

### GetPluginOrderPreParseSearchOk

`func (o *PluginRootResponse) GetPluginOrderPreParseSearchOk() (*string, bool)`

GetPluginOrderPreParseSearchOk returns a tuple with the PluginOrderPreParseSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreParseSearch

`func (o *PluginRootResponse) SetPluginOrderPreParseSearch(v string)`

SetPluginOrderPreParseSearch sets PluginOrderPreParseSearch field to given value.

### HasPluginOrderPreParseSearch

`func (o *PluginRootResponse) HasPluginOrderPreParseSearch() bool`

HasPluginOrderPreParseSearch returns a boolean if a field has been set.

### GetPluginOrderPreParseUnbind

`func (o *PluginRootResponse) GetPluginOrderPreParseUnbind() string`

GetPluginOrderPreParseUnbind returns the PluginOrderPreParseUnbind field if non-nil, zero value otherwise.

### GetPluginOrderPreParseUnbindOk

`func (o *PluginRootResponse) GetPluginOrderPreParseUnbindOk() (*string, bool)`

GetPluginOrderPreParseUnbindOk returns a tuple with the PluginOrderPreParseUnbind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreParseUnbind

`func (o *PluginRootResponse) SetPluginOrderPreParseUnbind(v string)`

SetPluginOrderPreParseUnbind sets PluginOrderPreParseUnbind field to given value.

### HasPluginOrderPreParseUnbind

`func (o *PluginRootResponse) HasPluginOrderPreParseUnbind() bool`

HasPluginOrderPreParseUnbind returns a boolean if a field has been set.

### GetPluginOrderPreOperationAdd

`func (o *PluginRootResponse) GetPluginOrderPreOperationAdd() string`

GetPluginOrderPreOperationAdd returns the PluginOrderPreOperationAdd field if non-nil, zero value otherwise.

### GetPluginOrderPreOperationAddOk

`func (o *PluginRootResponse) GetPluginOrderPreOperationAddOk() (*string, bool)`

GetPluginOrderPreOperationAddOk returns a tuple with the PluginOrderPreOperationAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreOperationAdd

`func (o *PluginRootResponse) SetPluginOrderPreOperationAdd(v string)`

SetPluginOrderPreOperationAdd sets PluginOrderPreOperationAdd field to given value.

### HasPluginOrderPreOperationAdd

`func (o *PluginRootResponse) HasPluginOrderPreOperationAdd() bool`

HasPluginOrderPreOperationAdd returns a boolean if a field has been set.

### GetPluginOrderPreOperationBind

`func (o *PluginRootResponse) GetPluginOrderPreOperationBind() string`

GetPluginOrderPreOperationBind returns the PluginOrderPreOperationBind field if non-nil, zero value otherwise.

### GetPluginOrderPreOperationBindOk

`func (o *PluginRootResponse) GetPluginOrderPreOperationBindOk() (*string, bool)`

GetPluginOrderPreOperationBindOk returns a tuple with the PluginOrderPreOperationBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreOperationBind

`func (o *PluginRootResponse) SetPluginOrderPreOperationBind(v string)`

SetPluginOrderPreOperationBind sets PluginOrderPreOperationBind field to given value.

### HasPluginOrderPreOperationBind

`func (o *PluginRootResponse) HasPluginOrderPreOperationBind() bool`

HasPluginOrderPreOperationBind returns a boolean if a field has been set.

### GetPluginOrderPreOperationCompare

`func (o *PluginRootResponse) GetPluginOrderPreOperationCompare() string`

GetPluginOrderPreOperationCompare returns the PluginOrderPreOperationCompare field if non-nil, zero value otherwise.

### GetPluginOrderPreOperationCompareOk

`func (o *PluginRootResponse) GetPluginOrderPreOperationCompareOk() (*string, bool)`

GetPluginOrderPreOperationCompareOk returns a tuple with the PluginOrderPreOperationCompare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreOperationCompare

`func (o *PluginRootResponse) SetPluginOrderPreOperationCompare(v string)`

SetPluginOrderPreOperationCompare sets PluginOrderPreOperationCompare field to given value.

### HasPluginOrderPreOperationCompare

`func (o *PluginRootResponse) HasPluginOrderPreOperationCompare() bool`

HasPluginOrderPreOperationCompare returns a boolean if a field has been set.

### GetPluginOrderPreOperationDelete

`func (o *PluginRootResponse) GetPluginOrderPreOperationDelete() string`

GetPluginOrderPreOperationDelete returns the PluginOrderPreOperationDelete field if non-nil, zero value otherwise.

### GetPluginOrderPreOperationDeleteOk

`func (o *PluginRootResponse) GetPluginOrderPreOperationDeleteOk() (*string, bool)`

GetPluginOrderPreOperationDeleteOk returns a tuple with the PluginOrderPreOperationDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreOperationDelete

`func (o *PluginRootResponse) SetPluginOrderPreOperationDelete(v string)`

SetPluginOrderPreOperationDelete sets PluginOrderPreOperationDelete field to given value.

### HasPluginOrderPreOperationDelete

`func (o *PluginRootResponse) HasPluginOrderPreOperationDelete() bool`

HasPluginOrderPreOperationDelete returns a boolean if a field has been set.

### GetPluginOrderPreOperationExtended

`func (o *PluginRootResponse) GetPluginOrderPreOperationExtended() string`

GetPluginOrderPreOperationExtended returns the PluginOrderPreOperationExtended field if non-nil, zero value otherwise.

### GetPluginOrderPreOperationExtendedOk

`func (o *PluginRootResponse) GetPluginOrderPreOperationExtendedOk() (*string, bool)`

GetPluginOrderPreOperationExtendedOk returns a tuple with the PluginOrderPreOperationExtended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreOperationExtended

`func (o *PluginRootResponse) SetPluginOrderPreOperationExtended(v string)`

SetPluginOrderPreOperationExtended sets PluginOrderPreOperationExtended field to given value.

### HasPluginOrderPreOperationExtended

`func (o *PluginRootResponse) HasPluginOrderPreOperationExtended() bool`

HasPluginOrderPreOperationExtended returns a boolean if a field has been set.

### GetPluginOrderPreOperationModify

`func (o *PluginRootResponse) GetPluginOrderPreOperationModify() string`

GetPluginOrderPreOperationModify returns the PluginOrderPreOperationModify field if non-nil, zero value otherwise.

### GetPluginOrderPreOperationModifyOk

`func (o *PluginRootResponse) GetPluginOrderPreOperationModifyOk() (*string, bool)`

GetPluginOrderPreOperationModifyOk returns a tuple with the PluginOrderPreOperationModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreOperationModify

`func (o *PluginRootResponse) SetPluginOrderPreOperationModify(v string)`

SetPluginOrderPreOperationModify sets PluginOrderPreOperationModify field to given value.

### HasPluginOrderPreOperationModify

`func (o *PluginRootResponse) HasPluginOrderPreOperationModify() bool`

HasPluginOrderPreOperationModify returns a boolean if a field has been set.

### GetPluginOrderPreOperationModifyDN

`func (o *PluginRootResponse) GetPluginOrderPreOperationModifyDN() string`

GetPluginOrderPreOperationModifyDN returns the PluginOrderPreOperationModifyDN field if non-nil, zero value otherwise.

### GetPluginOrderPreOperationModifyDNOk

`func (o *PluginRootResponse) GetPluginOrderPreOperationModifyDNOk() (*string, bool)`

GetPluginOrderPreOperationModifyDNOk returns a tuple with the PluginOrderPreOperationModifyDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreOperationModifyDN

`func (o *PluginRootResponse) SetPluginOrderPreOperationModifyDN(v string)`

SetPluginOrderPreOperationModifyDN sets PluginOrderPreOperationModifyDN field to given value.

### HasPluginOrderPreOperationModifyDN

`func (o *PluginRootResponse) HasPluginOrderPreOperationModifyDN() bool`

HasPluginOrderPreOperationModifyDN returns a boolean if a field has been set.

### GetPluginOrderPreOperationSearch

`func (o *PluginRootResponse) GetPluginOrderPreOperationSearch() string`

GetPluginOrderPreOperationSearch returns the PluginOrderPreOperationSearch field if non-nil, zero value otherwise.

### GetPluginOrderPreOperationSearchOk

`func (o *PluginRootResponse) GetPluginOrderPreOperationSearchOk() (*string, bool)`

GetPluginOrderPreOperationSearchOk returns a tuple with the PluginOrderPreOperationSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPreOperationSearch

`func (o *PluginRootResponse) SetPluginOrderPreOperationSearch(v string)`

SetPluginOrderPreOperationSearch sets PluginOrderPreOperationSearch field to given value.

### HasPluginOrderPreOperationSearch

`func (o *PluginRootResponse) HasPluginOrderPreOperationSearch() bool`

HasPluginOrderPreOperationSearch returns a boolean if a field has been set.

### GetPluginOrderPostOperationAbandon

`func (o *PluginRootResponse) GetPluginOrderPostOperationAbandon() string`

GetPluginOrderPostOperationAbandon returns the PluginOrderPostOperationAbandon field if non-nil, zero value otherwise.

### GetPluginOrderPostOperationAbandonOk

`func (o *PluginRootResponse) GetPluginOrderPostOperationAbandonOk() (*string, bool)`

GetPluginOrderPostOperationAbandonOk returns a tuple with the PluginOrderPostOperationAbandon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostOperationAbandon

`func (o *PluginRootResponse) SetPluginOrderPostOperationAbandon(v string)`

SetPluginOrderPostOperationAbandon sets PluginOrderPostOperationAbandon field to given value.

### HasPluginOrderPostOperationAbandon

`func (o *PluginRootResponse) HasPluginOrderPostOperationAbandon() bool`

HasPluginOrderPostOperationAbandon returns a boolean if a field has been set.

### GetPluginOrderPostOperationAdd

`func (o *PluginRootResponse) GetPluginOrderPostOperationAdd() string`

GetPluginOrderPostOperationAdd returns the PluginOrderPostOperationAdd field if non-nil, zero value otherwise.

### GetPluginOrderPostOperationAddOk

`func (o *PluginRootResponse) GetPluginOrderPostOperationAddOk() (*string, bool)`

GetPluginOrderPostOperationAddOk returns a tuple with the PluginOrderPostOperationAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostOperationAdd

`func (o *PluginRootResponse) SetPluginOrderPostOperationAdd(v string)`

SetPluginOrderPostOperationAdd sets PluginOrderPostOperationAdd field to given value.

### HasPluginOrderPostOperationAdd

`func (o *PluginRootResponse) HasPluginOrderPostOperationAdd() bool`

HasPluginOrderPostOperationAdd returns a boolean if a field has been set.

### GetPluginOrderPostOperationBind

`func (o *PluginRootResponse) GetPluginOrderPostOperationBind() string`

GetPluginOrderPostOperationBind returns the PluginOrderPostOperationBind field if non-nil, zero value otherwise.

### GetPluginOrderPostOperationBindOk

`func (o *PluginRootResponse) GetPluginOrderPostOperationBindOk() (*string, bool)`

GetPluginOrderPostOperationBindOk returns a tuple with the PluginOrderPostOperationBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostOperationBind

`func (o *PluginRootResponse) SetPluginOrderPostOperationBind(v string)`

SetPluginOrderPostOperationBind sets PluginOrderPostOperationBind field to given value.

### HasPluginOrderPostOperationBind

`func (o *PluginRootResponse) HasPluginOrderPostOperationBind() bool`

HasPluginOrderPostOperationBind returns a boolean if a field has been set.

### GetPluginOrderPostOperationCompare

`func (o *PluginRootResponse) GetPluginOrderPostOperationCompare() string`

GetPluginOrderPostOperationCompare returns the PluginOrderPostOperationCompare field if non-nil, zero value otherwise.

### GetPluginOrderPostOperationCompareOk

`func (o *PluginRootResponse) GetPluginOrderPostOperationCompareOk() (*string, bool)`

GetPluginOrderPostOperationCompareOk returns a tuple with the PluginOrderPostOperationCompare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostOperationCompare

`func (o *PluginRootResponse) SetPluginOrderPostOperationCompare(v string)`

SetPluginOrderPostOperationCompare sets PluginOrderPostOperationCompare field to given value.

### HasPluginOrderPostOperationCompare

`func (o *PluginRootResponse) HasPluginOrderPostOperationCompare() bool`

HasPluginOrderPostOperationCompare returns a boolean if a field has been set.

### GetPluginOrderPostOperationDelete

`func (o *PluginRootResponse) GetPluginOrderPostOperationDelete() string`

GetPluginOrderPostOperationDelete returns the PluginOrderPostOperationDelete field if non-nil, zero value otherwise.

### GetPluginOrderPostOperationDeleteOk

`func (o *PluginRootResponse) GetPluginOrderPostOperationDeleteOk() (*string, bool)`

GetPluginOrderPostOperationDeleteOk returns a tuple with the PluginOrderPostOperationDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostOperationDelete

`func (o *PluginRootResponse) SetPluginOrderPostOperationDelete(v string)`

SetPluginOrderPostOperationDelete sets PluginOrderPostOperationDelete field to given value.

### HasPluginOrderPostOperationDelete

`func (o *PluginRootResponse) HasPluginOrderPostOperationDelete() bool`

HasPluginOrderPostOperationDelete returns a boolean if a field has been set.

### GetPluginOrderPostOperationExtended

`func (o *PluginRootResponse) GetPluginOrderPostOperationExtended() string`

GetPluginOrderPostOperationExtended returns the PluginOrderPostOperationExtended field if non-nil, zero value otherwise.

### GetPluginOrderPostOperationExtendedOk

`func (o *PluginRootResponse) GetPluginOrderPostOperationExtendedOk() (*string, bool)`

GetPluginOrderPostOperationExtendedOk returns a tuple with the PluginOrderPostOperationExtended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostOperationExtended

`func (o *PluginRootResponse) SetPluginOrderPostOperationExtended(v string)`

SetPluginOrderPostOperationExtended sets PluginOrderPostOperationExtended field to given value.

### HasPluginOrderPostOperationExtended

`func (o *PluginRootResponse) HasPluginOrderPostOperationExtended() bool`

HasPluginOrderPostOperationExtended returns a boolean if a field has been set.

### GetPluginOrderPostOperationModify

`func (o *PluginRootResponse) GetPluginOrderPostOperationModify() string`

GetPluginOrderPostOperationModify returns the PluginOrderPostOperationModify field if non-nil, zero value otherwise.

### GetPluginOrderPostOperationModifyOk

`func (o *PluginRootResponse) GetPluginOrderPostOperationModifyOk() (*string, bool)`

GetPluginOrderPostOperationModifyOk returns a tuple with the PluginOrderPostOperationModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostOperationModify

`func (o *PluginRootResponse) SetPluginOrderPostOperationModify(v string)`

SetPluginOrderPostOperationModify sets PluginOrderPostOperationModify field to given value.

### HasPluginOrderPostOperationModify

`func (o *PluginRootResponse) HasPluginOrderPostOperationModify() bool`

HasPluginOrderPostOperationModify returns a boolean if a field has been set.

### GetPluginOrderPostOperationModifyDN

`func (o *PluginRootResponse) GetPluginOrderPostOperationModifyDN() string`

GetPluginOrderPostOperationModifyDN returns the PluginOrderPostOperationModifyDN field if non-nil, zero value otherwise.

### GetPluginOrderPostOperationModifyDNOk

`func (o *PluginRootResponse) GetPluginOrderPostOperationModifyDNOk() (*string, bool)`

GetPluginOrderPostOperationModifyDNOk returns a tuple with the PluginOrderPostOperationModifyDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostOperationModifyDN

`func (o *PluginRootResponse) SetPluginOrderPostOperationModifyDN(v string)`

SetPluginOrderPostOperationModifyDN sets PluginOrderPostOperationModifyDN field to given value.

### HasPluginOrderPostOperationModifyDN

`func (o *PluginRootResponse) HasPluginOrderPostOperationModifyDN() bool`

HasPluginOrderPostOperationModifyDN returns a boolean if a field has been set.

### GetPluginOrderPostOperationSearch

`func (o *PluginRootResponse) GetPluginOrderPostOperationSearch() string`

GetPluginOrderPostOperationSearch returns the PluginOrderPostOperationSearch field if non-nil, zero value otherwise.

### GetPluginOrderPostOperationSearchOk

`func (o *PluginRootResponse) GetPluginOrderPostOperationSearchOk() (*string, bool)`

GetPluginOrderPostOperationSearchOk returns a tuple with the PluginOrderPostOperationSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostOperationSearch

`func (o *PluginRootResponse) SetPluginOrderPostOperationSearch(v string)`

SetPluginOrderPostOperationSearch sets PluginOrderPostOperationSearch field to given value.

### HasPluginOrderPostOperationSearch

`func (o *PluginRootResponse) HasPluginOrderPostOperationSearch() bool`

HasPluginOrderPostOperationSearch returns a boolean if a field has been set.

### GetPluginOrderPostOperationUnbind

`func (o *PluginRootResponse) GetPluginOrderPostOperationUnbind() string`

GetPluginOrderPostOperationUnbind returns the PluginOrderPostOperationUnbind field if non-nil, zero value otherwise.

### GetPluginOrderPostOperationUnbindOk

`func (o *PluginRootResponse) GetPluginOrderPostOperationUnbindOk() (*string, bool)`

GetPluginOrderPostOperationUnbindOk returns a tuple with the PluginOrderPostOperationUnbind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostOperationUnbind

`func (o *PluginRootResponse) SetPluginOrderPostOperationUnbind(v string)`

SetPluginOrderPostOperationUnbind sets PluginOrderPostOperationUnbind field to given value.

### HasPluginOrderPostOperationUnbind

`func (o *PluginRootResponse) HasPluginOrderPostOperationUnbind() bool`

HasPluginOrderPostOperationUnbind returns a boolean if a field has been set.

### GetPluginOrderPostResponseAdd

`func (o *PluginRootResponse) GetPluginOrderPostResponseAdd() string`

GetPluginOrderPostResponseAdd returns the PluginOrderPostResponseAdd field if non-nil, zero value otherwise.

### GetPluginOrderPostResponseAddOk

`func (o *PluginRootResponse) GetPluginOrderPostResponseAddOk() (*string, bool)`

GetPluginOrderPostResponseAddOk returns a tuple with the PluginOrderPostResponseAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostResponseAdd

`func (o *PluginRootResponse) SetPluginOrderPostResponseAdd(v string)`

SetPluginOrderPostResponseAdd sets PluginOrderPostResponseAdd field to given value.

### HasPluginOrderPostResponseAdd

`func (o *PluginRootResponse) HasPluginOrderPostResponseAdd() bool`

HasPluginOrderPostResponseAdd returns a boolean if a field has been set.

### GetPluginOrderPostResponseBind

`func (o *PluginRootResponse) GetPluginOrderPostResponseBind() string`

GetPluginOrderPostResponseBind returns the PluginOrderPostResponseBind field if non-nil, zero value otherwise.

### GetPluginOrderPostResponseBindOk

`func (o *PluginRootResponse) GetPluginOrderPostResponseBindOk() (*string, bool)`

GetPluginOrderPostResponseBindOk returns a tuple with the PluginOrderPostResponseBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostResponseBind

`func (o *PluginRootResponse) SetPluginOrderPostResponseBind(v string)`

SetPluginOrderPostResponseBind sets PluginOrderPostResponseBind field to given value.

### HasPluginOrderPostResponseBind

`func (o *PluginRootResponse) HasPluginOrderPostResponseBind() bool`

HasPluginOrderPostResponseBind returns a boolean if a field has been set.

### GetPluginOrderPostResponseCompare

`func (o *PluginRootResponse) GetPluginOrderPostResponseCompare() string`

GetPluginOrderPostResponseCompare returns the PluginOrderPostResponseCompare field if non-nil, zero value otherwise.

### GetPluginOrderPostResponseCompareOk

`func (o *PluginRootResponse) GetPluginOrderPostResponseCompareOk() (*string, bool)`

GetPluginOrderPostResponseCompareOk returns a tuple with the PluginOrderPostResponseCompare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostResponseCompare

`func (o *PluginRootResponse) SetPluginOrderPostResponseCompare(v string)`

SetPluginOrderPostResponseCompare sets PluginOrderPostResponseCompare field to given value.

### HasPluginOrderPostResponseCompare

`func (o *PluginRootResponse) HasPluginOrderPostResponseCompare() bool`

HasPluginOrderPostResponseCompare returns a boolean if a field has been set.

### GetPluginOrderPostResponseDelete

`func (o *PluginRootResponse) GetPluginOrderPostResponseDelete() string`

GetPluginOrderPostResponseDelete returns the PluginOrderPostResponseDelete field if non-nil, zero value otherwise.

### GetPluginOrderPostResponseDeleteOk

`func (o *PluginRootResponse) GetPluginOrderPostResponseDeleteOk() (*string, bool)`

GetPluginOrderPostResponseDeleteOk returns a tuple with the PluginOrderPostResponseDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostResponseDelete

`func (o *PluginRootResponse) SetPluginOrderPostResponseDelete(v string)`

SetPluginOrderPostResponseDelete sets PluginOrderPostResponseDelete field to given value.

### HasPluginOrderPostResponseDelete

`func (o *PluginRootResponse) HasPluginOrderPostResponseDelete() bool`

HasPluginOrderPostResponseDelete returns a boolean if a field has been set.

### GetPluginOrderPostResponseExtended

`func (o *PluginRootResponse) GetPluginOrderPostResponseExtended() string`

GetPluginOrderPostResponseExtended returns the PluginOrderPostResponseExtended field if non-nil, zero value otherwise.

### GetPluginOrderPostResponseExtendedOk

`func (o *PluginRootResponse) GetPluginOrderPostResponseExtendedOk() (*string, bool)`

GetPluginOrderPostResponseExtendedOk returns a tuple with the PluginOrderPostResponseExtended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostResponseExtended

`func (o *PluginRootResponse) SetPluginOrderPostResponseExtended(v string)`

SetPluginOrderPostResponseExtended sets PluginOrderPostResponseExtended field to given value.

### HasPluginOrderPostResponseExtended

`func (o *PluginRootResponse) HasPluginOrderPostResponseExtended() bool`

HasPluginOrderPostResponseExtended returns a boolean if a field has been set.

### GetPluginOrderPostResponseModify

`func (o *PluginRootResponse) GetPluginOrderPostResponseModify() string`

GetPluginOrderPostResponseModify returns the PluginOrderPostResponseModify field if non-nil, zero value otherwise.

### GetPluginOrderPostResponseModifyOk

`func (o *PluginRootResponse) GetPluginOrderPostResponseModifyOk() (*string, bool)`

GetPluginOrderPostResponseModifyOk returns a tuple with the PluginOrderPostResponseModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostResponseModify

`func (o *PluginRootResponse) SetPluginOrderPostResponseModify(v string)`

SetPluginOrderPostResponseModify sets PluginOrderPostResponseModify field to given value.

### HasPluginOrderPostResponseModify

`func (o *PluginRootResponse) HasPluginOrderPostResponseModify() bool`

HasPluginOrderPostResponseModify returns a boolean if a field has been set.

### GetPluginOrderPostResponseModifyDN

`func (o *PluginRootResponse) GetPluginOrderPostResponseModifyDN() string`

GetPluginOrderPostResponseModifyDN returns the PluginOrderPostResponseModifyDN field if non-nil, zero value otherwise.

### GetPluginOrderPostResponseModifyDNOk

`func (o *PluginRootResponse) GetPluginOrderPostResponseModifyDNOk() (*string, bool)`

GetPluginOrderPostResponseModifyDNOk returns a tuple with the PluginOrderPostResponseModifyDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostResponseModifyDN

`func (o *PluginRootResponse) SetPluginOrderPostResponseModifyDN(v string)`

SetPluginOrderPostResponseModifyDN sets PluginOrderPostResponseModifyDN field to given value.

### HasPluginOrderPostResponseModifyDN

`func (o *PluginRootResponse) HasPluginOrderPostResponseModifyDN() bool`

HasPluginOrderPostResponseModifyDN returns a boolean if a field has been set.

### GetPluginOrderPostSynchronizationAdd

`func (o *PluginRootResponse) GetPluginOrderPostSynchronizationAdd() string`

GetPluginOrderPostSynchronizationAdd returns the PluginOrderPostSynchronizationAdd field if non-nil, zero value otherwise.

### GetPluginOrderPostSynchronizationAddOk

`func (o *PluginRootResponse) GetPluginOrderPostSynchronizationAddOk() (*string, bool)`

GetPluginOrderPostSynchronizationAddOk returns a tuple with the PluginOrderPostSynchronizationAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostSynchronizationAdd

`func (o *PluginRootResponse) SetPluginOrderPostSynchronizationAdd(v string)`

SetPluginOrderPostSynchronizationAdd sets PluginOrderPostSynchronizationAdd field to given value.

### HasPluginOrderPostSynchronizationAdd

`func (o *PluginRootResponse) HasPluginOrderPostSynchronizationAdd() bool`

HasPluginOrderPostSynchronizationAdd returns a boolean if a field has been set.

### GetPluginOrderPostSynchronizationDelete

`func (o *PluginRootResponse) GetPluginOrderPostSynchronizationDelete() string`

GetPluginOrderPostSynchronizationDelete returns the PluginOrderPostSynchronizationDelete field if non-nil, zero value otherwise.

### GetPluginOrderPostSynchronizationDeleteOk

`func (o *PluginRootResponse) GetPluginOrderPostSynchronizationDeleteOk() (*string, bool)`

GetPluginOrderPostSynchronizationDeleteOk returns a tuple with the PluginOrderPostSynchronizationDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostSynchronizationDelete

`func (o *PluginRootResponse) SetPluginOrderPostSynchronizationDelete(v string)`

SetPluginOrderPostSynchronizationDelete sets PluginOrderPostSynchronizationDelete field to given value.

### HasPluginOrderPostSynchronizationDelete

`func (o *PluginRootResponse) HasPluginOrderPostSynchronizationDelete() bool`

HasPluginOrderPostSynchronizationDelete returns a boolean if a field has been set.

### GetPluginOrderPostSynchronizationModify

`func (o *PluginRootResponse) GetPluginOrderPostSynchronizationModify() string`

GetPluginOrderPostSynchronizationModify returns the PluginOrderPostSynchronizationModify field if non-nil, zero value otherwise.

### GetPluginOrderPostSynchronizationModifyOk

`func (o *PluginRootResponse) GetPluginOrderPostSynchronizationModifyOk() (*string, bool)`

GetPluginOrderPostSynchronizationModifyOk returns a tuple with the PluginOrderPostSynchronizationModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostSynchronizationModify

`func (o *PluginRootResponse) SetPluginOrderPostSynchronizationModify(v string)`

SetPluginOrderPostSynchronizationModify sets PluginOrderPostSynchronizationModify field to given value.

### HasPluginOrderPostSynchronizationModify

`func (o *PluginRootResponse) HasPluginOrderPostSynchronizationModify() bool`

HasPluginOrderPostSynchronizationModify returns a boolean if a field has been set.

### GetPluginOrderPostSynchronizationModifyDN

`func (o *PluginRootResponse) GetPluginOrderPostSynchronizationModifyDN() string`

GetPluginOrderPostSynchronizationModifyDN returns the PluginOrderPostSynchronizationModifyDN field if non-nil, zero value otherwise.

### GetPluginOrderPostSynchronizationModifyDNOk

`func (o *PluginRootResponse) GetPluginOrderPostSynchronizationModifyDNOk() (*string, bool)`

GetPluginOrderPostSynchronizationModifyDNOk returns a tuple with the PluginOrderPostSynchronizationModifyDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostSynchronizationModifyDN

`func (o *PluginRootResponse) SetPluginOrderPostSynchronizationModifyDN(v string)`

SetPluginOrderPostSynchronizationModifyDN sets PluginOrderPostSynchronizationModifyDN field to given value.

### HasPluginOrderPostSynchronizationModifyDN

`func (o *PluginRootResponse) HasPluginOrderPostSynchronizationModifyDN() bool`

HasPluginOrderPostSynchronizationModifyDN returns a boolean if a field has been set.

### GetPluginOrderPostResponseSearch

`func (o *PluginRootResponse) GetPluginOrderPostResponseSearch() string`

GetPluginOrderPostResponseSearch returns the PluginOrderPostResponseSearch field if non-nil, zero value otherwise.

### GetPluginOrderPostResponseSearchOk

`func (o *PluginRootResponse) GetPluginOrderPostResponseSearchOk() (*string, bool)`

GetPluginOrderPostResponseSearchOk returns a tuple with the PluginOrderPostResponseSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderPostResponseSearch

`func (o *PluginRootResponse) SetPluginOrderPostResponseSearch(v string)`

SetPluginOrderPostResponseSearch sets PluginOrderPostResponseSearch field to given value.

### HasPluginOrderPostResponseSearch

`func (o *PluginRootResponse) HasPluginOrderPostResponseSearch() bool`

HasPluginOrderPostResponseSearch returns a boolean if a field has been set.

### GetPluginOrderSearchResultEntry

`func (o *PluginRootResponse) GetPluginOrderSearchResultEntry() string`

GetPluginOrderSearchResultEntry returns the PluginOrderSearchResultEntry field if non-nil, zero value otherwise.

### GetPluginOrderSearchResultEntryOk

`func (o *PluginRootResponse) GetPluginOrderSearchResultEntryOk() (*string, bool)`

GetPluginOrderSearchResultEntryOk returns a tuple with the PluginOrderSearchResultEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderSearchResultEntry

`func (o *PluginRootResponse) SetPluginOrderSearchResultEntry(v string)`

SetPluginOrderSearchResultEntry sets PluginOrderSearchResultEntry field to given value.

### HasPluginOrderSearchResultEntry

`func (o *PluginRootResponse) HasPluginOrderSearchResultEntry() bool`

HasPluginOrderSearchResultEntry returns a boolean if a field has been set.

### GetPluginOrderSearchResultReference

`func (o *PluginRootResponse) GetPluginOrderSearchResultReference() string`

GetPluginOrderSearchResultReference returns the PluginOrderSearchResultReference field if non-nil, zero value otherwise.

### GetPluginOrderSearchResultReferenceOk

`func (o *PluginRootResponse) GetPluginOrderSearchResultReferenceOk() (*string, bool)`

GetPluginOrderSearchResultReferenceOk returns a tuple with the PluginOrderSearchResultReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderSearchResultReference

`func (o *PluginRootResponse) SetPluginOrderSearchResultReference(v string)`

SetPluginOrderSearchResultReference sets PluginOrderSearchResultReference field to given value.

### HasPluginOrderSearchResultReference

`func (o *PluginRootResponse) HasPluginOrderSearchResultReference() bool`

HasPluginOrderSearchResultReference returns a boolean if a field has been set.

### GetPluginOrderSubordinateModifyDN

`func (o *PluginRootResponse) GetPluginOrderSubordinateModifyDN() string`

GetPluginOrderSubordinateModifyDN returns the PluginOrderSubordinateModifyDN field if non-nil, zero value otherwise.

### GetPluginOrderSubordinateModifyDNOk

`func (o *PluginRootResponse) GetPluginOrderSubordinateModifyDNOk() (*string, bool)`

GetPluginOrderSubordinateModifyDNOk returns a tuple with the PluginOrderSubordinateModifyDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderSubordinateModifyDN

`func (o *PluginRootResponse) SetPluginOrderSubordinateModifyDN(v string)`

SetPluginOrderSubordinateModifyDN sets PluginOrderSubordinateModifyDN field to given value.

### HasPluginOrderSubordinateModifyDN

`func (o *PluginRootResponse) HasPluginOrderSubordinateModifyDN() bool`

HasPluginOrderSubordinateModifyDN returns a boolean if a field has been set.

### GetPluginOrderIntermediateResponse

`func (o *PluginRootResponse) GetPluginOrderIntermediateResponse() string`

GetPluginOrderIntermediateResponse returns the PluginOrderIntermediateResponse field if non-nil, zero value otherwise.

### GetPluginOrderIntermediateResponseOk

`func (o *PluginRootResponse) GetPluginOrderIntermediateResponseOk() (*string, bool)`

GetPluginOrderIntermediateResponseOk returns a tuple with the PluginOrderIntermediateResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOrderIntermediateResponse

`func (o *PluginRootResponse) SetPluginOrderIntermediateResponse(v string)`

SetPluginOrderIntermediateResponse sets PluginOrderIntermediateResponse field to given value.

### HasPluginOrderIntermediateResponse

`func (o *PluginRootResponse) HasPluginOrderIntermediateResponse() bool`

HasPluginOrderIntermediateResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


