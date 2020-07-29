/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// UserInfo holds the information about the user needed to implement the user.Info interface.
type V1beta1UserInfo struct {

	// Any additional information provided by the authenticator.
	Extra map[string][]string `json:"extra,omitempty"`

	// The names of groups this user is a part of.
	Groups []string `json:"groups,omitempty"`

	// A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
	Uid string `json:"uid,omitempty"`

	// The name that uniquely identifies this user among all active users.
	Username string `json:"username,omitempty"`
}
