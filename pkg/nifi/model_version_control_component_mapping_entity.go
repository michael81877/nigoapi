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

type VersionControlComponentMappingEntity struct {
	// The mapping of Versioned Component Identifiers to instance ID's
	VersionControlComponentMapping map[string]string `json:"versionControlComponentMapping,omitempty"`
	// The revision of the Process Group
	ProcessGroupRevision *RevisionDto `json:"processGroupRevision,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
	// The Version Control information
	VersionControlInformation *VersionControlInformationDto `json:"versionControlInformation,omitempty"`
}
