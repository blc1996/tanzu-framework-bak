// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu-private/core/apis/client/v1alpha1"
	"github.com/vmware-tanzu-private/core/pkg/v1/client"
	"github.com/vmware-tanzu-private/tkg-cli/pkg/tkgctl"
	"github.com/vmware-tanzu-private/tkg-cli/pkg/types"
)

func newTKGCtlClient() (tkgctl.TKGClient, error) {
	tkgConfigDir, err := getTKGConfigDir()
	if err != nil {
		return nil, errors.Wrap(err, "unable to get default TKG config directory")
	}

	return tkgctl.New(tkgctl.Options{
		ConfigDir: tkgConfigDir,
		CustomizerOptions: types.CustomizerOptions{
			RegionManagerFactory: NewFactory(),
		},
		LogOptions: tkgctl.LoggingOptions{Verbosity: logLevel, File: logFile},
	})
}

func getTKGConfigDir() (string, error) {
	tanzuConfigDir, err := client.LocalDir()
	if err != nil {
		return "", errors.Wrap(err, "unable to get home directory")
	}
	return filepath.Join(tanzuConfigDir, "tkg"), nil
}

func runForCurrentMC(fn func(currServ *v1alpha1.Server) error) error {
	currServ, err := client.GetCurrentServer()
	if err != nil {
		return err
	}

	if !currServ.IsManagementCluster() {
		return errors.Errorf("Current server [%s] is not a management cluster", currServ.Name)
	}

	return fn(currServ)
}
