// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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

// This file was autogenerated by openapi-gen. Do not edit it manually!

package apis

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsnodevmattachment/v1alpha1.CnsNodeVmAttachment":       schema_pkg_apis_cns_v1alpha1_CnsNodeVmAttachment(ref),
		"sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsnodevmattachment/v1alpha1.CnsNodeVmAttachmentSpec":   schema_pkg_apis_cns_v1alpha1_CnsNodeVmAttachmentSpec(ref),
		"sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsnodevmattachment/v1alpha1.CnsNodeVmAttachmentStatus": schema_pkg_apis_cns_v1alpha1_CnsNodeVmAttachmentStatus(ref),
		"sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsvolumemetadata/v1alpha1.CnsVolumeMetadata":           schema_pkg_apis_cns_v1alpha1_CnsVolumeMetadata(ref),
		"sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsvolumemetadata/v1alpha1.CnsVolumeMetadataSpec":       schema_pkg_apis_cns_v1alpha1_CnsVolumeMetadataSpec(ref),
		"sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsvolumemetadata/v1alpha1.CnsVolumeMetadataStatus":     schema_pkg_apis_cns_v1alpha1_CnsVolumeMetadataStatus(ref),
		"sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsregistervolume/v1alpha1.CnsRegisterVolume":           schema_pkg_apis_cnsregistervolume_v1alpha1_CnsRegisterVolume(ref),
		"sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsregistervolume/v1alpha1.CnsRegisterVolumeSpec":       schema_pkg_apis_cnsregistervolume_v1alpha1_CnsRegisterVolumeSpec(ref),
		"sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsregistervolume/v1alpha1.CnsRegisterVolumeStatus":     schema_pkg_apis_cnsregistervolume_v1alpha1_CnsRegisterVolumeStatus(ref),
	}
}

func schema_pkg_apis_cns_v1alpha1_CnsNodeVmAttachment(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CnsNodeVmAttachment is the Schema for the cnsnodevmattachments API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsnodevmattachment/v1alpha1.CnsNodeVmAttachmentSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsnodevmattachment/v1alpha1.CnsNodeVmAttachmentStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsnodevmattachment/v1alpha1.CnsNodeVmAttachmentSpec", "sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsnodevmattachment/v1alpha1.CnsNodeVmAttachmentStatus"},
	}
}

func schema_pkg_apis_cns_v1alpha1_CnsNodeVmAttachmentSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CnsNodeVmAttachmentSpec defines the desired state of CnsNodeVmAttachment",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_cns_v1alpha1_CnsNodeVmAttachmentStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CnsNodeVmAttachmentStatus defines the observed state of CnsNodeVmAttachment",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_cns_v1alpha1_CnsVolumeMetadata(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CnsVolumeMetadata is the Schema for the cnsvolumemetadata API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsvolumemetadata/v1alpha1.CnsVolumeMetadataSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsvolumemetadata/v1alpha1.CnsVolumeMetadataStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsvolumemetadata/v1alpha1.CnsVolumeMetadataSpec", "sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsvolumemetadata/v1alpha1.CnsVolumeMetadataStatus"},
	}
}

func schema_pkg_apis_cns_v1alpha1_CnsVolumeMetadataSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CnsVolumeMetadataSpec defines the desired state of CnsVolumeMetadata",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_cns_v1alpha1_CnsVolumeMetadataStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CnsVolumeMetadataStatus defines the observed state of CnsVolumeMetadata",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_cnsregistervolume_v1alpha1_CnsRegisterVolume(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CnsRegisterVolume is the Schema for the cnsregistervolumes API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsregistervolume/v1alpha1.CnsRegisterVolumeSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsregistervolume/v1alpha1.CnsRegisterVolumeStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsregistervolume/v1alpha1.CnsRegisterVolumeSpec", "sigs.k8s.io/vsphere-csi-driver/pkg/apis/cnsoperator/cnsregistervolume/v1alpha1.CnsRegisterVolumeStatus"},
	}
}

func schema_pkg_apis_cnsregistervolume_v1alpha1_CnsRegisterVolumeSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CnsRegisterVolumeSpec defines the desired state of CnsRegisterVolume",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_cnsregistervolume_v1alpha1_CnsRegisterVolumeStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CnsRegisterVolumeStatus defines the observed state of CnsRegisterVolume",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
