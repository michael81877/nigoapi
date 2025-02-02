/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type VersionedFlow struct {
	// An WebLink to this entity.
	Link *JaxbLink `json:"link,omitempty"`
	// An ID to uniquely identify this object.
	Identifier string `json:"identifier,omitempty"`
	// The name of the item.
	Name string `json:"name"`
	// A description of the item.
	Description string `json:"description,omitempty"`
	// The identifier of the bucket this items belongs to. This cannot be changed after the item is created.
	BucketIdentifier string `json:"bucketIdentifier"`
	// The name of the bucket this items belongs to.
	BucketName string `json:"bucketName,omitempty"`
	// The timestamp of when the item was created, as milliseconds since epoch.
	CreatedTimestamp int64 `json:"createdTimestamp,omitempty"`
	// The timestamp of when the item was last modified, as milliseconds since epoch.
	ModifiedTimestamp int64 `json:"modifiedTimestamp,omitempty"`
	// The type of item.
	Type_ string `json:"type"`
	// The access that the current user has to the bucket containing this item.
	Permissions *Permissions `json:"permissions,omitempty"`
	// The number of versions of this flow.
	VersionCount int64 `json:"versionCount,omitempty"`
	// The revision of this entity used for optimistic-locking during updates.
	Revision *RevisionInfo `json:"revision,omitempty"`
}
