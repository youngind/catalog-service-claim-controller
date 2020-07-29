/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// ScaleStatus represents the current status of a scale subresource.
type AppsV1beta1ScaleStatus struct {

	// actual number of observed instances of the scaled object.
	Replicas int32 `json:"replicas"`

	// label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	Selector map[string]string `json:"selector,omitempty"`

	// label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	TargetSelector string `json:"targetSelector,omitempty"`
}
