# ConnectionEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | [***RevisionDto**](RevisionDTO.md) | The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses. | [optional] [default to null]
**Id** | **string** | The id of the component. | [optional] [default to null]
**Uri** | **string** | The URI for futures requests to the component. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) | The position of this component in the UI if applicable. | [optional] [default to null]
**Permissions** | [***PermissionsDto**](PermissionsDTO.md) | The permissions for this component. | [optional] [default to null]
**Bulletins** | [**[]BulletinEntity**](BulletinEntity.md) | The bulletins for this component. | [optional] [default to null]
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]
**Component** | [***ConnectionDto**](ConnectionDTO.md) |  | [optional] [default to null]
**Status** | [***ConnectionStatusDto**](ConnectionStatusDTO.md) | The status of the connection. | [optional] [default to null]
**Bends** | [**[]PositionDto**](PositionDTO.md) | The bend points on the connection. | [optional] [default to null]
**LabelIndex** | **int32** | The index of the bend point where to place the connection label. | [optional] [default to null]
**GetzIndex** | **int64** | The z index of the connection. | [optional] [default to null]
**SourceId** | **string** | The identifier of the source of this connection. | [optional] [default to null]
**SourceGroupId** | **string** | The identifier of the group of the source of this connection. | [optional] [default to null]
**SourceType** | **string** | The type of component the source connectable is. | [default to null]
**DestinationId** | **string** | The identifier of the destination of this connection. | [optional] [default to null]
**DestinationGroupId** | **string** | The identifier of the group of the destination of this connection. | [optional] [default to null]
**DestinationType** | **string** | The type of component the destination connectable is. | [default to null]

[[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to README]](../pkg/nifi/README.md)

