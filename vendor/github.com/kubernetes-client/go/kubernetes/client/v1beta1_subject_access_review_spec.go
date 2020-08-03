/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
type V1beta1SubjectAccessReviewSpec struct {

	// Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
	Extra map[string][]string `json:"extra,omitempty"`

	// Groups is the groups you're testing for.
	Group []string `json:"group,omitempty"`

	// NonResourceAttributes describes information for a non-resource access request
	NonResourceAttributes *V1beta1NonResourceAttributes `json:"nonResourceAttributes,omitempty"`

	// ResourceAuthorizationAttributes describes information for a resource access request
	ResourceAttributes *V1beta1ResourceAttributes `json:"resourceAttributes,omitempty"`

	// UID information about the requesting user.
	Uid string `json:"uid,omitempty"`

	// User is the user you're testing for. If you specify \"User\" but not \"Group\", then is it interpreted as \"What if User were not a member of any groups
	User string `json:"user,omitempty"`
}
