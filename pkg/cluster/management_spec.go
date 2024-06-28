package cluster

import (
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// ManagementSpec contains the components needed to manage a cluster's management components.
type ManagementSpec struct {
	*Config
	ManagementComponents *ManagementComponents
	Bundles              *releasev1alpha1.Bundles
	EKSARelease          *releasev1alpha1.EKSARelease
}

// NewManagementSpec returns a new ManagementSpec.
func NewManagementSpec(config *Config, managementComponents *ManagementComponents, bundles *releasev1alpha1.Bundles, eksaRelease *releasev1alpha1.EKSARelease) *ManagementSpec {
	if eksaRelease != nil {
		config.Cluster.SetManagementComponentsVersion(eksaRelease.Spec.Version)
	}
	return &ManagementSpec{
		Config:               config,
		ManagementComponents: managementComponents,
		Bundles:              bundles,
		EKSARelease:          eksaRelease,
	}
}

// ManagementSpecFromClusterSpec returns a ManagementSpec from a ClusterSpec.
func ManagementSpecFromClusterSpec(cluster *Spec) *ManagementSpec {
	return NewManagementSpec(cluster.Config, cluster.ManagementComponents, cluster.Bundles, cluster.EKSARelease)
}
