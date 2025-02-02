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

type PortDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string `json:"parentGroupId,omitempty"`
	// The position of this component in the UI if applicable.
	Position *PositionDto `json:"position,omitempty"`
	// The name of the port.
	Name string `json:"name,omitempty"`
	// The comments for the port.
	Comments string `json:"comments,omitempty"`
	// The state of the port.
	State string `json:"state,omitempty"`
	// The type of port.
	Type_ string `json:"type,omitempty"`
	// Whether the port has incoming or output connections to a remote NiFi. This is only applicable when the port is allowed to be accessed remotely.
	Transmitting bool `json:"transmitting,omitempty"`
	// The number of tasks that should be concurrently scheduled for the port.
	ConcurrentlySchedulableTaskCount int32 `json:"concurrentlySchedulableTaskCount,omitempty"`
	// The users that are allowed to access the port.
	UserAccessControl []string `json:"userAccessControl,omitempty"`
	// The user groups that are allowed to access the port.
	GroupAccessControl []string `json:"groupAccessControl,omitempty"`
	// Whether this port can be accessed remotely via Site-to-Site protocol.
	AllowRemoteAccess bool `json:"allowRemoteAccess,omitempty"`
	// Gets the validation errors from this port. These validation errors represent the problems with the port that must be resolved before it can be started.
	ValidationErrors []string `json:"validationErrors,omitempty"`
}
