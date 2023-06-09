package v1alpha1

import (
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// Conditions and condition Reasons for the Cluster object.

const (
	WaitingForCAPIClusterInitializedReason = "WaitingForCAPIClusterInitialized"

	WaitingForControlPlaneNodesReadyReason = "WaitingForControlPlaneNodesReady"

	WorkersReadyConditon clusterv1.ConditionType = "WorkersReady"

	WaitingForWorkerReadyReason = "WaitingForWorkerReady"

	DefaultCNIConfiguredCondition clusterv1.ConditionType = "DefaultCNIConfigured"
)
