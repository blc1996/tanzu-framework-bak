module github.com/vmware-tanzu-private/core/addons

go 1.14

require (
	github.com/go-logr/logr v0.3.0
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.3
	github.com/pkg/errors v0.9.1
	github.com/vmware-tanzu-private/core v0.0.0-20201211200158-5838874f2c38
	github.com/vmware-tanzu/carvel-kapp-controller v0.13.0
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800
	sigs.k8s.io/cluster-api v0.3.11-0.20201204161359-8437691189ad
	sigs.k8s.io/controller-runtime v0.7.0
)

replace github.com/vmware-tanzu-private/core => ../
