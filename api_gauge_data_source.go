/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// GaugeDataSourceApiService GaugeDataSourceApi service
type GaugeDataSourceApiService service

type ApiAddGaugeDataSourceRequest struct {
	ctx                       context.Context
	ApiService                *GaugeDataSourceApiService
	addGaugeDataSourceRequest *AddGaugeDataSourceRequest
}

// Create a new Gauge Data Source in the config
func (r ApiAddGaugeDataSourceRequest) AddGaugeDataSourceRequest(addGaugeDataSourceRequest AddGaugeDataSourceRequest) ApiAddGaugeDataSourceRequest {
	r.addGaugeDataSourceRequest = &addGaugeDataSourceRequest
	return r
}

func (r ApiAddGaugeDataSourceRequest) Execute() (*AddGaugeDataSource200Response, *http.Response, error) {
	return r.ApiService.AddGaugeDataSourceExecute(r)
}

/*
AddGaugeDataSource Add a new Gauge Data Source to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddGaugeDataSourceRequest
*/
func (a *GaugeDataSourceApiService) AddGaugeDataSource(ctx context.Context) ApiAddGaugeDataSourceRequest {
	return ApiAddGaugeDataSourceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddGaugeDataSource200Response
func (a *GaugeDataSourceApiService) AddGaugeDataSourceExecute(r ApiAddGaugeDataSourceRequest) (*AddGaugeDataSource200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddGaugeDataSource200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GaugeDataSourceApiService.AddGaugeDataSource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gauge-data-sources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addGaugeDataSourceRequest == nil {
		return localVarReturnValue, nil, reportError("addGaugeDataSourceRequest is required and must be specified")
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
	localVarPostBody = r.addGaugeDataSourceRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeleteGaugeDataSourceRequest struct {
	ctx                 context.Context
	ApiService          *GaugeDataSourceApiService
	gaugeDataSourceName string
}

func (r ApiDeleteGaugeDataSourceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGaugeDataSourceExecute(r)
}

/*
DeleteGaugeDataSource Delete a Gauge Data Source

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gaugeDataSourceName Name of the Gauge Data Source
	@return ApiDeleteGaugeDataSourceRequest
*/
func (a *GaugeDataSourceApiService) DeleteGaugeDataSource(ctx context.Context, gaugeDataSourceName string) ApiDeleteGaugeDataSourceRequest {
	return ApiDeleteGaugeDataSourceRequest{
		ApiService:          a,
		ctx:                 ctx,
		gaugeDataSourceName: gaugeDataSourceName,
	}
}

// Execute executes the request
func (a *GaugeDataSourceApiService) DeleteGaugeDataSourceExecute(r ApiDeleteGaugeDataSourceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GaugeDataSourceApiService.DeleteGaugeDataSource")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gauge-data-sources/{gauge-data-source-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"gauge-data-source-name"+"}", url.PathEscape(parameterToString(r.gaugeDataSourceName, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetGaugeDataSourceRequest struct {
	ctx                 context.Context
	ApiService          *GaugeDataSourceApiService
	gaugeDataSourceName string
}

func (r ApiGetGaugeDataSourceRequest) Execute() (*AddGaugeDataSource200Response, *http.Response, error) {
	return r.ApiService.GetGaugeDataSourceExecute(r)
}

/*
GetGaugeDataSource Returns a single Gauge Data Source

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gaugeDataSourceName Name of the Gauge Data Source
	@return ApiGetGaugeDataSourceRequest
*/
func (a *GaugeDataSourceApiService) GetGaugeDataSource(ctx context.Context, gaugeDataSourceName string) ApiGetGaugeDataSourceRequest {
	return ApiGetGaugeDataSourceRequest{
		ApiService:          a,
		ctx:                 ctx,
		gaugeDataSourceName: gaugeDataSourceName,
	}
}

// Execute executes the request
//
//	@return AddGaugeDataSource200Response
func (a *GaugeDataSourceApiService) GetGaugeDataSourceExecute(r ApiGetGaugeDataSourceRequest) (*AddGaugeDataSource200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddGaugeDataSource200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GaugeDataSourceApiService.GetGaugeDataSource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gauge-data-sources/{gauge-data-source-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"gauge-data-source-name"+"}", url.PathEscape(parameterToString(r.gaugeDataSourceName, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdateGaugeDataSourceRequest struct {
	ctx                 context.Context
	ApiService          *GaugeDataSourceApiService
	gaugeDataSourceName string
	updateRequest       *UpdateRequest
}

// Update an existing Gauge Data Source
func (r ApiUpdateGaugeDataSourceRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateGaugeDataSourceRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateGaugeDataSourceRequest) Execute() (*AddGaugeDataSource200Response, *http.Response, error) {
	return r.ApiService.UpdateGaugeDataSourceExecute(r)
}

/*
UpdateGaugeDataSource Update an existing Gauge Data Source by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gaugeDataSourceName Name of the Gauge Data Source
	@return ApiUpdateGaugeDataSourceRequest
*/
func (a *GaugeDataSourceApiService) UpdateGaugeDataSource(ctx context.Context, gaugeDataSourceName string) ApiUpdateGaugeDataSourceRequest {
	return ApiUpdateGaugeDataSourceRequest{
		ApiService:          a,
		ctx:                 ctx,
		gaugeDataSourceName: gaugeDataSourceName,
	}
}

// Execute executes the request
//
//	@return AddGaugeDataSource200Response
func (a *GaugeDataSourceApiService) UpdateGaugeDataSourceExecute(r ApiUpdateGaugeDataSourceRequest) (*AddGaugeDataSource200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddGaugeDataSource200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GaugeDataSourceApiService.UpdateGaugeDataSource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gauge-data-sources/{gauge-data-source-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"gauge-data-source-name"+"}", url.PathEscape(parameterToString(r.gaugeDataSourceName, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
