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

// PrometheusMonitorAttributeMetricAPIService PrometheusMonitorAttributeMetricAPI service
type PrometheusMonitorAttributeMetricAPIService service

type ApiAddPrometheusMonitorAttributeMetricRequest struct {
	ctx                                        context.Context
	ApiService                                 *PrometheusMonitorAttributeMetricAPIService
	httpServletExtensionName                   string
	addPrometheusMonitorAttributeMetricRequest *AddPrometheusMonitorAttributeMetricRequest
}

// Create a new Prometheus Monitor Attribute Metric in the config
func (r ApiAddPrometheusMonitorAttributeMetricRequest) AddPrometheusMonitorAttributeMetricRequest(addPrometheusMonitorAttributeMetricRequest AddPrometheusMonitorAttributeMetricRequest) ApiAddPrometheusMonitorAttributeMetricRequest {
	r.addPrometheusMonitorAttributeMetricRequest = &addPrometheusMonitorAttributeMetricRequest
	return r
}

func (r ApiAddPrometheusMonitorAttributeMetricRequest) Execute() (*PrometheusMonitorAttributeMetricResponse, *http.Response, error) {
	return r.ApiService.AddPrometheusMonitorAttributeMetricExecute(r)
}

/*
AddPrometheusMonitorAttributeMetric Add a new Prometheus Monitor Attribute Metric to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param httpServletExtensionName Name of the HTTP Servlet Extension
	@return ApiAddPrometheusMonitorAttributeMetricRequest
*/
func (a *PrometheusMonitorAttributeMetricAPIService) AddPrometheusMonitorAttributeMetric(ctx context.Context, httpServletExtensionName string) ApiAddPrometheusMonitorAttributeMetricRequest {
	return ApiAddPrometheusMonitorAttributeMetricRequest{
		ApiService:               a,
		ctx:                      ctx,
		httpServletExtensionName: httpServletExtensionName,
	}
}

// Execute executes the request
//
//	@return PrometheusMonitorAttributeMetricResponse
func (a *PrometheusMonitorAttributeMetricAPIService) AddPrometheusMonitorAttributeMetricExecute(r ApiAddPrometheusMonitorAttributeMetricRequest) (*PrometheusMonitorAttributeMetricResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PrometheusMonitorAttributeMetricResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrometheusMonitorAttributeMetricAPIService.AddPrometheusMonitorAttributeMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-extensions/{http-servlet-extension-name}/prometheus-monitor-attribute-metrics"
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-extension-name"+"}", url.PathEscape(parameterValueToString(r.httpServletExtensionName, "httpServletExtensionName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addPrometheusMonitorAttributeMetricRequest == nil {
		return localVarReturnValue, nil, reportError("addPrometheusMonitorAttributeMetricRequest is required and must be specified")
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
	localVarPostBody = r.addPrometheusMonitorAttributeMetricRequest
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

type ApiDeletePrometheusMonitorAttributeMetricRequest struct {
	ctx                                  context.Context
	ApiService                           *PrometheusMonitorAttributeMetricAPIService
	prometheusMonitorAttributeMetricName string
	httpServletExtensionName             string
}

func (r ApiDeletePrometheusMonitorAttributeMetricRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePrometheusMonitorAttributeMetricExecute(r)
}

/*
DeletePrometheusMonitorAttributeMetric Delete a Prometheus Monitor Attribute Metric

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param prometheusMonitorAttributeMetricName Name of the Prometheus Monitor Attribute Metric
	@param httpServletExtensionName Name of the HTTP Servlet Extension
	@return ApiDeletePrometheusMonitorAttributeMetricRequest
*/
func (a *PrometheusMonitorAttributeMetricAPIService) DeletePrometheusMonitorAttributeMetric(ctx context.Context, prometheusMonitorAttributeMetricName string, httpServletExtensionName string) ApiDeletePrometheusMonitorAttributeMetricRequest {
	return ApiDeletePrometheusMonitorAttributeMetricRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		prometheusMonitorAttributeMetricName: prometheusMonitorAttributeMetricName,
		httpServletExtensionName:             httpServletExtensionName,
	}
}

