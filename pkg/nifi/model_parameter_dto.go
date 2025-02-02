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

type ParameterDto struct {
	// The name of the Parameter
	Name string `json:"name,omitempty"`
	// The description of the Parameter
	Description *string `json:"description,omitempty"`
	// Whether or not the Parameter is sensitive
	Sensitive bool `json:"sensitive,omitempty"`
	// The value of the Parameter
	Value *string `json:"value"`
	// The set of all components in the flow that are referencing this Parameter
	ReferencingComponents []AffectedComponentEntity `json:"referencingComponents,omitempty"`
}
