//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by prerelease-lifecycle-gen. DO NOT EDIT.

package v1alpha3

import (
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *DeviceClass) APILifecycleIntroduced() (major, minor int) {
	return 1, 31
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *DeviceClass) APILifecycleDeprecated() (major, minor int) {
	return 1, 34
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *DeviceClass) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1beta1", Kind: "DeviceClass"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *DeviceClass) APILifecycleRemoved() (major, minor int) {
	return 1, 37
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *DeviceClassList) APILifecycleIntroduced() (major, minor int) {
	return 1, 31
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *DeviceClassList) APILifecycleDeprecated() (major, minor int) {
	return 1, 34
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *DeviceClassList) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1beta1", Kind: "DeviceClassList"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *DeviceClassList) APILifecycleRemoved() (major, minor int) {
	return 1, 37
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *ResourceClaim) APILifecycleIntroduced() (major, minor int) {
	return 1, 31
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *ResourceClaim) APILifecycleDeprecated() (major, minor int) {
	return 1, 34
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *ResourceClaim) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1beta1", Kind: "ResourceClaim"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *ResourceClaim) APILifecycleRemoved() (major, minor int) {
	return 1, 37
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *ResourceClaimList) APILifecycleIntroduced() (major, minor int) {
	return 1, 31
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *ResourceClaimList) APILifecycleDeprecated() (major, minor int) {
	return 1, 34
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *ResourceClaimList) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1beta1", Kind: "ResourceClaimList"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *ResourceClaimList) APILifecycleRemoved() (major, minor int) {
	return 1, 37
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *ResourceClaimTemplate) APILifecycleIntroduced() (major, minor int) {
	return 1, 31
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *ResourceClaimTemplate) APILifecycleDeprecated() (major, minor int) {
	return 1, 34
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *ResourceClaimTemplate) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1beta1", Kind: "ResourceClaimTemplate"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *ResourceClaimTemplate) APILifecycleRemoved() (major, minor int) {
	return 1, 37
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *ResourceClaimTemplateList) APILifecycleIntroduced() (major, minor int) {
	return 1, 31
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *ResourceClaimTemplateList) APILifecycleDeprecated() (major, minor int) {
	return 1, 34
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *ResourceClaimTemplateList) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1beta1", Kind: "ResourceClaimTemplateList"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *ResourceClaimTemplateList) APILifecycleRemoved() (major, minor int) {
	return 1, 37
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *ResourceSlice) APILifecycleIntroduced() (major, minor int) {
	return 1, 31
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *ResourceSlice) APILifecycleDeprecated() (major, minor int) {
	return 1, 34
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *ResourceSlice) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1beta1", Kind: "ResourceSlice"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *ResourceSlice) APILifecycleRemoved() (major, minor int) {
	return 1, 37
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *ResourceSliceList) APILifecycleIntroduced() (major, minor int) {
	return 1, 31
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *ResourceSliceList) APILifecycleDeprecated() (major, minor int) {
	return 1, 34
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *ResourceSliceList) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1beta1", Kind: "ResourceSliceList"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *ResourceSliceList) APILifecycleRemoved() (major, minor int) {
	return 1, 37
}
