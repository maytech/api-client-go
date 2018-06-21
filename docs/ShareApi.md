# \ShareApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesReturnMakedirIdPost**](ShareApi.md#FilesReturnMakedirIdPost) | **Post** /files-return/makedir/{id} | Return files makedir
[**FilesReturnMetadataIdGet**](ShareApi.md#FilesReturnMetadataIdGet) | **Get** /files-return/metadata/{id} | Get return files metadata
[**FilesReturnSendPost**](ShareApi.md#FilesReturnSendPost) | **Post** /files-return/send | Send return files share
[**FilesReturnUploadLinkIdPost**](ShareApi.md#FilesReturnUploadLinkIdPost) | **Post** /files-return/upload-link/{id} | Get return files upload link
[**QuicklinkCreatePost**](ShareApi.md#QuicklinkCreatePost) | **Post** /quicklink/create | Create quicklink
[**QuicklinkLoginPinPost**](ShareApi.md#QuicklinkLoginPinPost) | **Post** /quicklink/login-pin | Login with PIN on quicklink
[**QuicklinkRevokeIdGet**](ShareApi.md#QuicklinkRevokeIdGet) | **Get** /quicklink/revoke/{id} | Revoke quicklink
[**ShareCreatePost**](ShareApi.md#ShareCreatePost) | **Post** /share/create | Create share
[**ShareDownloadIdGet**](ShareApi.md#ShareDownloadIdGet) | **Get** /share/download/{id} | Download share files
[**ShareDownloadInfoIdGet**](ShareApi.md#ShareDownloadInfoIdGet) | **Get** /share/download-info/{id} | Share download info
[**ShareDownloadLinkIdGet**](ShareApi.md#ShareDownloadLinkIdGet) | **Get** /share/download-link/{id} | Download link
[**ShareDownloadLinkIdPost**](ShareApi.md#ShareDownloadLinkIdPost) | **Post** /share/download-link/{id} | Download link
[**ShareFilesIdGet**](ShareApi.md#ShareFilesIdGet) | **Get** /share/files/{id} | Share files
[**ShareLoginPinPost**](ShareApi.md#ShareLoginPinPost) | **Post** /share/login-pin | Login with pin on share
[**SharePreviewIdGet**](ShareApi.md#SharePreviewIdGet) | **Get** /share/preview/{id} | Preview share file
[**ShareRecipientsGet**](ShareApi.md#ShareRecipientsGet) | **Get** /share/recipients | Share recipinets
[**ShareRequestPost**](ShareApi.md#ShareRequestPost) | **Post** /share/request | Request share
[**ShareRevokeIdGet**](ShareApi.md#ShareRevokeIdGet) | **Get** /share/revoke/{id} | Revoke share
[**ShareSendRequestIdPost**](ShareApi.md#ShareSendRequestIdPost) | **Post** /share/send-request/{id} | DEPRECATED! Use /share/request instead.
[**TrackingGet**](ShareApi.md#TrackingGet) | **Get** /tracking/ | List share actions
[**TrackingIdGet**](ShareApi.md#TrackingIdGet) | **Get** /tracking/{id} | List share actions


# **FilesReturnMakedirIdPost**
> FilesReturnMakedirResps FilesReturnMakedirIdPost(ctx, id, body)
Return files makedir

Create folder via return files 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Share action ID | 
  **body** | [**FilesReturnMakedirReq**](FilesReturnMakedirReq.md)|  | 

### Return type

[**FilesReturnMakedirResps**](FilesReturnMakedirResps.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesReturnMetadataIdGet**
> FilesReturnMetadataResp FilesReturnMetadataIdGet(ctx, id)
Get return files metadata

Get return files metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share action ID | 

### Return type

[**FilesReturnMetadataResp**](FilesReturnMetadataResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesReturnSendPost**
> JobResp FilesReturnSendPost(ctx, body)
Send return files share

Create and send return files share 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FilesReturnSendReq**](FilesReturnSendReq.md)|  | 

### Return type

[**JobResp**](JobResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesReturnUploadLinkIdPost**
> FilesReturnUploadLinkResp FilesReturnUploadLinkIdPost(ctx, id, body)
Get return files upload link

Get return files upload link 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Share action ID | 
  **body** | [**FilesReturnUploadLinkReq**](FilesReturnUploadLinkReq.md)|  | 

### Return type

[**FilesReturnUploadLinkResp**](FilesReturnUploadLinkResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuicklinkCreatePost**
> QuicklinkCreateResp QuicklinkCreatePost(ctx, body)
Create quicklink

Creating public share with no recipients 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**QuicklinkCreateReq**](QuicklinkCreateReq.md)|  | 

### Return type

[**QuicklinkCreateResp**](QuicklinkCreateResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuicklinkLoginPinPost**
> QuicklinkLoginPinPost(ctx, body)
Login with PIN on quicklink

Login on quicklink protected with PIN 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**QuicklinkLoginPinReq**](QuicklinkLoginPinReq.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuicklinkRevokeIdGet**
> IdResp QuicklinkRevokeIdGet(ctx, id)
Revoke quicklink

Revoke quicklink 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share action ID | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareCreatePost**
> ShareCreateResp ShareCreatePost(ctx, body)
Create share

Create share 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ShareCreateReq**](ShareCreateReq.md)|  | 

### Return type

[**ShareCreateResp**](ShareCreateResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareDownloadIdGet**
> ShareDownloadIdGet(ctx, id, optional)
Download share files

Download share files content 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Download link ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| Download link ID | 
 **files** | [**[]string**](string.md)| File ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareDownloadInfoIdGet**
> ShareDownloadInfoResp ShareDownloadInfoIdGet(ctx, id)
Share download info

Get share download info 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share action ID | 

### Return type

[**ShareDownloadInfoResp**](ShareDownloadInfoResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareDownloadLinkIdGet**
> IdResp ShareDownloadLinkIdGet(ctx, id)
Download link

Download link for all files 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share action ID | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareDownloadLinkIdPost**
> IdResp ShareDownloadLinkIdPost(ctx, id, optional)
Download link

Download link for selected files 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share action ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| Share action ID | 
 **body** | [**ShareDownloadLinkReq**](ShareDownloadLinkReq.md)|  | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareFilesIdGet**
> []ShareFilesRespItems ShareFilesIdGet(ctx, id)
Share files

List share files 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share action ID | 

### Return type

[**[]ShareFilesRespItems**](ShareFilesRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareLoginPinPost**
> ShareLoginPinPost(ctx, body)
Login with pin on share

Login with pin on share protected with pin 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ShareLoginPinReq**](ShareLoginPinReq.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SharePreviewIdGet**
> SharePreviewIdGet(ctx, id, optional)
Preview share file

Preview share file 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share file ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| Share file ID | 
 **size** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareRecipientsGet**
> ShareRecipientsResp ShareRecipientsGet(ctx, optional)
Share recipinets

List share available recipients 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string**| Start of recipient email | 

### Return type

[**ShareRecipientsResp**](ShareRecipientsResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareRequestPost**
> ShareRequestResp ShareRequestPost(ctx, body)
Request share

Request share 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ShareRequestReq**](ShareRequestReq.md)|  | 

### Return type

[**ShareRequestResp**](ShareRequestResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareRevokeIdGet**
> IdResp ShareRevokeIdGet(ctx, id)
Revoke share

Revoke share 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share action ID | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShareSendRequestIdPost**
> JobResp ShareSendRequestIdPost(ctx, id, optional)
DEPRECATED! Use /share/request instead.

Send share request email 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Share Request ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| Share Request ID | 
 **body** | [**ShareSendRequestReq**](ShareSendRequestReq.md)|  | 

### Return type

[**JobResp**](JobResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackingGet**
> []TrackingRespItems TrackingGet(ctx, optional)
List share actions

List share actions 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **float32**| Share date timestamp | 

### Return type

[**[]TrackingRespItems**](TrackingRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackingIdGet**
> []TrackingIdRespItems TrackingIdGet(ctx, id, optional)
List share actions

List share actions 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| User id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| User id | 
 **date** | **float32**| Share date timestamp | 

### Return type

[**[]TrackingIdRespItems**](TrackingIdRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

