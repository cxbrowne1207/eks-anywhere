package cluster

import (
	"github.com/aws/eks-anywhere/pkg/constants"
	"github.com/aws/eks-anywhere/pkg/manifests"
	"github.com/aws/eks-anywhere/pkg/manifests/bundles"
	"github.com/aws/eks-anywhere/pkg/version"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	"github.com/pkg/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ReleaseBuilder struct {
	reader              manifests.FileReader
	cliVersion          version.Info
	bundlesManifestURL  string
	releasesManifestURL string
}

// ReleaseBuilderOpt allows to configure [ReleaseBuilder].
type ReleaseBuilderOpt func(*ReleaseBuilder)

// WithReleasesManifest configures the URL to read the Releases manifest.
func WithReleasesManifest(url string) ReleaseBuilderOpt {
	return func(b *ReleaseBuilder) {
		b.releasesManifestURL = url
	}
}

// WithOverrideBundlesManifest configures the URL to read the Bundles manifest.
// This overrides the Bundles declared in the Releases so reading the Releases
// manifest is skipped.
func WithOverrideBundlesManifest(url string) ReleaseBuilderOpt {
	return func(b *ReleaseBuilder) {
		b.bundlesManifestURL = url
	}
}

// NewReleaseBuilder returns a new cluster.ReleaseBuilder
func NewReleaseBuilder(reader manifests.FileReader, cliVersion version.Info, opts ...ReleaseBuilderOpt) ReleaseBuilder {
	rb := &ReleaseBuilder{
		reader:     reader,
		cliVersion: cliVersion,
	}

	for _, opt := range opts {
		opt(rb)
	}

	return *rb
}

// Build contructs a cluster.Release by reading the bundles from the
// provided bundlesManifestURL and cliVersion.
func (rb *ReleaseBuilder) Build() (*Release, error) {
	mReader := rb.createManifestReader()
	release, err := rb.getEksaRelease(mReader)
	if err != nil {
		return nil, err
	}

	bundlesManifest, err := rb.getBundles(mReader)
	if err != nil {
		return nil, errors.Wrapf(err, "getting Bundles file")
	}

	eksaRelease := buildEKSARelease(release, bundlesManifest)

	return &Release{
		Bundles:     bundlesManifest,
		EKSARelease: eksaRelease,
	}, nil
}

func (rb *ReleaseBuilder) createManifestReader() *manifests.Reader {
	var opts []manifests.ReaderOpt
	if rb.releasesManifestURL != "" {
		opts = append(opts, manifests.WithReleasesManifest(rb.releasesManifestURL))
	}
	return manifests.NewReader(rb.reader, opts...)
}

func (rb *ReleaseBuilder) getBundles(manifestReader *manifests.Reader) (*releasev1.Bundles, error) {
	bundlesURL := rb.bundlesManifestURL
	if bundlesURL == "" {
		return manifestReader.ReadBundlesForVersion(rb.cliVersion.GitVersion)
	}

	return bundles.Read(rb.reader, bundlesURL)
}

func (rb *ReleaseBuilder) getEksaRelease(mReader *manifests.Reader) (*releasev1.EksARelease, error) {
	if rb.bundlesManifestURL == "" {
		// this shouldn't return an error at this point due to getBundles performing similar operations prior to this call
		release, err := mReader.ReadReleaseForVersion(rb.cliVersion.GitVersion)
		if err != nil {
			return nil, err
		}
		return release, nil
	}

	// When using bundles-override or a custom bundle, a fake EksaRelease can be used since using a custom bundle
	// is like creating a new EKS-A version.
	return &releasev1.EksARelease{
		Version: rb.cliVersion.GitVersion,
	}, nil
}

// buildEKSARelease builds an EKSARelease from a Release and a Bundles.
func buildEKSARelease(release *releasev1.EksARelease, bundle *releasev1.Bundles) *releasev1.EKSARelease {
	eksaRelease := &releasev1.EKSARelease{
		TypeMeta: v1.TypeMeta{
			Kind:       releasev1.EKSAReleaseKind,
			APIVersion: releasev1.SchemeBuilder.GroupVersion.String(),
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      releasev1.GenerateEKSAReleaseName(release.Version),
			Namespace: constants.EksaSystemNamespace,
		},
		Spec: releasev1.EKSAReleaseSpec{
			ReleaseDate:       release.Date,
			Version:           release.Version,
			GitCommit:         release.GitCommit,
			BundleManifestURL: release.BundleManifestUrl,
			BundlesRef: releasev1.BundlesRef{
				APIVersion: releasev1.GroupVersion.String(),
				Name:       bundle.Name,
				Namespace:  bundle.Namespace,
			},
		},
	}
	return eksaRelease
}
