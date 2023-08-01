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

// SearchReferenceCriteriaApiService SearchReferenceCriteriaApi service
type SearchReferenceCriteriaApiService service

type ApiAddSearchReferenceCriteriaRequest struct {
	ctx                               context.Context
	ApiService                        *SearchReferenceCriteriaApiService
	addSearchReferenceCriteriaRequest *AddSearchReferenceCriteriaRequest
}

// Create a new Search Reference Criteria in the config
func (r ApiAddSearchReferenceCriteriaRequest) AddSearchReferenceCriteriaRequest(addSearchReferenceCriteriaRequest AddSearchReferenceCriteriaRequest) ApiAddSearchReferenceCriteriaRequest {
	r.addSearchReferenceCriteriaRequest = &addSearchReferenceCriteriaRequest
	return r
}

func (r ApiAddSearchReferenceCriteriaRequest) Execute() (*AddSearchReferenceCriteria200Response, *http.Response, error) {
	return r.ApiService.AddSearchReferenceCriteriaExecute(r)
}

/*
AddSearchReferenceCriteria Add a new Search Reference Criteria to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddSearchReferenceCriteriaRequest
*/
func (a *SearchReferenceCriteriaApiService) AddSearchReferenceCriteria(ctx context.Context) ApiAddSearchReferenceCriteriaRequest {
	return ApiAddSearchReferenceCriteriaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddSearchReferenceCriteria200Response
func (a *SearchReferenceCriteriaApiService) AddSearchReferenceCriteriaExecute(r ApiAddSearchReferenceCriteriaRequest) (*AddSearchReferenceCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddSearchReferenceCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchReferenceCriteriaApiService.AddSearchReferenceCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search-reference-criteria"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addSearchReferenceCriteriaRequest == nil {
		return localVarReturnValue, nil, reportError("addSearchReferenceCriteriaRequest is required and must be specified")
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
	localVarPostBody = r.addSearchReferenceCriteriaRequest
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

type ApiDeleteSearchReferenceCriteriaRequest struct {
	ctx                         context.Context
	ApiService                  *SearchReferenceCriteriaApiService
	searchReferenceCriteriaName string
}

func (r ApiDeleteSearchReferenceCriteriaRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSearchReferenceCriteriaExecute(r)
}

/*
DeleteSearchReferenceCriteria Delete a Search Reference Criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param searchReferenceCriteriaName Name of the Search Reference Criteria
	@return ApiDeleteSearchReferenceCriteriaRequest
*/
func (a *SearchReferenceCriteriaApiService) DeleteSearchReferenceCriteria(ctx context.Context, searchReferenceCriteriaName string) ApiDeleteSearchReferenceCriteriaRequest {
	return ApiDeleteSearchReferenceCriteriaRequest{
		ApiService:                  a,
		ctx:                         ctx,
		searchReferenceCriteriaName: searchReferenceCriteriaName,
	}
}

// Execute executes the request
func (a *SearchReferenceCriteriaApiService) DeleteSearchReferenceCriteriaExecute(r ApiDeleteSearchReferenceCriteriaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchReferenceCriteriaApiService.DeleteSearchReferenceCriteria")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search-reference-criteria/{search-reference-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"search-reference-criteria-name"+"}", url.PathEscape(parameterValueToString(r.searchReferenceCriteriaName, "searchReferenceCriteriaName")), -1)

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

type ApiGetSearchReferenceCriteriaRequest struct {
	ctx                         context.Context
	ApiService                  *SearchReferenceCriteriaApiService
	searchReferenceCriteriaName string
}

func (r ApiGetSearchReferenceCriteriaRequest) Execute() (*AddSearchReferenceCriteria200Response, *http.Response, error) {
	return r.ApiService.GetSearchReferenceCriteriaExecute(r)
}

/*
GetSearchReferenceCriteria Returns a single Search Reference Criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param searchReferenceCriteriaName Name of the Search Reference Criteria
	@return ApiGetSearchReferenceCriteriaRequest
*/
func (a *SearchReferenceCriteriaApiService) GetSearchReferenceCriteria(ctx context.Context, searchReferenceCriteriaName string) ApiGetSearchReferenceCriteriaRequest {
	return ApiGetSearchReferenceCriteriaRequest{
		ApiService:                  a,
		ctx:                         ctx,
		searchReferenceCriteriaName: searchReferenceCriteriaName,
	}
}

// Execute executes the request
//
//	@return AddSearchReferenceCriteria200Response
func (a *SearchReferenceCriteriaApiService) GetSearchReferenceCriteriaExecute(r ApiGetSearchReferenceCriteriaRequest) (*AddSearchReferenceCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddSearchReferenceCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchReferenceCriteriaApiService.GetSearchReferenceCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search-reference-criteria/{search-reference-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"search-reference-criteria-name"+"}", url.PathEscape(parameterValueToString(r.searchReferenceCriteriaName, "searchReferenceCriteriaName")), -1)

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

type ApiListSearchReferenceCriteriaRequest struct {
	ctx        context.Context
	ApiService *SearchReferenceCriteriaApiService
	filter     *string
}

// SCIM filter
func (r ApiListSearchReferenceCriteriaRequest) Filter(filter string) ApiListSearchReferenceCriteriaRequest {
	r.filter = &filter
	return r
}

func (r ApiListSearchReferenceCriteriaRequest) Execute() (*SearchReferenceCriteriaListResponse, *http.Response, error) {
	return r.ApiService.ListSearchReferenceCriteriaExecute(r)
}

/*
ListSearchReferenceCriteria Returns a list of all Search Reference Criteria objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListSearchReferenceCriteriaRequest
*/
func (a *SearchReferenceCriteriaApiService) ListSearchReferenceCriteria(ctx context.Context) ApiListSearchReferenceCriteriaRequest {
	return ApiListSearchReferenceCriteriaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchReferenceCriteriaListResponse
func (a *SearchReferenceCriteriaApiService) ListSearchReferenceCriteriaExecute(r ApiListSearchReferenceCriteriaRequest) (*SearchReferenceCriteriaListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchReferenceCriteriaListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchReferenceCriteriaApiService.ListSearchReferenceCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search-reference-criteria"

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

type ApiUpdateSearchReferenceCriteriaRequest struct {
	ctx                         context.Context
	ApiService                  *SearchReferenceCriteriaApiService
	searchReferenceCriteriaName string
	updateRequest               *UpdateRequest
}

// Update an existing Search Reference Criteria
func (r ApiUpdateSearchReferenceCriteriaRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateSearchReferenceCriteriaRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateSearchReferenceCriteriaRequest) Execute() (*AddSearchReferenceCriteria200Response, *http.Response, error) {
	return r.ApiService.UpdateSearchReferenceCriteriaExecute(r)
}

/*
UpdateSearchReferenceCriteria Update an existing Search Reference Criteria by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param searchReferenceCriteriaName Name of the Search Reference Criteria
	@return ApiUpdateSearchReferenceCriteriaRequest
*/
func (a *SearchReferenceCriteriaApiService) UpdateSearchReferenceCriteria(ctx context.Context, searchReferenceCriteriaName string) ApiUpdateSearchReferenceCriteriaRequest {
	return ApiUpdateSearchReferenceCriteriaRequest{
		ApiService:                  a,
		ctx:                         ctx,
		searchReferenceCriteriaName: searchReferenceCriteriaName,
	}
}

// Execute executes the request
//
//	@return AddSearchReferenceCriteria200Response
func (a *SearchReferenceCriteriaApiService) UpdateSearchReferenceCriteriaExecute(r ApiUpdateSearchReferenceCriteriaRequest) (*AddSearchReferenceCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddSearchReferenceCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchReferenceCriteriaApiService.UpdateSearchReferenceCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search-reference-criteria/{search-reference-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"search-reference-criteria-name"+"}", url.PathEscape(parameterValueToString(r.searchReferenceCriteriaName, "searchReferenceCriteriaName")), -1)

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