// Execute executes the request
func (a *PrometheusMonitorAttributeMetricAPIService) DeletePrometheusMonitorAttributeMetricExecute(r ApiDeletePrometheusMonitorAttributeMetricRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrometheusMonitorAttributeMetricAPIService.DeletePrometheusMonitorAttributeMetric")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-extensions/{http-servlet-extension-name}/prometheus-monitor-attribute-metrics/{prometheus-monitor-attribute-metric-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"prometheus-monitor-attribute-metric-name"+"}", url.PathEscape(parameterValueToString(r.prometheusMonitorAttributeMetricName, "prometheusMonitorAttributeMetricName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-extension-name"+"}", url.PathEscape(parameterValueToString(r.httpServletExtensionName, "httpServletExtensionName")), -1)

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

type ApiGetPrometheusMonitorAttributeMetricRequest struct {
	ctx                                  context.Context
	ApiService                           *PrometheusMonitorAttributeMetricAPIService
	prometheusMonitorAttributeMetricName string
	httpServletExtensionName             string
}

func (r ApiGetPrometheusMonitorAttributeMetricRequest) Execute() (*PrometheusMonitorAttributeMetricResponse, *http.Response, error) {
	return r.ApiService.GetPrometheusMonitorAttributeMetricExecute(r)
}

/*
GetPrometheusMonitorAttributeMetric Returns a single Prometheus Monitor Attribute Metric

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param prometheusMonitorAttributeMetricName Name of the Prometheus Monitor Attribute Metric
	@param httpServletExtensionName Name of the HTTP Servlet Extension
	@return ApiGetPrometheusMonitorAttributeMetricRequest
*/
func (a *PrometheusMonitorAttributeMetricAPIService) GetPrometheusMonitorAttributeMetric(ctx context.Context, prometheusMonitorAttributeMetricName string, httpServletExtensionName string) ApiGetPrometheusMonitorAttributeMetricRequest {
	return ApiGetPrometheusMonitorAttributeMetricRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		prometheusMonitorAttributeMetricName: prometheusMonitorAttributeMetricName,
		httpServletExtensionName:             httpServletExtensionName,
	}
}

