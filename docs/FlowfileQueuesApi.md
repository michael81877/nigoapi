# \FlowfileQueuesApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDropRequest**](FlowfileQueuesApi.md#CreateDropRequest) | **Post** /flowfile-queues/{id}/drop-requests | Creates a request to drop the contents of the queue in this connection.
[**CreateFlowFileListing**](FlowfileQueuesApi.md#CreateFlowFileListing) | **Post** /flowfile-queues/{id}/listing-requests | Lists the contents of the queue in this connection.
[**DeleteListingRequest**](FlowfileQueuesApi.md#DeleteListingRequest) | **Delete** /flowfile-queues/{id}/listing-requests/{listing-request-id} | Cancels and/or removes a request to list the contents of this connection.
[**DownloadFlowFileContent**](FlowfileQueuesApi.md#DownloadFlowFileContent) | **Get** /flowfile-queues/{id}/flowfiles/{flowfile-uuid}/content | Gets the content for a FlowFile in a Connection.
[**GetDropRequest**](FlowfileQueuesApi.md#GetDropRequest) | **Get** /flowfile-queues/{id}/drop-requests/{drop-request-id} | Gets the current status of a drop request for the specified connection.
[**GetFlowFile**](FlowfileQueuesApi.md#GetFlowFile) | **Get** /flowfile-queues/{id}/flowfiles/{flowfile-uuid} | Gets a FlowFile from a Connection.
[**GetListingRequest**](FlowfileQueuesApi.md#GetListingRequest) | **Get** /flowfile-queues/{id}/listing-requests/{listing-request-id} | Gets the current status of a listing request for the specified connection.
[**RemoveDropRequest**](FlowfileQueuesApi.md#RemoveDropRequest) | **Delete** /flowfile-queues/{id}/drop-requests/{drop-request-id} | Cancels and/or removes a request to drop the contents of this connection.


# **CreateDropRequest**
> DropRequestEntity CreateDropRequest(ctx, id)
Creates a request to drop the contents of the queue in this connection.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 

### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to README]](../pkg/nifi/README.md)

# **CreateFlowFileListing**
> ListingRequestEntity CreateFlowFileListing(ctx, id)
Lists the contents of the queue in this connection.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 

### Return type

[**ListingRequestEntity**](ListingRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to README]](../pkg/nifi/README.md)

# **DeleteListingRequest**
> ListingRequestEntity DeleteListingRequest(ctx, id, listingRequestId)
Cancels and/or removes a request to list the contents of this connection.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 
  **listingRequestId** | **string**| The listing request id. | 

### Return type

[**ListingRequestEntity**](ListingRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to README]](../pkg/nifi/README.md)

# **DownloadFlowFileContent**
> StreamingOutput DownloadFlowFileContent(ctx, id, flowfileUuid, optional)
Gets the content for a FlowFile in a Connection.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 
  **flowfileUuid** | **string**| The flowfile uuid. | 
 **optional** | ***FlowfileQueuesApiDownloadFlowFileContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowfileQueuesApiDownloadFlowFileContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **clusterNodeId** | **optional.String**| The id of the node where the content exists if clustered. | 

### Return type

[**StreamingOutput**](StreamingOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to README]](../pkg/nifi/README.md)

# **GetDropRequest**
> DropRequestEntity GetDropRequest(ctx, id, dropRequestId)
Gets the current status of a drop request for the specified connection.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 
  **dropRequestId** | **string**| The drop request id. | 

### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to README]](../pkg/nifi/README.md)

# **GetFlowFile**
> FlowFileEntity GetFlowFile(ctx, id, flowfileUuid, optional)
Gets a FlowFile from a Connection.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 
  **flowfileUuid** | **string**| The flowfile uuid. | 
 **optional** | ***FlowfileQueuesApiGetFlowFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowfileQueuesApiGetFlowFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterNodeId** | **optional.String**| The id of the node where the content exists if clustered. | 

### Return type

[**FlowFileEntity**](FlowFileEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to README]](../pkg/nifi/README.md)

# **GetListingRequest**
> ListingRequestEntity GetListingRequest(ctx, id, listingRequestId)
Gets the current status of a listing request for the specified connection.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 
  **listingRequestId** | **string**| The listing request id. | 

### Return type

[**ListingRequestEntity**](ListingRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to README]](../pkg/nifi/README.md)

# **RemoveDropRequest**
> DropRequestEntity RemoveDropRequest(ctx, id, dropRequestId)
Cancels and/or removes a request to drop the contents of this connection.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 
  **dropRequestId** | **string**| The drop request id. | 

### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to README]](../pkg/nifi/README.md)
