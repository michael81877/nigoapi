/*
 * Apache NiFi Registry REST API
 *
 * The REST API provides an interface to a registry with operations for saving, versioning, reading NiFi flows and components.
 *
 * API version: 0.8.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package registry

// The revision information for an entity managed through the REST API.
type RevisionInfo struct {
	// A client identifier used to make a request. By including a client identifier, the API can allow multiple requests without needing the current revision. Due to the asynchronous nature of requests/responses this was implemented to allow the client to make numerous requests without having to wait for the previous response to come back.
	ClientId string `json:"clientId,omitempty"`
	// NiFi Registry employs an optimistic locking strategy where the client must include a revision in their request when performing an update. In a response to a mutable flow request, this field represents the updated base version.
	Version int64 `json:"version,omitempty"`
	// The user that last modified the entity.
	LastModifier string `json:"lastModifier,omitempty"`
}