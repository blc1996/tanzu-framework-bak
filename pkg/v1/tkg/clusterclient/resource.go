// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package clusterclient

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	betav1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	extensionsV1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	capav1alpha3 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	capzv1alpha3 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	capvv1alpha3 "sigs.k8s.io/cluster-api-provider-vsphere/api/v1alpha3"
	capiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	clusterctlv1 "sigs.k8s.io/cluster-api/cmd/clusterctl/api/v1alpha3"
	controlplanev1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1alpha3"
	addonsv1 "sigs.k8s.io/cluster-api/exp/addons/api/v1alpha3"
	capdv1alpha3 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1alpha3"
	"sigs.k8s.io/cluster-api/util/conditions"
	crtclient "sigs.k8s.io/controller-runtime/pkg/client"

	runv1alpha1 "github.com/vmware-tanzu-private/core/apis/run/v1alpha1"
	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/api/tmc/v1alpha1"
	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/log"
)

// PostVerifyrFunc is a function which should be used as closure
type PostVerifyrFunc func(obj runtime.Object) error

func (c *client) ListResources(resourceReference interface{}, option ...crtclient.ListOption) error {
	obj, err := c.getRuntimeObject(resourceReference)
	if err != nil {
		return err
	}
	if err := c.clientSet.List(context.TODO(), obj, option...); err != nil {
		return errors.Wrapf(err, "failed to list %v", reflect.TypeOf(resourceReference))
	}
	return nil
}

func (c *client) DeleteResource(resourceReference interface{}) error {
	obj, err := c.getRuntimeObject(resourceReference)
	if err != nil {
		return err
	}
	if err := c.clientSet.Delete(ctx, obj); err != nil {
		return errors.Wrapf(err, "failed to delete %v", reflect.TypeOf(resourceReference))
	}
	return nil
}

func (c *client) CreateResource(resourceReference interface{}, resourceName, namespace string, opts ...crtclient.CreateOption) error {
	obj, err := c.getRuntimeObject(resourceReference)
	if err != nil {
		return err
	}

	if err := c.clientSet.Create(ctx, obj, opts...); err != nil {
		return errors.Wrapf(err, "error while creating object for %q %s/%s",
			obj.GetObjectKind(), namespace, resourceName)
	}
	return nil
}

func (c *client) UpdateResource(resourceReference interface{}, resourceName, namespace string, opts ...crtclient.UpdateOption) error {
	obj, err := c.getRuntimeObject(resourceReference)
	if err != nil {
		return err
	}

	if err := c.clientSet.Update(ctx, obj, opts...); err != nil {
		return errors.Wrapf(err, "error while creating object for %q %s/%s",
			obj.GetObjectKind(), namespace, resourceName)
	}
	return nil
}

func (c *client) PatchResource(resourceReference interface{}, resourceName, namespace, patchJSONString string, patchType types.PatchType, pollOptions *PollOptions) error {
	// if pollOptions are provided use the polling and wait for the result/error/timeout
	// else use normal get
	if pollOptions != nil {
		log.V(6).Infof("Applying patch to resource %s of type %s ...", resourceName, reflect.TypeOf(resourceReference))
		_, err := c.poller.PollImmediateWithGetter(pollOptions.Interval, pollOptions.Timeout, func() (interface{}, error) {
			return nil, c.patchResource(resourceReference, resourceName, namespace, patchJSONString, patchType)
		})
		return err
	}

	return c.patchResource(resourceReference, resourceName, namespace, patchJSONString, patchType)
}

func (c *client) patchResource(resourceReference interface{}, resourceName, namespace, patchJSONString string, patchType types.PatchType) error {
	patch := crtclient.RawPatch(patchType, []byte(patchJSONString))

	obj, err := c.getRuntimeObject(resourceReference)
	if err != nil {
		return err
	}
	clusterObjKey := crtclient.ObjectKey{
		Namespace: namespace,
		Name:      resourceName,
	}

	if err := c.clientSet.Get(ctx, clusterObjKey, obj); err != nil {
		return errors.Wrapf(err, "error reading %q %s/%s",
			obj.GetObjectKind(), namespace, resourceName)
	}

	if err := c.clientSet.Patch(ctx, obj, patch); err != nil {
		return errors.Wrapf(err, "error while applying patch for %q %s/%s",
			obj.GetObjectKind(), namespace, resourceName)
	}
	return nil
}

func (c *client) get(objectName, namespace string, o interface{}, postVerify PostVerifyrFunc) error {
	obj, err := c.getRuntimeObject(o)
	if err != nil {
		return err
	}
	objKey := crtclient.ObjectKey{Name: objectName, Namespace: namespace}
	if err := c.clientSet.Get(ctx, objKey, obj); err != nil {
		return err
	}
	if postVerify != nil {
		return postVerify(obj)
	}
	return nil
}

