/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.9
// +k8s:prerelease-lifecycle-gen:deprecated=1.21
// +k8s:prerelease-lifecycle-gen:replacement=storage.k8s.io,v1,VolumeAttachment

// VolumeAttachment captures the intent to attach or detach the specified volume
// to/from the specified node.
//
// VolumeAttachment objects are non-namespaced.
type VolumeAttachment struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// spec represents specification of the desired attach/detach volume behavior.
	// Populated by the Kubernetes system.
	Spec VolumeAttachmentSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	// status represents status of the VolumeAttachment request.
	// Populated by the entity completing the attach or detach
	// operation, i.e. the external-attacher.
	// +optional
	Status VolumeAttachmentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.9
// +k8s:prerelease-lifecycle-gen:deprecated=1.21
// +k8s:prerelease-lifecycle-gen:replacement=storage.k8s.io,v1,VolumeAttachmentList

// VolumeAttachmentList is a collection of VolumeAttachment objects.
type VolumeAttachmentList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of VolumeAttachments
	Items []VolumeAttachment `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpec struct {
	// attacher indicates the name of the volume driver that MUST handle this
	// request. This is the name returned by GetPluginName().
	Attacher string `json:"attacher" protobuf:"bytes,1,opt,name=attacher"`

	// source represents the volume that should be attached.
	Source VolumeAttachmentSource `json:"source" protobuf:"bytes,2,opt,name=source"`

	// nodeName represents the node that the volume should be attached to.
	NodeName string `json:"nodeName" protobuf:"bytes,3,opt,name=nodeName"`
}

// VolumeAttachmentSource represents a volume that should be attached.
// Right now only PersistentVolumes can be attached via external attacher,
// in the future we may allow also inline volumes in pods.
// Exactly one member can be set.
type VolumeAttachmentSource struct {
	// persistentVolumeName represents the name of the persistent volume to attach.
	// +optional
	PersistentVolumeName *string `json:"persistentVolumeName,omitempty" protobuf:"bytes,1,opt,name=persistentVolumeName"`

	// inlineVolumeSpec contains all the information necessary to attach
	// a persistent volume defined by a pod's inline VolumeSource. This field
	// is populated only for the CSIMigration feature. It contains
	// translated fields from a pod's inline VolumeSource to a
	// PersistentVolumeSpec. This field is alpha-level and is only
	// honored by servers that enabled the CSIMigration feature.
	// +optional
	InlineVolumeSpec *v1.PersistentVolumeSpec `json:"inlineVolumeSpec,omitempty" protobuf:"bytes,2,opt,name=inlineVolumeSpec"`
}

// VolumeAttachmentStatus is the status of a VolumeAttachment request.
type VolumeAttachmentStatus struct {
	// attached indicates the volume is successfully attached.
	// This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	Attached bool `json:"attached" protobuf:"varint,1,opt,name=attached"`

	// attachmentMetadata is populated with any
	// information returned by the attach operation, upon successful attach, that must be passed
	// into subsequent WaitForAttach or Mount calls.
	// This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	// +optional
	AttachmentMetadata map[string]string `json:"attachmentMetadata,omitempty" protobuf:"bytes,2,rep,name=attachmentMetadata"`

	// attachError represents the last error encountered during attach operation, if any.
	// This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	// +optional
	AttachError *VolumeError `json:"attachError,omitempty" protobuf:"bytes,3,opt,name=attachError,casttype=VolumeError"`

	// detachError represents the last error encountered during detach operation, if any.
	// This field must only be set by the entity completing the detach
	// operation, i.e. the external-attacher.
	// +optional
	DetachError *VolumeError `json:"detachError,omitempty" protobuf:"bytes,4,opt,name=detachError,casttype=VolumeError"`
}

// VolumeError captures an error encountered during a volume operation.
type VolumeError struct {
	// time represents the time the error was encountered.
	// +optional
	Time metav1.Time `json:"time,omitempty" protobuf:"bytes,1,opt,name=time"`

	// message represents the error encountered during Attach or Detach operation.
	// This string maybe logged, so it should not contain sensitive
	// information.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`

	// errorCode is a numeric gRPC code representing the error encountered during Attach or Detach operations.
	//
	// This is an optional, alpha field that requires the MutableCSINodeAllocatableCount feature gate being enabled to be set.
	//
	// +featureGate=MutableCSINodeAllocatableCount
	// +optional
	ErrorCode *int32 `json:"errorCode,omitempty" protobuf:"varint,3,opt,name=errorCode"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.19
// +k8s:prerelease-lifecycle-gen:deprecated=1.21
// +k8s:prerelease-lifecycle-gen:replacement=storage.k8s.io,v1beta1,CSIStorageCapacity

// CSIStorageCapacity stores the result of one CSI GetCapacity call.
// For a given StorageClass, this describes the available capacity in a
// particular topology segment.  This can be used when considering where to
// instantiate new PersistentVolumes.
//
// For example this can express things like:
// - StorageClass "standard" has "1234 GiB" available in "topology.kubernetes.io/zone=us-east1"
// - StorageClass "localssd" has "10 GiB" available in "kubernetes.io/hostname=knode-abc123"
//
// The following three cases all imply that no capacity is available for
// a certain combination:
// - no object exists with suitable topology and storage class name
// - such an object exists, but the capacity is unset
// - such an object exists, but the capacity is zero
//
// The producer of these objects can decide which approach is more suitable.
//
// They are consumed by the kube-scheduler when a CSI driver opts into
// capacity-aware scheduling with CSIDriverSpec.StorageCapacity. The scheduler
// compares the MaximumVolumeSize against the requested size of pending volumes
// to filter out unsuitable nodes. If MaximumVolumeSize is unset, it falls back
// to a comparison against the less precise Capacity. If that is also unset,
// the scheduler assumes that capacity is insufficient and tries some other
// node.
type CSIStorageCapacity struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata. The name has no particular meaning. It must be
	// be a DNS subdomain (dots allowed, 253 characters). To ensure that
	// there are no conflicts with other CSI drivers on the cluster, the recommendation
	// is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends
	// with the unique CSI driver name.
	//
	// Objects are namespaced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// nodeTopology defines which nodes have access to the storage
	// for which capacity was reported. If not set, the storage is
	// not accessible from any node in the cluster. If empty, the
	// storage is accessible from all nodes. This field is
	// immutable.
	//
	// +optional
	NodeTopology *metav1.LabelSelector `json:"nodeTopology,omitempty" protobuf:"bytes,2,opt,name=nodeTopology"`

	// storageClassName represents the name of the StorageClass that the reported capacity applies to.
	// It must meet the same requirements as the name of a StorageClass
	// object (non-empty, DNS subdomain). If that object no longer exists,
	// the CSIStorageCapacity object is obsolete and should be removed by its
	// creator.
	// This field is immutable.
	StorageClassName string `json:"storageClassName" protobuf:"bytes,3,name=storageClassName"`

	// capacity is the value reported by the CSI driver in its GetCapacityResponse
	// for a GetCapacityRequest with topology and parameters that match the
	// previous fields.
	//
	// The semantic is currently (CSI spec 1.2) defined as:
	// The available capacity, in bytes, of the storage that can be used
	// to provision volumes. If not set, that information is currently
	// unavailable.
	//
	// +optional
	Capacity *resource.Quantity `json:"capacity,omitempty" protobuf:"bytes,4,opt,name=capacity"`

	// maximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse
	// for a GetCapacityRequest with topology and parameters that match the
	// previous fields.
	//
	// This is defined since CSI spec 1.4.0 as the largest size
	// that may be used in a
	// CreateVolumeRequest.capacity_range.required_bytes field to
	// create a volume with the same parameters as those in
	// GetCapacityRequest. The corresponding value in the Kubernetes
	// API is ResourceRequirements.Requests in a volume claim.
	//
	// +optional
	MaximumVolumeSize *resource.Quantity `json:"maximumVolumeSize,omitempty" protobuf:"bytes,5,opt,name=maximumVolumeSize"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.19
