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

type ControllerServiceDefinition struct {
	// The class name of the service API
	ClassName string `json:"className,omitempty"`
	// The group id of the service API
	GroupId string `json:"groupId,omitempty"`
	// The artifact id of the service API
	ArtifactId string `json:"artifactId,omitempty"`
	// The version of the service API
	Version string `json:"version,omitempty"`
}
