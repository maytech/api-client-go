# \FileApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileAddTagIdPost**](FileApi.md#FileAddTagIdPost) | **Post** /file/add-tag/{id} | Add file tag
[**FileCopyPost**](FileApi.md#FileCopyPost) | **Post** /file/copy | Copy files
[**FileCsvIdGet**](FileApi.md#FileCsvIdGet) | **Get** /file/csv/{id} | Download CSV file with Folder Content
[**FileDeletePost**](FileApi.md#FileDeletePost) | **Post** /file/delete | Delete files
[**FileDiffIdGet**](FileApi.md#FileDiffIdGet) | **Get** /file/diff/{id} | File diff
[**FileDownloadIdGet**](FileApi.md#FileDownloadIdGet) | **Get** /file/download/{id} | Download file
[**FileDownloadLinkPost**](FileApi.md#FileDownloadLinkPost) | **Post** /file/download-link | Get download link
[**FileInfoIdGet**](FileApi.md#FileInfoIdGet) | **Get** /file/info/{id} | Get file info
[**FileMakedirPost**](FileApi.md#FileMakedirPost) | **Post** /file/makedir | Create new folder
[**FileMetadataGet**](FileApi.md#FileMetadataGet) | **Get** /file/metadata | Get all file metadata
[**FileMetadataIdGet**](FileApi.md#FileMetadataIdGet) | **Get** /file/metadata/{id} | Get file metadata
[**FileMetadataPost**](FileApi.md#FileMetadataPost) | **Post** /file/metadata | Set file metadata
[**FileModifyPost**](FileApi.md#FileModifyPost) | **Post** /file/modify | Get file modify link
[**FileMovePost**](FileApi.md#FileMovePost) | **Post** /file/move | Move files
[**FilePreviewIdGet**](FileApi.md#FilePreviewIdGet) | **Get** /file/preview/{id} | File preview
[**FileRenameIdPost**](FileApi.md#FileRenameIdPost) | **Post** /file/rename/{id} | Rename file or folder
[**FileSearchPost**](FileApi.md#FileSearchPost) | **Post** /file/search | Search files
[**FileSizeIdGet**](FileApi.md#FileSizeIdGet) | **Get** /file/size/{id} | Get file size
[**FileTagsIdGet**](FileApi.md#FileTagsIdGet) | **Get** /file/tags/{id} | File tags
[**FileWopiTokenIdGet**](FileApi.md#FileWopiTokenIdGet) | **Get** /file/wopi-token/{id} | Get wopi token for file


# **FileAddTagIdPost**
> FileTagResp FileAddTagIdPost(ctx, id, body)
Add file tag

Add tag to file 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 
  **body** | [**FileAddTagReq**](FileAddTagReq.md)|  | 

### Return type

[**FileTagResp**](FileTagResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileCopyPost**
> JobResp FileCopyPost(ctx, body)
Copy files

Copy files 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**CopyMoveFilesReq**](CopyMoveFilesReq.md)|  | 

### Return type

[**JobResp**](JobResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileCsvIdGet**
> FileCsvResp FileCsvIdGet(ctx, id)
Download CSV file with Folder Content

Download CSV file content with Folder Content. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| File action ID | 

### Return type

[**FileCsvResp**](FileCsvResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileDeletePost**
> IdsResp FileDeletePost(ctx, body)
Delete files

Delete number of files 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdsReq**](IdsReq.md)|  | 

### Return type

[**IdsResp**](IdsResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileDiffIdGet**
> FileDiffResp FileDiffIdGet(ctx, id, from, optional)
File diff

Get file changes diff 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 
  **from** | **float32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| ID of a file | 
 **from** | **float32**|  | 
 **to** | **float32**|  | 

### Return type

[**FileDiffResp**](FileDiffResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileDownloadIdGet**
> FileDownloadIdGet(ctx, id)
Download file

Download file content 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileDownloadLinkPost**
> IdResp FileDownloadLinkPost(ctx, body)
Get download link

Get files download link 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdsReq**](IdsReq.md)|  | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileInfoIdGet**
> FileInfoResp FileInfoIdGet(ctx, id)
Get file info

Get file info 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 

### Return type

[**FileInfoResp**](FileInfoResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileMakedirPost**
> FileResp FileMakedirPost(ctx, body)
Create new folder

Create new folder 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**MakeDirReq**](MakeDirReq.md)|  | 

### Return type

[**FileResp**](FileResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileMetadataGet**
> FileMetadataGetResp FileMetadataGet(ctx, optional)
Get all file metadata

Get all file metadata in current user home folder 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | **float32**| Return content flag | [default to 1]

### Return type

[**FileMetadataGetResp**](FileMetadataGetResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileMetadataIdGet**
> FileMetadataGetResp FileMetadataIdGet(ctx, id, optional)
Get file metadata

Get file metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| ID of a file | 
 **content** | **float32**| Return content flag | [default to 1]

### Return type

[**FileMetadataGetResp**](FileMetadataGetResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileMetadataPost**
> FileMetadataPostResp FileMetadataPost(ctx, body)
Set file metadata

Set file metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FileMetadataPostReq**](FileMetadataPostReq.md)|  | 

### Return type

[**FileMetadataPostResp**](FileMetadataPostResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileModifyPost**
> FileModifyResp FileModifyPost(ctx, body)
Get file modify link

Get file modify link 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FileModifyReq**](FileModifyReq.md)|  | 

### Return type

[**FileModifyResp**](FileModifyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileMovePost**
> IdsResp FileMovePost(ctx, body)
Move files

Move files 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**CopyMoveFilesReq**](CopyMoveFilesReq.md)|  | 

### Return type

[**IdsResp**](IdsResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilePreviewIdGet**
> FilePreviewResp FilePreviewIdGet(ctx, id)
File preview

Get file preview metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 

### Return type

[**FilePreviewResp**](FilePreviewResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileRenameIdPost**
> FileRenameResp FileRenameIdPost(ctx, id, body)
Rename file or folder

Rename file or folder 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 
  **body** | [**FileRenameReq**](FileRenameReq.md)|  | 

### Return type

[**FileRenameResp**](FileRenameResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileSearchPost**
> []FileResp FileSearchPost(ctx, body)
Search files

Search files 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**SearchReq**](SearchReq.md)|  | 

### Return type

[**[]FileResp**](FileResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileSizeIdGet**
> FileSizeResp FileSizeIdGet(ctx, id)
Get file size

Get file size 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 

### Return type

[**FileSizeResp**](FileSizeResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileTagsIdGet**
> []FileTagResp FileTagsIdGet(ctx, id)
File tags

List file tags 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 

### Return type

[**[]FileTagResp**](FileTagResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileWopiTokenIdGet**
> IdResp FileWopiTokenIdGet(ctx, id)
Get wopi token for file

Get wopi token for file 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