// +k8s:prerelease-lifecycle-gen:deprecated=1.21
// +k8s:prerelease-lifecycle-gen:replacement=storage.k8s.io,v1beta1,CSIStorageCapacityList

// CSIStorageCapacityList is a collection of CSIStorageCapacity objects.
type CSIStorageCapacityList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of CSIStorageCapacity objects.
	Items []CSIStorageCapacity `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:prerelease-lifecycle-gen:introduced=1.29
// +k8s:prerelease-lifecycle-gen:deprecated=1.32
// +k8s:prerelease-lifecycle-gen:replacement=storage.k8s.io,v1,VolumeAttributesClass
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeAttributesClass represents a specification of mutable volume attributes
// defined by the CSI driver. The class can be specified during dynamic provisioning
// of PersistentVolumeClaims, and changed in the PersistentVolumeClaim spec after provisioning.
type VolumeAttributesClass struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Name of the CSI driver
	// This field is immutable.
	DriverName string `json:"driverName" protobuf:"bytes,2,opt,name=driverName"`

	// parameters hold volume attributes defined by the CSI driver. These values
	// are opaque to the Kubernetes and are passed directly to the CSI driver.
	// The underlying storage provider supports changing these attributes on an
	// existing volume, however the parameters field itself is immutable. To
	// invoke a volume update, a new VolumeAttributesClass should be created with
	// new parameters, and the PersistentVolumeClaim should be updated to reference
	// the new VolumeAttributesClass.
	//
	// This field is required and must contain at least one key/value pair.
	// The keys cannot be empty, and the maximum number of parameters is 512, with
	// a cumulative max size of 256K. If the CSI driver rejects invalid parameters,
	// the target PersistentVolumeClaim will be set to an "Infeasible" state in the
	// modifyVolumeStatus field.
	Parameters map[string]string `json:"parameters,omitempty" protobuf:"bytes,3,rep,name=parameters"`
}

// +k8s:prerelease-lifecycle-gen:introduced=1.29
// +k8s:prerelease-lifecycle-gen:deprecated=1.32
// +k8s:prerelease-lifecycle-gen:replacement=storage.k8s.io,v1,VolumeAttributesClassList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeAttributesClassList is a collection of VolumeAttributesClass objects.
type VolumeAttributesClassList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of VolumeAttributesClass objects.
	Items []VolumeAttributesClass `json:"items" protobuf:"bytes,2,rep,name=items"`
}
