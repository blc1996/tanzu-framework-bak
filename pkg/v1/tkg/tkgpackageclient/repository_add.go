// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package tkgpackageclient

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pkg/errors"

	kappctrl "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	kappipkg "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/packaging/v1alpha1"
	versions "github.com/vmware-tanzu/carvel-vendir/pkg/vendir/versions/v1alpha1"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/tkgpackagedatamodel"
)

// AddRepository validates the provided input and adds the package repository CR to the cluster
func (p *pkgClient) AddRepository(o *tkgpackagedatamodel.RepositoryOptions) error {
	if err := p.validateRepositoryAdd(o.RepositoryName, o.RepositoryURL, o.Namespace); err != nil {
		return err
	}

	if o.CreateNamespace {
		if err := p.createNamespace(o.Namespace); err != nil {
			return err
		}
	}

	newPackageRepo, err := p.newPackageRepository(o.RepositoryName, o.RepositoryURL, o.Namespace)
	if err != nil {
		return err
	}

	if err := p.kappClient.CreatePackageRepository(newPackageRepo); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to create package repository '%s' in namespace '%s'", o.RepositoryName, o.Namespace))
	}

	return nil
}

// newPackageRepository creates a new instance of the PackageRepository object
func (p *pkgClient) newPackageRepository(repositoryName, repositoryImg, namespace string) (*kappipkg.PackageRepository, error) {
	pkgr := &kappipkg.PackageRepository{
		TypeMeta:   metav1.TypeMeta{APIVersion: tkgpackagedatamodel.DefaultAPIVersion, Kind: tkgpackagedatamodel.KindPackageRepository},
		ObjectMeta: metav1.ObjectMeta{Name: repositoryName, Namespace: namespace},
		Spec: kappipkg.PackageRepositorySpec{Fetch: &kappipkg.PackageRepositoryFetch{
			ImgpkgBundle: &kappctrl.AppFetchImgpkgBundle{Image: repositoryImg},
		}},
	}

	_, tag, err := parseRegistryImageUrl(repositoryImg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse OCI registry URL")
	}

	if tag == "" {
		pkgr.Spec.Fetch.ImgpkgBundle.TagSelection = &versions.VersionSelection{
			Semver: &versions.VersionSelectionSemver{
				Constraints: defaultImageTagConstraint,
			},
		}
	}
	return pkgr, nil
}

// validateRepositoryAdd ensures that another repository (with the same name or same OCI registry URL) does not already exist in the cluster
func (p *pkgClient) validateRepositoryAdd(repositoryName, repositoryImg, namespace string) error {
	repositoryList, err := p.kappClient.ListPackageRepositories(namespace)
	if err != nil {
		return errors.Wrap(err, "failed to list package repositories")
	}

	for _, repository := range repositoryList.Items { //nolint:gocritic
		if repository.Name == repositoryName {
			return errors.New("repository with the same name already exists")
		}

		if repository.Spec.Fetch != nil && repository.Spec.Fetch.ImgpkgBundle != nil &&
			repository.Spec.Fetch.ImgpkgBundle.Image == repositoryImg {
			return errors.New("repository with the same OCI registry URL already exists")
		}
	}

	return nil
}
