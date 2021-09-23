// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package tkgpackageclient

import (
	"fmt"
	versions "github.com/vmware-tanzu/carvel-vendir/pkg/vendir/versions/v1alpha1"

	"github.com/pkg/errors"
	k8serror "k8s.io/apimachinery/pkg/api/errors"

	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/tkgpackagedatamodel"
)

func (p *pkgClient) UpdateRepository(o *tkgpackagedatamodel.RepositoryOptions) error {
	existingRepository, err := p.kappClient.GetPackageRepository(o.RepositoryName, o.Namespace)
	if err != nil && !k8serror.IsNotFound(err) {
		return err
	}

	if existingRepository != nil {
		repositoryToUpdate := existingRepository.DeepCopy()
		// TODO: add package repository update validation, i.e. updated namespaced name and imageURL can't be duplicated

		_, tag, err := ParseImageUrl(o.RepositoryURL)
		if err != nil {
			return errors.Wrap(err, "invalid repository image URL")
		}

		found, err := checkPackageRepositoryTagSelection()
		if err != nil {
			return errors.Wrap(err, "failed to check package repository resource version")
		}

		repositoryToUpdate.Spec.Fetch.ImgpkgBundle.Image = o.RepositoryURL
		if found && tag == ""  {
			repositoryToUpdate.Spec.Fetch.ImgpkgBundle.TagSelection = &versions.VersionSelection{
				Semver: &versions.VersionSelectionSemver{
					Constraints: defaultImageTagConstraint,
				},
			}
		}

		if err := p.kappClient.UpdatePackageRepository(repositoryToUpdate); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to update package repository '%s' in namespace '%s'", o.RepositoryName, o.Namespace))
		}
	} else if o.CreateRepository {
		if err := p.AddRepository(o); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to create package repository '%s' in namespace '%s'", o.RepositoryName, o.Namespace))
		}
	} else {
		return errors.Wrap(err, fmt.Sprintf("failed to find package repository '%s' in namespace '%s'", o.RepositoryName, o.Namespace))
	}

	return nil
}