func (c *client) list(clusterName, namespace string, o interface{}, postVerify PostVerifyrFunc) error {
	obj, err := c.getRuntimeObject(o)
	if err != nil {
		return err
	}

	selectors := []crtclient.ListOption{
		crtclient.InNamespace(namespace),
		crtclient.MatchingLabels(map[string]string{capi.ClusterLabelName: clusterName}),
	}

	if err := c.clientSet.List(ctx, obj, selectors...); err != nil {
		return err
	}

	if postVerify != nil {
		return postVerify(obj)
	}
	return nil
}

func (c *client) getRuntimeObject(o interface{}) (runtime.Object, error) { //nolint:gocyclo,funlen
	switch obj := o.(type) {
	case *corev1.Namespace:
		return obj, nil
	case *corev1.Secret:
		return obj, nil
	case *corev1.SecretList:
		return obj, nil
	case *capi.Cluster:
		return obj, nil
	case *appsv1.Deployment:
		return obj, nil
	case *clusterctlv1.ProviderList:
		return obj, nil
	case *capi.ClusterList:
		return obj, nil
	case *capi.Machine:
		return obj, nil
	case *capi.MachineHealthCheck:
		return obj, nil
	case *capi.MachineHealthCheckList:
		return obj, nil
	case *capiv1alpha2.MachineList:
		return obj, nil
	case *capi.MachineList:
		return obj, nil
	case *capi.MachineDeploymentList:
		return obj, nil
	case *capi.MachineDeployment:
		return obj, nil
	case *controlplanev1.KubeadmControlPlane:
		return obj, nil
	case *controlplanev1.KubeadmControlPlaneList:
		return obj, nil
	case *unstructured.Unstructured:
		return obj, nil
	case *betav1.CronJobList:
		return obj, nil
	case *betav1.CronJob:
		return obj, nil
	case *capvv1alpha3.VSphereCluster:
		return obj, nil
	case *capvv1alpha3.VSphereMachineTemplate:
		return obj, nil
	case *capvv1alpha3.VSphereClusterList:
		return obj, nil
	case *capav1alpha3.AWSMachineTemplate:
		return obj, nil
	case *capdv1alpha3.DockerMachineTemplate:
		return obj, nil
	case *capav1alpha3.AWSCluster:
		return obj, nil
	case *capiv1alpha2.MachineDeploymentList:
		return obj, nil
	case *appsv1.DaemonSet:
		return obj, nil
	case *corev1.ConfigMap:
		return obj, nil
	case *v1alpha1.Extension:
		return obj, nil
	case *v1alpha1.ExtensionList:
		return obj, nil
	case *extensionsV1.CustomResourceDefinition:
		return obj, nil
	case *capzv1alpha3.AzureMachineTemplate:
		return obj, nil
	case *capzv1alpha3.AzureCluster:
		return obj, nil
	case *corev1.ServiceAccount:
		return obj, nil
	case *rbacv1.ClusterRole:
		return obj, nil
	case *rbacv1.ClusterRoleBinding:
		return obj, nil
	case *addonsv1.ClusterResourceSet:
		return obj, nil
	case *addonsv1.ClusterResourceSetList:
		return obj, nil
	case *runv1alpha1.TanzuKubernetesReleaseList:
		return obj, nil
	case *runv1alpha1.TanzuKubernetesRelease:
		return obj, nil
	default:
		return nil, errors.New("invalid object type")
	}
}

// VerifyClusterInitialized verifies the cluster is initialized or not (this is required before reading the kubeconfig secret)
func VerifyClusterInitialized(obj runtime.Object) error {
	switch cluster := obj.(type) {
	case *capi.Cluster:
		errList := []error{}
		if !conditions.IsTrue(cluster, capi.ControlPlaneReadyCondition) {
			errList = append(errList, errors.New("cluster control plane is still being initialized"))
		}

		// Nb. We are verifying infrastructure ready at this stage because it provides an early signal that the infrastructure provided is
		// properly working, but this is not strictly required for getting the kubeconfig secret
		if !conditions.IsTrue(cluster, capi.InfrastructureReadyCondition) {
			errList = append(errList, errors.New("cluster infrastructure is still being provisioned"))
		}
		return kerrors.NewAggregate(errList)
	default:
		return errors.Errorf("invalid type: %s during VerifyClusterInitialized", reflect.TypeOf(cluster))
	}
}

// VerifyClusterReady verifies the cluster is ready or not (this is required before starting the move operation)
func VerifyClusterReady(obj runtime.Object) error {
	// Nb. Currently there is no difference between VerifyClusterReady and VerifyClusterInitialized unless WorkersReady condition
	// would be added to cluster `Ready` condition aggregation.
	return VerifyClusterInitialized(obj)
}

