# \PGPApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactPgpKeyIdGet**](PGPApi.md#ContactPgpKeyIdGet) | **Get** /contact/pgp-key/{id} | Contact PGP key
[**KeyRequestMetadataIdGet**](PGPApi.md#KeyRequestMetadataIdGet) | **Get** /key-request/metadata/{id} | PGP key request metadata
[**KeyRequestRespondIdPost**](PGPApi.md#KeyRequestRespondIdPost) | **Post** /key-request/respond/{id} | Respond to PGP key request
[**PgpKeyCreatePost**](PGPApi.md#PgpKeyCreatePost) | **Post** /pgp-key/create | Create PGP key
[**PgpKeyDeleteIdGet**](PGPApi.md#PgpKeyDeleteIdGet) | **Get** /pgp-key/delete/{id} | Delete PGP key
[**PgpKeyEditIdPost**](PGPApi.md#PgpKeyEditIdPost) | **Post** /pgp-key/edit/{id} | Edit PGP key
[**PgpKeyMetadataIdGet**](PGPApi.md#PgpKeyMetadataIdGet) | **Get** /pgp-key/metadata/{id} | PGP key metadata
[**PgpKeyRecipientsPost**](PGPApi.md#PgpKeyRecipientsPost) | **Post** /pgp-key/recipients | Recipients PGP keys
[**PgpKeyRequestIdsGet**](PGPApi.md#PgpKeyRequestIdsGet) | **Get** /pgp-key/request/{ids[]} | Request PGP key
[**UserPgpKeyIdGet**](PGPApi.md#UserPgpKeyIdGet) | **Get** /user/pgp-key/{id} | User PGP key


# **ContactPgpKeyIdGet**
> PgpKeyResp ContactPgpKeyIdGet(ctx, id)
Contact PGP key

Get contact PGP key. Not used. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a contact | 

### Return type

[**PgpKeyResp**](PgpKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KeyRequestMetadataIdGet**
> KeyRequestMetadataResp KeyRequestMetadataIdGet(ctx, id)
PGP key request metadata

Get PGP key request metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of the PGP key request | 

### Return type

[**KeyRequestMetadataResp**](KeyRequestMetadataResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KeyRequestRespondIdPost**
> JobResp KeyRequestRespondIdPost(ctx, id, body)
Respond to PGP key request

Respond to PGP key request with PGP key 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a PGP key request | 
  **body** | [**KeyRequestRespondReq**](KeyRequestRespondReq.md)|  | 

### Return type

[**JobResp**](JobResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PgpKeyCreatePost**
> PgpKeyResp PgpKeyCreatePost(ctx, body)
Create PGP key

Create new PGP key 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**PgpCreateReq**](PgpCreateReq.md)|  | 

### Return type

[**PgpKeyResp**](PgpKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PgpKeyDeleteIdGet**
> IdResp PgpKeyDeleteIdGet(ctx, id)
Delete PGP key

Delete PGP key 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a PGP key | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PgpKeyEditIdPost**
> PgpKeyResp PgpKeyEditIdPost(ctx, id, optional)
Edit PGP key

Edit PGP key 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a PGP key | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| ID of a PGP key | 
 **body** | [**PgpEditReq**](PgpEditReq.md)|  | 

### Return type

[**PgpKeyResp**](PgpKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PgpKeyMetadataIdGet**
> PgpKeyResp PgpKeyMetadataIdGet(ctx, id)
PGP key metadata

Get PGP key metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a PGP key | 

### Return type

[**PgpKeyResp**](PgpKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PgpKeyRecipientsPost**
> []PgpKeyRecipientsRespItems PgpKeyRecipientsPost(ctx, body)
Recipients PGP keys

Get public PGP keys for given recipients 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdsReq**](IdsReq.md)|  | 

### Return type

[**[]PgpKeyRecipientsRespItems**](PgpKeyRecipientsRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PgpKeyRequestIdsGet**
> JobResp PgpKeyRequestIdsGet(ctx, ids)
Request PGP key

Send request PGP key email to user or contact 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ids** | [**string**](.md)| ID of contact | 

### Return type

[**JobResp**](JobResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPgpKeyIdGet**
> PgpKeyResp UserPgpKeyIdGet(ctx, id)
User PGP key

Get user PGP key 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a user | 

### Return type

[**PgpKeyResp**](PgpKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

