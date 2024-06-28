package cli_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/aws/eks-anywhere/pkg/cli"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/constants"
)

func TestRunUpgradeClusterDefaulter(t *testing.T) {
	g := NewWithT(t)

	c := baseCluster()

	clusterSpec := &cluster.Spec{
		ManagementSpec: &cluster.ManagementSpec{
			Config: &cluster.Config{
				Cluster: c,
			},
		},
	}
	mhcDefaulter := cluster.NewMachineHealthCheckDefaulter(constants.DefaultNodeStartupTimeout, constants.DefaultUnhealthyMachineTimeout, intstr.Parse(constants.DefaultMaxUnhealthy), intstr.Parse(constants.DefaultWorkerMaxUnhealthy))

	upgradeClusterDefaulter := cli.NewUpgradeClusterDefaulter(mhcDefaulter)
	clusterSpec, err := upgradeClusterDefaulter.Run(context.Background(), clusterSpec)

	g.Expect(err).To(BeNil())
	g.Expect(clusterSpec.Cluster.Spec.MachineHealthCheck).ToNot(BeNil())
}
