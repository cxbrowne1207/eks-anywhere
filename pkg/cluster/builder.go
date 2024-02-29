package cluster

import (
	eksdv1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
	"github.com/pkg/errors"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/constants"
	"github.com/aws/eks-anywhere/pkg/logger"
	"github.com/aws/eks-anywhere/pkg/manifests"
	"github.com/aws/eks-anywhere/pkg/manifests/bundles"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// FileSpecBuilder allows to build [Spec] by reading from files.
type FileSpecBuilder struct {
	// TODO(g-gaston): this is very much a CLI thing. Move to `pkg/cli` when available.
	releaseBuilder ReleaseBuilder
	reader         manifests.FileReader
}

// NewFileSpecBuilder builds a new [FileSpecBuilder].
// cliVersion is used to chose the right Bundles from the the Release manifest.
func NewFileSpecBuilder(reader manifests.FileReader, releaseBuilder ReleaseBuilder) FileSpecBuilder {
	return FileSpecBuilder{
		releaseBuilder: releaseBuilder,
		reader:         reader,
	}
}

// Build constructs a new [Spec] by reading the cluster config in yaml from a file and
// Releases, Bundles and EKS-D manifests from the configured URLs.
func (b FileSpecBuilder) Build(clusterConfigURL string) (*Spec, error) {
	config, err := b.getConfig(clusterConfigURL)
	if err != nil {
		return nil, err
	}

	release, err := b.releaseBuilder.Build()
	if err != nil {
		return nil, errors.Wrapf(err, "getting Release files")
	}

	bundlesManifest := release.Bundles
	bundle := bundlesManifest.Spec.VersionsBundles[0]
	var infraProviderName, infraProviderVersion string
	switch config.Cluster.Spec.DatacenterRef.Kind {
	case v1alpha1.CloudStackDatacenterKind:
		infraProviderName = "Cluster API Provider CloudStack"
		infraProviderVersion = bundle.CloudStack.Version
	case v1alpha1.DockerDatacenterKind:
		infraProviderName = "Cluster API Provider Docker"
		infraProviderVersion = bundle.Docker.Version
	case v1alpha1.NutanixDatacenterKind:
		infraProviderName = "Cluster API Provider Nutanix"
		infraProviderVersion = bundle.Nutanix.Version
	case v1alpha1.SnowDatacenterKind:
		infraProviderName = "Cluster API Provider AWS Snow"
		infraProviderVersion = bundle.Snow.Version
	case v1alpha1.TinkerbellDatacenterKind:
		infraProviderName = "Cluster API Provider Tinkerbell"
		infraProviderVersion = bundle.Tinkerbell.Version
	case v1alpha1.VSphereDatacenterKind:
		infraProviderName = "Cluster API Provider VSphere"
		infraProviderVersion = bundle.VSphere.Version
	}
	logger.V(4).Info(
		"Using CAPI provider versions",
		"Core Cluster API", bundle.ClusterAPI.Version,
		"Kubeadm Bootstrap", bundle.Bootstrap.Version,
		"Kubeadm Control Plane", bundle.ControlPlane.Version,
		"External etcd Bootstrap", bundle.ExternalEtcdBootstrap.Version,
		"External etcd Controller", bundle.ExternalEtcdController.Version,
		infraProviderName, infraProviderVersion,
	)
	bundlesManifest.Namespace = constants.EksaSystemNamespace

	configManager, err := NewDefaultConfigManager()
	if err != nil {
		return nil, err
	}

	config.Cluster.Spec.BundlesRef = nil

	if err = configManager.SetDefaults(config); err != nil {
		return nil, err
	}

	eksdReleases, err := getAllEksdReleases(config.Cluster, bundlesManifest, b.reader)
	if err != nil {
		return nil, err
	}

	releaseVersion := v1alpha1.EksaVersion(release.EKSARelease.Spec.Version)
	config.Cluster.Spec.EksaVersion = &releaseVersion

	return NewSpec(config, bundlesManifest, eksdReleases, release.EKSARelease)
}

func getAllEksdReleases(cluster *v1alpha1.Cluster, bundlesManifest *releasev1.Bundles, reader bundles.Reader) ([]eksdv1.Release, error) {
	versions := cluster.KubernetesVersions()
	m := make([]eksdv1.Release, 0, len(versions))
	for _, version := range versions {
		eksd, err := getEksdReleases(version, bundlesManifest, reader)
		if err != nil {
			return nil, err
		}
		m = append(m, *eksd)
	}
	return m, nil
}

func getEksdReleases(version v1alpha1.KubernetesVersion, bundlesManifest *releasev1.Bundles, reader bundles.Reader) (*eksdv1.Release, error) {
	versionsBundle, err := GetVersionsBundle(version, bundlesManifest)
	if err != nil {
		return nil, err
	}

	eksd, err := bundles.ReadEKSD(reader, *versionsBundle)
	if err != nil {
		return nil, err
	}

	return eksd, nil
}

func (b FileSpecBuilder) getConfig(clusterConfigURL string) (*Config, error) {
	yaml, err := b.reader.ReadFile(clusterConfigURL)
	if err != nil {
		return nil, errors.Wrapf(err, "reading cluster config file")
	}

	return ParseConfig(yaml)
}
