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

// ResultCriteriaAPIService ResultCriteriaAPI service
type ResultCriteriaAPIService service

type ApiAddResultCriteriaRequest struct {
	ctx                      context.Context
	ApiService               *ResultCriteriaAPIService
	addResultCriteriaRequest *AddResultCriteriaRequest
}

// Create a new Result Criteria in the config
func (r ApiAddResultCriteriaRequest) AddResultCriteriaRequest(addResultCriteriaRequest AddResultCriteriaRequest) ApiAddResultCriteriaRequest {
	r.addResultCriteriaRequest = &addResultCriteriaRequest
	return r
}

func (r ApiAddResultCriteriaRequest) Execute() (*AddResultCriteria200Response, *http.Response, error) {
	return r.ApiService.AddResultCriteriaExecute(r)
}

/*
AddResultCriteria Add a new Result Criteria to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddResultCriteriaRequest
*/
func (a *ResultCriteriaAPIService) AddResultCriteria(ctx context.Context) ApiAddResultCriteriaRequest {
	return ApiAddResultCriteriaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddResultCriteria200Response
func (a *ResultCriteriaAPIService) AddResultCriteriaExecute(r ApiAddResultCriteriaRequest) (*AddResultCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddResultCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResultCriteriaAPIService.AddResultCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/result-criteria"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addResultCriteriaRequest == nil {
		return localVarReturnValue, nil, reportError("addResultCriteriaRequest is required and must be specified")
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
	localVarPostBody = r.addResultCriteriaRequest
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

type ApiDeleteResultCriteriaRequest struct {
	ctx                context.Context
	ApiService         *ResultCriteriaAPIService
	resultCriteriaName string
}

func (r ApiDeleteResultCriteriaRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteResultCriteriaExecute(r)
}

/*
DeleteResultCriteria Delete a Result Criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resultCriteriaName Name of the Result Criteria
	@return ApiDeleteResultCriteriaRequest
*/
func (a *ResultCriteriaAPIService) DeleteResultCriteria(ctx context.Context, resultCriteriaName string) ApiDeleteResultCriteriaRequest {
	return ApiDeleteResultCriteriaRequest{
		ApiService:         a,
		ctx:                ctx,
		resultCriteriaName: resultCriteriaName,
	}
}

// Execute executes the request
func (a *ResultCriteriaAPIService) DeleteResultCriteriaExecute(r ApiDeleteResultCriteriaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResultCriteriaAPIService.DeleteResultCriteria")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/result-criteria/{result-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"result-criteria-name"+"}", url.PathEscape(parameterValueToString(r.resultCriteriaName, "resultCriteriaName")), -1)

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

type ApiGetResultCriteriaRequest struct {
	ctx                context.Context
	ApiService         *ResultCriteriaAPIService
	resultCriteriaName string
}

func (r ApiGetResultCriteriaRequest) Execute() (*AddResultCriteria200Response, *http.Response, error) {
	return r.ApiService.GetResultCriteriaExecute(r)
}

/*
GetResultCriteria Returns a single Result Criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resultCriteriaName Name of the Result Criteria
	@return ApiGetResultCriteriaRequest
*/
func (a *ResultCriteriaAPIService) GetResultCriteria(ctx context.Context, resultCriteriaName string) ApiGetResultCriteriaRequest {
	return ApiGetResultCriteriaRequest{
		ApiService:         a,
		ctx:                ctx,
		resultCriteriaName: resultCriteriaName,
	}
}

// Execute executes the request
//
//	@return AddResultCriteria200Response
func (a *ResultCriteriaAPIService) GetResultCriteriaExecute(r ApiGetResultCriteriaRequest) (*AddResultCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddResultCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResultCriteriaAPIService.GetResultCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/result-criteria/{result-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"result-criteria-name"+"}", url.PathEscape(parameterValueToString(r.resultCriteriaName, "resultCriteriaName")), -1)

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

type ApiListResultCriteriaRequest struct {
	ctx        context.Context
	ApiService *ResultCriteriaAPIService
	filter     *string
}

// SCIM filter
func (r ApiListResultCriteriaRequest) Filter(filter string) ApiListResultCriteriaRequest {
	r.filter = &filter
	return r
}

func (r ApiListResultCriteriaRequest) Execute() (*ResultCriteriaListResponse, *http.Response, error) {
	return r.ApiService.ListResultCriteriaExecute(r)
}

/*
ListResultCriteria Returns a list of all Result Criteria objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListResultCriteriaRequest
*/
func (a *ResultCriteriaAPIService) ListResultCriteria(ctx context.Context) ApiListResultCriteriaRequest {
	return ApiListResultCriteriaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ResultCriteriaListResponse
func (a *ResultCriteriaAPIService) ListResultCriteriaExecute(r ApiListResultCriteriaRequest) (*ResultCriteriaListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultCriteriaListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResultCriteriaAPIService.ListResultCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/result-criteria"

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

type ApiUpdateResultCriteriaRequest struct {
	ctx                context.Context
	ApiService         *ResultCriteriaAPIService
	resultCriteriaName string
	updateRequest      *UpdateRequest
}

// Update an existing Result Criteria
func (r ApiUpdateResultCriteriaRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateResultCriteriaRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateResultCriteriaRequest) Execute() (*AddResultCriteria200Response, *http.Response, error) {
	return r.ApiService.UpdateResultCriteriaExecute(r)
}

/*
UpdateResultCriteria Update an existing Result Criteria by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resultCriteriaName Name of the Result Criteria
	@return ApiUpdateResultCriteriaRequest
*/
func (a *ResultCriteriaAPIService) UpdateResultCriteria(ctx context.Context, resultCriteriaName string) ApiUpdateResultCriteriaRequest {
	return ApiUpdateResultCriteriaRequest{
		ApiService:         a,
		ctx:                ctx,
		resultCriteriaName: resultCriteriaName,
	}
}

// Execute executes the request
//
//	@return AddResultCriteria200Response
func (a *ResultCriteriaAPIService) UpdateResultCriteriaExecute(r ApiUpdateResultCriteriaRequest) (*AddResultCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddResultCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResultCriteriaAPIService.UpdateResultCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/result-criteria/{result-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"result-criteria-name"+"}", url.PathEscape(parameterValueToString(r.resultCriteriaName, "resultCriteriaName")), -1)

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
