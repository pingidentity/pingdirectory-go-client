/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// CorrelatedLdapDataViewAPIService CorrelatedLdapDataViewAPI service
type CorrelatedLdapDataViewAPIService service

type ApiAddCorrelatedLdapDataViewRequest struct {
	ctx                              context.Context
	ApiService                       *CorrelatedLdapDataViewAPIService
	scimResourceTypeName             string
	addCorrelatedLdapDataViewRequest *AddCorrelatedLdapDataViewRequest
}

// Create a new Correlated LDAP Data View in the config
func (r ApiAddCorrelatedLdapDataViewRequest) AddCorrelatedLdapDataViewRequest(addCorrelatedLdapDataViewRequest AddCorrelatedLdapDataViewRequest) ApiAddCorrelatedLdapDataViewRequest {
	r.addCorrelatedLdapDataViewRequest = &addCorrelatedLdapDataViewRequest
	return r
}

func (r ApiAddCorrelatedLdapDataViewRequest) Execute() (*CorrelatedLdapDataViewResponse, *http.Response, error) {
	return r.ApiService.AddCorrelatedLdapDataViewExecute(r)
}

/*
AddCorrelatedLdapDataView Add a new Correlated LDAP Data View to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param scimResourceTypeName Name of the SCIM Resource Type
	@return ApiAddCorrelatedLdapDataViewRequest
*/
func (a *CorrelatedLdapDataViewAPIService) AddCorrelatedLdapDataView(ctx context.Context, scimResourceTypeName string) ApiAddCorrelatedLdapDataViewRequest {
	return ApiAddCorrelatedLdapDataViewRequest{
		ApiService:           a,
		ctx:                  ctx,
		scimResourceTypeName: scimResourceTypeName,
	}
}

