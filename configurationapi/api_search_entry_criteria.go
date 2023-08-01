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

// SearchEntryCriteriaApiService SearchEntryCriteriaApi service
type SearchEntryCriteriaApiService service

type ApiAddSearchEntryCriteriaRequest struct {
	ctx                           context.Context
	ApiService                    *SearchEntryCriteriaApiService
	addSearchEntryCriteriaRequest *AddSearchEntryCriteriaRequest
}

// Create a new Search Entry Criteria in the config
func (r ApiAddSearchEntryCriteriaRequest) AddSearchEntryCriteriaRequest(addSearchEntryCriteriaRequest AddSearchEntryCriteriaRequest) ApiAddSearchEntryCriteriaRequest {
	r.addSearchEntryCriteriaRequest = &addSearchEntryCriteriaRequest
	return r
}

func (r ApiAddSearchEntryCriteriaRequest) Execute() (*AddSearchEntryCriteria200Response, *http.Response, error) {
	return r.ApiService.AddSearchEntryCriteriaExecute(r)
}

/*
AddSearchEntryCriteria Add a new Search Entry Criteria to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddSearchEntryCriteriaRequest
*/
func (a *SearchEntryCriteriaApiService) AddSearchEntryCriteria(ctx context.Context) ApiAddSearchEntryCriteriaRequest {
	return ApiAddSearchEntryCriteriaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddSearchEntryCriteria200Response
func (a *SearchEntryCriteriaApiService) AddSearchEntryCriteriaExecute(r ApiAddSearchEntryCriteriaRequest) (*AddSearchEntryCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddSearchEntryCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchEntryCriteriaApiService.AddSearchEntryCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search-entry-criteria"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addSearchEntryCriteriaRequest == nil {
		return localVarReturnValue, nil, reportError("addSearchEntryCriteriaRequest is required and must be specified")
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
	localVarPostBody = r.addSearchEntryCriteriaRequest
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

type ApiDeleteSearchEntryCriteriaRequest struct {
	ctx                     context.Context
	ApiService              *SearchEntryCriteriaApiService
	searchEntryCriteriaName string
}

func (r ApiDeleteSearchEntryCriteriaRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSearchEntryCriteriaExecute(r)
}

/*
DeleteSearchEntryCriteria Delete a Search Entry Criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param searchEntryCriteriaName Name of the Search Entry Criteria
	@return ApiDeleteSearchEntryCriteriaRequest
*/
func (a *SearchEntryCriteriaApiService) DeleteSearchEntryCriteria(ctx context.Context, searchEntryCriteriaName string) ApiDeleteSearchEntryCriteriaRequest {
	return ApiDeleteSearchEntryCriteriaRequest{
		ApiService:              a,
		ctx:                     ctx,
		searchEntryCriteriaName: searchEntryCriteriaName,
	}
}

// Execute executes the request
func (a *SearchEntryCriteriaApiService) DeleteSearchEntryCriteriaExecute(r ApiDeleteSearchEntryCriteriaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchEntryCriteriaApiService.DeleteSearchEntryCriteria")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search-entry-criteria/{search-entry-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"search-entry-criteria-name"+"}", url.PathEscape(parameterValueToString(r.searchEntryCriteriaName, "searchEntryCriteriaName")), -1)

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

type ApiGetSearchEntryCriteriaRequest struct {
	ctx                     context.Context
	ApiService              *SearchEntryCriteriaApiService
	searchEntryCriteriaName string
}

func (r ApiGetSearchEntryCriteriaRequest) Execute() (*AddSearchEntryCriteria200Response, *http.Response, error) {
	return r.ApiService.GetSearchEntryCriteriaExecute(r)
}

/*
GetSearchEntryCriteria Returns a single Search Entry Criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param searchEntryCriteriaName Name of the Search Entry Criteria
	@return ApiGetSearchEntryCriteriaRequest
*/
func (a *SearchEntryCriteriaApiService) GetSearchEntryCriteria(ctx context.Context, searchEntryCriteriaName string) ApiGetSearchEntryCriteriaRequest {
	return ApiGetSearchEntryCriteriaRequest{
		ApiService:              a,
		ctx:                     ctx,
		searchEntryCriteriaName: searchEntryCriteriaName,
	}
}

// Execute executes the request
//
//	@return AddSearchEntryCriteria200Response
func (a *SearchEntryCriteriaApiService) GetSearchEntryCriteriaExecute(r ApiGetSearchEntryCriteriaRequest) (*AddSearchEntryCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddSearchEntryCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchEntryCriteriaApiService.GetSearchEntryCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search-entry-criteria/{search-entry-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"search-entry-criteria-name"+"}", url.PathEscape(parameterValueToString(r.searchEntryCriteriaName, "searchEntryCriteriaName")), -1)

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

type ApiListSearchEntryCriteriaRequest struct {
	ctx        context.Context
	ApiService *SearchEntryCriteriaApiService
	filter     *string
}

// SCIM filter
func (r ApiListSearchEntryCriteriaRequest) Filter(filter string) ApiListSearchEntryCriteriaRequest {
	r.filter = &filter
	return r
}

func (r ApiListSearchEntryCriteriaRequest) Execute() (*SearchEntryCriteriaListResponse, *http.Response, error) {
	return r.ApiService.ListSearchEntryCriteriaExecute(r)
}

/*
ListSearchEntryCriteria Returns a list of all Search Entry Criteria objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListSearchEntryCriteriaRequest
*/
func (a *SearchEntryCriteriaApiService) ListSearchEntryCriteria(ctx context.Context) ApiListSearchEntryCriteriaRequest {
	return ApiListSearchEntryCriteriaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchEntryCriteriaListResponse
func (a *SearchEntryCriteriaApiService) ListSearchEntryCriteriaExecute(r ApiListSearchEntryCriteriaRequest) (*SearchEntryCriteriaListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchEntryCriteriaListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchEntryCriteriaApiService.ListSearchEntryCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search-entry-criteria"

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

type ApiUpdateSearchEntryCriteriaRequest struct {
	ctx                     context.Context
	ApiService              *SearchEntryCriteriaApiService
	searchEntryCriteriaName string
	updateRequest           *UpdateRequest
}

// Update an existing Search Entry Criteria
func (r ApiUpdateSearchEntryCriteriaRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateSearchEntryCriteriaRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateSearchEntryCriteriaRequest) Execute() (*AddSearchEntryCriteria200Response, *http.Response, error) {
	return r.ApiService.UpdateSearchEntryCriteriaExecute(r)
}

/*
UpdateSearchEntryCriteria Update an existing Search Entry Criteria by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param searchEntryCriteriaName Name of the Search Entry Criteria
	@return ApiUpdateSearchEntryCriteriaRequest
*/
func (a *SearchEntryCriteriaApiService) UpdateSearchEntryCriteria(ctx context.Context, searchEntryCriteriaName string) ApiUpdateSearchEntryCriteriaRequest {
	return ApiUpdateSearchEntryCriteriaRequest{
		ApiService:              a,
		ctx:                     ctx,
		searchEntryCriteriaName: searchEntryCriteriaName,
	}
}

// Execute executes the request
//
//	@return AddSearchEntryCriteria200Response
func (a *SearchEntryCriteriaApiService) UpdateSearchEntryCriteriaExecute(r ApiUpdateSearchEntryCriteriaRequest) (*AddSearchEntryCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddSearchEntryCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchEntryCriteriaApiService.UpdateSearchEntryCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search-entry-criteria/{search-entry-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"search-entry-criteria-name"+"}", url.PathEscape(parameterValueToString(r.searchEntryCriteriaName, "searchEntryCriteriaName")), -1)

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
