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

type ComponentDifferenceGroup struct {
	// The id of the component whose changes are grouped together.
	ComponentId string `json:"componentId,omitempty"`
	// The name of the component whose changes are grouped together.
	ComponentName string `json:"componentName,omitempty"`
	// The type of component these changes relate to.
	ComponentType string `json:"componentType,omitempty"`
	// The process group id for this component.
	ProcessGroupId string `json:"processGroupId,omitempty"`
	// The list of changes related to this component between the 2 versions.
	Differences []ComponentDifference `json:"differences,omitempty"`
}