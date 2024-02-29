package cluster

import (
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type ManagementConfig struct {
	*Config
	Bundles     *releasev1alpha1.Bundles
	EKSARelease *releasev1alpha1.EKSARelease
}

func NewManagementConfig(eksaRelease *releasev1alpha1.EKSARelease, bundles *releasev1alpha1.Bundles, config *Config) *ManagementConfig {
	return &ManagementConfig{
		EKSARelease: eksaRelease,
		Bundles:     bundles,
		Config:      config,
	}
}
