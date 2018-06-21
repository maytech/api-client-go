# \ContactApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactCreatePost**](ContactApi.md#ContactCreatePost) | **Post** /contact/create | Create contact
[**ContactDeletePost**](ContactApi.md#ContactDeletePost) | **Post** /contact/delete | Delete contact
[**ContactEditIdPost**](ContactApi.md#ContactEditIdPost) | **Post** /contact/edit/{id} | Contact metadata
[**ContactGet**](ContactApi.md#ContactGet) | **Get** /contact | List user contacts
[**ContactGroupGet**](ContactApi.md#ContactGroupGet) | **Get** /contact/group | All contact groups
[**ContactMetadataIdGet**](ContactApi.md#ContactMetadataIdGet) | **Get** /contact/metadata/{id} | Contact metadata
[**ContactPgpKeyIdGet**](ContactApi.md#ContactPgpKeyIdGet) | **Get** /contact/pgp-key/{id} | Contact PGP key
[**ContactUpgradeIdGet**](ContactApi.md#ContactUpgradeIdGet) | **Get** /contact/upgrade/{id} | Upgrade contact


# **ContactCreatePost**
> ContactResp ContactCreatePost(ctx, body)
Create contact

Create new contact 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ContactCreateReq**](ContactCreateReq.md)|  | 

### Return type

[**ContactResp**](ContactResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactDeletePost**
> []ContactDeleteRespItems ContactDeletePost(ctx, body)
Delete contact

Delete contact 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdsReq**](IdsReq.md)| IDs of a contacts | 

### Return type

[**[]ContactDeleteRespItems**](ContactDeleteRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactEditIdPost**
> ContactResp ContactEditIdPost(ctx, id, optional)
Contact metadata

Get contact metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a contact | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| ID of a contact | 
 **body** | [**ContactEditResp**](ContactEditResp.md)|  | 

### Return type

[**ContactResp**](ContactResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactGet**
> []ContactResp ContactGet(ctx, )
List user contacts

List user contacts 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ContactResp**](ContactResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactGroupGet**
> []ContactGroupRespItems ContactGroupGet(ctx, )
All contact groups

List all contact groups 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ContactGroupRespItems**](ContactGroupRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactMetadataIdGet**
> ContactResp ContactMetadataIdGet(ctx, id)
Contact metadata

Get contact metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a contact | 

### Return type

[**ContactResp**](ContactResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **ContactUpgradeIdGet**
> UserResp ContactUpgradeIdGet(ctx, id)
Upgrade contact

Upgrade contact to user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a contact | 

### Return type

[**UserResp**](UserResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