// VerifyMachinesReady verifies the machine are ready or not (this is required before starting the move operation)
func VerifyMachinesReady(obj runtime.Object) error {
	switch machines := obj.(type) {
	case *capi.MachineList:
		errList := []error{}
		// Checking all the machine have a NodeRef
		// Nb. NodeRef is considered a better signal than InfrastructureReady, because it ensures the node in the workload cluster is up and running.
		for i := range machines.Items {
			if machines.Items[i].Status.NodeRef == nil {
				errList = append(errList, errors.Errorf("machine %s is still being provisioned", machines.Items[i].Name))
			}
		}
		return kerrors.NewAggregate(errList)
	default:
		return errors.Errorf("invalid type: %s during VerifyMachinesReady", reflect.TypeOf(machines))
	}
}

// VerifyKubeadmControlPlaneReplicas verifies the KubeadmControlPlane has all the required replicas (this is required before starting the move operation)
func VerifyKubeadmControlPlaneReplicas(obj runtime.Object) error {
	switch kcps := obj.(type) {
	case *controlplanev1.KubeadmControlPlaneList:
		errList := []error{}
		for i := range kcps.Items {
			var desiredReplica int32 = 1
			if kcps.Items[i].Spec.Replicas != nil {
				desiredReplica = *kcps.Items[i].Spec.Replicas
			}
			if desiredReplica != kcps.Items[i].Status.ReadyReplicas {
				errList = append(errList, errors.Errorf("control-plane is still creating replicas, DesiredReplicas=%v Replicas=%v ReadyReplicas=%v UpdatedReplicas=%v",
					desiredReplica, kcps.Items[i].Status.Replicas, kcps.Items[i].Status.ReadyReplicas, kcps.Items[i].Status.UpdatedReplicas))
			}
		}
		return kerrors.NewAggregate(errList)
	default:
		return errors.Errorf("invalid type: %s during VerifyKubeadmControlPlaneReplicas", reflect.TypeOf(kcps))
	}
}

// VerifyMachineDeploymentsReplicas verifies the MachineDeployment has all the required replicas (this is required before starting the move operation)
func VerifyMachineDeploymentsReplicas(obj runtime.Object) error {
	switch deployments := obj.(type) {
	case *capi.MachineDeploymentList:
		errList := []error{}
		for i := range deployments.Items {
			var desiredReplica int32 = 1
			if deployments.Items[i].Spec.Replicas != nil {
				desiredReplica = *deployments.Items[i].Spec.Replicas
			}
			if desiredReplica != deployments.Items[i].Status.ReadyReplicas {
				errList = append(errList, errors.Errorf("worker nodes are still being created for MachineDeployment '%s', DesiredReplicas=%v Replicas=%v ReadyReplicas=%v UpdatedReplicas=%v",
					deployments.Items[i].Name, desiredReplica, deployments.Items[i].Status.Replicas, deployments.Items[i].Status.ReadyReplicas, deployments.Items[i].Status.UpdatedReplicas))
			}
		}
		return kerrors.NewAggregate(errList)
	default:
		return errors.Errorf("invalid type: %s during VerifyMachineDeploymentsReplicas", reflect.TypeOf(deployments))
	}
}

// VerifyDeploymentAvailable verifies the deployment has at least one replica running under it or not
func VerifyDeploymentAvailable(obj runtime.Object) error {
	switch deployment := obj.(type) {
	case *appsv1.Deployment:
		if deployment.Status.AvailableReplicas < 1 {
			return errors.Errorf("pods are not yet running for deployment '%s' in namespace '%s'", deployment.Name, deployment.Namespace)
		}
	default:
		return errors.Errorf("invalid type: %s during VerifyDeploymentAvailable", reflect.TypeOf(deployment))
	}
	return nil
}

// VerifyAutoscalerDeploymentAvailable verifies autoscaler deployment's availability
func VerifyAutoscalerDeploymentAvailable(obj runtime.Object) error {
	switch deployment := obj.(type) {
	case *appsv1.Deployment:
		if *deployment.Spec.Replicas != deployment.Status.AvailableReplicas || *deployment.Spec.Replicas != deployment.Status.UpdatedReplicas || *deployment.Spec.Replicas != deployment.Status.Replicas {
			return errors.Errorf("pods are not yet running for deployment '%s' in namespace '%s'", deployment.Name, deployment.Namespace)
		}
	default:
		return errors.Errorf("invalid type: %s during VerifyAutoscalerDeploymentAvailable", reflect.TypeOf(deployment))
	}
	return nil
}

// VerifyCRSAppliedSuccessfully verifies that all CRS objects are applied successfully after cluster creation
func VerifyCRSAppliedSuccessfully(obj runtime.Object) error {
	switch crsList := obj.(type) {
	case *addonsv1.ClusterResourceSetList:
		errList := []error{}
		for i := range crsList.Items {
			if !conditions.IsTrue(crsList.Items[i].DeepCopy(), addonsv1.ResourcesAppliedCondition) {
				errList = append(errList, errors.Errorf("ClusterResourceSet %s is not yet applied", crsList.Items[i].Name))
			}
		}
		return kerrors.NewAggregate(errList)
	default:
		return errors.Errorf("invalid type: %s during VerifyCRSAppliedSuccessfully", reflect.TypeOf(crsList))
	}
}
