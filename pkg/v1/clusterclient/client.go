package clusterclient

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	runv1alpha1 "github.com/vmware-tanzu-private/core/apis/run/v1alpha1"
	"github.com/vmware-tanzu-private/core/pkg/v1/tkr/pkg/constants"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	crtclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

const (
	InfrastructureRefVSphere = "VSphereCluster"
	InfrastructureRefAWS     = "AWSCluster"
	InfrastructureRefAzure   = "AzureCluster"
)

type client struct {
	crtClient crtclient.Client
}

func NewClusterClient(kubeConfigPath string, context string) (Client, error) {

	var scheme = runtime.NewScheme()

	_ = runv1alpha1.AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)
	_ = capi.AddToScheme(scheme)

	config, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load kubeconfig from %s", kubeConfigPath)
	}
	if context != "" {
		config.CurrentContext = context
	}
	rawConfig, err := clientcmd.Write(*config)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to write config")
	}

	restConfig, err := clientcmd.RESTConfigFromKubeConfig(rawConfig)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to set up rest config")
	}

	mapper, err := apiutil.NewDynamicRESTMapper(restConfig, apiutil.WithLazyDiscovery)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to set up rest mapper")
	}

	crtClient, err := crtclient.New(restConfig, crtclient.Options{Scheme: scheme, Mapper: mapper})
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create cluster client")
	}

	return &client{
		crtClient: crtClient,
	}, nil
}

func (c *client) GetTanzuKubernetesReleases(tkrName string) ([]runv1alpha1.TanzuKubernetesRelease, error) {

	tkrList := &runv1alpha1.TanzuKubernetesReleaseList{}
	err := c.crtClient.List(context.Background(), tkrList)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list current TKRs")
	}
	if tkrName == "" {
		return tkrList.Items, nil
	}

	result := []runv1alpha1.TanzuKubernetesRelease{}
	for _, tkr := range tkrList.Items {
		if strings.HasPrefix(tkr.Name, tkrName) {
			result = append(result, tkr)
		}
	}
	return result, nil
}
func (c *client) GetBomConfigMap(tkrNameLabel string) (corev1.ConfigMap, error) {

	selectors := []crtclient.ListOption{
		crtclient.InNamespace(constants.TKRNamespace),
		crtclient.MatchingLabels(map[string]string{constants.BomConfigMapTKRLabel: tkrNameLabel}),
	}

	cmList := &corev1.ConfigMapList{}
	err := c.crtClient.List(context.Background(), cmList, selectors...)
	if err != nil {
		return corev1.ConfigMap{}, errors.Wrap(err, "failed to list current TKRs")
	}
	if len(cmList.Items) != 1 {
		return corev1.ConfigMap{}, errors.Wrapf(err, "failed to find the BOM ConfigMap matching the label %s: %v", tkrNameLabel, err)
	}

	return cmList.Items[0], nil
}

func (c *client) GetClusterInfrastructure() (string, error) {
	clusters := &capi.ClusterList{}

	selectors := []crtclient.ListOption{
		crtclient.MatchingLabels(map[string]string{constants.ManagememtClusterRoleLabel: ""}),
	}
	err := c.crtClient.List(context.Background(), clusters, selectors...)
	if err != nil || len(clusters.Items) != 1 {
		return "", errors.Wrap(err, "unable to get current management cluster")
	}

	return clusters.Items[0].Spec.InfrastructureRef.Kind, nil
}