// Execute executes the request
//
//	@return PrometheusMonitorAttributeMetricResponse
func (a *PrometheusMonitorAttributeMetricAPIService) GetPrometheusMonitorAttributeMetricExecute(r ApiGetPrometheusMonitorAttributeMetricRequest) (*PrometheusMonitorAttributeMetricResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PrometheusMonitorAttributeMetricResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrometheusMonitorAttributeMetricAPIService.GetPrometheusMonitorAttributeMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-extensions/{http-servlet-extension-name}/prometheus-monitor-attribute-metrics/{prometheus-monitor-attribute-metric-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"prometheus-monitor-attribute-metric-name"+"}", url.PathEscape(parameterValueToString(r.prometheusMonitorAttributeMetricName, "prometheusMonitorAttributeMetricName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-extension-name"+"}", url.PathEscape(parameterValueToString(r.httpServletExtensionName, "httpServletExtensionName")), -1)

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

type ApiListPrometheusMonitorAttributeMetricsRequest struct {
	ctx                      context.Context
	ApiService               *PrometheusMonitorAttributeMetricAPIService
	httpServletExtensionName string
	filter                   *string
}

// SCIM filter
func (r ApiListPrometheusMonitorAttributeMetricsRequest) Filter(filter string) ApiListPrometheusMonitorAttributeMetricsRequest {
	r.filter = &filter
	return r
}

func (r ApiListPrometheusMonitorAttributeMetricsRequest) Execute() (*PrometheusMonitorAttributeMetricListResponse, *http.Response, error) {
	return r.ApiService.ListPrometheusMonitorAttributeMetricsExecute(r)
}

/*
ListPrometheusMonitorAttributeMetrics Returns a list of all Prometheus Monitor Attribute Metric objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param httpServletExtensionName Name of the HTTP Servlet Extension
	@return ApiListPrometheusMonitorAttributeMetricsRequest
*/
func (a *PrometheusMonitorAttributeMetricAPIService) ListPrometheusMonitorAttributeMetrics(ctx context.Context, httpServletExtensionName string) ApiListPrometheusMonitorAttributeMetricsRequest {
	return ApiListPrometheusMonitorAttributeMetricsRequest{
		ApiService:               a,
		ctx:                      ctx,
		httpServletExtensionName: httpServletExtensionName,
	}
}

// Execute executes the request
//
//	@return PrometheusMonitorAttributeMetricListResponse
func (a *PrometheusMonitorAttributeMetricAPIService) ListPrometheusMonitorAttributeMetricsExecute(r ApiListPrometheusMonitorAttributeMetricsRequest) (*PrometheusMonitorAttributeMetricListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PrometheusMonitorAttributeMetricListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrometheusMonitorAttributeMetricAPIService.ListPrometheusMonitorAttributeMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-extensions/{http-servlet-extension-name}/prometheus-monitor-attribute-metrics"
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-extension-name"+"}", url.PathEscape(parameterValueToString(r.httpServletExtensionName, "httpServletExtensionName")), -1)

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

type ApiUpdatePrometheusMonitorAttributeMetricRequest struct {
	ctx                                  context.Context
	ApiService                           *PrometheusMonitorAttributeMetricAPIService
	prometheusMonitorAttributeMetricName string
	httpServletExtensionName             string
	updateRequest                        *UpdateRequest
}

// Update an existing Prometheus Monitor Attribute Metric
func (r ApiUpdatePrometheusMonitorAttributeMetricRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdatePrometheusMonitorAttributeMetricRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdatePrometheusMonitorAttributeMetricRequest) Execute() (*PrometheusMonitorAttributeMetricResponse, *http.Response, error) {
	return r.ApiService.UpdatePrometheusMonitorAttributeMetricExecute(r)
}

/*
UpdatePrometheusMonitorAttributeMetric Update an existing Prometheus Monitor Attribute Metric by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param prometheusMonitorAttributeMetricName Name of the Prometheus Monitor Attribute Metric
	@param httpServletExtensionName Name of the HTTP Servlet Extension
	@return ApiUpdatePrometheusMonitorAttributeMetricRequest
*/
func (a *PrometheusMonitorAttributeMetricAPIService) UpdatePrometheusMonitorAttributeMetric(ctx context.Context, prometheusMonitorAttributeMetricName string, httpServletExtensionName string) ApiUpdatePrometheusMonitorAttributeMetricRequest {
	return ApiUpdatePrometheusMonitorAttributeMetricRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		prometheusMonitorAttributeMetricName: prometheusMonitorAttributeMetricName,
		httpServletExtensionName:             httpServletExtensionName,
	}
}

// Execute executes the request
//
//	@return PrometheusMonitorAttributeMetricResponse
func (a *PrometheusMonitorAttributeMetricAPIService) UpdatePrometheusMonitorAttributeMetricExecute(r ApiUpdatePrometheusMonitorAttributeMetricRequest) (*PrometheusMonitorAttributeMetricResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PrometheusMonitorAttributeMetricResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrometheusMonitorAttributeMetricAPIService.UpdatePrometheusMonitorAttributeMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-extensions/{http-servlet-extension-name}/prometheus-monitor-attribute-metrics/{prometheus-monitor-attribute-metric-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"prometheus-monitor-attribute-metric-name"+"}", url.PathEscape(parameterValueToString(r.prometheusMonitorAttributeMetricName, "prometheusMonitorAttributeMetricName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-extension-name"+"}", url.PathEscape(parameterValueToString(r.httpServletExtensionName, "httpServletExtensionName")), -1)

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
