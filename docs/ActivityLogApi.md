# \ActivityLogApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackingActivityGet**](ActivityLogApi.md#TrackingActivityGet) | **Get** /tracking/activity | Activity log
[**TrackingCsvGet**](ActivityLogApi.md#TrackingCsvGet) | **Get** /tracking/csv | Download CSV file with Activity Log
[**TrackingDownloadsIdGet**](ActivityLogApi.md#TrackingDownloadsIdGet) | **Get** /tracking/downloads/{id} | Share action file downloads
[**TrackingFilesIdGet**](ActivityLogApi.md#TrackingFilesIdGet) | **Get** /tracking/files/{id} | Share action files


# **TrackingActivityGet**
> []TrackingActivityRespItems TrackingActivityGet(ctx, optional)
Activity log

List activity (action) log. For details - https://dev.maytech.net/wiki/display/ISV3/Activity+Log 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| Log id | 
 **userId** | [**string**](.md)| User id | 
 **limit** | **float32**| Rows per page | [default to 100]
 **from** | **float32**| UTC timestamp | [default to 0]
 **to** | **float32**| UTC timestamp | 

### Return type

[**[]TrackingActivityRespItems**](TrackingActivityRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackingCsvGet**
> []TrackingCsvRespItems TrackingCsvGet(ctx, optional)
Download CSV file with Activity Log

Download CSV file content with Activity Log. For details - https://dev.maytech.net/wiki/display/ISV3/Activity+Log 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**string**](.md)| User id | 
 **from** | **float32**| UTC timestamp | 
 **to** | **float32**| UTC timestamp | 

### Return type

[**[]TrackingCsvRespItems**](TrackingCSVRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackingDownloadsIdGet**
> []TrackingDownloadsRespItems TrackingDownloadsIdGet(ctx, id)
Share action file downloads

List share action file downloads 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Shared file ID | 

### Return type

[**[]TrackingDownloadsRespItems**](TrackingDownloadsRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackingFilesIdGet**
> []TrackingFilesRespItems TrackingFilesIdGet(ctx, id)
Share action files

List share action files 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share action ID | 

### Return type

[**[]TrackingFilesRespItems**](TrackingFilesRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

