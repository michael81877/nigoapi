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

type CurrentUserEntity struct {
	// The user identity being serialized.
	Identity string `json:"identity,omitempty"`
	// Whether the current user is anonymous.
	Anonymous bool `json:"anonymous,omitempty"`
	// Permissions for querying provenance.
	ProvenancePermissions *PermissionsDto `json:"provenancePermissions,omitempty"`
	// Permissions for accessing counters.
	CountersPermissions *PermissionsDto `json:"countersPermissions,omitempty"`
	// Permissions for accessing tenants.
	TenantsPermissions *PermissionsDto `json:"tenantsPermissions,omitempty"`
	// Permissions for accessing the controller.
	ControllerPermissions *PermissionsDto `json:"controllerPermissions,omitempty"`
	// Permissions for accessing the policies.
	PoliciesPermissions *PermissionsDto `json:"policiesPermissions,omitempty"`
	// Permissions for accessing system.
	SystemPermissions *PermissionsDto `json:"systemPermissions,omitempty"`
	// Permissions for accessing parameter contexts.
	ParameterContextPermissions *PermissionsDto `json:"parameterContextPermissions,omitempty"`
	// Permissions for accessing restricted components. Note: the read permission are not used and will always be false.
	RestrictedComponentsPermissions *PermissionsDto `json:"restrictedComponentsPermissions,omitempty"`
	// Permissions for specific component restrictions.
	ComponentRestrictionPermissions []ComponentRestrictionPermissionDto `json:"componentRestrictionPermissions,omitempty"`
	// Whether the current user can version flows.
	CanVersionFlows bool `json:"canVersionFlows,omitempty"`
}
