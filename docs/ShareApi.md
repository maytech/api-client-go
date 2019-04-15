# \ShareApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesReturnMakedirIdPost**](ShareApi.md#FilesReturnMakedirIdPost) | **Post** /files-return/makedir/{id} | Create a directory for returned files
[**FilesReturnMetadataIdGet**](ShareApi.md#FilesReturnMetadataIdGet) | **Get** /files-return/metadata/{id} | Get return files metadata
[**FilesReturnSendPost**](ShareApi.md#FilesReturnSendPost) | **Post** /files-return/send | Return files in the created share
[**FilesReturnUploadLinkIdPost**](ShareApi.md#FilesReturnUploadLinkIdPost) | **Post** /files-return/upload-link/{id} | Get return files upload link
[**QuicklinkCreatePost**](ShareApi.md#QuicklinkCreatePost) | **Post** /quicklink/create | Create a quicklink
[**QuicklinkLoginPinPost**](ShareApi.md#QuicklinkLoginPinPost) | **Post** /quicklink/login-pin | Log in with PIN to access a quicklink
[**QuicklinkRevokeIdGet**](ShareApi.md#QuicklinkRevokeIdGet) | **Get** /quicklink/revoke/{id} | Revoke a quicklink
[**ShareCreatePost**](ShareApi.md#ShareCreatePost) | **Post** /share/create | Create a file share
[**ShareDownloadIdGet**](ShareApi.md#ShareDownloadIdGet) | **Get** /share/download/{id} | Download share files
[**ShareDownloadInfoIdGet**](ShareApi.md#ShareDownloadInfoIdGet) | **Get** /share/download-info/{id} | Get share download info
[**ShareDownloadLinkIdGet**](ShareApi.md#ShareDownloadLinkIdGet) | **Get** /share/download-link/{id} | Get download link for all files
[**ShareDownloadLinkIdPost**](ShareApi.md#ShareDownloadLinkIdPost) | **Post** /share/download-link/{id} | Get download link for specified files
[**ShareFilesIdGet**](ShareApi.md#ShareFilesIdGet) | **Get** /share/files/{id} | List shared files
[**ShareLoginPinPost**](ShareApi.md#ShareLoginPinPost) | **Post** /share/login-pin | Log in with PIN to access a share
[**SharePreviewIdGet**](ShareApi.md#SharePreviewIdGet) | **Get** /share/preview/{id} | Preview a shared file
[**ShareRecipientsGet**](ShareApi.md#ShareRecipientsGet) | **Get** /share/recipients | List all contacts for the share
[**ShareRequestPost**](ShareApi.md#ShareRequestPost) | **Post** /share/request | Send a request to share files
[**ShareRevokeIdGet**](ShareApi.md#ShareRevokeIdGet) | **Get** /share/revoke/{id} | Revoke a share
[**ShareSendRequestIdPost**](ShareApi.md#ShareSendRequestIdPost) | **Post** /share/send-request/{id} | Request files. Use /share/request API call instead.
[**TrackingGet**](ShareApi.md#TrackingGet) | **Get** /tracking/ | List share actions metadata for all users
[**TrackingIdGet**](ShareApi.md#TrackingIdGet) | **Get** /tracking/{id} | List share actions metadata for a user


# **FilesReturnMakedirIdPost**
> FilesReturnMakedirResps FilesReturnMakedirIdPost(ctx, id, body)
Create a directory for returned files

Add a directory for retuned files in the sender’s file tree. In order to return a file tree structure, you should duplicate it on the server side, that is send this API call to create each folder separately. This call returns the ID response that can be used as the parent ID for creating a new folder in the already created folder. 

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

Get the metadata of returning files. 

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
> FilesReturnSendResp FilesReturnSendPost(ctx, body)
Return files in the created share

Return files to the share sender. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FilesReturnSendReq**](FilesReturnSendReq.md)|  | 

### Return type

[**FilesReturnSendResp**](FilesReturnSendResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesReturnUploadLinkIdPost**
> FilesReturnUploadLinkResp FilesReturnUploadLinkIdPost(ctx, id, body)
Get return files upload link

Get the upload link to return files to the share sender. 

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
Create a quicklink

Create a quicklink with the usage behaviour the same as a public share. 

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
Log in with PIN to access a quicklink

Log in using a previously generated PIN that was sent while creating the quicklink to have access to it. 

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
Revoke a quicklink

Disable access to a quicklink. 

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
Create a file share

Create a share with files. File share types:           - P (public) - Anyone can download. Limited tracking - IP address only.           - T (tracked) - Any registered user/recipient can download. Full tracking.           - C (restricted) - Only the registered email recipient(s) can download. Full tracking. 

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
> ShareDownloadIdGet(ctx, id)
Download share files

Download files requested in the share/download-link call. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Download link ID | 

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
Get share download info

Get the share metadata by the given share action iD. File share types:           - P (public) - Anyone can download. Limited tracking - IP address only.           - T (tracked) - Any registered user/recipient can download. Full tracking.           - C (restricted) - Only the registered email recipient(s) can download. Full tracking. 

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
Get download link for all files

Retrieve a link for downloading all files. 

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
Get download link for specified files

Retrieve a link for downloading specified files. 

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
List shared files

Retrieve a list of shared files. 

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
Log in with PIN to access a share

Log in using a previously generated PIN that was sent while creating the share to have access to it. 

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
Preview a shared file

Preview a shared file by given share file ID. 

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
List all contacts for the share

Retrieve a list of all recipients that can receive the share. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string**| Search emails by the specified beginning. | 

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
Send a request to share files

Send a file sharing request. 

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
Revoke a share

Revoke access to a share 

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
Request files. Use /share/request API call instead.

Send a file sharing request. 

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
List share actions metadata for all users

Retrieve the metadata of created shares and quicklinks of all users that can be managed by a logged-in user. 

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
List share actions metadata for a user

Retrieve the metadata of created shares and quicklinks of a user by given user ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| User ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| User ID | 
 **date** | **float32**| Share date timestamp | 

### Return type

[**[]TrackingIdRespItems**](TrackingIdRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