// Execute executes the request
//
//	@return CorrelatedLdapDataViewResponse
func (a *CorrelatedLdapDataViewAPIService) AddCorrelatedLdapDataViewExecute(r ApiAddCorrelatedLdapDataViewRequest) (*CorrelatedLdapDataViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CorrelatedLdapDataViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrelatedLdapDataViewAPIService.AddCorrelatedLdapDataView")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views"
	localVarPath = strings.Replace(localVarPath, "{"+"scim-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.scimResourceTypeName, "scimResourceTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addCorrelatedLdapDataViewRequest == nil {
		return localVarReturnValue, nil, reportError("addCorrelatedLdapDataViewRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addCorrelatedLdapDataViewRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteCorrelatedLdapDataViewRequest struct {
	ctx                        context.Context
	ApiService                 *CorrelatedLdapDataViewAPIService
	correlatedLdapDataViewName string
	scimResourceTypeName       string
}

func (r ApiDeleteCorrelatedLdapDataViewRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCorrelatedLdapDataViewExecute(r)
}

/*
DeleteCorrelatedLdapDataView Delete a Correlated LDAP Data View

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param correlatedLdapDataViewName Name of the Correlated LDAP Data View
	@param scimResourceTypeName Name of the SCIM Resource Type
	@return ApiDeleteCorrelatedLdapDataViewRequest
*/
func (a *CorrelatedLdapDataViewAPIService) DeleteCorrelatedLdapDataView(ctx context.Context, correlatedLdapDataViewName string, scimResourceTypeName string) ApiDeleteCorrelatedLdapDataViewRequest {
	return ApiDeleteCorrelatedLdapDataViewRequest{
		ApiService:                 a,
		ctx:                        ctx,
		correlatedLdapDataViewName: correlatedLdapDataViewName,
		scimResourceTypeName:       scimResourceTypeName,
	}
}

// Execute executes the request
func (a *CorrelatedLdapDataViewAPIService) DeleteCorrelatedLdapDataViewExecute(r ApiDeleteCorrelatedLdapDataViewRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrelatedLdapDataViewAPIService.DeleteCorrelatedLdapDataView")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"correlated-ldap-data-view-name"+"}", url.PathEscape(parameterValueToString(r.correlatedLdapDataViewName, "correlatedLdapDataViewName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scim-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.scimResourceTypeName, "scimResourceTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCorrelatedLdapDataViewRequest struct {
	ctx                        context.Context
	ApiService                 *CorrelatedLdapDataViewAPIService
	correlatedLdapDataViewName string
	scimResourceTypeName       string
}

func (r ApiGetCorrelatedLdapDataViewRequest) Execute() (*CorrelatedLdapDataViewResponse, *http.Response, error) {
	return r.ApiService.GetCorrelatedLdapDataViewExecute(r)
}

/*
GetCorrelatedLdapDataView Returns a single Correlated LDAP Data View

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param correlatedLdapDataViewName Name of the Correlated LDAP Data View
	@param scimResourceTypeName Name of the SCIM Resource Type
	@return ApiGetCorrelatedLdapDataViewRequest
*/
func (a *CorrelatedLdapDataViewAPIService) GetCorrelatedLdapDataView(ctx context.Context, correlatedLdapDataViewName string, scimResourceTypeName string) ApiGetCorrelatedLdapDataViewRequest {
	return ApiGetCorrelatedLdapDataViewRequest{
		ApiService:                 a,
		ctx:                        ctx,
		correlatedLdapDataViewName: correlatedLdapDataViewName,
		scimResourceTypeName:       scimResourceTypeName,
	}
}

// Execute executes the request
//
//	@return CorrelatedLdapDataViewResponse
func (a *CorrelatedLdapDataViewAPIService) GetCorrelatedLdapDataViewExecute(r ApiGetCorrelatedLdapDataViewRequest) (*CorrelatedLdapDataViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CorrelatedLdapDataViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrelatedLdapDataViewAPIService.GetCorrelatedLdapDataView")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"correlated-ldap-data-view-name"+"}", url.PathEscape(parameterValueToString(r.correlatedLdapDataViewName, "correlatedLdapDataViewName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scim-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.scimResourceTypeName, "scimResourceTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListCorrelatedLdapDataViewsRequest struct {
	ctx                  context.Context
	ApiService           *CorrelatedLdapDataViewAPIService
	scimResourceTypeName string
	filter               *string
}

// SCIM filter
func (r ApiListCorrelatedLdapDataViewsRequest) Filter(filter string) ApiListCorrelatedLdapDataViewsRequest {
	r.filter = &filter
	return r
}

func (r ApiListCorrelatedLdapDataViewsRequest) Execute() (*CorrelatedLdapDataViewListResponse, *http.Response, error) {
	return r.ApiService.ListCorrelatedLdapDataViewsExecute(r)
}

/*
ListCorrelatedLdapDataViews Returns a list of all Correlated LDAP Data View objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param scimResourceTypeName Name of the SCIM Resource Type
	@return ApiListCorrelatedLdapDataViewsRequest
*/
func (a *CorrelatedLdapDataViewAPIService) ListCorrelatedLdapDataViews(ctx context.Context, scimResourceTypeName string) ApiListCorrelatedLdapDataViewsRequest {
	return ApiListCorrelatedLdapDataViewsRequest{
		ApiService:           a,
		ctx:                  ctx,
		scimResourceTypeName: scimResourceTypeName,
	}
}

// Execute executes the request
//
//	@return CorrelatedLdapDataViewListResponse
func (a *CorrelatedLdapDataViewAPIService) ListCorrelatedLdapDataViewsExecute(r ApiListCorrelatedLdapDataViewsRequest) (*CorrelatedLdapDataViewListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CorrelatedLdapDataViewListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrelatedLdapDataViewAPIService.ListCorrelatedLdapDataViews")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views"
	localVarPath = strings.Replace(localVarPath, "{"+"scim-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.scimResourceTypeName, "scimResourceTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateCorrelatedLdapDataViewRequest struct {
	ctx                        context.Context
	ApiService                 *CorrelatedLdapDataViewAPIService
	correlatedLdapDataViewName string
	scimResourceTypeName       string
	updateRequest              *UpdateRequest
}

// Update an existing Correlated LDAP Data View
func (r ApiUpdateCorrelatedLdapDataViewRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateCorrelatedLdapDataViewRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateCorrelatedLdapDataViewRequest) Execute() (*CorrelatedLdapDataViewResponse, *http.Response, error) {
	return r.ApiService.UpdateCorrelatedLdapDataViewExecute(r)
}

/*
UpdateCorrelatedLdapDataView Update an existing Correlated LDAP Data View by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param correlatedLdapDataViewName Name of the Correlated LDAP Data View
	@param scimResourceTypeName Name of the SCIM Resource Type
	@return ApiUpdateCorrelatedLdapDataViewRequest
*/
func (a *CorrelatedLdapDataViewAPIService) UpdateCorrelatedLdapDataView(ctx context.Context, correlatedLdapDataViewName string, scimResourceTypeName string) ApiUpdateCorrelatedLdapDataViewRequest {
	return ApiUpdateCorrelatedLdapDataViewRequest{
		ApiService:                 a,
		ctx:                        ctx,
		correlatedLdapDataViewName: correlatedLdapDataViewName,
		scimResourceTypeName:       scimResourceTypeName,
	}
}

// Execute executes the request
//
//	@return CorrelatedLdapDataViewResponse
func (a *CorrelatedLdapDataViewAPIService) UpdateCorrelatedLdapDataViewExecute(r ApiUpdateCorrelatedLdapDataViewRequest) (*CorrelatedLdapDataViewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CorrelatedLdapDataViewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CorrelatedLdapDataViewAPIService.UpdateCorrelatedLdapDataView")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"correlated-ldap-data-view-name"+"}", url.PathEscape(parameterValueToString(r.correlatedLdapDataViewName, "correlatedLdapDataViewName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scim-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.scimResourceTypeName, "scimResourceTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRequest == nil {
		return localVarReturnValue, nil, reportError("updateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
