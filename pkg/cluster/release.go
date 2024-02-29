// pkg/cluster/release.go
package cluster

import releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"

// Release represents a grouping of EKS-A release objects.
type Release struct {
	Bundles     *releasev1alpha1.Bundles
	EKSARelease *releasev1alpha1.EKSARelease
}

// NewRelease returns a new cluster.Release object.
func NewRelease(bundles *releasev1alpha1.Bundles, eksaRelease *releasev1alpha1.EKSARelease) *Release {
	return &Release{
		Bundles:     bundles,
		EKSARelease: eksaRelease,
	}
}
