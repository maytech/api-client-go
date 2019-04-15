# \ActivityLogApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackingActivityGet**](ActivityLogApi.md#TrackingActivityGet) | **Get** /tracking/activity | Get activity log
[**TrackingCsvGet**](ActivityLogApi.md#TrackingCsvGet) | **Get** /tracking/csv | Download CSV file with activity log
[**TrackingDownloadsIdGet**](ActivityLogApi.md#TrackingDownloadsIdGet) | **Get** /tracking/downloads/{id} | Get shared file downloads
[**TrackingFilesIdGet**](ActivityLogApi.md#TrackingFilesIdGet) | **Get** /tracking/files/{id} | Get share action files


# **TrackingActivityGet**
> []TrackingActivityRespItems TrackingActivityGet(ctx, optional)
Get activity log

Retrieve the history of actions by the specified user (all visible users) for a given period. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| Log ID | 
 **userId** | [**string**](.md)| User ID | 
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
Download CSV file with activity log

Download a file with full activity log in the CSV format by given user ID. If the user ID is not specified, the activity of all manageable users should be displayed. The content of the file will display all actions performed in the account for a specified period. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**string**](.md)| User ID | 
 **from** | **float32**| UTC timestamp | 
 **to** | **float32**| UTC timestamp | 

### Return type

[**[]TrackingCsvRespItems**](TrackingCSVRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: raw

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackingDownloadsIdGet**
> []TrackingDownloadsRespItems TrackingDownloadsIdGet(ctx, id)
Get shared file downloads

Retrieve information about download actions of the file. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Shared file ID - File ID from /tacking/files/id | 

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
Get share action files

Retrieve a list of shared files by specified share action ID with the number of downloads. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share ID | 

### Return type

[**[]TrackingFilesRespItems**](TrackingFilesRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

