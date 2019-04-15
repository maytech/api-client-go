# \FileApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileAddTagIdPost**](FileApi.md#FileAddTagIdPost) | **Post** /file/add-tag/{id} | Add a file tag
[**FileCopyPost**](FileApi.md#FileCopyPost) | **Post** /file/copy | Copy files
[**FileCsvIdGet**](FileApi.md#FileCsvIdGet) | **Get** /file/csv/{id} | Download CSV file with Folder Content
[**FileDeletePost**](FileApi.md#FileDeletePost) | **Post** /file/delete | Delete files
[**FileDiffIdGet**](FileApi.md#FileDiffIdGet) | **Get** /file/diff/{id} | Display changes of the file
[**FileDownloadIdGet**](FileApi.md#FileDownloadIdGet) | **Get** /file/download/{id} | Download file
[**FileDownloadLinkPost**](FileApi.md#FileDownloadLinkPost) | **Post** /file/download-link | Get download link
[**FileInfoIdGet**](FileApi.md#FileInfoIdGet) | **Get** /file/info/{id} | Get file info
[**FileMakedirPost**](FileApi.md#FileMakedirPost) | **Post** /file/makedir | Create a new folder
[**FileMetadataGet**](FileApi.md#FileMetadataGet) | **Get** /file/metadata | Get metadata of files
[**FileMetadataIdGet**](FileApi.md#FileMetadataIdGet) | **Get** /file/metadata/{id} | Get all file metadata
[**FileMetadataPost**](FileApi.md#FileMetadataPost) | **Post** /file/metadata | Modify file metadata
[**FileModifyPost**](FileApi.md#FileModifyPost) | **Post** /file/modify | Get file modification link
[**FileMovePost**](FileApi.md#FileMovePost) | **Post** /file/move | Move files
[**FilePreviewIdGet**](FileApi.md#FilePreviewIdGet) | **Get** /file/preview/{id} | Get a file preview
[**FileRenameIdPost**](FileApi.md#FileRenameIdPost) | **Post** /file/rename/{id} | Rename a file
[**FileSearchPost**](FileApi.md#FileSearchPost) | **Post** /file/search | Search files
[**FileSizeIdGet**](FileApi.md#FileSizeIdGet) | **Get** /file/size/{id} | Get file size
[**FileTagsIdGet**](FileApi.md#FileTagsIdGet) | **Get** /file/tags/{id} | Get a list of file tags


# **FileAddTagIdPost**
> FileTagResp FileAddTagIdPost(ctx, id, body)
Add a file tag

Add a tag to filter the file list. 

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

Creates a copy of a file or a folder. The original version of the file will not be changed. On success 202 response it returns “job_id”. To check the result, see the API call “job/status\". 

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

Return CSV file with information containing file metadata. It contains the path for each file displaying the hierarchy of files. This API returns the content-type: ”raw\". 

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

Move a file or folder to Trash. If you delete less than 10 files, it returns 200 and the IDs of the deleted files. In case you delete more than 10 files, the API returns 202 response. 

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
Display changes of the file

Get file changes for specified preiod. 

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

Download a file or files by given ID. Multiple file download returns files in Zip format. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file download link | 

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

Get a unique key for downloading files. 

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

Retrieve the file information by the specified ID. 

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
Create a new folder

Create a new folder 

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
Get metadata of files

Get information about files in the current User Home folder. 

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
Get all file metadata

Retrieve all metadata associated with a given file. If ID endpoint is not provided, get the metadata of a current Use Home folder. 

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
Modify file metadata

Update file metadata with the given payload. 

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
Get file modification link

Get file modification link 

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

Move a file or folder from one location to another. 

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
Get a file preview

Retrieve a file preview by the given ID of the file. 

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
Rename a file

Change a file name. 

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

Search any file by the given directory. 

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

Get details about the file size by the given ID. 

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
Get a list of file tags

Retrieve a list of available file tags. 

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

