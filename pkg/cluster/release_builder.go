package cluster

import (
	"github.com/aws/eks-anywhere/pkg/manifests"
	"github.com/aws/eks-anywhere/pkg/version"
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type ReleaseBuilder struct {
	reader             manifests.FileReader
	cliVersion         version.Info
	bundlesManifestURL string
}

// NewReleaseBuilder returns a new cluster.ReleaseBuilder
func (rb *ReleaseBuilder) NewReleaseBuilder(reader manifests.FileReader, cliVersion version.Info, bundlesManifestURL string) *releasev1alpha1.Release {

}

// Build contructs a cluster.Release by reading the bundles from the
// provided bundlesManifestURL and cliVersion.
func (rb *ReleaseBuilder) Build() *releasev1alpha1.Release {
}
