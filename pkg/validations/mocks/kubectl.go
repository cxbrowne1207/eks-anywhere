// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/validations/kubectl.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	kubernetes "github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	executables "github.com/aws/eks-anywhere/pkg/executables"
	types "github.com/aws/eks-anywhere/pkg/types"
	v1alpha10 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// MockKubectlClient is a mock of KubectlClient interface.
type MockKubectlClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubectlClientMockRecorder
}

// MockKubectlClientMockRecorder is the mock recorder for MockKubectlClient.
type MockKubectlClientMockRecorder struct {
	mock *MockKubectlClient
}

// NewMockKubectlClient creates a new mock instance.
func NewMockKubectlClient(ctrl *gomock.Controller) *MockKubectlClient {
	mock := &MockKubectlClient{ctrl: ctrl}
	mock.recorder = &MockKubectlClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubectlClient) EXPECT() *MockKubectlClientMockRecorder {
	return m.recorder
}

// GetBundles mocks base method.
func (m *MockKubectlClient) GetBundles(ctx context.Context, kubeconfigFile, name, namespace string) (*v1alpha10.Bundles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBundles", ctx, kubeconfigFile, name, namespace)
	ret0, _ := ret[0].(*v1alpha10.Bundles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBundles indicates an expected call of GetBundles.
func (mr *MockKubectlClientMockRecorder) GetBundles(ctx, kubeconfigFile, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBundles", reflect.TypeOf((*MockKubectlClient)(nil).GetBundles), ctx, kubeconfigFile, name, namespace)
}

// GetClusters mocks base method.
func (m *MockKubectlClient) GetClusters(ctx context.Context, cluster *types.Cluster) ([]types.CAPICluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusters", ctx, cluster)
	ret0, _ := ret[0].([]types.CAPICluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusters indicates an expected call of GetClusters.
func (mr *MockKubectlClientMockRecorder) GetClusters(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusters", reflect.TypeOf((*MockKubectlClient)(nil).GetClusters), ctx, cluster)
}

// GetEksaAWSIamConfig mocks base method.
func (m *MockKubectlClient) GetEksaAWSIamConfig(ctx context.Context, awsIamConfigName, kubeconfigFile, namespace string) (*v1alpha1.AWSIamConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaAWSIamConfig", ctx, awsIamConfigName, kubeconfigFile, namespace)
	ret0, _ := ret[0].(*v1alpha1.AWSIamConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaAWSIamConfig indicates an expected call of GetEksaAWSIamConfig.
func (mr *MockKubectlClientMockRecorder) GetEksaAWSIamConfig(ctx, awsIamConfigName, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaAWSIamConfig", reflect.TypeOf((*MockKubectlClient)(nil).GetEksaAWSIamConfig), ctx, awsIamConfigName, kubeconfigFile, namespace)
}

// GetEksaCluster mocks base method.
func (m *MockKubectlClient) GetEksaCluster(ctx context.Context, cluster *types.Cluster, clusterName string) (*v1alpha1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaCluster", ctx, cluster, clusterName)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaCluster indicates an expected call of GetEksaCluster.
func (mr *MockKubectlClientMockRecorder) GetEksaCluster(ctx, cluster, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaCluster", reflect.TypeOf((*MockKubectlClient)(nil).GetEksaCluster), ctx, cluster, clusterName)
}

// GetEksaFluxConfig mocks base method.
func (m *MockKubectlClient) GetEksaFluxConfig(ctx context.Context, fluxConfigName, kubeconfigFile, namespace string) (*v1alpha1.FluxConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaFluxConfig", ctx, fluxConfigName, kubeconfigFile, namespace)
	ret0, _ := ret[0].(*v1alpha1.FluxConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaFluxConfig indicates an expected call of GetEksaFluxConfig.
func (mr *MockKubectlClientMockRecorder) GetEksaFluxConfig(ctx, fluxConfigName, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaFluxConfig", reflect.TypeOf((*MockKubectlClient)(nil).GetEksaFluxConfig), ctx, fluxConfigName, kubeconfigFile, namespace)
}

// GetEksaGitOpsConfig mocks base method.
func (m *MockKubectlClient) GetEksaGitOpsConfig(ctx context.Context, gitOpsConfigName, kubeconfigFile, namespace string) (*v1alpha1.GitOpsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaGitOpsConfig", ctx, gitOpsConfigName, kubeconfigFile, namespace)
	ret0, _ := ret[0].(*v1alpha1.GitOpsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaGitOpsConfig indicates an expected call of GetEksaGitOpsConfig.
func (mr *MockKubectlClientMockRecorder) GetEksaGitOpsConfig(ctx, gitOpsConfigName, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaGitOpsConfig", reflect.TypeOf((*MockKubectlClient)(nil).GetEksaGitOpsConfig), ctx, gitOpsConfigName, kubeconfigFile, namespace)
}

// GetEksaOIDCConfig mocks base method.
func (m *MockKubectlClient) GetEksaOIDCConfig(ctx context.Context, oidcConfigName, kubeconfigFile, namespace string) (*v1alpha1.OIDCConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaOIDCConfig", ctx, oidcConfigName, kubeconfigFile, namespace)
	ret0, _ := ret[0].(*v1alpha1.OIDCConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaOIDCConfig indicates an expected call of GetEksaOIDCConfig.
func (mr *MockKubectlClientMockRecorder) GetEksaOIDCConfig(ctx, oidcConfigName, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaOIDCConfig", reflect.TypeOf((*MockKubectlClient)(nil).GetEksaOIDCConfig), ctx, oidcConfigName, kubeconfigFile, namespace)
}

// GetEksaTinkerbellDatacenterConfig mocks base method.
func (m *MockKubectlClient) GetEksaTinkerbellDatacenterConfig(ctx context.Context, tinkerbellDatacenterConfigName, kubeconfigFile, namespace string) (*v1alpha1.TinkerbellDatacenterConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaTinkerbellDatacenterConfig", ctx, tinkerbellDatacenterConfigName, kubeconfigFile, namespace)
	ret0, _ := ret[0].(*v1alpha1.TinkerbellDatacenterConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaTinkerbellDatacenterConfig indicates an expected call of GetEksaTinkerbellDatacenterConfig.
func (mr *MockKubectlClientMockRecorder) GetEksaTinkerbellDatacenterConfig(ctx, tinkerbellDatacenterConfigName, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaTinkerbellDatacenterConfig", reflect.TypeOf((*MockKubectlClient)(nil).GetEksaTinkerbellDatacenterConfig), ctx, tinkerbellDatacenterConfigName, kubeconfigFile, namespace)
}

// GetEksaTinkerbellMachineConfig mocks base method.
func (m *MockKubectlClient) GetEksaTinkerbellMachineConfig(ctx context.Context, tinkerbellMachineConfigName, kubeconfigFile, namespace string) (*v1alpha1.TinkerbellMachineConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaTinkerbellMachineConfig", ctx, tinkerbellMachineConfigName, kubeconfigFile, namespace)
	ret0, _ := ret[0].(*v1alpha1.TinkerbellMachineConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaTinkerbellMachineConfig indicates an expected call of GetEksaTinkerbellMachineConfig.
func (mr *MockKubectlClientMockRecorder) GetEksaTinkerbellMachineConfig(ctx, tinkerbellMachineConfigName, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaTinkerbellMachineConfig", reflect.TypeOf((*MockKubectlClient)(nil).GetEksaTinkerbellMachineConfig), ctx, tinkerbellMachineConfigName, kubeconfigFile, namespace)
}

// GetEksaVSphereDatacenterConfig mocks base method.
func (m *MockKubectlClient) GetEksaVSphereDatacenterConfig(ctx context.Context, vsphereDatacenterConfigName, kubeconfigFile, namespace string) (*v1alpha1.VSphereDatacenterConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaVSphereDatacenterConfig", ctx, vsphereDatacenterConfigName, kubeconfigFile, namespace)
	ret0, _ := ret[0].(*v1alpha1.VSphereDatacenterConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaVSphereDatacenterConfig indicates an expected call of GetEksaVSphereDatacenterConfig.
func (mr *MockKubectlClientMockRecorder) GetEksaVSphereDatacenterConfig(ctx, vsphereDatacenterConfigName, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaVSphereDatacenterConfig", reflect.TypeOf((*MockKubectlClient)(nil).GetEksaVSphereDatacenterConfig), ctx, vsphereDatacenterConfigName, kubeconfigFile, namespace)
}

// GetObject mocks base method.
func (m *MockKubectlClient) GetObject(ctx context.Context, resourceType, name, namespace, kubeconfig string, obj runtime.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", ctx, resourceType, name, namespace, kubeconfig, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetObject indicates an expected call of GetObject.
func (mr *MockKubectlClientMockRecorder) GetObject(ctx, resourceType, name, namespace, kubeconfig, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockKubectlClient)(nil).GetObject), ctx, resourceType, name, namespace, kubeconfig, obj)
}

// List mocks base method.
func (m *MockKubectlClient) List(ctx context.Context, kubeconfig string, list kubernetes.ObjectList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, kubeconfig, list)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockKubectlClientMockRecorder) List(ctx, kubeconfig, list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockKubectlClient)(nil).List), ctx, kubeconfig, list)
}

// SearchIdentityProviderConfig mocks base method.
func (m *MockKubectlClient) SearchIdentityProviderConfig(ctx context.Context, ipName, kind, kubeconfigFile, namespace string) ([]*v1alpha1.VSphereDatacenterConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIdentityProviderConfig", ctx, ipName, kind, kubeconfigFile, namespace)
	ret0, _ := ret[0].([]*v1alpha1.VSphereDatacenterConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIdentityProviderConfig indicates an expected call of SearchIdentityProviderConfig.
func (mr *MockKubectlClientMockRecorder) SearchIdentityProviderConfig(ctx, ipName, kind, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIdentityProviderConfig", reflect.TypeOf((*MockKubectlClient)(nil).SearchIdentityProviderConfig), ctx, ipName, kind, kubeconfigFile, namespace)
}

// ValidateClustersCRD mocks base method.
func (m *MockKubectlClient) ValidateClustersCRD(ctx context.Context, cluster *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateClustersCRD", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateClustersCRD indicates an expected call of ValidateClustersCRD.
func (mr *MockKubectlClientMockRecorder) ValidateClustersCRD(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateClustersCRD", reflect.TypeOf((*MockKubectlClient)(nil).ValidateClustersCRD), ctx, cluster)
}

// ValidateControlPlaneNodes mocks base method.
func (m *MockKubectlClient) ValidateControlPlaneNodes(ctx context.Context, cluster *types.Cluster, clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateControlPlaneNodes", ctx, cluster, clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateControlPlaneNodes indicates an expected call of ValidateControlPlaneNodes.
func (mr *MockKubectlClientMockRecorder) ValidateControlPlaneNodes(ctx, cluster, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateControlPlaneNodes", reflect.TypeOf((*MockKubectlClient)(nil).ValidateControlPlaneNodes), ctx, cluster, clusterName)
}

// ValidateEKSAClustersCRD mocks base method.
func (m *MockKubectlClient) ValidateEKSAClustersCRD(ctx context.Context, cluster *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateEKSAClustersCRD", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateEKSAClustersCRD indicates an expected call of ValidateEKSAClustersCRD.
func (mr *MockKubectlClientMockRecorder) ValidateEKSAClustersCRD(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateEKSAClustersCRD", reflect.TypeOf((*MockKubectlClient)(nil).ValidateEKSAClustersCRD), ctx, cluster)
}

// ValidateNodes mocks base method.
func (m *MockKubectlClient) ValidateNodes(ctx context.Context, kubeconfig string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateNodes", ctx, kubeconfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateNodes indicates an expected call of ValidateNodes.
func (mr *MockKubectlClientMockRecorder) ValidateNodes(ctx, kubeconfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateNodes", reflect.TypeOf((*MockKubectlClient)(nil).ValidateNodes), ctx, kubeconfig)
}

// ValidateWorkerNodes mocks base method.
func (m *MockKubectlClient) ValidateWorkerNodes(ctx context.Context, clusterName, kubeconfig string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateWorkerNodes", ctx, clusterName, kubeconfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateWorkerNodes indicates an expected call of ValidateWorkerNodes.
func (mr *MockKubectlClientMockRecorder) ValidateWorkerNodes(ctx, clusterName, kubeconfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateWorkerNodes", reflect.TypeOf((*MockKubectlClient)(nil).ValidateWorkerNodes), ctx, clusterName, kubeconfig)
}

// Version mocks base method.
func (m *MockKubectlClient) Version(ctx context.Context, cluster *types.Cluster) (*executables.VersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", ctx, cluster)
	ret0, _ := ret[0].(*executables.VersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockKubectlClientMockRecorder) Version(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockKubectlClient)(nil).Version), ctx, cluster)
}
